// Copyright (C) MongoDB, Inc. 2019-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package util

import (
	"testing"

	"github.com/mongodb/mongo-tools-common/testtype"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTimestampGreaterThan(t *testing.T) {
	testtype.SkipUnlessTestType(t, testtype.UnitTestType)
	reference := primitive.Timestamp{T: 5, I: 0}

	Convey("With some sample values", t, func() {
		Convey("different T's should compare correctly", func() {
			So(TimestampGreaterThan(primitive.Timestamp{T: 1000, I: 0}, reference), ShouldBeTrue)
			So(TimestampGreaterThan(reference, primitive.Timestamp{T: 1000, I: 0}), ShouldBeFalse)
		})

		Convey("matching T's should compare correctly", func() {
			So(TimestampGreaterThan(primitive.Timestamp{T: 5, I: 1}, reference), ShouldBeTrue)
			So(TimestampGreaterThan(reference, primitive.Timestamp{T: 5, I: 1}), ShouldBeFalse)
		})

		Convey("equal timestamps should compare correctly", func() {
			So(TimestampGreaterThan(reference, reference), ShouldBeFalse)
		})
	})
}

func TestCompareTimestamps(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs primitive.Timestamp
		expected int
	}{
		{"equal", primitive.Timestamp{5, 5}, primitive.Timestamp{5, 5}, 0},
		{"lhs T greater", primitive.Timestamp{10, 5}, primitive.Timestamp{5, 5}, 1},
		{"lhs I greater", primitive.Timestamp{5, 10}, primitive.Timestamp{5, 5}, 1},
		{"lhs both greater", primitive.Timestamp{10, 10}, primitive.Timestamp{5, 5}, 1},
		{"rhs T greater", primitive.Timestamp{5, 5}, primitive.Timestamp{10, 5}, -1},
		{"rhs I greater", primitive.Timestamp{5, 5}, primitive.Timestamp{5, 10}, -1},
		{"rhs both greater", primitive.Timestamp{5, 5}, primitive.Timestamp{10, 10}, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if res := CompareTimestamps(tc.lhs, tc.rhs); res != tc.expected {
				t.Fatalf("result mismatch; expected %d, got %d", tc.expected, res)
			}
		})
	}
}
