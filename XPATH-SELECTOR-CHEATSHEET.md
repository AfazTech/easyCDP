# XPath Cheat Sheet :)

## Basic Nodes

| XPath    | Meaning                                    | Example      |
| -------- | ------------------------------------------ | ------------ |
| `//node` | Select all nodes with name `node`          | `//div`      |
| `/node`  | Select `node` children of the current node | `/html/body` |
| `.`      | Select the current node                    | `.`          |
| `..`     | Select the parent node                     | `..`         |
| `@attr`  | Select attribute named `attr`              | `//@href`    |

---

## Predicates (Conditions)

| XPath                              | Meaning                                               | Example                            |
| ---------------------------------- | ----------------------------------------------------- | ---------------------------------- |
| `//div[@id='header']`              | Select `div` with `id` attribute equal to `header`    | `//div[@id='header']`              |
| `//input[@type='checkbox']`        | Select input elements with type checkbox              | `//input[@type='checkbox']`        |
| `//a[@class='btn active']`         | Select `<a>` with exact class attribute               | `//a[@class='btn active']`         |
| `//p[text()='Hello']`              | Select `<p>` with exact text content                  | `//p[text()='Hello']`              |
| `//div[contains(@class,'active')]` | Select div with class containing 'active'             | `//div[contains(@class,'active')]` |
| `//span[starts-with(@id,'item')]`  | Select span with id starting with 'item'              | `//span[starts-with(@id,'item')]`  |
| `//input[@checked]`                | Select inputs with `checked` attribute (present)      | `//input[@checked]`                |
| `//a[ends-with(@href,'.pdf')]`     | Select `<a>` with href ending with '.pdf' (XPath 2.0) | `//a[ends-with(@href,'.pdf')]`     |

---

## Axes (Relationships)

| XPath                        | Meaning                               | Example                      |
| ---------------------------- | ------------------------------------- | ---------------------------- |
| `//div/child::p`             | Select `p` children of all `div`      | `//div/child::p`             |
| `//div/parent::*`            | Select the parent of all `div`        | `//div/parent::*`            |
| `//div/following-sibling::p` | Select all `p` siblings after `div`   | `//div/following-sibling::p` |
| `//p/preceding-sibling::div` | Select all `div` siblings before `p`  | `//p/preceding-sibling::div` |
| `//div/ancestor::body`       | Select ancestor `body` of `div`       | `//div/ancestor::body`       |
| `//div/descendant::a`        | Select all `<a>` descendants of `div` | `//div/descendant::a`        |
| `//div/self::div`            | Select current node if it is `div`    | `//div/self::div`            |

---

## Position and Index

| XPath                        | Meaning                                  | Example                      |
| ---------------------------- | ---------------------------------------- | ---------------------------- |
| `(//li)[1]`                  | Select first `<li>` in the document      | `(//li)[1]`                  |
| `//ul/li[1]`                 | Select first `<li>` child of each `<ul>` | `//ul/li[1]`                 |
| `//tr[last()]`               | Select last `<tr>`                       | `//tr[last()]`               |
| `//td[position() <= 3]`      | Select first 3 `<td>`                    | `//td[position() <= 3]`      |
| `//li[position() mod 2 = 1]` | Select odd position `<li>` elements      | `//li[position() mod 2 = 1]` |

---

## Logical Operators

| XPath                                         | Meaning                                   | Example                                       |
| --------------------------------------------- | ----------------------------------------- | --------------------------------------------- |
| `//input[@type='text' and @name='user']`      | Select input with type text AND name user | `//input[@type='text' and @name='user']`      |
| `//button[@id='submit' or @name='submitBtn']` | Select button with id OR name submit      | `//button[@id='submit' or @name='submitBtn']` |
| `//div[not(@class='hidden')]`                 | Select div without class hidden           | `//div[not(@class='hidden')]`                 |
| `//p[text()='foo' or text()='bar']`           | Select `<p>` with text 'foo' or 'bar'     | `//p[text()='foo' or text()='bar']`           |

---

## Functions on Text and Attributes

| XPath                                                                                    | Meaning                                    | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ------------------------------------------ | ---------------------------------------------------------------------------------------- |
| `//a[contains(text(),'Next')]`                                                           | Select `<a>` containing text 'Next'        | `//a[contains(text(),'Next')]`                                                           |
| `//p[normalize-space(text())='Hello']`                                                   | Select `<p>` with trimmed text 'Hello'     | `//p[normalize-space(text())='Hello']`                                                   |
| `//a[translate(@class,'ABCDEFGHIJKLMNOPQRSTUVWXYZ','abcdefghijklmnopqrstuvwxyz')='btn']` | Case-insensitive match class equals 'btn'  | `//a[translate(@class,'ABCDEFGHIJKLMNOPQRSTUVWXYZ','abcdefghijklmnopqrstuvwxyz')='btn']` |
| `//img[@src and string-length(@src) > 0]`                                                | Select images with non-empty src attribute | `//img[@src and string-length(@src) > 0]`                                                |
| `//a[count(@*)=2]`                                                                       | Select `<a>` with exactly two attributes   | `//a[count(@*)=2]`                                                                       |

---

## Wildcards

| XPath      | Meaning                    | Example    |
| ---------- | -------------------------- | ---------- |
| `//*`      | Select all elements        | `//*`      |
| `//div/*`  | Select all children of div | `//div/*`  |
| `//@*`     | Select all attributes      | `//@*`     |
| `//text()` | Select all text nodes      | `//text()` |

---

## Combining XPath Expressions

| XPath                                             | Meaning                          | Example       |                |
| ------------------------------------------------- | -------------------------------- | ------------- | -------------- |
| `//div \| //span`                                 | Select all div **or** span nodes | \`//div       | //span\`       |
| `(//div)[position()=1] \| (//span)[position()=1]` | Select first div or first span   | \`(//div)\[1] | (//span)\[1]\` |

---

## Advanced and Misc

| XPath                                                               | Meaning                                             | Example                                                             |
| ------------------------------------------------------------------- | --------------------------------------------------- | ------------------------------------------------------------------- |
| `//div[.//a]`                                                       | Select div containing at least one `<a>` descendant | `//div[.//a]`                                                       |
| `//div[not(.//a)]`                                                  | Select div with no `<a>` descendants                | `//div[not(.//a)]`                                                  |
| `//a[@href][not(starts-with(@href,'#'))]`                           | Select `<a>` with href not starting with #          | `//a[@href][not(starts-with(@href,'#'))]`                           |
| `//a[contains(concat(' ', normalize-space(@class), ' '), ' btn ')]` | Select `<a>` containing class 'btn' exactly         | `//a[contains(concat(' ', normalize-space(@class), ' '), ' btn ')]` |
| `//node()[last()]`                                                  | Select last node of any kind                        | `//node()[last()]`                                                  |

---

## Namespaces (XPath 1.0 / 2.0)

| XPath                                                 | Meaning                                                     | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------- |
| `//*[namespace-uri()='http://www.w3.org/1999/xhtml']` | Select elements in XHTML namespace                          | `//*[namespace-uri()='http://www.w3.org/1999/xhtml']` |
| `//xhtml:div`                                         | Select `div` in XHTML namespace (prefix must be registered) | `//xhtml:div`                                         |
