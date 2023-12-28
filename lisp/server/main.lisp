(ql:quickload "hunchentoot")

(defvar *implants* ())

(hunchentoot:start 
  (make-instance 'hunchentoot:easy-acceptor :port 3000))
(hunchentoot:define-easy-handler (foo :uri "/") ()
  (let ((request-type (hunchentoot:request-method hunchentoot:*request*))
        (postdata (hunchentoot:post-parameters* hunchentoot:*request*))
        (implantname (hunchentoot:post-parameter "implantname" hunchentoot:*request*)))
    (cond ((eq request-type :get) (format t "GOT GET~%"));; handle get request
          ((eq request-type :post)
            (format t "~a~%" postdata)
            (format t "~a~%" (assoc 'implantname postdata))
            (format t "~a~%" implantname)))))
        ;;    (let* ((data-string (hunchentoot:post-parameters* hunchentoot:*request*)))
        ;;           (format t "~a~%" data-string)
                  ;;(json-obj (jsown:parse data-string))) ;; use jsown to parse the string
             ;; play with json-obj
             ;;  data-string))))) ;; return the original post data string, so that the save() in backbone.js will be notified about the success.