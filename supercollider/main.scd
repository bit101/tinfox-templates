(
var cleanup = {
  "cleanup".postln;
  ServerTree.removeAll;
  ServerBoot.removeAll;
  ServerQuit.removeAll;
  CmdPeriod.run;
};

var makeBusses = {
  "makeBusses".postln;
  ~out = 0;
  s.newBusAllocators;
  ~reverbBus = Bus.audio(s, 2);
  ~delayBus = Bus.audio(s, 2);
  ~masterBus = Bus.audio(s, 2);
};

var makeSynthDefs = {
  "makeSynthDefs".postln;
  thisProcess.interpreter.executeFile("/home/keith/code/supercollider/lib/synthdefs.scd");
  thisProcess.interpreter.executeFile("/home/keith/code/supercollider/lib/effects.scd");
  ~keyboardSynth = \pulse;
};

var makeEffects = {
  "makeEffects".postln;
  ~master = Synth(\masterAmp, [in: ~masterBus, out: ~out, amp: 0.5]);
  ~delay = Synth(\delay, [in: ~delayBus, out: ~masterBus, delay: 0.15, decay: 2, drive: 0.3]);
  ~reverb = Synth(\reverb, [in: ~reverbBus, out: ~delayBus, room: 0.7]);
};

var postDone = {
  "=================================".postln;
  "= READY!!!!!".postln;
  "=================================".postln;
};

s.waitForBoot({
  cleanup.();
  makeBusses.();
  makeSynthDefs.();
  s.sync;
  thisProcess.interpreter.executeFile("midi.scd");
  ~setupMidi.(false);
  ~addMidiEvents.();
  ServerTree.add(makeEffects);
  ServerTree.add(~addMidiEvents);
  ServerTree.add(postDone);

});
)

// can change this easily.
~keyboardSynth = \kick;

(
Pbindef(\main,
  \instrument, \pulse,
  \dur, 0.25,
  \octave, 4,
  \degree, Prand((1..8), inf),
  \amp, 0.5,
  \out, ~reverbBus,
);
)

Pbindef(\main).play
Pbindef(\main).stop


