#include <bits/stdc++.h>
#include <set>
using namespace std;

#define IOfast                        \
    ios_base::sync_with_stdio(false); \
    cin.tie(NULL);                    \
    cout.tie(NULL);
#define ll long long
#define ld long double
#define pb push_back
#define mp make_pair
#define ff first
#define ss second
#define nl '\n'
#define all(x) x.begin(), x.end()
#define rall(x) x.rbegin(), x.rend()
#define sz(x) (int)x.size()
#define rep(i, a, b) for (int i = a; i < b; i++)
#define vll vector<ll>
#define pii pair<int, int>
#define pll pair<ll, ll>
#define yes cout << "YES\n"
#define no cout << "NO\n"
#define vout(v)           \
    for (auto i : v)      \
        cout << i << " "; \
    cout << nl;

template<class T> bool ckmin(T&a, T&b) { bool B = a > b; a = min(a,b); return B; }
template<class T> bool ckmax(T&a, T&b) { bool B = a < b; a = max(a,b); return B; }

const int MOD = 1e9 + 7;

void solve(){
}

int main()
{
    IOfast;
    int t = 1;
    cin >> t;
    while (t--) solve();
}
