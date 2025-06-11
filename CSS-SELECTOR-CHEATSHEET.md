# CSS Selectors Cheat Sheet :)

## Basic Selectors

| Selector  | Meaning                               | Example        |
| --------- | ------------------------------------- | -------------- |
| `*`       | Select all elements                   | `*`            |
| `element` | Select all `<element>`                | `div`          |
| `#id`     | Select element with specific `id`     | `#header`      |
| `.class`  | Select elements with specific `class` | `.btn-primary` |

---

## Combinators

| Selector | Meaning                               | Example   |
| -------- | ------------------------------------- | --------- |
| `A B`    | Select all `B` inside `A`             | `div p`   |
| `A > B`  | Select `B` directly inside `A`        | `ul > li` |
| `A + B`  | Select `B` immediately after `A`      | `h1 + p`  |
| `A ~ B`  | Select all `B` after `A` (same level) | `div ~ p` |

---

## Pseudo-classes

| Selector               | Meaning                              | Example                  |
| ---------------------- | ------------------------------------ | ------------------------ |
| `:first-child`         | First child of its parent            | `li:first-child`         |
| `:last-child`          | Last child of its parent             | `p:last-child`           |
| `:nth-child(n)`        | nth child of its parent              | `tr:nth-child(odd)`      |
| `:nth-last-child(n)`   | nth child from the end               | `li:nth-last-child(2)`   |
| `:only-child`          | Single child of its parent           | `div:only-child`         |
| `:first-of-type`       | First element of its type            | `h1:first-of-type`       |
| `:last-of-type`        | Last element of its type             | `p:last-of-type`         |
| `:nth-of-type(n)`      | nth element of its type              | `td:nth-of-type(3)`      |
| `:nth-last-of-type(n)` | nth element of its type from the end | `tr:nth-last-of-type(2)` |
| `:empty`               | Select elements with no children     | `div:empty`              |
| `:not(selector)`       | Select elements that don't match     | `a:not(.active)`         |
| `:checked`             | Checked inputs (checkbox, radio)     | `input:checked`          |
| `:disabled`            | Disabled form elements               | `button:disabled`        |
| `:enabled`             | Enabled form elements                | `input:enabled`          |

---

## Attribute Selectors

| Selector        | Meaning                                      | Example                                         |
| --------------- | -------------------------------------------- | ----------------------------------------------- |
| `[attr]`        | Elements with `attr`                         | `[type]`                                        |
| `[attr=value]`  | Elements with `attr` equal to `value`        | `[type="checkbox"]`                             |
| `[attr~=value]` | `attr` contains value (space-separated list) | `[class~="btn"]`                                |
| `[attr|=value]` | `attr` is exactly value or starts with `value-` | `[lang|="en"]`                                |
| `[attr^=value]` | `attr` starts with value                     | `[href^="https"]`                               |
| `[attr$=value]` | `attr` ends with value                       | `[src$=".png"]`                                 |
| `[attr*=value]` | `attr` contains value anywhere               | `[title*="warning"]`                            |

---

## Group Selectors

| Selector | Meaning                 | Example      |
| -------- | ----------------------- | ------------ |
| `A, B`   | Select both `A` and `B` | `h1, h2, h3` |
