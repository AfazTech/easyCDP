# CSS Selectors Cheat Sheet :)

## Basic Selectors

| Selector  | Meaning                                 | Example           |
| --------- | --------------------------------------- | ----------------- |
| `*`       | Select all elements                     | `*`               |
| `element` | Select all `<element>`                  | `div`             |
| `#id`     | Select element with specific `id`       | `#header`         |
| `.class`  | Select elements with specific `class`   | `.btn-primary`    |
| `[attr]`  | Select elements with specific attribute | `input[disabled]` |

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
| `:hover`               | When mouse is over element           | `button:hover`           |
| `:focus`               | Element in focus (input, button)     | `input:focus`            |
| `:visited`             | Visited links                        | `a:visited`              |
| `:link`                | Unvisited links                      | `a:link`                 |
| `:target`              | Targeted element by URL hash         | `#section:target`        |
| `:root`                | The root element of the document     | `:root`                  |
| `:lang(lang)`          | Elements with language attribute     | `p:lang(en)`             |

---

## Attribute Selectors

| Selector        | Meaning                                                                                              | Example                                         |          |          |
| --------------- | ---------------------------------------------------------------------------------------------------- | ----------------------------------------------- | -------- | -------- |
| `[attr]`        | Elements with `attr`                                                                                 | `[type]`                                        |          |          |
| `[attr=value]`  | Elements with `attr` equal to `value`                                                                | `[type="checkbox"]`                             |          |          |
| `[attr~=value]` | `attr` contains value (space-separated list)                                                         | `[class~="btn"]`                                |          |          |
| \`\[attr        | =value]\`                                                                                            | `attr` is exactly value or starts with `value-` | \`\[lang | ="en"]\` |
| `[attr^=value]` | `attr` starts with value                                                                             | `[href^="https"]`                               |          |          |
| `[attr$=value]` | `attr` ends with value                                                                               | `[src$=".png"]`                                 |          |          |
| `[attr*=value]` | `attr` contains value anywhere                                                                       | `[title*="warning"]`                            |          |          |
| `[attr!=value]` | Select elements whose `attr` is not equal to `value` (not standard, supported by some preprocessors) | `[type!=radio]` (use JS or \:not() for real)    |          |          |

---

## Group Selectors

| Selector | Meaning                 | Example      |
| -------- | ----------------------- | ------------ |
| `A, B`   | Select both `A` and `B` | `h1, h2, h3` |

---

## Structural Pseudo-classes (advanced)

| Selector             | Meaning                          | Example                |
| -------------------- | -------------------------------- | ---------------------- |
| `:nth-child(An+B)`   | Select elements based on pattern | `li:nth-child(3n+1)`   |
| `:nth-last-child(n)` | nth child from the end           | `li:nth-last-child(1)` |
| `:only-of-type`      | Only element of its type         | `p:only-of-type`       |

---

## UI Element States

| Selector         | Meaning                            | Example               |
| ---------------- | ---------------------------------- | --------------------- |
| `:disabled`      | Disabled form controls             | `input:disabled`      |
| `:enabled`       | Enabled form controls              | `input:enabled`       |
| `:checked`       | Checked checkboxes/radios          | `input:checked`       |
| `:indeterminate` | Indeterminate state checkbox/radio | `input:indeterminate` |

---

## Pseudo-elements

| Selector         | Meaning                       | Example           |
| ---------------- | ----------------------------- | ----------------- |
| `::before`       | Insert content before element | `p::before`       |
| `::after`        | Insert content after element  | `p::after`        |
| `::first-letter` | First letter of element       | `p::first-letter` |
| `::first-line`   | First line of element         | `p::first-line`   |
| `::selection`    | Selected portion of text      | `::selection`     |

---

## Logical Operators in Selectors

| Selector   | Meaning                                   | Example          |
| ---------- | ----------------------------------------- | ---------------- |
| `E:not(s)` | Select elements `E` that do NOT match `s` | `a:not(.active)` |
| `E1, E2`   | Select all `E1` and `E2`                  | `div, p`         |
| `E1 E2`    | Select all `E2` descendants of `E1`       | `ul li`          |
| `E1 > E2`  | Select all `E2` children of `E1`          | `ul > li`        |

---

## Wildcards & Universal

| Selector        | Meaning      | Example               |          |          |
| --------------- | ------------ | --------------------- | -------- | -------- |
| `*`             | All elements | `*`                   |          |          |
| \`\[attr        | =value]\`    | Exact or prefix match | \`\[lang | ="en"]\` |
| `[attr^=value]` | Starts with  | `[href^="https"]`     |          |          |
