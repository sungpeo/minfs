// This file is part of MinFS
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

/*
 * MinFS - fuse driver for Object Storage (C) 2016 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package minfs

// Operation -
type Operation struct {
	Error chan error
}

// MoveOperation - Move source object to target object. Copy source to target, delete the source.
type MoveOperation struct {
	*Operation

	Source string
	Target string
}

func newMoveOp(sourcePath, targetPath string) MoveOperation {
	return MoveOperation{
		Source: sourcePath,
		Target: targetPath,
		Operation: &Operation{
			Error: make(chan error),
		},
	}
}

// CopyOperation - Copy source object to target.
type CopyOperation struct {
	*Operation

	Source string
	Target string
}

// PutOperation - Copy source file to target.
type PutOperation struct {
	*Operation

	Length int64

	Source string
	Target string
}

func newPutOp(sourcePath string, targetPath string, length int64) PutOperation {
	return PutOperation{
		Source: sourcePath,
		Target: targetPath,
		Length: int64(length),
		Operation: &Operation{
			Error: make(chan error),
		},
	}
}
