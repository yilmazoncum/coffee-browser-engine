
# Coffee Browser Engine

Its my browser engine like WebKit, Blink, Gecko.

Coffee was made for personal development purposes. So it only has very simple features

## Parsing

Coffee parser a simple Html like this.

```html
<html lang="en">
    <body class="something">
        <h1>Coffee Engine</h1>
        <div id="container">
            <span></span>
            <div class="container1">
                <p id="paragraph1"> 
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
                </p>
            </div>
            <div class="container2">
                <p id="paragraph2"> 
                    Nullam vel ornare magna. Vivamus venenatis diam id tellus fermentum aliquet.
                </p>
            </div>
        </div>
    </body>
</html>   
```

and if you want, you can add style on them with css:

```CSS
html {
    background-color: rgb(255, 128, 0);
    width: 2000px;
    height: 1000px; 
}

body {
    display: block;
    background-color: rgb(99, 35, 0);
    width: 1800px;
    height: 800px;
    margin-top: 100px;
    margin-left: 100px; 
}

h1 {
    margin-top: 60px;
    margin-left: 100px;
}

.container1 {
    background-color: rgb(0, 126, 164);
    width: 700px;
    height: 600px;
    top: 100px;
    margin-left: 100px;
    margin-left: 100px;
}
.container2 {
    background-color: rgb(0, 66, 45);
    width: 700px;
    height: 600px;
    top: 100px;
    left: 1000px;
}

#paragraph1 {
    background-color: rgb(218, 22, 176);
    width: 600px;
    height: 100px;
    top: 30px;
    left: 50px;
}

#paragraph2 {
    background-color: rgb(215, 25, 25);
    width: 600px;
    height: 100px;
    top: 30px;
    left: 50px;
}   
```

For now, Engine only support a few css tag like:
```
top,bottom,left,right,margin,margin-top,margin-bottom,margin-left,margin-right
```

## Flow
Coffee follow the Webkit like flow

![alt text](https://web-dev.imgix.net/image/T4FyVKpzu4WKF1kBNvXepbi08t52/S9TJhnMX1cu1vrYuQRqM.png?auto=format&w=741)

You can run ```go run main.go``` and Coffee render a png image after all of this steps and it show this image on your default browser

image like this:
![download](https://github.com/yilmazoncum/coffee-browser-engine/assets/55596806/6a782548-c381-49ce-b48d-9ec0c6b5ff96)


## References 

- https://web.dev/howbrowserswork/

- https://limpet.net/mbrubeck/2014/08/08/toy-layout-engine-1.html

- https://github.com/thomscoder/pandora
