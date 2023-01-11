template<typename T>
T lcm(T x,T y) {
  return (x / __gcd(x,y)) * y;
}

