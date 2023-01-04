#include <bits/stdc++.h>
#include <cmath>
using namespace std;

#define IOfast                        \
    ios_base::sync_with_stdio(false); \
    cin.tie(NULL);                    \
    cout.tie(NULL);
#define ankit main
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

/********** Debug Stuff ********/
#define deb(...) log(#__VA_ARGS__, __VA_ARGS__)
template<typename ...Args>
void log(string vars,Args&&... values) {
    cout << vars << " = ";
    string delim = "";
    (...,(cout << delim << values,delim = ", "));
}
/************/

const int MOD = 1e9 + 7;

void solve(){
    int x = 50;
    string y = "ankit";
    double z = 60.0;
    deb(x,y,z);
}

int ankit()
{
    IOfast;
    int t = 1;
    cin >> t;
    while (t--) solve();
}
