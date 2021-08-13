Slack Panel
===========

This is the code for a little Slack quick-status setter. The idea is that you
put this on a Raspberry Pi and then attach [a
touchscreen](https://www.adafruit.com/product/1601). Once that's done, this
service gets set up to run on boot, and provides a small web server. Set up
chromium to open a specific page in kiosk mode and boom, a wee little display
you can use to set your status quickly in Slack.

Oh, and if you've got buttons on your screen hat then you can specify what they
do, too.

## Todo

 * [ ] basic skeleton
 * [ ] define config format
 * [ ] route for adding slack accounts
 * [ ] route for oauth2 callback
 * [ ] route for button display 
 * [ ] goroutine for handling slack events
 
There's probably lots of other things, but that's the bare skeleton.
