package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"vdo-be/pkg/api"
)

func decode[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var req T

	contentType := r.Header.Get("Content-Type")
	switch {
	case contentType == mimeApplicationJSON:
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB
	case strings.HasPrefix(contentType, mimeMultipartFormData):
		limitFileSize := int64(50 << 20) // 50 MB
		r.Body = http.MaxBytesReader(w, r.Body, limitFileSize)
		err := r.ParseMultipartForm(limitFileSize)
		if errors.As(err, new(*http.MaxBytesError)) {
			return req, errExceedFileSize
		}
	}

	err := SetRequestValue(r, &req)
	return req, err
}

func SetRequestValue[T any](r *http.Request, req *T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	values := reflect.ValueOf(req)
	if values.Kind() != reflect.Pointer {
		return errNotStructPointer // Should not be possible, Constraint by type
	}
	fields := reflect.TypeOf(req).Elem()

	num := fields.NumField()
	for i := 0; i < num; i++ {

		field := fields.Field(i)
		value := values.Elem().Field(i)

		val, ok := getRequestValueByTag(r, field)
		if !ok || val == "" {
			continue
		}

		err = setFieldValue(value, val)
		if err != nil {
			return err
		}
	}

	if !methodHaveBody(r.Method) {
		return nil
	}

	contentType := r.Header.Get("Content-Type")
	switch {
	case contentType == mimeApplicationJSON:
		return json.NewDecoder(r.Body).Decode(req)
	case strings.HasPrefix(contentType, mimeMultipartFormData):
		return setStructFile(r, req)
	}
	return nil
}

func getRequestValueByTag(r *http.Request, field reflect.StructField) (string, bool) {
	if tag, ok := field.Tag.Lookup("path"); ok {
		return r.PathValue(tag), true
	}
	if tag, ok := field.Tag.Lookup("header"); ok {
		return r.Header.Get(tag), true
	}
	if tag, ok := field.Tag.Lookup("query"); ok {
		return r.URL.Query().Get(tag), true
	}
	if tag, ok := field.Tag.Lookup("form"); ok {
		return r.PostFormValue(tag), true
	}
	return "", false
}

func setFieldValue(value reflect.Value, val string) error {
	if !value.CanSet() {
		return nil
	}

	switch value.Kind() {
	case reflect.String:
		value.SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		value.SetFloat(v)
	case reflect.Bool:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		value.SetBool(v)
	case reflect.Pointer:
		zero := reflect.New(value.Type().Elem())
		value.Set(zero)

		decoder, ok := value.Interface().(rpcDecoder)
		if ok {
			return decoder.DecodeRPC([]byte(val))
		}
	case reflect.Struct:
		decoder, ok := value.Addr().Interface().(rpcDecoder)
		if ok {
			return decoder.DecodeRPC([]byte(val))
		}
	}
	return nil
}

func setStructFile[T any](r *http.Request, req *T) (err error) {

	values := reflect.ValueOf(req)
	fields := reflect.TypeOf(req).Elem()

	num := fields.NumField()
	for i := 0; i < num; i++ {

		field := fields.Field(i)
		value := values.Elem().Field(i)

		tag, ok := field.Tag.Lookup("file")
		if !ok {
			continue
		}

		if reflect.TypeOf(api.File{}) != value.Type() {
			return errNotFileType
		}

		f, h, err := r.FormFile(tag)
		if err != nil {
			return err
		}
		file := api.File{
			File:       f,
			FileHeader: h,
		}

		v := reflect.ValueOf(file)
		value.Set(v)
	}

	return nil
}

func methodHaveBody(m string) bool {
	return m == http.MethodPost || m == http.MethodPut || m == http.MethodPatch
}

type rpcDecoder interface {
	DecodeRPC([]byte) error
}

const (
	mimeApplicationJSON   = "application/json"
	mimeMultipartFormData = "multipart/form-data"
)

var (
	errExceedFileSize   = errors.New("file size limit exceeded")
	errNotFileType      = errors.New("struct field is not api.File")
	errNotStructPointer = errors.New("not a pointer to struct")
)
