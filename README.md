# i3-scratch-list
List i3 scratchpad windows

# usage
```# gives a list of windows and ids:
   i3-scratch-list```
   
 ```# gives a list of window titles:
    i3-scratch-list -noid```

# installation

If you have your go bin in your $PATH:

 go get github.com/draxil/i3-scratch-list 

Otherwise download this repository and go and from the directory:

 go build .
 cp i3-scratch-list /somewhere/in/your/path

# my rofi / dmenu formula for finding a window in your scratch

  i3-scratch-list  | rofi -dmenu | perl -nE 'my($id) = split("-");  $id =~ tr/ //d; if($id){say $id}' | xargs -i{} i3-msg "[id=\""{}"\"] scratchpad show"
