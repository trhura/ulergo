(defn proper-divisors [num]
  nil)

(defn divisible [num div]
  (zero? (rem num div)))

(defn is-prime? [num]
  (cond (<= num 1) false
        (== num 2) true
        :else (not-any? #(divisible num %) (range 2 (inc (Math/sqrt num))))))

(defn next-prime [num]
  (if (is-prime? (inc num))
    num
    (next-prime (inc num))))
