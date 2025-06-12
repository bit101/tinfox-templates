FRAMES=60
# ZEROS = number of digits in FRAMES
ZEROS=2
CYCLIC=true
WIDTH=1280
HEIGHT=720
VERBOSE=false
CONSOLE=false
ANTIALIAS=true
FPS=30

# single or video
default: single

################
# Single frame
################
single: clean pov-single
	@eog ./out/out.png

pov-single: 
	@povray \
		main.pov \
		Antialias=$(ANTIALIAS) \
		Display=false \
		Verbose=$(VERBOSE) \
		All_Console=$(CONSOLE) \
		+W$(WIDTH) +H$(HEIGHT) \
		Output_File_Name=./out/out.png

################
# Animation
################
video: clean pov-anim ffmpeg
	@mpv -loop ./out/out.mp4

ffmpeg: 
	@ffmpeg \
		-y \
		-framerate $(FPS) \
		-i ./out/frames/frame_%0$(ZEROS)d.png \
		-s:v $(WIDTH)x$(HEIGHT) \
		-c:v libx264 \
		-profile:v high \
		-crf 20 \
		-pix_fmt yuv420p \
		./out/out.mp4

pov-anim: 
	@povray \
		main.pov \
		Antialias=$(ANTIALIAS) \
		Display=false \
		+W$(WIDTH) +H$(HEIGHT) \
		Final_Frame=$(FRAMES) \
		Cyclic_Animation=$(CYCLIC) \
		Verbose=$(VERBOSE) \
		All_Console=$(CONSOLE) \
		Output_File_Name=./out/frames/frame_.png

################
# Other
################
clean:
	@rm -rf out/frames
	@mkdir -p out/frames

