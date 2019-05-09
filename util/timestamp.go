// Copyright (C) MongoDB, Inc. 2019-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TimestampGreaterThan returns true if lhs comes after rhs, false otherwise.
func TimestampGreaterThan(lhs, rhs primitive.Timestamp) bool {
	return lhs.T > rhs.T || lhs.T == rhs.T && lhs.I > rhs.I
}

// CompareTimestamps returns a positive number if lhs > rhs, a negative number if lhs < rhs, and zero if lhs == rhs.
func CompareTimestamps(lhs, rhs primitive.Timestamp) int {
	if lhs.T > rhs.T {
		return 1
	}
	if lhs.T < rhs.T {
		return -1
	}
	if lhs.I > rhs.I {
		return 1
	}
	if lhs.I < rhs.I {
		return -1
	}
	return 0
}
