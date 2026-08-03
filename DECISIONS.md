# Design Decisions

This document explains the implementation choices made while porting the original JavaScript library to Go.

---

## API Design

The original library exposes a single JavaScript function that behaves differently depending on the input type.

Go is statically typed and does not support this style naturally.

Instead, the implementation provides two explicit functions:

- Parse(string)
- Format(float64)

This follows Go conventions while preserving the original behavior.

---

## Numeric Types

JavaScript represents numbers using the Number type.

Go distinguishes integer and floating-point types.

Formatting operations use float64, while Parse returns int64 to represent byte counts.

---

## Error Handling

The JavaScript implementation throws exceptions for invalid inputs.

Go returns an error value following idiomatic Go practices.

---

## Regular Expressions

The original regular expression was translated using Go's regexp package while preserving matching behavior.

---

## Unit Mapping

Units are represented internally using a lookup table.

Supported units:

- B
- KB
- MB
- GB
- TB
- PB

Each unit uses binary multiples (1024).

---

## Testing

Unit tests were written in Go to validate behavior equivalent to the original implementation.

Additional edge cases were added to improve confidence in the port.

---

## Known Differences

The Go implementation follows idiomatic Go design and therefore differs slightly from JavaScript in API style while maintaining equivalent functionality.
