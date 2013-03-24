

package fgms


// from .hxx
const SG_180 = 180.0
const SG_PI = 3.1415926535
const SG_RADIANS_TO_DEGREES = (SG_180/SG_PI)
const SG_DEGREES_TO_RADIANS = (SG_PI/SG_180)
const SG_FEET_TO_METER    = 0.3048

//from c.xx
/**
 *  High-precision versions of the above produced with an arbitrary
 * precision calculator (the compiler might lose a few bits in the FPU
 * operations).  These are specified to 81 bits of mantissa, which is
 * higher than any FPU known to me:
 */
const double SQUASH  = 0.9966471893352525192801545;
const double STRETCH = 1.0033640898209764189003079;
const double POLRAD  = 6356752.3142451794975639668;

// Radians To Nautical Miles 
const SG_RAD_TO_NM  = 3437.7467707849392526

// Nautical Miles in a Meter 
const SG_NM_TO_METER  = 1852.0000

// Meters to Feet 
const SG_METER_TO_FEET  = 3.28083989501312335958

// PI2 
const SGD_PI_2    = 1.57079632679489661923



type Point3D struct {
	x float32 
	y float32 
	z float32 
}