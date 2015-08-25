# redditproto

Protocol Buffer definitions for Reddit's JSON api types. For
more information about what these messages represent, see
[Reddit's docs](https://github.com/reddit/reddit/wiki/JSON).

To generate the code for the protobuffers for your language, you can most likely
do the following:

    [your package manager] install protobuf-compiler
    protoc --[your lang code]_out=. *.proto

Common examples:

    protoc --cpp_out=. *.proto
    protoc --java_out=. *.proto
    protoc --python_out=. *.proto

See ````protoc --help```` for further guidance.

The golang generated code is included to make this package go gettable.
