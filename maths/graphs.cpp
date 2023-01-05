// credit: kal013 (codeforces)
#include <vector>

namespace Graphs{
struct DAG{
	int n;
	vector<vector<int>> adj;

	DAG(int _n = 0) { init(_n); }

	DAG(const vector<vector<int>>& _adj) {
		init(_adj);
	}

	inline void init(int _n){
		n = _n;
		adj.assign(n, {});
	}
	inline void init(const vector<vector<int>>& _adj){
		n = static_cast<int>(_adj.size());
		adj = _adj;
	}
	inline void add_edge(int a, int b){
		adj[a].push_back(b);
	}
	void dfs_topo(int u, vector<int>& order, vector<bool>& visited) {
		visited[u] = true;
		for(int v: adj[u]){
			if (!visited[v]){
				dfs_topo(v, order, visited);
			}
		}
		order.push_back(u);
	}
	vector<int> toposort() {
		vector<bool> visited(n, false);
		vector<int> ans;
		for(int i = 0; i < n; ++i){
			if (!visited[i]){
				dfs_topo(i, ans, visited);
			}
		}
		reverse(ans.begin(), ans.end());
		return ans;
	}
	void dfs_scc(int v, int& index, vector<int>& r_index, vector<int>& S, int& c){ // D.J. Pearce algo 3.
		bool root = true;
		r_index[v] = index++;
		for(int w: adj[v]){
			if (r_index[w] == -1) dfs_scc(w, index, r_index, S, c);
			if (r_index[w] < r_index[v]){
				r_index[v] = r_index[w];
				root = false;
			}
		}
		if (root){
			--index;
			while (!S.empty() && r_index[v] <= r_index[S.back()]){
				int w = S.back(); S.pop_back();
				r_index[w] = c;
				--index;
			}
			r_index[v] = c;
			--c;
		}
		else{
			S.push_back(v);
		}
	}
	vector<int> scc_condense() { // returns comp s.t. comp[i] == comp[j] iff i, j in same scc and if i ---> j then comp[i] < comp[j]
		vector<int> comp(n, -1);
		int index = 1, c = n - 1;
		vector<int> S;
		for(int i = 0; i < n; ++i){
			if (comp[i] == -1){
			dfs_scc(i, index, comp, S, c);
			}
		}
		for(auto &e: comp){
			e -= c + 1;
		}
		return comp;
	}
	vector<vector<int>> scc_condense_adj(const vector<int>& comp) {
		const int m = (*max_element(comp.begin(), comp.end())) + 1;
		vector<vector<int>> ans(m);
		for(int u = 0; u < n; ++u){
			ans[comp[u]].push_back(u);
		}
		vector<bool> taken(m, false);
		for(int i = 0; i < m; ++i){
			vector<int> dp;
			for(int u: ans[i]){
				for(int v: adj[u]){
					int V = comp[v];
					if (V == i || taken[V])
					continue;
					dp.push_back(V);
					taken[V] = true;
				}
			}
			for(int V: dp){
				taken[V] = false;
			}
			ans[i].swap(dp);
		}
		return ans;
	}
	DAG scc_condense_graph(const vector<int>& comp) {
		DAG(scc_condense_adj(comp));
	}
	DAG scc_condense_graph() {
		return scc_condense_graph(scc_condense());
	}
	vector<vector<int>> scc_condense_adj() {
		return scc_condense_adj(scc_condense());
	}
};
 
 
struct TwoSat{
	int n;
	DAG g;
	TwoSat(int _n = 0) {init(_n);}
	 
	void init(int _n){
		n = _n;
		g.init(_n);
	}
	 
	inline void set(int x, int value_x) {
		assert (0 <= x && x < n); 
		assert (0 <= value_x && value_x <= 1);
		g.add_edge(2 * x + (value_x ^ 1), 2 * x + value_x);
	}
	 
	inline void add_or_clause(int x, int value_x, int y, int value_y){
		assert(0 <= x && x < n && 0 <= y && y < n);
		assert(0 <= value_x && value_x <= 1 && 0 <= value_y && value_y <= 1);
		 
		g.add_edge(2 * x + (value_x ^ 1), 2 * y + value_y);
		g.add_edge(2 * y + (value_y ^ 1), 2 * x + value_x);
	}
	 
	inline void add_implication(int x, int value_x, int y, int value_y){
		assert(0 <= x && x < n && 0 <= y && y < n);
		assert(0 <= value_x && value_x <= 1 && 0 <= value_y && value_y <= 1);		
		 
		g.add_edge(2 * x + value_x, 2 * y + value_y);
	}
	 
	inline vector<int> solve() {
		int cnt;
		vector<int> c = g.scc_condense();
		vector<int> res(n);
		for (int i = 0; i < n; i++) {
			if (c[2 * i] == c[2 * i + 1]) {
				return vector<int>();
			}
			res[i] = (c[2 * i] < c[2 * i + 1]);
		}
		return res;
	}
};
};
