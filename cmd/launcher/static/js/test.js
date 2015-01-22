// notes:
//  - never use == use === and "use strict"
//  - var a = {}; if(typeof a.b === 'undefined') {} //check if a is undefined
//  - var a = val1 || val2 // a will have first val(1,2) that is not false
//  - .prototype = can add methods/data to all instances of an object. can use this inside of prototype
// - all functions have an implicit 'arguments' param/member; basically an array.
// - call vs apply -
// - image - progammatic version of <image> in html, for js. image.src = path_to_image and context.drawImage(img, x, y)
// - context had transform/rotate/etc a la gl.  start with context.save(). follow w/ context.restore(). (a la gl push/pop)
//

// - localhost webserver: dir you want to serve>python -m SimpleHttpServer, need index.html page in dir
var canvas = document.querySelector("#myCanvas");var context = canvas.getContext("2d");

// code to make window fullscreen at all times/orientations
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

window.onresize = function()
{
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
}

window.onorientationchange = function()
{
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
}
// end code to make window fullscreen at all times/orientations

var circlePos = [200, 100];
var lastTime = new Date();

function update()
{
  // getting dt (in millisecs)
  var newTime = new Date();
  var dt = (newTime - lastTime) / 1000;
  lastTime = newTime;

  window.requestAnimationFrame(update); // gives you ONE animation frame, needs to be recalled every frame.  Pauses when your window isn't in focus.

  context.fillStyle = "rgba(254, 128, 128, 1";
  context.fillRect(0, 0, canvas.width, canvas.height);

  context.strokeStyle = "#0F";
  context.fillStyle = "#0D7";
  context.lineWidth = 7;
  context.beginPath();
  context.moveTo(0, 0);
  context.lineTo(100, 100);
  context.lineTo(100, 200);
  context.closePath();
  context.stroke();  // outlines shape.
  context.fill();    // fills shape.

  context.beginPath();
  context.arc(circlePos[0], circlePos[1], 50, Math.PI * 2, false);
  context.closePath();
  context.stroke();  // outlines shape.
  context.fill();    // fills shape.

  /* circlePos[0] += dt * 50;  // updating pos based on dt */
}

// adding touch control
window.addEventListener("touchmove", function(e)
{
  // console.log(e)
  // console.log("Touch moved!");
  e.preventDefault();  // for games you prolly want to prevent defaults behavior for most events.
  var touch = e.changedTouches[0];
  circlePos = [touch.clientX, touch.clientY];
});

update();

// invoke a non-func and pass in namespace you want as param
(function(MyNamespace)
{
  "use strict";  // only use strict inside of namespace, not globally.
  MyNamespace.someFunction = function() {
    // define a function on our namespace
  };
})(window.MyNamespace = window.MyNamespace || {} );
// use window.MyNamespace to get "undefined" val instead of an error
