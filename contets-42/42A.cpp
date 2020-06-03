include <iostream>
using namespace std;

int main() {
    int n, V, b;
    int a[10000+5];
    cin>>n>>V;
    for (int i=0; i<n; i++){
            cin>>a[i];
        }
    float min_bi = 1e6+5;
    for (int i=0; i<n; i++){
            cin>>b;
            if (float(b) / a[i]< min_bi){
                        min_bi = float(b) / a[i];
                    }
        }
    float result = 0.0;
    for (int i=0;i <n; i++){
            result += float(min_bi) * a[i];
        }
    printf("%.4f", min(float(V), result)); 
    return 0;
}
