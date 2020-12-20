//Package vsr provides structures and methods for developing an 
//information retrieval system using a vector space model.
package vsr

//Vector represents a hash map vector for a document. It stores
//every term in the document and the number of times the term occurs.
type Vector struct {
    HashMap map[string]float64
}


//Add will increment the count of a token in the vector.
//If the term is not present in the vector then the term is added
//and the count is initialized to 1.
func (v *Vector) Add(token string) {
    count, present := v.HashMap[token]
    if present {
        v.HashMap[token] = count + 1
    } else {
        v.HashMap[token] = 1
    }
}


