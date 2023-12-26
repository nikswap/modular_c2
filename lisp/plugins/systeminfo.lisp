(defun run-this ()
  (format t "~a~%" (list (machine-instance) (machine-type) (machine-version) (software-type) (software-version))))

(run-this)
