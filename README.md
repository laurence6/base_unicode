# base_unicode

base_unicode encodes arbitrary data into Unicode characters. Unicode character table can be specified by user.

Note: This scheme is not efficient for storage space. The output is a little shorter than that of base64 (7% shorter using default Unicode character table).

## Unicode character table

The user can specify a Unicode character table for encoding and decoding. A default table with Chinese characters is embedded, and its content can be found in file `default_table.txt`.

Recommended table length is a power of two.

Please do not use characters between 0x20 ~ 0x60. (This may be fixed in the future.)

## Usage

```
Usage: ./base_unicode [OPTIONS]... [FILE]
If FILE is empty or '-', read from standard input.
  -d
  -decode
        decode data
  -t string
  -table string
        path of table (if empty, use embedded table)
  -w int
  -wrap int
        wrap lines after n characters (0 to disable wrap) (default 80)
```

## License

Copyright (C) 2017-2017  Laurence Liu <liuxy6@gmail.com>

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
