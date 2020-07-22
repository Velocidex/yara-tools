# Yara tools

Command line tool to manipulate yara rules.

## Clean

For very large rulesets it is sometimes more convenient to remove
un-necessary data from the rules:

1. This makes the ruleset much smaller.
2. It might remove sensitive information about the rules themselves.


Both of these are necessary prior to using large rulesets in
Velociraptor because we dont want to push un-necessary data to many
thousands of end points. Also some of the yara rules contain sensitive
information in the metadata or comments.

Strip the yara rule:

```
yara_tools clean my_file.yara > my_cleaned_yara.yara
```
