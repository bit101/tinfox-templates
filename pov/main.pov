#version 3.7;
#include "stdinc.inc"
#include "textures.inc"
#include "stones.inc"
#include "../lib/objects.inc"

global_settings { assumed_gamma 1 }

light_source {
  <-4, 4, -3>
    White
}

light_source {
  <8, 8, 0>
    White 
}

camera {
  right x * image_width / image_height
  location <2, 2, -4>
  look_at  <0, 0, 0>
}

box {
  <-1, -1, -1>
  <1, 1, 1>

  texture {
    T_Stone10
  }
  rotate clock*360
}

sky_sphere {
  pigment {
    gradient y
    color_map {
      [0 color Red]
      [1 color Blue]
    }
    scale 2
    translate -1
  }
}
