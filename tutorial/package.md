1. 不同的包名只能放到不同的目录下，同一个目录下只能有一个包名，自定义包名应与目录名一致
2. 被导出的方法名，首字母需大写


eg.

mkdir packages
mdkir custommodule

pageckage custommodule

in main.go

import "goapp/packages/custommodule"