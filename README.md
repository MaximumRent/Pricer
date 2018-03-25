# Pricer
Pricer is simple parser for real estate sites. Written on Golang.
# GOPATH
GOPATH installed on project root - /path-to-project/pricer
# Database
Here is used NoSQL database for data storaging.

# Page parsing
Further will be described structure of site parsing and tags, which will be used on every parsed page.
-> HREF TAG          (index 0)
--> PRICE TAG        (index 1)
--> STREET TAG       (index 2) (which will be translated to coordinates)
--> TOTAL SPACE TAG  (index 3)
--> LIVING SPACE TAG (index 4)
--> ROOMS COUNT TAG  (index 5)
--> FLOOR TAG        (index 6)
Configuration described in parse_config.json file.