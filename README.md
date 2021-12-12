# rpic - raspberry pi controller

raspberry pi's have been leveraged for many projects like a controller for home lighting or as an ad blocking proxy server, however, since the base system for raspberry pi is typically linux based, there needs to be a level of technical ability to perform most simple tasks like a graceful shutdown.

raspberry pi controller (rpic) is a remote management system for your raspberry pi device built on top of raspbian, the goal is to provide a set of administrative functions and api's to simplify the usage of a raspberry pi.

rpic will be based on raspbian initially which is the most widely used distribution for raspberry pi, but could be ported to other linux distributions.

## features

* configure system
* power down/restart system
* upgrade system
* upgrade software
* monitor device
* diagnostics
* backup/restore

### configure system

could be general system settings like configuring the network, resolution output settings, sleep/suspend options, or it could be application related settings.

### power down/restart system

rasbperry pi does not provide a power button, so it's common to do a hard shutdown which is effectively pulling the power, however, this is not a graceful shutdown and could impact the file system and data on the microsd card.

starting up a system requires that the cable be plugged in or for the raspberry pi to have a hardware button to toggle the state.

### upgrade system

this can be considered firmware for all intensive purposes, but raspbian operating system and dependent libraries require upgrade from time to time for security purposes or to address general bugs.

`there should be an online and offline method.`

`there should also be an auto-update process.`

### upgrade software

software applications that run on top of raspbian requires updates, this may or may not require restart of the device.

`there should be an online and offline method.`

### monitor device

storage capacity and utilization, network traffic information, temperature, cpu and memory utilization, file system integrity, general log information, etc.

### diagnostics

troubleshoot and find issues with your device.  there could be a more advanced for of diagnostics that allows the device to contact a trusted remote service and upload log and dump information for better diagnostics.

### backup/restore

backup state and restore from past points in time.

## security

since most of these commands require super user privileges on raspbian, this could have some serious security implications, so there needs to be some level of security to access these capabilities like a user/password for rpic.  this password can be configured and there needs to be a way to reset the system in case the password is forgotten. 

## interface

rpic will provide a restful api which can be accessed by developers to integrate with their application or to write custom clients.  the most basic client will be included as a web interface which can be enabled or disabled.

an ios client may be added, however, intrinsically, most applications will want to have the functionality embedded into their own user interface as opposed to having to use multiple interfaces.
