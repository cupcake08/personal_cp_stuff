namespace LIS {
  const int INF = 1e9;
  // credit: https://cp-algorithms.com/sequences/longest_increasing_subsequence.html
  int lis(vi const& a) {
    int n = sz(a);
    vi d(n + 1, INF);
    d[0] = -INF;
    for(int i = 0; i < n; i++) {
      int j = upper_bound(all(d), a[i]) - d.begin();
      if(d[j - 1] < a[i] && a[i] < d[j]) {
        d[j] = a[i];
      }
    }
    int ans = 0;
    for(int i=0;i<=n;i++) {
      if(d[i] < INF) {
        ans = i;
      }
    }
    return ans;
  }
}

