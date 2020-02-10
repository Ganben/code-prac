; simple fun
(defun quadratic-roots (a b c)
    "Returns the roots of a quadratic equation aX^2 + bX + c = 0"
    (let ((discriminant (- (* b b) (* 4 a c))))
      (values (/ (+ (- b) (sqrt discriminant)) (* 2 a))
              (/ (- (- b) (sqrt discriminant)) (* 2 a)))))
; recursive 
(defun factorial (n)
  (cond ((zerop n) 1)
        (t (*      ; * is the last function called 
            n
            (factorial (- n 1))))))

; simple looping 
(defun simple-adding-machine-1 ()
  (let ((sum 0)
        next)
    (loop
      (setq next (read))
      (cond ((numberp next)
             (incf sum next))
            ((eq '= next)
             (print sum)
             (return))
            (t
             (format t "~&~A ignored!~%" next))))
    (values)))

;; data defs

(defvar *checks* (make-array 100 :adjustable t :fill-pointer 0)
  "A vector of checks.")

(defconstant +first-check-number+ 100 
  "The number of the first check.")

(defvar *next-check-number* +first-check-number+ 
  "The number of the next check.")

(defvar *payees* (make-hash-table :test #'equal) 
  "Payees with checks paid to each.")

(defstruct check
  number date amount payee memo)

(defun current-date-string ()
  "Returns current date as a string."
  (multiple-value-bind (sec min hr day mon yr dow dst-p tz)
                       (get-decoded-time)
    (declare (ignore sec min hr dow dst-p tz))
    (format nil "~A-~A-~A" yr mon day)))

(defun write-check (amount payee memo)
  "Writes the next check in sequence."
  (let ((new-check (make-check 
                    :number *next-check-number*
                    :date (current-date-string)
                    :amount amount
                    :payee payee
                    :memo memo)))
    (incf *next-check-number*)
    (vector-push-extend new-check *checks*)
    (push new-check (gethash payee *payees*))
    new-check))

(defun get-check (number)
  "Returns a check given its number, or NIL if no such check."
  (when (and (<= +first-check-number+ number) (< number *next-check-number*))
    (aref *checks* (- number +first-check-number+))))

(defun void-check (number)
  "Voids a check and returns T.  Returns NIL if no such check."
  (let ((check (get-check number)))
    (when check
      (setf (gethash (check-payee check) *payees*)
            (delete check (gethash (check-payee check) *payees*)))
      (setf (aref *checks* (- number +first-check-number+)) nil)
      t)))

(defun list-checks (payee)
  "Lists all of the checks written to payee."
  (gethash payee *payees*))

(defun list-all-checks ()
  "Lists all checks written."
  (coerce *checks* 'list))

(defun sum-checks ()
  (let ((sum 0))
    (map nil #'(lambda (check)
                 (when check
                   (incf sum (check-amount check))))
         *checks*)
    sum))

(defun list-payees ()
  "Lists all payees."
  (let ((payees ()))
    (maphash #'(lambda (key value)
                 (declare (ignore value))
                 (push key payees))
             *payees*)
    payees))

(dolist (item '(1 2 4 5 9 17 25))
    (format t "~&~D is~:[n't~;~] a perfect square.~%" item (integerp (sqrt item))))

(dotimes (n 11)
    (print n) (prin1 (* n n)))

(do ((which 1 (1+ which))
       (list '(foo bar baz qux) (rest list)))
      ((null list) 'done)
    (format t "~&Item ~D is ~S.~%" which (first list)))

(defstruct (ship
              (:print-function
               (lambda (struct stream depth)
                 (declare (ignore depth))
                 (print-unreadable-object (struct stream) 
                   (format stream "ship ~A of ~A at (~D, ~D) moving (~D, ~D)"
                           (ship-name struct)
                           (ship-player struct)
                           (ship-x-pos struct)
                           (ship-y-pos struct)
                           (ship-x-vel struct)
                           (ship-y-vel struct))))))
    (name "unnamed")
    player
    (x-pos 0.0)
    (y-pos 0.0)
    (x-vel 0.0)
    (y-vel 0.0))

(defclass 3d-point ()
  ((x :accessor point-x)
   (y :accessor point-y)
   (z :accessor point-z)))

(defclass 3d-point-2 ()
    ((x :reader get-x :writer set-x)
     (y :reader get-y :writer set-y)
     (z :reader get-z :writer set-z)))

(defclass sphere ()
    (x :accessor x)
    (y :accessor y)
    (z :accessor z)
    (radius :accessor radius)
    (volume :reader volume)
    (translate :writer translate))

; Volume from Radius 
(defmethod volume ((object sphere))
  (* 4/3 pi (expt (radius object) 3)))

; Default slot reader (illustration only) 
(defmethod slot-reader ((object object-class))
  (slot-value object 'slot-name))

; Default slot writer (illustration only) 
(defmethod slot-writer (new-value (object object-class))
  (setf (slot-value object 'slot-name) new-value))

(defclass 2d-object () ())

(defclass 2d-centered-object (2d-object)
  (x :accessor x)
  (y :accessor y)
  (orientation :accessor orientation)

(defclass oval (2d-centered-object)
  (axis-1 :accessor axis-1)
  (axis-2 :accessor axis-2))

(defclass regular-polygon (2d-centered-object)
  (n-sides :accessor number-of-sides)
  (size :accessor size))

(defclass 3d-point ()
  ((x :accessor point-x :initform 0 :initarg :x
      :documentation "x coordinate" :type real)
   (y :accessor point-y :initform 0 :initarg :y
      :documentation "y coordinate" :type real)
   (z :accessor point-z :initform 0 :initarg :z
      :documentation "z coordinate" :type real)))
