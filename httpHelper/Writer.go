package httpHelper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

// Writer HTTP Writer
type Writer struct {
	HTTPWriter    http.ResponseWriter
	ContentType   ContentType
	safeResponses bool
}

// NewWriter new responder
func NewWriter(writer http.ResponseWriter, contentType ContentType) Writer {
	return Writer{writer, contentType, false}
}

// NewWriterJSON New responder JSON
func NewWriterJSON(writer http.ResponseWriter) Writer {
	return Writer{writer, ContentTypeJSON, false}
}

// NewWriterText New responder text
func NewWriterText(writer http.ResponseWriter) Writer {
	return Writer{writer, ContentTypeText, false}
}

func (r *Writer) SetContentType(contentType ContentType) {
	r.ContentType = contentType
}

func (r *Writer) SafeResponses() {
	r.safeResponses = true
}

func (r Writer) getError(err ErrorResponse) []byte {
	return r.getResponseText(err)
}

func (r Writer) getResponseText(resp interface{}) []byte {
	result := []byte{}
	if resp == nil {
		return []byte("")
	} else if IsErrorResponse(resp) && resp.(*ErrorResponse).Message == "" {
		resp.(*ErrorResponse).SetMessage(http.StatusText(resp.(*ErrorResponse).StatusCode))
	}

	if r.ContentType.Is(ContentTypeJSON) {
		j, _ := json.Marshal(resp)
		return j
	} else if r.ContentType.Is(ContentTypeText) {
		switch t := resp.(type) {
		case *ErrorResponse:
			return []byte(resp.(*ErrorResponse).Message)
		case string:
			return []byte(resp.(string))
		case bool:
			return []byte(strconv.FormatBool(resp.(bool)))
		case float32:
			return []byte(strconv.FormatFloat(resp.(float64), 'f', -1, 32))
		case float64:
			return []byte(strconv.FormatFloat(resp.(float64), 'f', -1, 64))
		case int:
			return []byte(strconv.FormatInt(int64(resp.(int)), 10))
		case int8:
			return []byte(strconv.FormatInt(int64(resp.(int8)), 10))
		case int16:
			return []byte(strconv.FormatInt(int64(resp.(int16)), 10))
		case int32:
			return []byte(strconv.FormatInt(int64(resp.(int32)), 10))
		case int64:
			return []byte(strconv.FormatInt(resp.(int64), 10))
		case uint:
			return []byte(strconv.FormatUint(uint64(resp.(uint)), 10))
		case uint8:
			return []byte(strconv.FormatUint(uint64(resp.(uint8)), 10))
		case uint16:
			return []byte(strconv.FormatUint(uint64(resp.(uint16)), 10))
		case uint32:
			return []byte(strconv.FormatUint(uint64(resp.(uint32)), 10))
		case uint64:
			return []byte(strconv.FormatUint(resp.(uint64), 10))
		case []byte:
			return resp.([]byte)
		case StringResponse:
			return []byte(resp.(StringResponse).String())
		default:
			var buf bytes.Buffer
			buf.WriteString("Unsupported text type ")
			buf.WriteString(reflect.TypeOf(t).String())
			buf.WriteString("; Use JSON response")
			return buf.Bytes()
		}
	}
	return result
}

func (r Writer) write(a interface{}, code int) {
	if code >= 200 && code < 300 && r.safeResponses {
		a = SafeResponse{StatusCode: code, Status: http.StatusText(code), Data: a}
	}
	r.HTTPWriter.WriteHeader(code)
	r.HTTPWriter.Write(r.getResponseText(a))
}

// Ok Return ok
func (r Writer) Ok(a interface{}) {
	r.write(a, http.StatusOK)
}

// Created created
func (r Writer) Created(a interface{}) {
	r.write(a, http.StatusCreated)
}

// Accepted accepted
func (r Writer) Accepted(a interface{}) {
	r.write(a, http.StatusAccepted)
}

// NonAuthoritativeInformation non authoritative information
func (r Writer) NonAuthoritativeInformation(a interface{}) {
	r.write(a, http.StatusNonAuthoritativeInfo)
}

// NoContent no content
func (r Writer) NoContent(a interface{}) {
	r.write(a, http.StatusNoContent)
}

// PartialContent partial content
func (r Writer) PartialContent(a interface{}) {
	r.write(a, http.StatusPartialContent)
}

// MultipleChoices multiple choices
func (r Writer) MultipleChoices(a interface{}) {
	r.write(a, http.StatusMultipleChoices)
}

// MovedPermanently moved permanently
func (r Writer) MovedPermanently(a interface{}) {
	r.write(a, http.StatusMovedPermanently)
}

// Found found
func (r Writer) Found(a interface{}) {
	r.write(a, http.StatusFound)
}

// SeeOther see other
func (r Writer) SeeOther(a interface{}) {
	r.write(a, http.StatusSeeOther)
}

// NotModified not modified
func (r Writer) NotModified(a interface{}) {
	r.write(a, http.StatusNotModified)
}

// UseProxy use proxy
func (r Writer) UseProxy(a interface{}) {
	r.write(a, http.StatusUseProxy)
}

// TemporaryRedirect temporary redirect
func (r Writer) TemporaryRedirect(a interface{}) {
	r.write(a, http.StatusTemporaryRedirect)
}

// Error Return custom error
func (r Writer) Error(err string, code int, errorData interface{}) {
	errResponse := &ErrorResponse{code, err, errorData}
	r.write(errResponse, code)
}

// BadRequest Respond to an HTTP request with BadRequest
func (r Writer) BadRequest(err string, data interface{}) {
	r.Error(err, http.StatusBadRequest, data)
}

// Unauthorized Return unauthorized
func (r Writer) Unauthorized(err string, data interface{}) {
	r.Error(err, http.StatusUnauthorized, data)
}

// PaymentRequired payment required
func (r Writer) PaymentRequired(err string, data interface{}) {
	r.Error(err, http.StatusPaymentRequired, data)
}

// NotFound return not found
func (r Writer) NotFound(err string, data interface{}) {
	r.Error(err, http.StatusNotFound, data)
}

// MethodNotAllowed method not allowed
func (r Writer) MethodNotAllowed(err string, data interface{}) {
	r.Error(err, http.StatusMethodNotAllowed, data)
}

// NotAcceptable not acceptable
func (r Writer) NotAcceptable(err string, data interface{}) {
	r.Error(err, http.StatusNotAcceptable, data)
}

// ProxyAuthenticationRequired proxy authentication required
func (r Writer) ProxyAuthenticationRequired(err string, data interface{}) {
	r.Error(err, http.StatusProxyAuthRequired, data)
}

// RequestTimeout request timeout
func (r Writer) RequestTimeout(err string, data interface{}) {
	r.Error(err, http.StatusRequestTimeout, data)
}

// Conflict conflict
func (r Writer) Conflict(err string, data interface{}) {
	r.Error(err, http.StatusConflict, data)
}

// Gone gone
func (r Writer) Gone(err string, data interface{}) {
	r.Error(err, http.StatusGone, data)
}

// LengthRequired length required
func (r Writer) LengthRequired(err string, data interface{}) {
	r.Error(err, http.StatusLengthRequired, data)
}

// PreconditionFailed precondition failed
func (r Writer) PreconditionFailed(err string, data interface{}) {
	r.Error(err, http.StatusPreconditionFailed, data)
}

// RequestEntityTooLarge request entity too large
func (r Writer) RequestEntityTooLarge(err string, data interface{}) {
	r.Error(err, http.StatusRequestEntityTooLarge, data)
}

// RequestURITooLong request URI too long
func (r Writer) RequestURITooLong(err string, data interface{}) {
	r.Error(err, http.StatusRequestURITooLong, data)
}

// UnsupportedMediaType unsupported media type
func (r Writer) UnsupportedMediaType(err string, data interface{}) {
	r.Error(err, http.StatusUnsupportedMediaType, data)
}

// RequestedRangeNotSatisfiable requested range not satisfied
func (r Writer) RequestedRangeNotSatisfiable(err string, data interface{}) {
	r.Error(err, http.StatusRequestedRangeNotSatisfiable, data)
}

// ExpectationFailed expectation failed
func (r Writer) ExpectationFailed(err string, data interface{}) {
	r.Error(err, http.StatusExpectationFailed, data)
}

// InternalServerError Return internal server error
func (r Writer) InternalServerError(err string, data interface{}) {
	r.Error(err, http.StatusInternalServerError, data)
}

// NotImplemented not implemented
func (r Writer) NotImplemented(err string, data interface{}) {
	r.Error(err, http.StatusNotImplemented, data)
}

// BadGateway bad gateway
func (r Writer) BadGateway(err string, data interface{}) {
	r.Error(err, http.StatusBadGateway, data)
}

// ServiceUnavailable service unavailable
func (r Writer) ServiceUnavailable(err string, data interface{}) {
	r.Error(err, http.StatusServiceUnavailable, data)
}

// GatewayTimeout gateway timeout
func (r Writer) GatewayTimeout(err string, data interface{}) {
	r.Error(err, http.StatusGatewayTimeout, data)
}

// HTTPVersionNotSupported HTTP version not supported
func (r Writer) HTTPVersionNotSupported(err string, data interface{}) {
	r.Error(err, http.StatusHTTPVersionNotSupported, data)
}
