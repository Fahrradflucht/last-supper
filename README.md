# Last Supper
Dummy image provider written in Go.

This is basically a rewrite of [imgsrc](https://github.com/mabels/ImgSrc)
with the only difference that it is written in Go and therefore not
requires a JVM and that it has a slightly saner API (`widthxheight` and not
`heightxwidth`).

## Notes:
- The included basicfont is fixed size and therefore not suitable for
  project.
- Instead on server start check if a (configurable ?) font is available
  and get it from the interwebs outherwise. <| Is offline support
  diserable? Maybe add option to set a fontpath.

## Feature Ideas:
- Make the basepath configurable
- Make the fontstorage path configurable
- Add transparancy as a query param
- Add font as a query param
- gzip

