## 0.4.1 (2020-07-14)

Misc:

- Improve example codes for understandability.

## 0.4.0 (2020-07-14)

Enhance: ([#3](https://github.com/progrhyme/go-lv/pull/3))

- Add fmt.Print() style functions which don't have suffix "f": Trace, Debug, Info, Notice, Warn, Error and Critical who takes `...interface{}` arguments.
- Add functions (*Logger) GetLevel and (*Logger) SetLevel to be usable as interface.

Change: ([#3](https://github.com/progrhyme/go-lv/pull/3))

- Obsolete interface Loggable.
- Instead, add three types of interface: Minimal, Standard and Granular so that one can choose log level granularity.

Bug Fix: ([#3](https://github.com/progrhyme/go-lv/pull/3))

- When flags `log.Llongfile` or `log.Lshortfile`, the caller information couldn't displayed properly.

## 0.3.1 (2020-07-09)

Modify:

- Add func `Criticalf()` to `Loggable` interface

## 0.3.0 (2020-07-09)

Enhance: ([#2](https://github.com/progrhyme/go-lv/pull/2))

- Add new level `LCritical`

Change: ([#2](https://github.com/progrhyme/go-lv/pull/2))

- Rename `LWarning` -> `LWarn`

## 0.2.0 (2020-07-08)

Changes: ([#1](https://github.com/progrhyme/go-lv/pull/1))

- Rename `Level` constants and their string representations. i.e. `Trace` -> `LTrace`, `"Trace"` -> `"TRACE"`
- Rename func `LabelToLevel()` -> `WordToLevel()` and make the argument case-insensitive

## 0.1.1 (2020-07-08)

Brush up documentation (No code change).

## 0.1.0 (2020-07-07)

Initial release.
