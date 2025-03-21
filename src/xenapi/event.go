/*
 * Copyright (c) Cloud Software Group, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *   1) Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *
 *   2) Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials
 *      provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package xenapi

import (
	"fmt"
	"time"
)

type EventRecord struct {
	// The record of the database object that was added, changed or deleted
	Snapshot RecordInterface `json:"snapshot,omitempty"`
	// An ID, monotonically increasing, and local to the current session
	ID int `json:"id,omitempty"`
	// The time at which the event occurred
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The name of the class of the object that changed
	Class string `json:"class,omitempty"`
	// The operation that was performed
	Operation EventOperation `json:"operation,omitempty"`
	// A reference to the object that changed
	Ref string `json:"ref,omitempty"`
	// The uuid of the object that changed
	ObjUUID string `json:"obj_uuid,omitempty"`
}

type EventRef string

type RecordInterface interface{}

type EventBatch struct {
	Token          string         `json:"token,omitempty"`
	ValidRefCounts map[string]int `json:"validRefCounts,omitempty"`
	Events         []EventRecord  `json:"events,omitempty"`
}

// Asynchronous event registration and handling
type event struct{}

var Event event

// Inject: Injects an artificial event on the given object and returns the corresponding ID in the form of a token, which can be used as a point of reference for database events. For example, to check whether an object has reached the right state before attempting an operation, one can inject an artificial event on the object and wait until the token returned by consecutive event.from calls is lexicographically greater than the one returned by event.inject.
// Version: tampa
func (event) Inject(session *Session, class string, ref string) (retval string, err error) {
	method := "event.inject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "class"), class)
	if err != nil {
		return
	}
	refArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "ref"), ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classArg, refArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// Inject3: Injects an artificial event on the given object and returns the corresponding ID in the form of a token, which can be used as a point of reference for database events. For example, to check whether an object has reached the right state before attempting an operation, one can inject an artificial event on the object and wait until the token returned by consecutive event.from calls is lexicographically greater than the one returned by event.inject.
// Version: tampa
func (event) Inject3(session *Session, class string, ref string) (retval string, err error) {
	method := "event.inject"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "class"), class)
	if err != nil {
		return
	}
	refArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "ref"), ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classArg, refArg)
	if err != nil {
		return
	}
	retval, err = deserializeString(method+" -> ", result)
	return
}

// GetCurrentID: Return the ID of the next event to be generated by the system
// Version: rio
func (event) GetCurrentID(session *Session) (retval int, err error) {
	method := "event.get_current_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// GetCurrentID1: Return the ID of the next event to be generated by the system
// Version: rio
func (event) GetCurrentID1(session *Session) (retval int, err error) {
	method := "event.get_current_id"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeInt(method+" -> ", result)
	return
}

// From: Blocking call which returns a new token and a (possibly empty) batch of events. The returned token can be used in subsequent calls to this function.
// Version: boston
//
// Errors:
// SESSION_NOT_REGISTERED - This session is not registered to receive events. You must call event.register before event.next. The session handle you are using is echoed.
// EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (event) From(session *Session, classes []string, token string, timeout float64) (retval EventBatch, err error) {
	method := "event.from"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	tokenArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token"), token)
	if err != nil {
		return
	}
	timeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "timeout"), timeout)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg, tokenArg, timeoutArg)
	if err != nil {
		return
	}
	retval, err = deserializeEventBatch(method+" -> ", result)
	return
}

// From4: Blocking call which returns a new token and a (possibly empty) batch of events. The returned token can be used in subsequent calls to this function.
// Version: boston
//
// Errors:
// SESSION_NOT_REGISTERED - This session is not registered to receive events. You must call event.register before event.next. The session handle you are using is echoed.
// EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (event) From4(session *Session, classes []string, token string, timeout float64) (retval EventBatch, err error) {
	method := "event.from"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	tokenArg, err := serializeString(fmt.Sprintf("%s(%s)", method, "token"), token)
	if err != nil {
		return
	}
	timeoutArg, err := serializeFloat(fmt.Sprintf("%s(%s)", method, "timeout"), timeout)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg, tokenArg, timeoutArg)
	if err != nil {
		return
	}
	retval, err = deserializeEventBatch(method+" -> ", result)
	return
}

// Next: Blocking call which returns a (possibly empty) batch of events. This method is only recommended for legacy use. New development should use event.from which supersedes this method.
// Version: rio
//
// Errors:
// SESSION_NOT_REGISTERED - This session is not registered to receive events. You must call event.register before event.next. The session handle you are using is echoed.
// EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (event) Next(session *Session) (retval []EventRecord, err error) {
	method := "event.next"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeEventRecordSet(method+" -> ", result)
	return
}

// Next1: Blocking call which returns a (possibly empty) batch of events. This method is only recommended for legacy use. New development should use event.from which supersedes this method.
// Version: rio
//
// Errors:
// SESSION_NOT_REGISTERED - This session is not registered to receive events. You must call event.register before event.next. The session handle you are using is echoed.
// EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (event) Next1(session *Session) (retval []EventRecord, err error) {
	method := "event.next"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg)
	if err != nil {
		return
	}
	retval, err = deserializeEventRecordSet(method+" -> ", result)
	return
}

// Unregister: Removes this session&apos;s registration with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) Unregister(session *Session, classes []string) (err error) {
	method := "event.unregister"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, classesArg)
	return
}

// AsyncUnregister: Removes this session&apos;s registration with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) AsyncUnregister(session *Session, classes []string) (retval TaskRef, err error) {
	method := "Async.event.unregister"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Unregister2: Removes this session&apos;s registration with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) Unregister2(session *Session, classes []string) (err error) {
	method := "event.unregister"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, classesArg)
	return
}

// AsyncUnregister2: Removes this session&apos;s registration with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) AsyncUnregister2(session *Session, classes []string) (retval TaskRef, err error) {
	method := "Async.event.unregister"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Register: Registers this session with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) Register(session *Session, classes []string) (err error) {
	method := "event.register"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, classesArg)
	return
}

// AsyncRegister: Registers this session with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) AsyncRegister(session *Session, classes []string) (retval TaskRef, err error) {
	method := "Async.event.register"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}

// Register2: Registers this session with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) Register2(session *Session, classes []string) (err error) {
	method := "event.register"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	_, err = session.client.sendCall(method, sessionIDArg, classesArg)
	return
}

// AsyncRegister2: Registers this session with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next.
// Version: rio
func (event) AsyncRegister2(session *Session, classes []string) (retval TaskRef, err error) {
	method := "Async.event.register"
	sessionIDArg, err := serializeSessionRef(fmt.Sprintf("%s(%s)", method, "session_id"), session.ref)
	if err != nil {
		return
	}
	classesArg, err := serializeStringSet(fmt.Sprintf("%s(%s)", method, "classes"), classes)
	if err != nil {
		return
	}
	result, err := session.client.sendCall(method, sessionIDArg, classesArg)
	if err != nil {
		return
	}
	retval, err = deserializeTaskRef(method+" -> ", result)
	return
}
