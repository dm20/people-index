# People Index
A simple Go service for creating and modifying key-value pairs that need to persist in memory.

Features:
  1) Create JSON files if not already created
  2) Edit existing or new JSON files by adding 'people' tags  
     - A person tag can be any key and associated value
  3) Look up person information by key
  4) Delete person information associated with a given key
  5) Add multiple values to one key

Include the test() function in people_index.go to demo the functionality. 

Uses <a href='https://github.com/Jeffail/gabs'> this JSON library</a>
