<h3>Part 1</h3>
<p>
<br>Using $y – y_1 = m(x – x_1)$ to find the intersection point, where m is slope.
<br>$y – y_1 = m_1*(x – x_1)$ -> first equation
<br>$y – y_2 = m_2*(x – x_2)$ -> second equation
<br>first equation minus second equation:
<br>$y_2 - y_1 = m_1*x - m_1* x_1 - m_2*x + m_2 * x_2$
<br>after some modification:
<br>$x = (y_1 - y_2 - m_1 * x_1 + m_2 * x_2) / (m_2 - m_1)$
<br>$y = m_1*(x – x_1) + y_1$
</p>

<h3>Part 2</h3>
<p>
<br>Assuming that if a thrown rock at time "t" hits another rock then:
<br>$x_r + t * v_xr = x_1 + t * v_x1$ so
<br>$t = (x_r - x_1) / (v_x1 - v_xr)$
<br>and comparing x with y there is equation:
<br>$(x_r - x_1) / (v_x1 - v_xr) = (y_r - y_1) / (v_y1 - v_yr)$
so:
<br>$(v_y2-v_y1) * x_r + (v_x1-v_x2) * y_r + (y_1-y_2) * v_xr + (x_1-x_2) * v_yr = x_2* v_y2 - y_2 * v_x2 - x_2 * v_y1 + y_1 * v_x1$
<br>Using this equation and data for 4 rocks it is possible to solve it using gaussian elimination.
</p>
