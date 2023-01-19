package adt

type Wind struct {
	direction float64
	speed float64 
}

func (w *Wind) Set( direction, speed float64) { 
	w.direction = direction 
	w.speed = speed 
} 

func (w Wind) WindData() (float64, float64) {
	 return w.direction, w.speed 
}
