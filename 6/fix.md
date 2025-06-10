# Enigma Machine - Bug Fixes

## Bugs Found and Fixed

### 1. Double-Stepping Timing Issue

**Bug**: Rotor stepping logic checked notch positions after stepping rotors, causing incorrect double-stepping behavior.

**Problem**: The middle rotor would not step properly when it should trigger double-stepping, breaking historical Enigma accuracy.

**Fix**: Modified `step_rotors()` method to check notch positions before stepping:
```ruby
# Check notch positions before stepping
middle_at_notch = @rotors[1].at_notch?
right_at_notch = @rotors[2].at_notch?

# Step based on pre-checked positions
@rotors[0].step if middle_at_notch
@rotors[1].step if right_at_notch || middle_at_notch
@rotors[2].step
```

### 2. Plugboard Reciprocity Issue

**Bug**: Plugboard swapping was only applied on input, not output, breaking the Enigma's reciprocal property.

**Problem**: Messages encrypted with plugboard settings could not be properly decrypted with the same settings.

**Fix**: Added plugboard application on output in `encrypt_char()` method:
```ruby
# Apply plugboard on output for reciprocity
c = plugboard_swap(c, @plugboard_pairs)
```