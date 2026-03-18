- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2020-03-12
- RFC PR: N/A
- Tracking Issue: N/A

# Proposal: Loose mode

## Background

Current [storage]'s pair handle behavior is inconsistent.

In all `parseStoragePairXXX` functions, we will ignore not supported pairs via only pick supported one:

```go
v, ok = values[ps.DirFunc]
if ok {
    result.HasDirFunc = true
    result.DirFunc = v.(types.ObjectFunc)
}
```

But in other pair related logic, like `storage_class` support, we also returned errors:

```go
func parseStorageClass(in storageclass.Type) (string, error) {
	switch in {
	case storageclass.Hot:
		return storageClassStandard, nil
	case storageclass.Warm:
		return storageClassStandardIA, nil
	default:
		return "", &services.PairError{
			Op:    "parse storage class",
			Err:   services.ErrStorageClassNotSupported,
			Pairs: []*types.Pair{{Key: ps.StorageClass, Value: in}},
		}
	}
}
```

So users could be confused how we handle our compatibility related issues.

## Proposal

So I propose that all a `loose` mode for all services. `loose` mode will be `off` as default, and services will return error when they reach incompatible place. And when `loose` is on, all incompatible error will be ignored.

For example:

We have a Storager who doesn't support `Size` pair in `Read`.

`loose` on: This error will be ignored.
`loose` off: Storager returns a compatibility related error.

Currently, we mixed compatibility error and other pair related error in `PairError`. We will add two different error: `ErrCapabilityInsufficient` and `ErrRestrictionDissatisfied`.

`ErrCapabilityInsufficient` means this service doesn't have this capability, and `ErrRestrictionDissatisfied` means this operation doesn't meat service's restriction. `ErrCapabilityInsufficient` could be ignored safely if you don't care much about service behavior consistency, and will be ignored in loose mode.

Based on these errors, we will have new error structs like `PairRequiredError` to carry error contexts:

```go
// NewPairRequiredError will create a new PairRequiredError.
func NewPairRequiredError(keys ...string) *PairRequiredError {
	return &PairRequiredError{
		Err:  ErrRestrictionDissatisfied,
		Keys: keys,
	}
}

// PairRequiredError means this operation has required pair but missing.
type PairRequiredError struct {
	Err error

	Keys []string
}

func (e *PairRequiredError) Error() string {
	return fmt.Sprintf("pair required, %v: %s", e.Keys, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *PairRequiredError) Unwrap() error {
	return e.Err
}
```

## Rationale

None.

## Compatibility

- More `ErrCapabilityInsufficient` could be returned as `loose` mode will be on as default
- Some error could be returned as other error structs instead of `PairError`
- `PairError` will be removed

## Implementation

Most of the work would be done by the author of this proposal.

[storage]: https://github.com/Xuanwo/storage


## License

Copyright 2024 go-storage authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
