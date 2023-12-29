(ql:quickload "hunchentoot")

(defparameter *implants* ())

(hunchentoot:start 
  (make-instance 'hunchentoot:easy-acceptor :port 3000))
(hunchentoot:define-easy-handler (foo :uri "/") ()
  (let ((request-type (hunchentoot:request-method hunchentoot:*request*))
        (postdata (hunchentoot:post-parameters* hunchentoot:*request*))
        (implantname (read-from-string (hunchentoot:post-parameter "implantname" hunchentoot:*request*))))
    (cond ((eq request-type :get) (format t "GOT GET~%"));; handle get request
          ((eq request-type :post)
            (format t "POSTDATA: ~a~%" postdata)
	   (format t "ALL IMPLANTS BEFORE: ~a~%" *implants*)
 	   (let ((implant (assoc implantname *implants*)))
	     (format t "IMPLANT: ~a~%" implant)
	     (if implant
		  (let ((plugins (cadadr implant)))
		    (setf *implants* (remove implant *implants*))
		    (push (list implantname (list (get-universal-time) plugins)) *implants*))
		 (push (list implantname (list (get-universal-time) '("(format t \"HI\"))"))) *implants*)))
	   (format t "ALL IMPLANTS AFTER: ~a~%" *implants*)
	   (car (cadadr (assoc implantname *implants*)))
	   ))))
        ;;    (let* ((data-string (hunchentoot:post-parameters* hunchentoot:*request*)))
        ;;           (format t "~a~%" data-string)
                  ;;(json-obj (jsown:parse data-string))) ;; use jsown to parse the string
             ;; play with json-obj
             ;;  data-string))))) ;; return the original post data string, so that the save() in backbone.js will be notified about the success.
;;(defun run-this ()(format t \"~a~%\" (list (machine-instance) (machine-type) (machine-version) (software-type) (software-version))))(run-this)
