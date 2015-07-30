# redditproto

This repo includes the protobuffer definitions for Reddit's JSON api types. For
more information about what these message represent, see
[Reddit's docs](https://github.com/reddit/reddit/wiki/JSON).

To generate the code for the protobuffers for your language, you can most likely
do the following:

    [your package manager] install protobuf-compiler
    protoc --[your lang code]_out=. reddit.proto

Common examples:

    protoc --cpp_out=. reddit.proto
    protoc --java_out=. reddit.proto
    protoc --python_out=. reddit.proto

See ````protoc --help```` for further guidance.

The golang generated code is included to make this package go gettable.
