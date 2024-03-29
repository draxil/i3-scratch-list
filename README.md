# i3-scratch-list
List i3 scratchpad windows

# Usage
``` 
# gives a list of windows ids and titles:
i3-scratch-list 
# gives a list of window titles:
i3-scratch-list -noid
```
## My rofi / dmenu formula for finding a window in your scratch

So the simple way is the window title only:
```
 i3-scratch-list -noid  | rofi -dmenu -i  -p "scratchpad" | xargs -i{} i3-msg "[title=\""{}"\"] scratchpad show"
```
Of course if you want to be able to pick from identical names you can include the ID, of course this is shown in the dmenu/rofi:
```
  i3-scratch-list  | rofi -dmenu -i -p "scratchpad" | perl -nE 'my($id) = split("-");  $id =~ tr/ //d; if($id){say $id}' | xargs -i{} i3-msg "[id=\""{}"\"] scratchpad show"
```

Better solutions in an issue please.

# Installation

## With go

If you have your `$GOBIN` in your `$PATH`:
```
 go install github.com/draxil/i3-scratch-list@latest
```

## github release builds

See the [releases page](https://github.com/draxil/i3-scratch-list/releases).

## Build
Otherwise download this repository and go, then from the directory:
```
 go build .
 cp i3-scratch-list /somewhere/in/your/path
```

Better packaging will come if anyone cares to pester for it.


# Credits

- /u/LionSuneater on reddit for reminding me that rofi/dmenu has a -p switch

