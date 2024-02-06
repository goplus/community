var Lt = typeof globalThis < "u" ? globalThis : typeof window < "u" ? window : typeof global < "u" ? global : typeof self < "u" ? self : {};
function rt(e, t) {
  return e(t = { exports: {} }, t.exports), t.exports;
}
var $t, Hr, xr = function(e) {
  return e && e.Math == Math && e;
}, O = xr(typeof globalThis == "object" && globalThis) || xr(typeof window == "object" && window) || xr(typeof self == "object" && self) || xr(typeof Lt == "object" && Lt) || /* @__PURE__ */ function() {
  return this;
}() || Function("return this")(), J = function(e) {
  try {
    return !!e();
  } catch {
    return !0;
  }
}, jn = !J(function() {
  var e = (function() {
  }).bind();
  return typeof e != "function" || e.hasOwnProperty("prototype");
}), zl = Function.prototype, os = zl.apply, ss = zl.call, Mt = typeof Reflect == "object" && Reflect.apply || (jn ? ss.bind(os) : function() {
  return ss.apply(os, arguments);
}), Ul = Function.prototype, Bd = Ul.bind, Si = Ul.call, Hd = jn && Bd.bind(Si, Si), z = jn ? function(e) {
  return e && Hd(e);
} : function(e) {
  return e && function() {
    return Si.apply(e, arguments);
  };
}, oe = function(e) {
  return typeof e == "function";
}, re = !J(function() {
  return Object.defineProperty({}, 1, { get: function() {
    return 7;
  } })[1] != 7;
}), Cr = Function.prototype.call, Z = jn ? Cr.bind(Cr) : function() {
  return Cr.apply(Cr, arguments);
}, cs = {}.propertyIsEnumerable, ls = Object.getOwnPropertyDescriptor, eo = { f: ls && !cs.call({ 1: 2 }, 1) ? function(e) {
  var t = ls(this, e);
  return !!t && t.enumerable;
} : cs }, Vt = function(e, t) {
  return { enumerable: !(1 & e), configurable: !(2 & e), writable: !(4 & e), value: t };
}, zd = z({}.toString), Ud = z("".slice), ht = function(e) {
  return Ud(zd(e), 8, -1);
}, Ma = O.Object, Wd = z("".split), to = J(function() {
  return !Ma("z").propertyIsEnumerable(0);
}) ? function(e) {
  return ht(e) == "String" ? Wd(e, "") : Ma(e);
} : Ma, qd = O.TypeError, an = function(e) {
  if (e == null)
    throw qd("Can't call method on " + e);
  return e;
}, Ve = function(e) {
  return to(an(e));
}, ye = function(e) {
  return typeof e == "object" ? e !== null : oe(e);
}, X = {}, us = function(e) {
  return oe(e) ? e : void 0;
}, ze = function(e, t) {
  return arguments.length < 2 ? us(X[e]) || us(O[e]) : X[e] && X[e][t] || O[e] && O[e][t];
}, de = z({}.isPrototypeOf), zr = ze("navigator", "userAgent") || "", fs = O.process, ds = O.Deno, ps = fs && fs.versions || ds && ds.version, gs = ps && ps.v8;
gs && (Hr = ($t = gs.split("."))[0] > 0 && $t[0] < 4 ? 1 : +($t[0] + $t[1])), !Hr && zr && (!($t = zr.match(/Edge\/(\d+)/)) || $t[1] >= 74) && ($t = zr.match(/Chrome\/(\d+)/)) && (Hr = +$t[1]);
var Tn = Hr, et = !!Object.getOwnPropertySymbols && !J(function() {
  var e = Symbol();
  return !String(e) || !(Object(e) instanceof Symbol) || !Symbol.sham && Tn && Tn < 41;
}), no = et && !Symbol.sham && typeof Symbol.iterator == "symbol", Gd = O.Object, $n = no ? function(e) {
  return typeof e == "symbol";
} : function(e) {
  var t = ze("Symbol");
  return oe(t) && de(t.prototype, Gd(e));
}, Kd = O.String, ur = function(e) {
  try {
    return Kd(e);
  } catch {
    return "Object";
  }
}, Zd = O.TypeError, ne = function(e) {
  if (oe(e))
    return e;
  throw Zd(ur(e) + " is not a function");
}, Kr = function(e, t) {
  var r = e[t];
  return r == null ? void 0 : ne(r);
}, Yd = O.TypeError, Xd = Object.defineProperty, hs = "__core-js_shared__", jt = O[hs] || function(e, t) {
  try {
    Xd(O, e, { value: t, configurable: !0, writable: !0 });
  } catch {
    O[e] = t;
  }
  return t;
}(hs, {}), zt = rt(function(e) {
  (e.exports = function(t, r) {
    return jt[t] || (jt[t] = r !== void 0 ? r : {});
  })("versions", []).push({ version: "3.22.6", mode: "pure", copyright: "© 2014-2022 Denis Pushkarev (zloirock.ru)", license: "https://github.com/zloirock/core-js/blob/v3.22.6/LICENSE", source: "https://github.com/zloirock/core-js" });
}), Vd = O.Object, lt = function(e) {
  return Vd(an(e));
}, Jd = z({}.hasOwnProperty), F = Object.hasOwn || function(e, t) {
  return Jd(lt(e), t);
}, Qd = 0, ep = Math.random(), tp = z(1 .toString), nr = function(e) {
  return "Symbol(" + (e === void 0 ? "" : e) + ")_" + tp(++Qd + ep, 36);
}, Zn = zt("wks"), Xt = O.Symbol, ms = Xt && Xt.for, np = no ? Xt : Xt && Xt.withoutSetter || nr, _e = function(e) {
  if (!F(Zn, e) || !et && typeof Zn[e] != "string") {
    var t = "Symbol." + e;
    et && F(Xt, e) ? Zn[e] = Xt[e] : Zn[e] = no && ms ? ms(t) : np(t);
  }
  return Zn[e];
}, rp = O.TypeError, ap = _e("toPrimitive"), ip = function(e, t) {
  if (!ye(e) || $n(e))
    return e;
  var r, n = Kr(e, ap);
  if (n) {
    if (t === void 0 && (t = "default"), r = Z(n, e, t), !ye(r) || $n(r))
      return r;
    throw rp("Can't convert object to primitive value");
  }
  return t === void 0 && (t = "number"), function(a, i) {
    var o, s;
    if (i === "string" && oe(o = a.toString) && !ye(s = Z(o, a)) || oe(o = a.valueOf) && !ye(s = Z(o, a)) || i !== "string" && oe(o = a.toString) && !ye(s = Z(o, a)))
      return s;
    throw Yd("Can't convert object to primitive value");
  }(e, t);
}, Jt = function(e) {
  var t = ip(e, "string");
  return $n(t) ? t : t + "";
}, Ai = O.document, op = ye(Ai) && ye(Ai.createElement), Wl = function(e) {
  return op ? Ai.createElement(e) : {};
}, ql = !re && !J(function() {
  return Object.defineProperty(Wl("div"), "a", { get: function() {
    return 7;
  } }).a != 7;
}), bs = Object.getOwnPropertyDescriptor, on = { f: re ? bs : function(e, t) {
  if (e = Ve(e), t = Jt(t), ql)
    try {
      return bs(e, t);
    } catch {
    }
  if (F(e, t))
    return Vt(!Z(eo.f, e, t), e[t]);
} }, sp = /#|\.prototype\./, fr = function(e, t) {
  var r = lp[cp(e)];
  return r == fp || r != up && (oe(t) ? J(t) : !!t);
}, cp = fr.normalize = function(e) {
  return String(e).replace(sp, ".").toLowerCase();
}, lp = fr.data = {}, up = fr.NATIVE = "N", fp = fr.POLYFILL = "P", dp = fr, pp = z(z.bind), Ze = function(e, t) {
  return ne(e), t === void 0 ? e : jn ? pp(e, t) : function() {
    return e.apply(t, arguments);
  };
}, Gl = re && J(function() {
  return Object.defineProperty(function() {
  }, "prototype", { value: 42, writable: !1 }).prototype != 42;
}), gp = O.String, hp = O.TypeError, K = function(e) {
  if (ye(e))
    return e;
  throw hp(gp(e) + " is not an object");
}, mp = O.TypeError, ja = Object.defineProperty, bp = Object.getOwnPropertyDescriptor, Da = "enumerable", Fa = "configurable", Ba = "writable", Je = { f: re ? Gl ? function(e, t, r) {
  if (K(e), t = Jt(t), K(r), typeof e == "function" && t === "prototype" && "value" in r && Ba in r && !r[Ba]) {
    var n = bp(e, t);
    n && n[Ba] && (e[t] = r.value, r = { configurable: Fa in r ? r[Fa] : n[Fa], enumerable: Da in r ? r[Da] : n[Da], writable: !1 });
  }
  return ja(e, t, r);
} : ja : function(e, t, r) {
  if (K(e), t = Jt(t), K(r), ql)
    try {
      return ja(e, t, r);
    } catch {
    }
  if ("get" in r || "set" in r)
    throw mp("Accessors not supported");
  return "value" in r && (e[t] = r.value), e;
} }, ot = re ? function(e, t, r) {
  return Je.f(e, t, Vt(1, r));
} : function(e, t, r) {
  return e[t] = r, e;
}, vp = on.f, yp = function(e) {
  var t = function(r, n, a) {
    if (this instanceof t) {
      switch (arguments.length) {
        case 0:
          return new e();
        case 1:
          return new e(r);
        case 2:
          return new e(r, n);
      }
      return new e(r, n, a);
    }
    return Mt(e, this, arguments);
  };
  return t.prototype = e.prototype, t;
}, L = function(e, t) {
  var r, n, a, i, o, s, c, u, l = e.target, f = e.global, p = e.stat, d = e.proto, g = f ? O : p ? O[l] : (O[l] || {}).prototype, _ = f ? X : X[l] || ot(X, l, {})[l], m = _.prototype;
  for (a in t)
    r = !dp(f ? a : l + (p ? "." : "#") + a, e.forced) && g && F(g, a), o = _[a], r && (s = e.dontCallGetSet ? (u = vp(g, a)) && u.value : g[a]), i = r && s ? s : t[a], r && typeof o == typeof i || (c = e.bind && r ? Ze(i, O) : e.wrap && r ? yp(i) : d && oe(i) ? z(i) : i, (e.sham || i && i.sham || o && o.sham) && ot(c, "sham", !0), ot(_, a, c), d && (F(X, n = l + "Prototype") || ot(X, n, {}), ot(X[n], a, i), e.real && m && !m[a] && ot(m, a, i)));
}, Qt = z([].slice), vs = O.Function, _p = z([].concat), kp = z([].join), Ha = {}, xi = jn ? vs.bind : function(e) {
  var t = ne(this), r = t.prototype, n = Qt(arguments, 1), a = function() {
    var i = _p(n, Qt(arguments));
    return this instanceof a ? function(o, s, c) {
      if (!F(Ha, s)) {
        for (var u = [], l = 0; l < s; l++)
          u[l] = "a[" + l + "]";
        Ha[s] = vs("C,a", "return new C(" + kp(u, ",") + ")");
      }
      return Ha[s](o, c);
    }(t, i.length, i) : t.apply(e, i);
  };
  return ye(r) && (a.prototype = r), a;
}, Kl = {};
Kl[_e("toStringTag")] = "z";
var ro = String(Kl) === "[object z]", wp = _e("toStringTag"), Ep = O.Object, Sp = ht(/* @__PURE__ */ function() {
  return arguments;
}()) == "Arguments", Dn = ro ? ht : function(e) {
  var t, r, n;
  return e === void 0 ? "Undefined" : e === null ? "Null" : typeof (r = function(a, i) {
    try {
      return a[i];
    } catch {
    }
  }(t = Ep(e), wp)) == "string" ? r : Sp ? ht(t) : (n = ht(t)) == "Object" && oe(t.callee) ? "Arguments" : n;
}, Ap = z(Function.toString);
oe(jt.inspectSource) || (jt.inspectSource = function(e) {
  return Ap(e);
});
var Zl = jt.inspectSource, Yl = function() {
}, xp = [], Xl = ze("Reflect", "construct"), ao = /^\s*(?:class|function)\b/, Cp = z(ao.exec), Tp = !ao.exec(Yl), Yn = function(e) {
  if (!oe(e))
    return !1;
  try {
    return Xl(Yl, xp, e), !0;
  } catch {
    return !1;
  }
}, Vl = function(e) {
  if (!oe(e))
    return !1;
  switch (Dn(e)) {
    case "AsyncFunction":
    case "GeneratorFunction":
    case "AsyncGeneratorFunction":
      return !1;
  }
  try {
    return Tp || !!Cp(ao, Zl(e));
  } catch {
    return !0;
  }
};
Vl.sham = !0;
var Tr, ia = !Xl || J(function() {
  var e;
  return Yn(Yn.call) || !Yn(Object) || !Yn(function() {
    e = !0;
  }) || e;
}) ? Vl : Yn, $p = O.TypeError, Zr = function(e) {
  if (ia(e))
    return e;
  throw $p(ur(e) + " is not a constructor");
}, Rp = Math.ceil, Op = Math.floor, Pp = Math.trunc || function(e) {
  var t = +e;
  return (t > 0 ? Op : Rp)(t);
}, dr = function(e) {
  var t = +e;
  return t != t || t === 0 ? 0 : Pp(t);
}, Lp = Math.max, Ip = Math.min, At = function(e, t) {
  var r = dr(e);
  return r < 0 ? Lp(r + t, 0) : Ip(r, t);
}, Np = Math.min, Jl = function(e) {
  return e > 0 ? Np(dr(e), 9007199254740991) : 0;
}, bt = function(e) {
  return Jl(e.length);
}, ys = function(e) {
  return function(t, r, n) {
    var a, i = Ve(t), o = bt(i), s = At(n, o);
    if (e && r != r) {
      for (; o > s; )
        if ((a = i[s++]) != a)
          return !0;
    } else
      for (; o > s; s++)
        if ((e || s in i) && i[s] === r)
          return e || s || 0;
    return !e && -1;
  };
}, io = { includes: ys(!0), indexOf: ys(!1) }, Fn = {}, Mp = io.indexOf, _s = z([].push), Ql = function(e, t) {
  var r, n = Ve(e), a = 0, i = [];
  for (r in n)
    !F(Fn, r) && F(n, r) && _s(i, r);
  for (; t.length > a; )
    F(n, r = t[a++]) && (~Mp(i, r) || _s(i, r));
  return i;
}, Yr = ["constructor", "hasOwnProperty", "isPrototypeOf", "propertyIsEnumerable", "toLocaleString", "toString", "valueOf"], pr = Object.keys || function(e) {
  return Ql(e, Yr);
}, jp = re && !Gl ? Object.defineProperties : function(e, t) {
  K(e);
  for (var r, n = Ve(t), a = pr(t), i = a.length, o = 0; i > o; )
    Je.f(e, r = a[o++], n[r]);
  return e;
}, oo = { f: jp }, Dp = ze("document", "documentElement"), ks = zt("keys"), oa = function(e) {
  return ks[e] || (ks[e] = nr(e));
}, Ci = "prototype", Ti = "script", eu = oa("IE_PROTO"), za = function() {
}, tu = function(e) {
  return "<" + Ti + ">" + e + "</" + Ti + ">";
}, ws = function(e) {
  e.write(tu("")), e.close();
  var t = e.parentWindow.Object;
  return e = null, t;
}, Ur = function() {
  try {
    Tr = new ActiveXObject("htmlfile");
  } catch {
  }
  var e, t, r;
  Ur = typeof document < "u" ? document.domain && Tr ? ws(Tr) : (t = Wl("iframe"), r = "java" + Ti + ":", t.style.display = "none", Dp.appendChild(t), t.src = String(r), (e = t.contentWindow.document).open(), e.write(tu("document.F=Object")), e.close(), e.F) : ws(Tr);
  for (var n = Yr.length; n--; )
    delete Ur[Ci][Yr[n]];
  return Ur();
};
Fn[eu] = !0;
var mt = Object.create || function(e, t) {
  var r;
  return e !== null ? (za[Ci] = K(e), r = new za(), za[Ci] = null, r[eu] = e) : r = Ur(), t === void 0 ? r : oo.f(r, t);
}, so = ze("Reflect", "construct"), Fp = Object.prototype, Bp = [].push, nu = J(function() {
  function e() {
  }
  return !(so(function() {
  }, [], e) instanceof e);
}), ru = !J(function() {
  so(function() {
  });
}), Es = nu || ru;
L({ target: "Reflect", stat: !0, forced: Es, sham: Es }, { construct: function(e, t) {
  Zr(e), K(t);
  var r = arguments.length < 3 ? e : Zr(arguments[2]);
  if (ru && !nu)
    return so(e, t, r);
  if (e == r) {
    switch (t.length) {
      case 0:
        return new e();
      case 1:
        return new e(t[0]);
      case 2:
        return new e(t[0], t[1]);
      case 3:
        return new e(t[0], t[1], t[2]);
      case 4:
        return new e(t[0], t[1], t[2], t[3]);
    }
    var n = [null];
    return Mt(Bp, n, t), new (Mt(xi, e, n))();
  }
  var a = r.prototype, i = mt(ye(a) ? a : Fp), o = Mt(e, i, t);
  return ye(o) ? o : i;
} });
var au = X.Reflect.construct, k = au, Ss = Je.f;
L({ target: "Object", stat: !0, forced: Object.defineProperty !== Ss, sham: !re }, { defineProperty: Ss });
var Hp = rt(function(e) {
  var t = X.Object, r = e.exports = function(n, a, i) {
    return t.defineProperty(n, a, i);
  };
  t.defineProperty.sham && (r.sham = !0);
}), iu = Hp, sa = iu;
function As(e, t) {
  for (var r = 0; r < t.length; r++) {
    var n = t[r];
    n.enumerable = n.enumerable || !1, n.configurable = !0, "value" in n && (n.writable = !0), sa(e, n.key, n);
  }
}
function M(e, t, r) {
  return t && As(e.prototype, t), r && As(e, r), sa(e, "prototype", { writable: !1 }), e;
}
function j(e, t) {
  if (!(e instanceof t))
    throw new TypeError("Cannot call a class as a function");
}
L({ target: "Object", stat: !0, sham: !re }, { create: mt });
var zp = X.Object, ou = function(e, t) {
  return zp.create(e, t);
}, su = ou, Up = O.String, Wp = O.TypeError, qp = Object.setPrototypeOf || ("__proto__" in {} ? function() {
  var e, t = !1, r = {};
  try {
    (e = z(Object.getOwnPropertyDescriptor(Object.prototype, "__proto__").set))(r, []), t = r instanceof Array;
  } catch {
  }
  return function(n, a) {
    return K(n), function(i) {
      if (typeof i == "object" || oe(i))
        return i;
      throw Wp("Can't set " + Up(i) + " as a prototype");
    }(a), t ? e(n, a) : n.__proto__ = a, n;
  };
}() : void 0);
L({ target: "Object", stat: !0 }, { setPrototypeOf: qp });
var cu = X.Object.setPrototypeOf;
function rr(e, t) {
  return rr = cu || function(r, n) {
    return r.__proto__ = n, r;
  }, rr(e, t);
}
function U(e, t) {
  if (typeof t != "function" && t !== null)
    throw new TypeError("Super expression must either be null or a function");
  e.prototype = su(t && t.prototype, { constructor: { value: e, writable: !0, configurable: !0 } }), sa(e, "prototype", { writable: !1 }), t && rr(e, t);
}
var en = Array.isArray || function(e) {
  return ht(e) == "Array";
}, Ft = function(e, t, r) {
  var n = Jt(t);
  n in e ? Je.f(e, n, Vt(0, r)) : e[n] = r;
}, Gp = _e("species"), xs = O.Array, co = function(e, t) {
  return new (function(r) {
    var n;
    return en(r) && (n = r.constructor, (ia(n) && (n === xs || en(n.prototype)) || ye(n) && (n = n[Gp]) === null) && (n = void 0)), n === void 0 ? xs : n;
  }(e))(t === 0 ? 0 : t);
}, Kp = _e("species"), gr = function(e) {
  return Tn >= 51 || !J(function() {
    var t = [];
    return (t.constructor = {})[Kp] = function() {
      return { foo: 1 };
    }, t[e](Boolean).foo !== 1;
  });
}, lu = _e("isConcatSpreadable"), Cs = 9007199254740991, Ts = "Maximum allowed index exceeded", $s = O.TypeError, Zp = Tn >= 51 || !J(function() {
  var e = [];
  return e[lu] = !1, e.concat()[0] !== e;
}), Yp = gr("concat"), Xp = function(e) {
  if (!ye(e))
    return !1;
  var t = e[lu];
  return t !== void 0 ? !!t : en(e);
};
L({ target: "Array", proto: !0, arity: 1, forced: !Zp || !Yp }, { concat: function(e) {
  var t, r, n, a, i, o = lt(this), s = co(o, 0), c = 0;
  for (t = -1, n = arguments.length; t < n; t++)
    if (Xp(i = t === -1 ? o : arguments[t])) {
      if (c + (a = bt(i)) > Cs)
        throw $s(Ts);
      for (r = 0; r < a; r++, c++)
        r in i && Ft(s, c, i[r]);
    } else {
      if (c >= Cs)
        throw $s(Ts);
      Ft(s, c++, i);
    }
  return s.length = c, s;
} });
var Xr, tr, Vr, Vp = O.String, tt = function(e) {
  if (Dn(e) === "Symbol")
    throw TypeError("Cannot convert a Symbol value to a string");
  return Vp(e);
}, Jp = Yr.concat("length", "prototype"), ar = { f: Object.getOwnPropertyNames || function(e) {
  return Ql(e, Jp);
} }, Qp = O.Array, eg = Math.max, uu = ar.f, fu = typeof window == "object" && window && Object.getOwnPropertyNames ? Object.getOwnPropertyNames(window) : [], tg = function(e) {
  try {
    return uu(e);
  } catch {
    return function(r, n, a) {
      for (var i = bt(r), o = At(n, i), s = At(a === void 0 ? i : a, i), c = Qp(eg(s - o, 0)), u = 0; o < s; o++, u++)
        Ft(c, u, r[o]);
      return c.length = u, c;
    }(fu);
  }
}, lo = { f: function(e) {
  return fu && ht(e) == "Window" ? tg(e) : uu(Ve(e));
} }, ca = { f: Object.getOwnPropertySymbols }, tn = function(e, t, r, n) {
  return n && n.enumerable ? e[t] = r : ot(e, t, r), e;
}, uo = { f: _e }, ng = Je.f, be = function(e) {
  var t = X.Symbol || (X.Symbol = {});
  F(t, e) || ng(t, e, { value: uo.f(e) });
}, du = function() {
  var e = ze("Symbol"), t = e && e.prototype, r = t && t.valueOf, n = _e("toPrimitive");
  t && !t[n] && tn(t, n, function(a) {
    return Z(r, this);
  }, { arity: 1 });
}, rg = ro ? {}.toString : function() {
  return "[object " + Dn(this) + "]";
}, ag = Je.f, Rs = _e("toStringTag"), Rn = function(e, t, r, n) {
  if (e) {
    var a = r ? e : e.prototype;
    F(a, Rs) || ag(a, Rs, { configurable: !0, value: t }), n && !ro && ot(a, "toString", rg);
  }
}, Os = O.WeakMap, ig = oe(Os) && /native code/.test(Zl(Os)), Ps = "Object already initialized", $i = O.TypeError, og = O.WeakMap;
if (ig || jt.state) {
  var Zt = jt.state || (jt.state = new og()), sg = z(Zt.get), Ls = z(Zt.has), cg = z(Zt.set);
  Xr = function(e, t) {
    if (Ls(Zt, e))
      throw new $i(Ps);
    return t.facade = e, cg(Zt, e, t), t;
  }, tr = function(e) {
    return sg(Zt, e) || {};
  }, Vr = function(e) {
    return Ls(Zt, e);
  };
} else {
  var bn = oa("state");
  Fn[bn] = !0, Xr = function(e, t) {
    if (F(e, bn))
      throw new $i(Ps);
    return t.facade = e, ot(e, bn, t), t;
  }, tr = function(e) {
    return F(e, bn) ? e[bn] : {};
  }, Vr = function(e) {
    return F(e, bn);
  };
}
var vt = { set: Xr, get: tr, has: Vr, enforce: function(e) {
  return Vr(e) ? tr(e) : Xr(e, {});
}, getterFor: function(e) {
  return function(t) {
    var r;
    if (!ye(t) || (r = tr(t)).type !== e)
      throw $i("Incompatible receiver, " + e + " required");
    return r;
  };
} }, Is = z([].push), Rt = function(e) {
  var t = e == 1, r = e == 2, n = e == 3, a = e == 4, i = e == 6, o = e == 7, s = e == 5 || i;
  return function(c, u, l, f) {
    for (var p, d, g = lt(c), _ = to(g), m = Ze(u, l), h = bt(_), v = 0, w = f || co, E = t ? w(c, h) : r || o ? w(c, 0) : void 0; h > v; v++)
      if ((s || v in _) && (d = m(p = _[v], v, g), e))
        if (t)
          E[v] = d;
        else if (d)
          switch (e) {
            case 3:
              return !0;
            case 5:
              return p;
            case 6:
              return v;
            case 2:
              Is(E, p);
          }
        else
          switch (e) {
            case 4:
              return !1;
            case 7:
              Is(E, p);
          }
    return i ? -1 : n || a ? a : E;
  };
}, sn = { forEach: Rt(0), map: Rt(1), filter: Rt(2), some: Rt(3), every: Rt(4), find: Rt(5), findIndex: Rt(6), filterReject: Rt(7) }, la = sn.forEach, We = oa("hidden"), Jr = "Symbol", ir = "prototype", lg = vt.set, Ns = vt.getterFor(Jr), st = Object[ir], En = O.Symbol, Qn = En && En[ir], ug = O.TypeError, Ua = O.QObject, pu = on.f, Pt = Je.f, gu = lo.f, fg = eo.f, hu = z([].push), xt = zt("symbols"), hr = zt("op-symbols"), dg = zt("wks"), Wa = !Ua || !Ua[ir] || !Ua[ir].findChild, Ri = re && J(function() {
  return mt(Pt({}, "a", { get: function() {
    return Pt(this, "a", { value: 7 }).a;
  } })).a != 7;
}) ? function(e, t, r) {
  var n = pu(st, t);
  n && delete st[t], Pt(e, t, r), n && e !== st && Pt(st, t, n);
} : Pt, qa = function(e, t) {
  var r = xt[e] = mt(Qn);
  return lg(r, { type: Jr, tag: e, description: t }), re || (r.description = t), r;
}, Qr = function(e, t, r) {
  e === st && Qr(hr, t, r), K(e);
  var n = Jt(t);
  return K(r), F(xt, n) ? (r.enumerable ? (F(e, We) && e[We][n] && (e[We][n] = !1), r = mt(r, { enumerable: Vt(0, !1) })) : (F(e, We) || Pt(e, We, Vt(1, {})), e[We][n] = !0), Ri(e, n, r)) : Pt(e, n, r);
}, Ga = function(e, t) {
  K(e);
  var r = Ve(t), n = pr(r).concat(bu(r));
  return la(n, function(a) {
    re && !Z(mu, r, a) || Qr(e, a, r[a]);
  }), e;
}, mu = function(e) {
  var t = Jt(e), r = Z(fg, this, t);
  return !(this === st && F(xt, t) && !F(hr, t)) && (!(r || !F(this, t) || !F(xt, t) || F(this, We) && this[We][t]) || r);
}, Ms = function(e, t) {
  var r = Ve(e), n = Jt(t);
  if (r !== st || !F(xt, n) || F(hr, n)) {
    var a = pu(r, n);
    return !a || !F(xt, n) || F(r, We) && r[We][n] || (a.enumerable = !0), a;
  }
}, js = function(e) {
  var t = gu(Ve(e)), r = [];
  return la(t, function(n) {
    F(xt, n) || F(Fn, n) || hu(r, n);
  }), r;
}, bu = function(e) {
  var t = e === st, r = gu(t ? hr : Ve(e)), n = [];
  return la(r, function(a) {
    !F(xt, a) || t && !F(st, a) || hu(n, xt[a]);
  }), n;
};
et || (Qn = (En = function() {
  if (de(Qn, this))
    throw ug("Symbol is not a constructor");
  var e = arguments.length && arguments[0] !== void 0 ? tt(arguments[0]) : void 0, t = nr(e), r = function(n) {
    this === st && Z(r, hr, n), F(this, We) && F(this[We], t) && (this[We][t] = !1), Ri(this, t, Vt(1, n));
  };
  return re && Wa && Ri(st, t, { configurable: !0, set: r }), qa(t, e);
})[ir], tn(Qn, "toString", function() {
  return Ns(this).tag;
}), tn(En, "withoutSetter", function(e) {
  return qa(nr(e), e);
}), eo.f = mu, Je.f = Qr, oo.f = Ga, on.f = Ms, ar.f = lo.f = js, ca.f = bu, uo.f = function(e) {
  return qa(_e(e), e);
}, re && Pt(Qn, "description", { configurable: !0, get: function() {
  return Ns(this).description;
} })), L({ global: !0, constructor: !0, wrap: !0, forced: !et, sham: !et }, { Symbol: En }), la(pr(dg), function(e) {
  be(e);
}), L({ target: Jr, stat: !0, forced: !et }, { useSetter: function() {
  Wa = !0;
}, useSimple: function() {
  Wa = !1;
} }), L({ target: "Object", stat: !0, forced: !et, sham: !re }, { create: function(e, t) {
  return t === void 0 ? mt(e) : Ga(mt(e), t);
}, defineProperty: Qr, defineProperties: Ga, getOwnPropertyDescriptor: Ms }), L({ target: "Object", stat: !0, forced: !et }, { getOwnPropertyNames: js }), du(), Rn(En, Jr), Fn[We] = !0;
var vu = et && !!Symbol.for && !!Symbol.keyFor, Ka = zt("string-to-symbol-registry"), pg = zt("symbol-to-string-registry");
L({ target: "Symbol", stat: !0, forced: !vu }, { for: function(e) {
  var t = tt(e);
  if (F(Ka, t))
    return Ka[t];
  var r = ze("Symbol")(t);
  return Ka[t] = r, pg[r] = t, r;
} });
var Ds = zt("symbol-to-string-registry");
L({ target: "Symbol", stat: !0, forced: !vu }, { keyFor: function(e) {
  if (!$n(e))
    throw TypeError(ur(e) + " is not a symbol");
  if (F(Ds, e))
    return Ds[e];
} });
var Dt = ze("JSON", "stringify"), $r = z(/./.exec), Fs = z("".charAt), gg = z("".charCodeAt), hg = z("".replace), mg = z(1 .toString), bg = /[\uD800-\uDFFF]/g, Bs = /^[\uD800-\uDBFF]$/, Hs = /^[\uDC00-\uDFFF]$/, zs = !et || J(function() {
  var e = ze("Symbol")();
  return Dt([e]) != "[null]" || Dt({ a: e }) != "{}" || Dt(Object(e)) != "{}";
}), Us = J(function() {
  return Dt("\uDF06\uD834") !== '"\\udf06\\ud834"' || Dt("\uDEAD") !== '"\\udead"';
}), vg = function(e, t) {
  var r = Qt(arguments), n = t;
  if ((ye(t) || e !== void 0) && !$n(e))
    return en(t) || (t = function(a, i) {
      if (oe(n) && (i = Z(n, this, a, i)), !$n(i))
        return i;
    }), r[1] = t, Mt(Dt, null, r);
}, yg = function(e, t, r) {
  var n = Fs(r, t - 1), a = Fs(r, t + 1);
  return $r(Bs, e) && !$r(Hs, a) || $r(Hs, e) && !$r(Bs, n) ? "\\u" + mg(gg(e, 0), 16) : e;
};
Dt && L({ target: "JSON", stat: !0, arity: 3, forced: zs || Us }, { stringify: function(e, t, r) {
  var n = Qt(arguments), a = Mt(zs ? vg : Dt, null, n);
  return Us && typeof a == "string" ? hg(a, bg, yg) : a;
} });
var _g = !et || J(function() {
  ca.f(1);
});
L({ target: "Object", stat: !0, forced: _g }, { getOwnPropertySymbols: function(e) {
  var t = ca.f;
  return t ? t(lt(e)) : [];
} }), be("asyncIterator"), be("hasInstance"), be("isConcatSpreadable"), be("iterator"), be("match"), be("matchAll"), be("replace"), be("search"), be("species"), be("split"), be("toPrimitive"), du(), be("toStringTag"), Rn(ze("Symbol"), "Symbol"), be("unscopables"), Rn(O.JSON, "JSON", !0);
var It, Ws, qs, kg = X.Symbol, Et = {}, yu = Function.prototype, wg = re && Object.getOwnPropertyDescriptor, Za = F(yu, "name"), _u = { EXISTS: Za, PROPER: Za && (function() {
}).name === "something", CONFIGURABLE: Za && (!re || re && wg(yu, "name").configurable) }, ku = !J(function() {
  function e() {
  }
  return e.prototype.constructor = null, Object.getPrototypeOf(new e()) !== e.prototype;
}), Gs = oa("IE_PROTO"), Oi = O.Object, Eg = Oi.prototype, On = ku ? Oi.getPrototypeOf : function(e) {
  var t = lt(e);
  if (F(t, Gs))
    return t[Gs];
  var r = t.constructor;
  return oe(r) && t instanceof r ? r.prototype : t instanceof Oi ? Eg : null;
}, Pi = _e("iterator"), wu = !1;
[].keys && ("next" in (qs = [].keys()) ? (Ws = On(On(qs))) !== Object.prototype && (It = Ws) : wu = !0);
var Sg = It == null || J(function() {
  var e = {};
  return It[Pi].call(e) !== e;
});
It = Sg ? {} : mt(It), oe(It[Pi]) || tn(It, Pi, function() {
  return this;
});
var Eu = { IteratorPrototype: It, BUGGY_SAFARI_ITERATORS: wu }, Ag = Eu.IteratorPrototype, xg = function() {
  return this;
}, Cg = _u.PROPER, Rr = Eu.BUGGY_SAFARI_ITERATORS, Ya = _e("iterator"), Ks = "keys", Or = "values", Zs = "entries", Tg = function() {
  return this;
}, fo = function(e, t, r, n, a, i, o) {
  (function(h, v, w, E) {
    var S = v + " Iterator";
    h.prototype = mt(Ag, { next: Vt(+!E, w) }), Rn(h, S, !1, !0), Et[S] = xg;
  })(r, t, n);
  var s, c, u, l = function(h) {
    if (h === a && _)
      return _;
    if (!Rr && h in d)
      return d[h];
    switch (h) {
      case Ks:
      case Or:
      case Zs:
        return function() {
          return new r(this, h);
        };
    }
    return function() {
      return new r(this);
    };
  }, f = t + " Iterator", p = !1, d = e.prototype, g = d[Ya] || d["@@iterator"] || a && d[a], _ = !Rr && g || l(a), m = t == "Array" && d.entries || g;
  if (m && (s = On(m.call(new e()))) !== Object.prototype && s.next && (Rn(s, f, !0, !0), Et[f] = Tg), Cg && a == Or && g && g.name !== Or && (p = !0, _ = function() {
    return Z(g, this);
  }), a)
    if (c = { values: l(Or), keys: i ? _ : l(Ks), entries: l(Zs) }, o)
      for (u in c)
        (Rr || p || !(u in d)) && tn(d, u, c[u]);
    else
      L({ target: t, proto: !0, forced: Rr || p }, c);
  return o && d[Ya] !== _ && tn(d, Ya, _, { name: a }), Et[t] = _, c;
};
Je.f;
var Su = "Array Iterator", $g = vt.set, Rg = vt.getterFor(Su);
fo(Array, "Array", function(e, t) {
  $g(this, { type: Su, target: Ve(e), index: 0, kind: t });
}, function() {
  var e = Rg(this), t = e.target, r = e.kind, n = e.index++;
  return !t || n >= t.length ? (e.target = void 0, { value: void 0, done: !0 }) : r == "keys" ? { value: n, done: !1 } : r == "values" ? { value: t[n], done: !1 } : { value: [n, t[n]], done: !1 };
}, "values"), Et.Arguments = Et.Array;
var Ys = _e("toStringTag");
for (var Xa in { CSSRuleList: 0, CSSStyleDeclaration: 0, CSSValueList: 0, ClientRectList: 0, DOMRectList: 0, DOMStringList: 0, DOMTokenList: 1, DataTransferItemList: 0, FileList: 0, HTMLAllCollection: 0, HTMLCollection: 0, HTMLFormElement: 0, HTMLSelectElement: 0, MediaList: 0, MimeTypeArray: 0, NamedNodeMap: 0, NodeList: 1, PaintRequestList: 0, Plugin: 0, PluginArray: 0, SVGLengthList: 0, SVGNumberList: 0, SVGPathSegList: 0, SVGPointList: 0, SVGStringList: 0, SVGTransformList: 0, SourceBufferList: 0, StyleSheetList: 0, TextTrackCueList: 0, TextTrackList: 0, TouchList: 0 }) {
  var Xs = O[Xa], Va = Xs && Xs.prototype;
  Va && Dn(Va) !== Ys && ot(Va, Ys, Xa), Et[Xa] = Et.Array;
}
var Au = kg, Og = Au;
be("asyncDispose"), be("dispose"), be("matcher"), be("metadata"), be("observable"), be("patternMatch"), be("replaceAll");
var Sn = Og, Pg = z("".charAt), Vs = z("".charCodeAt), Lg = z("".slice), Js = function(e) {
  return function(t, r) {
    var n, a, i = tt(an(t)), o = dr(r), s = i.length;
    return o < 0 || o >= s ? e ? "" : void 0 : (n = Vs(i, o)) < 55296 || n > 56319 || o + 1 === s || (a = Vs(i, o + 1)) < 56320 || a > 57343 ? e ? Pg(i, o) : n : e ? Lg(i, o, o + 2) : a - 56320 + (n - 55296 << 10) + 65536;
  };
}, Ig = { codeAt: Js(!1), charAt: Js(!0) }.charAt, xu = "String Iterator", Ng = vt.set, Mg = vt.getterFor(xu);
fo(String, "String", function(e) {
  Ng(this, { type: xu, string: tt(e), index: 0 });
}, function() {
  var e, t = Mg(this), r = t.string, n = t.index;
  return n >= r.length ? { value: void 0, done: !0 } : (e = Ig(r, n), t.index += e.length, { value: e, done: !1 });
});
var jg = uo.f("iterator");
function He(e) {
  return He = typeof Sn == "function" && typeof jg == "symbol" ? function(t) {
    return typeof t;
  } : function(t) {
    return t && typeof Sn == "function" && t.constructor === Sn && t !== Sn.prototype ? "symbol" : typeof t;
  }, He(e);
}
function Be(e) {
  if (e === void 0)
    throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
  return e;
}
function B(e, t) {
  if (t && (He(t) === "object" || typeof t == "function"))
    return t;
  if (t !== void 0)
    throw new TypeError("Derived constructors may only return object or undefined");
  return Be(e);
}
var Dg = J(function() {
  On(1);
});
L({ target: "Object", stat: !0, forced: Dg, sham: !ku }, { getPrototypeOf: function(e) {
  return On(lt(e));
} });
var Qs = X.Object.getPrototypeOf;
function T(e) {
  return T = cu ? Qs : function(t) {
    return t.__proto__ || Qs(t);
  }, T(e);
}
function R(e, t, r) {
  return t in e ? sa(e, t, { value: r, enumerable: !0, configurable: !0, writable: !0 }) : e[t] = r, e;
}
var Fg = function() {
  this.__data__ = [], this.size = 0;
}, ua = function(e, t) {
  return e === t || e != e && t != t;
}, fa = function(e, t) {
  for (var r = e.length; r--; )
    if (ua(e[r][0], t))
      return r;
  return -1;
}, Bg = Array.prototype.splice, Hg = function(e) {
  var t = this.__data__, r = fa(t, e);
  return !(r < 0) && (r == t.length - 1 ? t.pop() : Bg.call(t, r, 1), --this.size, !0);
}, zg = function(e) {
  var t = this.__data__, r = fa(t, e);
  return r < 0 ? void 0 : t[r][1];
}, Ug = function(e) {
  return fa(this.__data__, e) > -1;
}, Wg = function(e, t) {
  var r = this.__data__, n = fa(r, e);
  return n < 0 ? (++this.size, r.push([e, t])) : r[n][1] = t, this;
};
function yn(e) {
  var t = -1, r = e == null ? 0 : e.length;
  for (this.clear(); ++t < r; ) {
    var n = e[t];
    this.set(n[0], n[1]);
  }
}
yn.prototype.clear = Fg, yn.prototype.delete = Hg, yn.prototype.get = zg, yn.prototype.has = Ug, yn.prototype.set = Wg;
var da = yn, qg = function() {
  this.__data__ = new da(), this.size = 0;
}, Gg = function(e) {
  var t = this.__data__, r = t.delete(e);
  return this.size = t.size, r;
}, Kg = function(e) {
  return this.__data__.get(e);
}, Zg = function(e) {
  return this.__data__.has(e);
}, Cu = typeof Lt == "object" && Lt && Lt.Object === Object && Lt, Yg = typeof self == "object" && self && self.Object === Object && self, yt = Cu || Yg || Function("return this")(), Bt = yt.Symbol, Tu = Object.prototype, Xg = Tu.hasOwnProperty, Vg = Tu.toString, Xn = Bt ? Bt.toStringTag : void 0, Jg = function(e) {
  var t = Xg.call(e, Xn), r = e[Xn];
  try {
    e[Xn] = void 0;
    var n = !0;
  } catch {
  }
  var a = Vg.call(e);
  return n && (t ? e[Xn] = r : delete e[Xn]), a;
}, Qg = Object.prototype.toString, eh = function(e) {
  return Qg.call(e);
}, ec = Bt ? Bt.toStringTag : void 0, cn = function(e) {
  return e == null ? e === void 0 ? "[object Undefined]" : "[object Null]" : ec && ec in Object(e) ? Jg(e) : eh(e);
}, Ut = function(e) {
  var t = typeof e;
  return e != null && (t == "object" || t == "function");
}, po = function(e) {
  if (!Ut(e))
    return !1;
  var t = cn(e);
  return t == "[object Function]" || t == "[object GeneratorFunction]" || t == "[object AsyncFunction]" || t == "[object Proxy]";
}, Ja = yt["__core-js_shared__"], tc = function() {
  var e = /[^.]+$/.exec(Ja && Ja.keys && Ja.keys.IE_PROTO || "");
  return e ? "Symbol(src)_1." + e : "";
}(), th = function(e) {
  return !!tc && tc in e;
}, nh = Function.prototype.toString, ln = function(e) {
  if (e != null) {
    try {
      return nh.call(e);
    } catch {
    }
    try {
      return e + "";
    } catch {
    }
  }
  return "";
}, rh = /^\[object .+?Constructor\]$/, ah = Function.prototype, ih = Object.prototype, oh = ah.toString, sh = ih.hasOwnProperty, ch = RegExp("^" + oh.call(sh).replace(/[\\^$.*+?()[\]{}|]/g, "\\$&").replace(/hasOwnProperty|(function).*?(?=\\\()| for .+?(?=\\\])/g, "$1.*?") + "$"), lh = function(e) {
  return !(!Ut(e) || th(e)) && (po(e) ? ch : rh).test(ln(e));
}, uh = function(e, t) {
  return e == null ? void 0 : e[t];
}, un = function(e, t) {
  var r = uh(e, t);
  return lh(r) ? r : void 0;
}, or = un(yt, "Map"), sr = un(Object, "create"), fh = function() {
  this.__data__ = sr ? sr(null) : {}, this.size = 0;
}, dh = function(e) {
  var t = this.has(e) && delete this.__data__[e];
  return this.size -= t ? 1 : 0, t;
}, ph = Object.prototype.hasOwnProperty, gh = function(e) {
  var t = this.__data__;
  if (sr) {
    var r = t[e];
    return r === "__lodash_hash_undefined__" ? void 0 : r;
  }
  return ph.call(t, e) ? t[e] : void 0;
}, hh = Object.prototype.hasOwnProperty, mh = function(e) {
  var t = this.__data__;
  return sr ? t[e] !== void 0 : hh.call(t, e);
}, bh = function(e, t) {
  var r = this.__data__;
  return this.size += this.has(e) ? 0 : 1, r[e] = sr && t === void 0 ? "__lodash_hash_undefined__" : t, this;
};
function _n(e) {
  var t = -1, r = e == null ? 0 : e.length;
  for (this.clear(); ++t < r; ) {
    var n = e[t];
    this.set(n[0], n[1]);
  }
}
_n.prototype.clear = fh, _n.prototype.delete = dh, _n.prototype.get = gh, _n.prototype.has = mh, _n.prototype.set = bh;
var nc = _n, vh = function() {
  this.size = 0, this.__data__ = { hash: new nc(), map: new (or || da)(), string: new nc() };
}, yh = function(e) {
  var t = typeof e;
  return t == "string" || t == "number" || t == "symbol" || t == "boolean" ? e !== "__proto__" : e === null;
}, pa = function(e, t) {
  var r = e.__data__;
  return yh(t) ? r[typeof t == "string" ? "string" : "hash"] : r.map;
}, _h = function(e) {
  var t = pa(this, e).delete(e);
  return this.size -= t ? 1 : 0, t;
}, kh = function(e) {
  return pa(this, e).get(e);
}, wh = function(e) {
  return pa(this, e).has(e);
}, Eh = function(e, t) {
  var r = pa(this, e), n = r.size;
  return r.set(e, t), this.size += r.size == n ? 0 : 1, this;
};
function kn(e) {
  var t = -1, r = e == null ? 0 : e.length;
  for (this.clear(); ++t < r; ) {
    var n = e[t];
    this.set(n[0], n[1]);
  }
}
kn.prototype.clear = vh, kn.prototype.delete = _h, kn.prototype.get = kh, kn.prototype.has = wh, kn.prototype.set = Eh;
var Sh = kn, Ah = function(e, t) {
  var r = this.__data__;
  if (r instanceof da) {
    var n = r.__data__;
    if (!or || n.length < 199)
      return n.push([e, t]), this.size = ++r.size, this;
    r = this.__data__ = new Sh(n);
  }
  return r.set(e, t), this.size = r.size, this;
};
function wn(e) {
  var t = this.__data__ = new da(e);
  this.size = t.size;
}
wn.prototype.clear = qg, wn.prototype.delete = Gg, wn.prototype.get = Kg, wn.prototype.has = Zg, wn.prototype.set = Ah;
var $u = wn, ea = function() {
  try {
    var e = un(Object, "defineProperty");
    return e({}, "", {}), e;
  } catch {
  }
}(), go = function(e, t, r) {
  t == "__proto__" && ea ? ea(e, t, { configurable: !0, enumerable: !0, value: r, writable: !0 }) : e[t] = r;
}, Li = function(e, t, r) {
  (r !== void 0 && !ua(e[t], r) || r === void 0 && !(t in e)) && go(e, t, r);
}, xh = /* @__PURE__ */ function(e) {
  return function(t, r, n) {
    for (var a = -1, i = Object(t), o = n(t), s = o.length; s--; ) {
      var c = o[e ? s : ++a];
      if (r(i[c], c, i) === !1)
        break;
    }
    return t;
  };
}(), Ru = rt(function(e, t) {
  var r = t && !t.nodeType && t, n = r && e && !e.nodeType && e, a = n && n.exports === r ? yt.Buffer : void 0, i = a ? a.allocUnsafe : void 0;
  e.exports = function(o, s) {
    if (s)
      return o.slice();
    var c = o.length, u = i ? i(c) : new o.constructor(c);
    return o.copy(u), u;
  };
}), rc = yt.Uint8Array, ho = function(e) {
  var t = new e.constructor(e.byteLength);
  return new rc(t).set(new rc(e)), t;
}, Ou = function(e, t) {
  var r = t ? ho(e.buffer) : e.buffer;
  return new e.constructor(r, e.byteOffset, e.length);
}, Pu = function(e, t) {
  var r = -1, n = e.length;
  for (t || (t = Array(n)); ++r < n; )
    t[r] = e[r];
  return t;
}, ac = Object.create, Ch = /* @__PURE__ */ function() {
  function e() {
  }
  return function(t) {
    if (!Ut(t))
      return {};
    if (ac)
      return ac(t);
    e.prototype = t;
    var r = new e();
    return e.prototype = void 0, r;
  };
}(), Lu = function(e, t) {
  return function(r) {
    return e(t(r));
  };
}, mo = Lu(Object.getPrototypeOf, Object), Th = Object.prototype, bo = function(e) {
  var t = e && e.constructor;
  return e === (typeof t == "function" && t.prototype || Th);
}, Iu = function(e) {
  return typeof e.constructor != "function" || bo(e) ? {} : Ch(mo(e));
}, Wt = function(e) {
  return e != null && typeof e == "object";
}, ic = function(e) {
  return Wt(e) && cn(e) == "[object Arguments]";
}, Nu = Object.prototype, $h = Nu.hasOwnProperty, Rh = Nu.propertyIsEnumerable, Ii = ic(/* @__PURE__ */ function() {
  return arguments;
}()) ? ic : function(e) {
  return Wt(e) && $h.call(e, "callee") && !Rh.call(e, "callee");
}, Pn = Array.isArray, Mu = function(e) {
  return typeof e == "number" && e > -1 && e % 1 == 0 && e <= 9007199254740991;
}, ga = function(e) {
  return e != null && Mu(e.length) && !po(e);
}, Oh = function(e) {
  return Wt(e) && ga(e);
}, Ph = function() {
  return !1;
}, vo = rt(function(e, t) {
  var r = t && !t.nodeType && t, n = r && e && !e.nodeType && e, a = n && n.exports === r ? yt.Buffer : void 0, i = (a ? a.isBuffer : void 0) || Ph;
  e.exports = i;
}), Lh = Function.prototype, Ih = Object.prototype, ju = Lh.toString, Nh = Ih.hasOwnProperty, Mh = ju.call(Object), jh = function(e) {
  if (!Wt(e) || cn(e) != "[object Object]")
    return !1;
  var t = mo(e);
  if (t === null)
    return !0;
  var r = Nh.call(t, "constructor") && t.constructor;
  return typeof r == "function" && r instanceof r && ju.call(r) == Mh;
}, ie = {};
ie["[object Float32Array]"] = ie["[object Float64Array]"] = ie["[object Int8Array]"] = ie["[object Int16Array]"] = ie["[object Int32Array]"] = ie["[object Uint8Array]"] = ie["[object Uint8ClampedArray]"] = ie["[object Uint16Array]"] = ie["[object Uint32Array]"] = !0, ie["[object Arguments]"] = ie["[object Array]"] = ie["[object ArrayBuffer]"] = ie["[object Boolean]"] = ie["[object DataView]"] = ie["[object Date]"] = ie["[object Error]"] = ie["[object Function]"] = ie["[object Map]"] = ie["[object Number]"] = ie["[object Object]"] = ie["[object RegExp]"] = ie["[object Set]"] = ie["[object String]"] = ie["[object WeakMap]"] = !1;
var Dh = function(e) {
  return Wt(e) && Mu(e.length) && !!ie[cn(e)];
}, yo = function(e) {
  return function(t) {
    return e(t);
  };
}, Ln = rt(function(e, t) {
  var r = t && !t.nodeType && t, n = r && e && !e.nodeType && e, a = n && n.exports === r && Cu.process, i = function() {
    try {
      var o = n && n.require && n.require("util").types;
      return o || a && a.binding && a.binding("util");
    } catch {
    }
  }();
  e.exports = i;
}), oc = Ln && Ln.isTypedArray, Du = oc ? yo(oc) : Dh, Ni = function(e, t) {
  if ((t !== "constructor" || typeof e[t] != "function") && t != "__proto__")
    return e[t];
}, Fh = Object.prototype.hasOwnProperty, Fu = function(e, t, r) {
  var n = e[t];
  Fh.call(e, t) && ua(n, r) && (r !== void 0 || t in e) || go(e, t, r);
}, mr = function(e, t, r, n) {
  var a = !r;
  r || (r = {});
  for (var i = -1, o = t.length; ++i < o; ) {
    var s = t[i], c = n ? n(r[s], e[s], s, r, e) : void 0;
    c === void 0 && (c = e[s]), a ? go(r, s, c) : Fu(r, s, c);
  }
  return r;
}, Bh = function(e, t) {
  for (var r = -1, n = Array(e); ++r < e; )
    n[r] = t(r);
  return n;
}, Hh = /^(?:0|[1-9]\d*)$/, Bu = function(e, t) {
  var r = typeof e;
  return !!(t = t ?? 9007199254740991) && (r == "number" || r != "symbol" && Hh.test(e)) && e > -1 && e % 1 == 0 && e < t;
}, zh = Object.prototype.hasOwnProperty, Hu = function(e, t) {
  var r = Pn(e), n = !r && Ii(e), a = !r && !n && vo(e), i = !r && !n && !a && Du(e), o = r || n || a || i, s = o ? Bh(e.length, String) : [], c = s.length;
  for (var u in e)
    !t && !zh.call(e, u) || o && (u == "length" || a && (u == "offset" || u == "parent") || i && (u == "buffer" || u == "byteLength" || u == "byteOffset") || Bu(u, c)) || s.push(u);
  return s;
}, Uh = function(e) {
  var t = [];
  if (e != null)
    for (var r in Object(e))
      t.push(r);
  return t;
}, Wh = Object.prototype.hasOwnProperty, qh = function(e) {
  if (!Ut(e))
    return Uh(e);
  var t = bo(e), r = [];
  for (var n in e)
    (n != "constructor" || !t && Wh.call(e, n)) && r.push(n);
  return r;
}, br = function(e) {
  return ga(e) ? Hu(e, !0) : qh(e);
}, Gh = function(e) {
  return mr(e, br(e));
}, Kh = function(e, t, r, n, a, i, o) {
  var s = Ni(e, r), c = Ni(t, r), u = o.get(c);
  if (u)
    Li(e, r, u);
  else {
    var l = i ? i(s, c, r + "", e, t, o) : void 0, f = l === void 0;
    if (f) {
      var p = Pn(c), d = !p && vo(c), g = !p && !d && Du(c);
      l = c, p || d || g ? Pn(s) ? l = s : Oh(s) ? l = Pu(s) : d ? (f = !1, l = Ru(c, !0)) : g ? (f = !1, l = Ou(c, !0)) : l = [] : jh(c) || Ii(c) ? (l = s, Ii(s) ? l = Gh(s) : Ut(s) && !po(s) || (l = Iu(c))) : f = !1;
    }
    f && (o.set(c, l), a(l, c, n, i, o), o.delete(c)), Li(e, r, l);
  }
}, Zh = function e(t, r, n, a, i) {
  t !== r && xh(r, function(o, s) {
    if (i || (i = new $u()), Ut(o))
      Kh(t, r, s, n, e, a, i);
    else {
      var c = a ? a(Ni(t, s), o, s + "", t, r, i) : void 0;
      c === void 0 && (c = o), Li(t, s, c);
    }
  }, br);
}, zu = function(e) {
  return e;
}, Yh = function(e, t, r) {
  switch (r.length) {
    case 0:
      return e.call(t);
    case 1:
      return e.call(t, r[0]);
    case 2:
      return e.call(t, r[0], r[1]);
    case 3:
      return e.call(t, r[0], r[1], r[2]);
  }
  return e.apply(t, r);
}, sc = Math.max, Xh = function(e, t, r) {
  return t = sc(t === void 0 ? e.length - 1 : t, 0), function() {
    for (var n = arguments, a = -1, i = sc(n.length - t, 0), o = Array(i); ++a < i; )
      o[a] = n[t + a];
    a = -1;
    for (var s = Array(t + 1); ++a < t; )
      s[a] = n[a];
    return s[t] = r(o), Yh(e, this, s);
  };
}, Vh = function(e) {
  return function() {
    return e;
  };
}, Jh = ea ? function(e, t) {
  return ea(e, "toString", { configurable: !0, enumerable: !1, value: Vh(t), writable: !0 });
} : zu, Qh = Date.now, e1 = /* @__PURE__ */ function(e) {
  var t = 0, r = 0;
  return function() {
    var n = Qh(), a = 16 - (n - r);
    if (r = n, a > 0) {
      if (++t >= 800)
        return arguments[0];
    } else
      t = 0;
    return e.apply(void 0, arguments);
  };
}(Jh), t1 = function(e, t) {
  return e1(Xh(e, t, zu), e + "");
}, n1 = function(e, t, r) {
  if (!Ut(r))
    return !1;
  var n = typeof t;
  return !!(n == "number" ? ga(r) && Bu(t, r.length) : n == "string" && t in r) && ua(r[t], e);
}, r1 = function(e) {
  return t1(function(t, r) {
    var n = -1, a = r.length, i = a > 1 ? r[a - 1] : void 0, o = a > 2 ? r[2] : void 0;
    for (i = e.length > 3 && typeof i == "function" ? (a--, i) : void 0, o && n1(r[0], r[1], o) && (i = a < 3 ? void 0 : i, a = 1), t = Object(t); ++n < a; ) {
      var s = r[n];
      s && e(t, s, n, i);
    }
    return t;
  });
}, a1 = r1(function(e, t, r, n) {
  Zh(e, t, r, n);
}), i1 = a1, nt = iu;
L({ target: "Function", proto: !0, forced: Function.bind !== xi }, { bind: xi });
var Ie = function(e) {
  return X[e + "Prototype"];
}, o1 = Ie("Function").bind, Qa = Function.prototype, Uu = function(e) {
  var t = e.bind;
  return e === Qa || de(Qa, e) && t === Qa.bind ? o1 : t;
}, je = Uu, s1 = Ie("Array").concat, ei = Array.prototype, b = function(e) {
  var t = e.concat;
  return e === ei || de(ei, e) && t === ei.concat ? s1 : t;
}, ha = function(e, t) {
  var r = [][e];
  return !!r && J(function() {
    r.call(null, t || function() {
      return 1;
    }, 1);
  });
}, c1 = sn.forEach, cc = ha("forEach") ? [].forEach : function(e) {
  return c1(this, e, arguments.length > 1 ? arguments[1] : void 0);
};
L({ target: "Array", proto: !0, forced: [].forEach != cc }, { forEach: cc });
var l1 = Ie("Array").forEach, ti = Array.prototype, u1 = { DOMTokenList: !0, NodeList: !0 }, q = function(e) {
  var t = e.forEach;
  return e === ti || de(ti, e) && t === ti.forEach || F(u1, Dn(e)) ? l1 : t;
}, f1 = J(function() {
  pr(1);
});
L({ target: "Object", stat: !0, forced: f1 }, { keys: function(e) {
  return pr(lt(e));
} });
var ue = X.Object.keys, d1 = sn.filter, p1 = gr("filter");
L({ target: "Array", proto: !0, forced: !p1 }, { filter: function(e) {
  return d1(this, e, arguments.length > 1 ? arguments[1] : void 0);
} });
var g1 = Ie("Array").filter, ni = Array.prototype, Xe = function(e) {
  var t = e.filter;
  return e === ni || de(ni, e) && t === ni.filter ? g1 : t;
}, h1 = sn.findIndex, lc = "findIndex", uc = !0;
lc in [] && Array(1)[lc](function() {
  uc = !1;
}), L({ target: "Array", proto: !0, forced: uc }, { findIndex: function(e) {
  return h1(this, e, arguments.length > 1 ? arguments[1] : void 0);
} });
var m1 = Ie("Array").findIndex, ri = Array.prototype, Mi = function(e) {
  var t = e.findIndex;
  return e === ri || de(ri, e) && t === ri.findIndex ? m1 : t;
}, b1 = gr("splice"), v1 = O.TypeError, y1 = Math.max, _1 = Math.min;
L({ target: "Array", proto: !0, forced: !b1 }, { splice: function(e, t) {
  var r, n, a, i, o, s, c = lt(this), u = bt(c), l = At(e, u), f = arguments.length;
  if (f === 0 ? r = n = 0 : f === 1 ? (r = 0, n = u - l) : (r = f - 2, n = _1(y1(dr(t), 0), u - l)), u + r - n > 9007199254740991)
    throw v1("Maximum allowed length exceeded");
  for (a = co(c, n), i = 0; i < n; i++)
    (o = l + i) in c && Ft(a, i, c[o]);
  if (a.length = n, r < n) {
    for (i = l; i < u - n; i++)
      s = i + r, (o = i + n) in c ? c[s] = c[o] : delete c[s];
    for (i = u; i > u - n + r; i--)
      delete c[i - 1];
  } else if (r > n)
    for (i = u - n; i > l; i--)
      s = i + r - 1, (o = i + n - 1) in c ? c[s] = c[o] : delete c[s];
  for (i = 0; i < r; i++)
    c[i + l] = arguments[i + 2];
  return c.length = u - n + r, a;
} });
var k1 = Ie("Array").splice, ai = Array.prototype, w1 = function(e) {
  var t = e.splice;
  return e === ai || de(ai, e) && t === ai.splice ? k1 : t;
}, fc = !1, In = { SEN: "sentence", PAR: "paragraph", DEFAULT: "sentence" }, pe = function() {
  function e(t) {
    j(this, e), R(this, "$engine", void 0), R(this, "$locale", void 0), this.RULE = this.rule(t);
  }
  return M(e, [{ key: "getType", value: function() {
    return this.constructor.HOOK_TYPE || In.DEFAULT;
  } }, { key: "getName", value: function() {
    return this.constructor.HOOK_NAME;
  } }, { key: "afterInit", value: function(t) {
    typeof t == "function" && t();
  } }, { key: "setLocale", value: function(t) {
    this.$locale = t;
  } }, { key: "beforeMakeHtml", value: function(t) {
    return t;
  } }, { key: "makeHtml", value: function(t) {
    return t;
  } }, { key: "afterMakeHtml", value: function(t) {
    return t;
  } }, { key: "onKeyDown", value: function(t, r) {
  } }, { key: "getOnKeyDown", value: function() {
    return this.onKeyDown || !1;
  } }, { key: "getAttributesTest", value: function() {
    return /^(color|fontSize|font-size|id|title|class|target|underline|line-through|overline|sub|super)$/;
  } }, { key: "$testAttributes", value: function(t, r) {
    this.getAttributesTest().test(t) && r();
  } }, { key: "getAttributes", value: function(t) {
    return { attrs: {}, str: t };
  } }, { key: "test", value: function(t) {
    return !!this.RULE.reg && this.RULE.reg.test(t);
  } }, { key: "rule", value: function(t) {
    return { begin: "", end: "", content: "", reg: new RegExp("") };
  } }, { key: "mounted", value: function() {
  } }], [{ key: "getMathJaxConfig", value: function() {
    return fc;
  } }, { key: "setMathJaxConfig", value: function(t) {
    fc = t;
  } }]), e;
}();
R(pe, "HOOK_NAME", "default"), R(pe, "HOOK_TYPE", In.DEFAULT);
var E1 = sn.map, S1 = gr("map");
L({ target: "Array", proto: !0, forced: !S1 }, { map: function(e) {
  return E1(this, e, arguments.length > 1 ? arguments[1] : void 0);
} });
var Vn, A1 = Ie("Array").map, ii = Array.prototype, ve = function(e) {
  var t = e.map;
  return e === ii || de(ii, e) && t === ii.map ? A1 : t;
}, Nn = `	
\v\f\r                　\u2028\u2029\uFEFF`, dc = z("".replace), ta = "[" + Nn + "]", x1 = RegExp("^" + ta + ta + "*"), C1 = RegExp(ta + ta + "*$"), oi = function(e) {
  return function(t) {
    var r = tt(an(t));
    return 1 & e && (r = dc(r, x1, "")), 2 & e && (r = dc(r, C1, "")), r;
  };
}, _o = { start: oi(1), end: oi(2), trim: oi(3) }, T1 = _u.PROPER, $1 = _o.trim;
L({ target: "String", proto: !0, forced: (Vn = "trim", J(function() {
  return !!Nn[Vn]() || "​᠎"[Vn]() !== "​᠎" || T1 && Nn[Vn].name !== Vn;
})) }, { trim: function() {
  return $1(this);
} });
var R1 = Ie("String").trim, si = String.prototype, N = function(e) {
  var t = e.trim;
  return typeof e == "string" || e === si || de(si, e) && t === si.trim ? R1 : t;
}, O1 = _o.trim, Jn = O.parseInt, pc = O.Symbol, gc = pc && pc.iterator, Wu = /^[+-]?0x/i, P1 = z(Wu.exec), hc = Jn(Nn + "08") !== 8 || Jn(Nn + "0x16") !== 22 || gc && !J(function() {
  Jn(Object(gc));
}) ? function(e, t) {
  var r = O1(tt(e));
  return Jn(r, t >>> 0 || (P1(Wu, r) ? 16 : 10));
} : Jn;
L({ global: !0, forced: parseInt != hc }, { parseInt: hc });
var xn = X.parseInt;
function nn(e, t) {
  var r, n, a, i = arguments.length > 2 && arguments[2] !== void 0 && arguments[2];
  return /^\n/.test(e) ? i ? ((r = (n = e.match(/^\n+/g)) === null || n === void 0 || (a = n[0]) === null || a === void 0 ? void 0 : a.length) !== null && r !== void 0 ? r : 0) > 1 ? `

`.concat(t) : `
`.concat(t) : `

`.concat(t) : t;
}
function ji(e, t) {
  var r = (e.match(/\n/g) || []).length;
  return e !== "" && (r -= 2), r + t;
}
L({ target: "Array", stat: !0 }, { isArray: en });
var qu = X.Array.isArray, cr = qu;
function L1(e, t) {
  if (cr(t))
    return t;
}
function Gu(e) {
  return typeof localStorage < "u" && localStorage.getItem("cherry-".concat(e)) !== null;
}
function Ku() {
  var e = "false";
  return typeof localStorage < "u" && (e = localStorage.getItem("cherry-classicBr")), e === "true";
}
var xe = X.Object.getOwnPropertySymbols, Zu = on.f, I1 = J(function() {
  Zu(1);
});
L({ target: "Object", stat: !0, forced: !re || I1, sham: !re }, { getOwnPropertyDescriptor: function(e, t) {
  return Zu(Ve(e), t);
} });
var N1 = rt(function(e) {
  var t = X.Object, r = e.exports = function(n, a) {
    return t.getOwnPropertyDescriptor(n, a);
  };
  t.getOwnPropertyDescriptor.sham && (r.sham = !0);
}), Yu = N1, Te = Yu, M1 = z([].concat), j1 = ze("Reflect", "ownKeys") || function(e) {
  var t = ar.f(K(e)), r = ca.f;
  return r ? M1(t, r(e)) : t;
};
L({ target: "Object", stat: !0, sham: !re }, { getOwnPropertyDescriptors: function(e) {
  for (var t, r, n = Ve(e), a = on.f, i = j1(n), o = {}, s = 0; i.length > s; )
    (r = a(n, t = i[s++])) !== void 0 && Ft(o, t, r);
  return o;
} });
var Ce = X.Object.getOwnPropertyDescriptors, mc = oo.f;
L({ target: "Object", stat: !0, forced: Object.defineProperties !== mc, sham: !re }, { defineProperties: mc });
var D1 = rt(function(e) {
  var t = X.Object, r = e.exports = function(n, a) {
    return t.defineProperties(n, a);
  };
  t.defineProperties.sham && (r.sham = !0);
}), Ct = D1, F1 = O.RangeError, bc = String.fromCharCode, vc = String.fromCodePoint, B1 = z([].join), H1 = !!vc && vc.length != 1;
L({ target: "String", stat: !0, arity: 1, forced: H1 }, { fromCodePoint: function(e) {
  for (var t, r = [], n = arguments.length, a = 0; n > a; ) {
    if (t = +arguments[a++], At(t, 1114111) !== t)
      throw F1(t + " is not a valid code point");
    r[a] = t < 65536 ? bc(t) : bc(55296 + ((t -= 65536) >> 10), t % 1024 + 56320);
  }
  return B1(r, "");
} });
var Di = X.String.fromCodePoint, z1 = io.indexOf, Fi = z([].indexOf), yc = !!Fi && 1 / Fi([1], 1, -0) < 0, U1 = ha("indexOf");
L({ target: "Array", proto: !0, forced: yc || !U1 }, { indexOf: function(e) {
  var t = arguments.length > 1 ? arguments[1] : void 0;
  return yc ? Fi(this, e, t) || 0 : z1(this, e, t);
} });
var _c, kc, W1 = Ie("Array").indexOf, ci = Array.prototype, Xu = function(e) {
  var t = e.indexOf;
  return e === ci || de(ci, e) && t === ci.indexOf ? W1 : t;
}, Me = Xu;
function wc(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function vn(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = wc(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = wc(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
var Ec = { "<": "&lt;", ">": "&gt;", "&": "&amp;", '"': "&quot;", "'": "&#x27;" }, q1 = { lt: "<", gt: ">", amp: "&", quot: '"', apos: "'" }, ko = vn(vn(vn(vn(vn(vn({}, { 34: "&quot;", 38: "&amp;", 39: "&apos;", 60: "&lt;", 62: "&gt;" }), { 192: "&Agrave;", 193: "&Aacute;", 194: "&Acirc;", 195: "&Atilde;", 196: "&Auml;", 197: "&Aring;", 198: "&AElig;", 199: "&Ccedil;", 200: "&Egrave;", 201: "&Eacute;", 202: "&Ecirc;", 203: "&Euml;", 204: "&Igrave;", 205: "&Iacute;", 206: "&Icirc;", 207: "&Iuml;", 208: "&ETH;", 209: "&Ntilde;", 210: "&Ograve;", 211: "&Oacute;", 212: "&Ocirc;", 213: "&Otilde;", 214: "&Ouml;", 216: "&Oslash;", 217: "&Ugrave;", 218: "&Uacute;", 219: "&Ucirc;", 220: "&Uuml;", 221: "&Yacute;", 222: "&THORN;", 223: "&szlig;", 224: "&agrave;", 225: "&aacute;", 226: "&acirc;", 227: "&atilde;", 228: "&auml;", 229: "&aring;", 230: "&aelig;", 231: "&ccedil;", 232: "&egrave;", 233: "&eacute;", 234: "&ecirc;", 235: "&euml;", 236: "&igrave;", 237: "&iacute;", 238: "&icirc;", 239: "&iuml;", 240: "&eth;", 241: "&ntilde;", 242: "&ograve;", 243: "&oacute;", 244: "&ocirc;", 245: "&otilde;", 246: "&ouml;", 248: "&oslash;", 249: "&ugrave;", 250: "&uacute;", 251: "&ucirc;", 252: "&uuml;", 253: "&yacute;", 254: "&thorn;", 255: "&yuml;" }), { 160: "&nbsp;", 161: "&iexcl;", 162: "&cent;", 163: "&pound;", 164: "&curren;", 165: "&yen;", 166: "&brvbar;", 167: "&sect;", 168: "&uml;", 169: "&copy;", 170: "&ordf;", 171: "&laquo;", 172: "&not;", 173: "&shy;", 174: "&reg;", 175: "&macr;", 176: "&deg;", 177: "&plusmn;", 178: "&sup2;", 179: "&sup3;", 180: "&acute;", 181: "&micro;", 182: "&para;", 184: "&cedil;", 185: "&sup1;", 186: "&ordm;", 187: "&raquo;", 188: "&frac14;", 189: "&frac12;", 190: "&frac34;", 191: "&iquest;", 215: "&times;", 247: "&divide;" }), { 8704: "&forall;", 8706: "&part;", 8707: "&exist;", 8709: "&empty;", 8711: "&nabla;", 8712: "&isin;", 8713: "&notin;", 8715: "&ni;", 8719: "&prod;", 8721: "&sum;", 8722: "&minus;", 8727: "&lowast;", 8730: "&radic;", 8733: "&prop;", 8734: "&infin;", 8736: "&ang;", 8743: "&and;", 8744: "&or;", 8745: "&cap;", 8746: "&cup;", 8747: "&int;", 8756: "&there4;", 8764: "&sim;", 8773: "&cong;", 8776: "&asymp;", 8800: "&ne;", 8801: "&equiv;", 8804: "&le;", 8805: "&ge;", 8834: "&sub;", 8835: "&sup;", 8836: "&nsub;", 8838: "&sube;", 8839: "&supe;", 8853: "&oplus;", 8855: "&otimes;", 8869: "&perp;", 8901: "&sdot;" }), { 913: "&Alpha;", 914: "&Beta;", 915: "&Gamma;", 916: "&Delta;", 917: "&Epsilon;", 918: "&Zeta;", 919: "&Eta;", 920: "&Theta;", 921: "&Iota;", 922: "&Kappa;", 923: "&Lambda;", 924: "&Mu;", 925: "&Nu;", 926: "&Xi;", 927: "&Omicron;", 928: "&Pi;", 929: "&Rho;", 931: "&Sigma;", 932: "&Tau;", 933: "&Upsilon;", 934: "&Phi;", 935: "&Chi;", 936: "&Psi;", 937: "&Omega;", 945: "&alpha;", 946: "&beta;", 947: "&gamma;", 948: "&delta;", 949: "&epsilon;", 950: "&zeta;", 951: "&eta;", 952: "&theta;", 953: "&iota;", 954: "&kappa;", 955: "&lambda;", 956: "&mu;", 957: "&nu;", 958: "&xi;", 959: "&omicron;", 960: "&pi;", 961: "&rho;", 962: "&sigmaf;", 963: "&sigma;", 964: "&tau;", 965: "&upsilon;", 966: "&phi;", 967: "&chi;", 968: "&psi;", 969: "&omega;", 977: "&thetasym;", 978: "&upsih;", 982: "&piv;" }), { 338: "&OElig;", 339: "&oelig;", 352: "&Scaron;", 353: "&scaron;", 376: "&Yuml;", 402: "&fnof;", 710: "&circ;", 732: "&tilde;", 8194: "&ensp;", 8195: "&emsp;", 8201: "&thinsp;", 8204: "&zwnj;", 8205: "&zwj;", 8206: "&lrm;", 8207: "&rlm;", 8211: "&ndash;", 8212: "&mdash;", 8216: "&lsquo;", 8217: "&rsquo;", 8218: "&sbquo;", 8220: "&ldquo;", 8221: "&rdquo;", 8222: "&bdquo;", 8224: "&dagger;", 8225: "&Dagger;", 8226: "&bull;", 8230: "&hellip;", 8240: "&permil;", 8242: "&prime;", 8243: "&Prime;", 8249: "&lsaquo;", 8250: "&rsaquo;", 8254: "&oline;", 8364: "&euro;", 8482: "&trade;", 8592: "&larr;", 8593: "&uarr;", 8594: "&rarr;", 8595: "&darr;", 8596: "&harr;", 8629: "&crarr;", 8968: "&lceil;", 8969: "&rceil;", 8970: "&lfloor;", 8971: "&rfloor;", 9674: "&loz;", 9824: "&spades;", 9827: "&clubs;", 9829: "&hearts;", 9830: "&diams;" }), Sc = ue(ko), Ac = ve(Sc).call(Sc, function(e) {
  return ko[e].replace(/^&(\w+);$/g, function(t, r) {
    return r.toLowerCase();
  });
}), Bi = function(e) {
  return typeof e != "string" || e.length <= 0;
}, xc = function(e) {
  try {
    var t = Di(e);
    return !Bi(t);
  } catch {
    return !1;
  }
}, vr = ["h1|h2|h3|h4|h5|h6", "ul|ol|li|dd|dl|dt", "table|thead|tbody|tfoot|col|colgroup|th|td|tr", "div|article|section|footer|aside|details|summary|code|audio|video|canvas|figure", "address|center|cite|p|pre|blockquote|marquee|caption|figcaption|track|source|output|svg"].join("|"), G1 = ["span|a|link|b|s|i|del|u|em|strong|sup|sub|kbd", "nav|font|bdi|samp|map|area|small|time|bdo|var|wbr|meter|dfn", "ruby|rt|rp|mark|q|progress|input|textarea|select|ins"].join("|"), K1 = new RegExp(b(_c = b(kc = "^(".concat(vr, "|")).call(kc, G1, "|")).call(_c, "br|img|hr", ")( |$|/)"), "i");
function Mn(e, t) {
  return typeof e != "string" ? "" : t ? e.replace(/[<>&]/g, function(r) {
    return Ec[r] || r;
  }) : e.replace(/[<>&"']/g, function(r) {
    return Ec[r] || r;
  });
}
function qe(e, t) {
  if (typeof e != "string")
    return "";
  var r = Vu(e);
  return r = function(n) {
    return typeof n != "string" ? "" : n.replace(/&(\w+);?/g, function(a, i) {
      return q1[i] || a;
    });
  }(r), Mn(r, t);
}
function Vu(e) {
  return e.replace(/&#(\d+);?/g, function(t, r) {
    return ko[r] || t;
  });
}
function Z1(e) {
  var t = function(a) {
    return a.replace(/&#x([0-9a-f]+);?/gi, function(i, o) {
      var s = xn("0x".concat(o), 16);
      try {
        return Di(s);
      } catch {
        return i;
      }
    });
  }(function(a) {
    return a.replace(/&#(\d+);?/g, function(i, o) {
      try {
        return Di(o);
      } catch {
        return i;
      }
    });
  }(e)).match(/^\s*([\w\W]+?)(?=:)/i);
  if (!t)
    return !0;
  var r = ["javascript", "data"], n = t[1].replace(/[\s]/g, "");
  return Me(r).call(r, n.toLowerCase()) === -1;
}
function Nt(e) {
  return encodeURI(e).replace(/%25/g, "%");
}
function Y1(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Cc = 0, se = function(e) {
  U(r, pe);
  var t = Y1(r);
  function r() {
    var n, a = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { needCache: !1 }, i = a.needCache, o = a.defaultCache, s = o === void 0 ? {} : o;
    return j(this, r), (n = t.call(this, {})).needCache = !!i, n.sign = "", i && (n.cache = s || {}, n.cacheKey = "~~C".concat(Cc), Cc += 1), n;
  }
  return M(r, [{ key: "initBrReg", value: function() {
    var n = arguments.length > 0 && arguments[0] !== void 0 && arguments[0];
    this.classicBr = Gu("classicBr") ? Ku() : n, this.removeBrAfterBlock = null, this.removeBrBeforeBlock = null, this.removeNewlinesBetweenTags = null;
  } }, { key: "$cleanParagraph", value: function(n) {
    var a = n.replace(/^\n+/, "").replace(/\n+$/, "");
    return this.classicBr ? a : this.joinRawHtml(a).replace(/\n/g, "<br>").replace(/\r/g, `
`);
  } }, { key: "joinRawHtml", value: function(n) {
    if (!this.removeBrAfterBlock) {
      var a, i, o, s, c = (a = (i = this.$engine.htmlWhiteListAppend) === null || i === void 0 ? void 0 : i.split("|")) !== null && a !== void 0 ? a : [];
      c = Xe(o = ve(c).call(c, function(l) {
        return /[a-z-]+/gi.test(l) ? l : null;
      })).call(o, function(l) {
        return l !== null;
      });
      var u = b(c).call(c, vr).join("|");
      this.removeBrAfterBlock = new RegExp("<(".concat(u, ")(>| [^>]*?>)[^\\S\\n]*?\\n"), "ig"), this.removeBrBeforeBlock = new RegExp("\\n[^\\S\\n]*?<\\/(".concat(u, ")>[^\\S\\n]*?\\n"), "ig"), this.removeNewlinesBetweenTags = new RegExp(b(s = "<\\/(".concat(u, ")>[^\\S\\n]*?\\n([^\\S\\n]*?)<(")).call(s, u, ")(>| [^>]*?>)"), "ig");
    }
    return n.replace(this.removeBrAfterBlock, "<$1$2").replace(this.removeBrBeforeBlock, "</$1>").replace(this.removeNewlinesBetweenTags, "</$1>\r$2<$3$4");
  } }, { key: "toHtml", value: function(n, a) {
    return n;
  } }, { key: "makeHtml", value: function(n, a) {
    return a(n).html;
  } }, { key: "afterMakeHtml", value: function(n) {
    return this.restoreCache(n);
  } }, { key: "isContainsCache", value: function(n, a) {
    if (a) {
      var i = /^(\s*~~C\d+I\w+\$\s*)+$/g.test(n), o = new RegExp("~~C\\d+I".concat(r.IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX, "\\w+\\$"), "g").test(n);
      return i && !o;
    }
    return new RegExp("~~C\\d+I(?!".concat(r.IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX, ")\\w+\\$"), "g").test(n);
  } }, { key: "$splitHtmlByCache", value: function(n) {
    var a = new RegExp("\\n*~~C\\d+I(?!".concat(r.IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX, ")\\w+\\$\\n?"), "g");
    return { caches: n.match(a), contents: n.split(a) };
  } }, { key: "makeExcludingCached", value: function(n, a) {
    for (var i = this.$splitHtmlByCache(n), o = i.caches, s = i.contents, c = ve(s).call(s, a), u = "", l = 0; l < c.length; l++) {
      var f;
      u += c[l], o && o[l] && (u += N(f = o[l]).call(f));
    }
    return u;
  } }, { key: "getCacheWithSpace", value: function(n, a) {
    var i, o, s, c, u, l, f = arguments.length > 2 && arguments[2] !== void 0 && arguments[2], p = (i = (o = a.match(/^\n+/)) === null || o === void 0 ? void 0 : o[0]) !== null && i !== void 0 ? i : "", d = (s = (c = a.match(/\n+$/)) === null || c === void 0 ? void 0 : c[0]) !== null && s !== void 0 ? s : "";
    return f ? nn(a, n) : b(u = b(l = "".concat(p)).call(l, n)).call(u, d);
  } }, { key: "getLineCount", value: function(n) {
    var a, i, o, s = n, c = (a = (i = (arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "").match(/^\n+/g)) === null || i === void 0 || (o = i[0]) === null || o === void 0 ? void 0 : o.length) !== null && a !== void 0 ? a : 0;
    c = c === 1 ? 1 : 0, s = s.replace(/^\n+/g, "");
    var u = new RegExp(`
*~~C\\d+I(?:`.concat(r.IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX, ")?\\w+?_L(\\d+)\\$"), "g"), l = 0;
    return s = s.replace(u, function(f, p) {
      return l += xn(p, 10), f.replace(/^\n+/g, "");
    }), c + l + (s.match(/\n/g) || []).length + 1;
  } }, { key: "pushCache", value: function(n) {
    var a, i, o = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "", s = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : 0;
    if (this.needCache) {
      var c = o || this.$engine.md5(n);
      return this.cache[c] = { content: n, using: !0 }, b(a = b(i = "".concat(this.cacheKey, "I")).call(i, c, "_L")).call(a, s, "$");
    }
  } }, { key: "popCache", value: function(n) {
    if (this.needCache)
      return this.cache[n].content || "";
  } }, { key: "resetCache", value: function() {
    if (this.needCache) {
      for (var n = 0, a = ue(this.cache); n < a.length; n++) {
        var i = a[n];
        this.cache[i].using || delete this.cache[i];
      }
      for (var o = 0, s = ue(this.cache); o < s.length; o++) {
        var c = s[o];
        this.cache[c].using = !1;
      }
    }
  } }, { key: "restoreCache", value: function(n) {
    var a, i = this;
    if (!this.needCache)
      return n;
    var o = new RegExp(b(a = "".concat(this.cacheKey, "I((?:")).call(a, r.IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX, ")?\\w+)\\$"), "g"), s = n.replace(o, function(c, u) {
      return i.popCache(u.replace(/_L\d+$/, ""));
    });
    return this.resetCache(), s;
  } }, { key: "checkCache", value: function(n, a) {
    var i, o, s = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : 0;
    return this.sign = this.$engine.md5(n), this.cache[this.sign] ? (this.cache[this.sign].using = !0, b(i = b(o = "".concat(this.cacheKey, "I")).call(o, this.sign, "_L")).call(i, s, "$")) : this.toHtml(n, a);
  } }, { key: "mounted", value: function() {
  } }, { key: "signWithCache", value: function(n) {
    return !1;
  } }]), r;
}();
R(se, "HOOK_TYPE", In.PAR), R(se, "IN_PARAGRAPH_CACHE_KEY_PREFIX", "!"), R(se, "IN_PARAGRAPH_CACHE_KEY_PREFIX_REGEX", "\\!");
var Tc = J(function() {
  if (typeof ArrayBuffer == "function") {
    var e = new ArrayBuffer(8);
    Object.isExtensible(e) && Object.defineProperty(e, "a", { value: 8 });
  }
}), Pr = Object.isExtensible, li = J(function() {
  Pr(1);
}) || Tc ? function(e) {
  return !!ye(e) && (!Tc || ht(e) != "ArrayBuffer") && (!Pr || Pr(e));
} : Pr, X1 = !J(function() {
  return Object.isExtensible(Object.preventExtensions({}));
}), An = rt(function(e) {
  var t = Je.f, r = !1, n = nr("meta"), a = 0, i = function(s) {
    t(s, n, { value: { objectID: "O" + a++, weakData: {} } });
  }, o = e.exports = { enable: function() {
    o.enable = function() {
    }, r = !0;
    var s = ar.f, c = z([].splice), u = {};
    u[n] = 1, s(u).length && (ar.f = function(l) {
      for (var f = s(l), p = 0, d = f.length; p < d; p++)
        if (f[p] === n) {
          c(f, p, 1);
          break;
        }
      return f;
    }, L({ target: "Object", stat: !0, forced: !0 }, { getOwnPropertyNames: lo.f }));
  }, fastKey: function(s, c) {
    if (!ye(s))
      return typeof s == "symbol" ? s : (typeof s == "string" ? "S" : "P") + s;
    if (!F(s, n)) {
      if (!li(s))
        return "F";
      if (!c)
        return "E";
      i(s);
    }
    return s[n].objectID;
  }, getWeakData: function(s, c) {
    if (!F(s, n)) {
      if (!li(s))
        return !0;
      if (!c)
        return !1;
      i(s);
    }
    return s[n].weakData;
  }, onFreeze: function(s) {
    return X1 && r && li(s) && !F(s, n) && i(s), s;
  } };
  Fn[n] = !0;
});
An.enable, An.fastKey, An.getWeakData, An.onFreeze;
var V1 = _e("iterator"), J1 = Array.prototype, Ju = function(e) {
  return e !== void 0 && (Et.Array === e || J1[V1] === e);
}, Q1 = _e("iterator"), ma = function(e) {
  if (e != null)
    return Kr(e, Q1) || Kr(e, "@@iterator") || Et[Dn(e)];
}, em = O.TypeError, ba = function(e, t) {
  var r = arguments.length < 2 ? ma(e) : t;
  if (ne(r))
    return K(Z(r, e));
  throw em(ur(e) + " is not iterable");
}, Hi = function(e, t, r) {
  var n, a;
  K(e);
  try {
    if (!(n = Kr(e, "return"))) {
      if (t === "throw")
        throw r;
      return r;
    }
    n = Z(n, e);
  } catch (i) {
    a = !0, n = i;
  }
  if (t === "throw")
    throw r;
  if (a)
    throw n;
  return K(n), r;
}, tm = O.TypeError, Wr = function(e, t) {
  this.stopped = e, this.result = t;
}, $c = Wr.prototype, Pe = function(e, t, r) {
  var n, a, i, o, s, c, u, l = r && r.that, f = !(!r || !r.AS_ENTRIES), p = !(!r || !r.IS_ITERATOR), d = !(!r || !r.INTERRUPTED), g = Ze(t, l), _ = function(h) {
    return n && Hi(n, "normal", h), new Wr(!0, h);
  }, m = function(h) {
    return f ? (K(h), d ? g(h[0], h[1], _) : g(h[0], h[1])) : d ? g(h, _) : g(h);
  };
  if (p)
    n = e;
  else {
    if (!(a = ma(e)))
      throw tm(ur(e) + " is not iterable");
    if (Ju(a)) {
      for (i = 0, o = bt(e); o > i; i++)
        if ((s = m(e[i])) && de($c, s))
          return s;
      return new Wr(!1);
    }
    n = ba(e, a);
  }
  for (c = n.next; !(u = Z(c, n)).done; ) {
    try {
      s = m(u.value);
    } catch (h) {
      Hi(n, "throw", h);
    }
    if (typeof s == "object" && s && de($c, s))
      return s;
  }
  return new Wr(!1);
}, nm = O.TypeError, Qu = function(e, t) {
  if (de(t, e))
    return e;
  throw nm("Incorrect invocation");
}, rm = Je.f, am = sn.forEach, im = vt.set, om = vt.getterFor, Rc = function(e, t, r) {
  for (var n in t)
    r && r.unsafe && e[n] ? e[n] = t[n] : tn(e, n, t[n], r);
  return e;
}, Oc = _e("species"), sm = Je.f, Pc = An.fastKey, Lc = vt.set, ui = vt.getterFor, cm = { getConstructor: function(e, t, r, n) {
  var a = e(function(u, l) {
    Qu(u, i), Lc(u, { type: t, index: mt(null), first: void 0, last: void 0, size: 0 }), re || (u.size = 0), l != null && Pe(l, u[n], { that: u, AS_ENTRIES: r });
  }), i = a.prototype, o = ui(t), s = function(u, l, f) {
    var p, d, g = o(u), _ = c(u, l);
    return _ ? _.value = f : (g.last = _ = { index: d = Pc(l, !0), key: l, value: f, previous: p = g.last, next: void 0, removed: !1 }, g.first || (g.first = _), p && (p.next = _), re ? g.size++ : u.size++, d !== "F" && (g.index[d] = _)), u;
  }, c = function(u, l) {
    var f, p = o(u), d = Pc(l);
    if (d !== "F")
      return p.index[d];
    for (f = p.first; f; f = f.next)
      if (f.key == l)
        return f;
  };
  return Rc(i, { clear: function() {
    for (var u = o(this), l = u.index, f = u.first; f; )
      f.removed = !0, f.previous && (f.previous = f.previous.next = void 0), delete l[f.index], f = f.next;
    u.first = u.last = void 0, re ? u.size = 0 : this.size = 0;
  }, delete: function(u) {
    var l = this, f = o(l), p = c(l, u);
    if (p) {
      var d = p.next, g = p.previous;
      delete f.index[p.index], p.removed = !0, g && (g.next = d), d && (d.previous = g), f.first == p && (f.first = d), f.last == p && (f.last = g), re ? f.size-- : l.size--;
    }
    return !!p;
  }, forEach: function(u) {
    for (var l, f = o(this), p = Ze(u, arguments.length > 1 ? arguments[1] : void 0); l = l ? l.next : f.first; )
      for (p(l.value, l.key, this); l && l.removed; )
        l = l.previous;
  }, has: function(u) {
    return !!c(this, u);
  } }), Rc(i, r ? { get: function(u) {
    var l = c(this, u);
    return l && l.value;
  }, set: function(u, l) {
    return s(this, u === 0 ? 0 : u, l);
  } } : { add: function(u) {
    return s(this, u = u === 0 ? 0 : u, u);
  } }), re && sm(i, "size", { get: function() {
    return o(this).size;
  } }), a;
}, setStrong: function(e, t, r) {
  var n = t + " Iterator", a = ui(t), i = ui(n);
  fo(e, t, function(o, s) {
    Lc(this, { type: n, target: o, state: a(o), kind: s, last: void 0 });
  }, function() {
    for (var o = i(this), s = o.kind, c = o.last; c && c.removed; )
      c = c.previous;
    return o.target && (o.last = c = c ? c.next : o.state.first) ? s == "keys" ? { value: c.key, done: !1 } : s == "values" ? { value: c.value, done: !1 } : { value: [c.key, c.value], done: !1 } : (o.target = void 0, { value: void 0, done: !0 });
  }, r ? "entries" : "values", !r, !0), function(o) {
    var s = ze(o), c = Je.f;
    re && s && !s[Oc] && c(s, Oc, { configurable: !0, get: function() {
      return this;
    } });
  }(t);
} };
(function(e, t, r) {
  var n, a = e.indexOf("Map") !== -1, i = e.indexOf("Weak") !== -1, o = a ? "set" : "add", s = O[e], c = s && s.prototype, u = {};
  if (re && oe(s) && (i || c.forEach && !J(function() {
    new s().entries().next();
  }))) {
    var l = (n = t(function(p, d) {
      im(Qu(p, l), { type: e, collection: new s() }), d != null && Pe(d, p[o], { that: p, AS_ENTRIES: a });
    })).prototype, f = om(e);
    am(["add", "clear", "delete", "forEach", "get", "has", "set", "keys", "values", "entries"], function(p) {
      var d = p == "add" || p == "set";
      !(p in c) || i && p == "clear" || ot(l, p, function(g, _) {
        var m = f(this).collection;
        if (!d && i && !ye(g))
          return p == "get" && void 0;
        var h = m[p](g === 0 ? 0 : g, _);
        return d ? this : h;
      });
    }), i || rm(l, "size", { configurable: !0, get: function() {
      return f(this).collection.size;
    } });
  } else
    n = r.getConstructor(t, e, a, o), An.enable();
  Rn(n, e, !1, !0), u[e] = n, L({ global: !0, forced: !0 }, u), i || r.setStrong(n, e, a);
})("Map", function(e) {
  return function() {
    return e(this, arguments.length ? arguments[0] : void 0);
  };
}, cm);
var lm = X.Map, Ic = [].push;
L({ target: "Map", stat: !0, forced: !0 }, { from: function(e) {
  var t, r, n, a, i = arguments.length, o = i > 1 ? arguments[1] : void 0;
  return Zr(this), (t = o !== void 0) && ne(o), e == null ? new this() : (r = [], t ? (n = 0, a = Ze(o, i > 2 ? arguments[2] : void 0), Pe(e, function(s) {
    Z(Ic, r, a(s, n++));
  })) : Pe(e, Ic, { that: r }), new this(r));
} });
L({ target: "Map", stat: !0, forced: !0 }, { of: function() {
  return new this(Qt(arguments));
} });
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { deleteAll: function() {
  for (var e, t = K(this), r = ne(t.delete), n = !0, a = 0, i = arguments.length; a < i; a++)
    e = Z(r, t, arguments[a]), n = n && e;
  return !!n;
} });
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { emplace: function(e, t) {
  var r = K(this), n = ne(r.get), a = ne(r.has), i = ne(r.set), o = Z(a, r, e) && "update" in t ? t.update(Z(n, r, e), e, r) : t.insert(e, r);
  return Z(i, r, e, o), o;
} });
var pt = ba;
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { every: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0);
  return !Pe(r, function(a, i, o) {
    if (!n(i, a, t))
      return o();
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).stopped;
} });
var um = _e("species"), zi = function(e, t) {
  var r, n = K(e).constructor;
  return n === void 0 || (r = K(n)[um]) == null ? t : Zr(r);
};
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { filter: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0), a = new (zi(t, ze("Map")))(), i = ne(a.set);
  return Pe(r, function(o, s) {
    n(s, o, t) && Z(i, a, o, s);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0 }), a;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { find: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0);
  return Pe(r, function(a, i, o) {
    if (n(i, a, t))
      return o(i);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).result;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { findKey: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0);
  return Pe(r, function(a, i, o) {
    if (n(i, a, t))
      return o(a);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).result;
} });
var fm = z([].push);
L({ target: "Map", stat: !0, forced: !0 }, { groupBy: function(e, t) {
  ne(t);
  var r = ba(e), n = new this(), a = ne(n.has), i = ne(n.get), o = ne(n.set);
  return Pe(r, function(s) {
    var c = t(s);
    Z(a, n, c) ? fm(Z(i, n, c), s) : Z(o, n, c, [s]);
  }, { IS_ITERATOR: !0 }), n;
} });
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { includes: function(e) {
  return Pe(pt(K(this)), function(t, r, n) {
    if ((a = r) === (i = e) || a != a && i != i)
      return n();
    var a, i;
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).stopped;
} }), L({ target: "Map", stat: !0, forced: !0 }, { keyBy: function(e, t) {
  var r = new this();
  ne(t);
  var n = ne(r.set);
  return Pe(e, function(a) {
    Z(n, r, t(a), a);
  }), r;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { keyOf: function(e) {
  return Pe(pt(K(this)), function(t, r, n) {
    if (r === e)
      return n(t);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).result;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { mapKeys: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0), a = new (zi(t, ze("Map")))(), i = ne(a.set);
  return Pe(r, function(o, s) {
    Z(i, a, n(s, o, t), s);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0 }), a;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { mapValues: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0), a = new (zi(t, ze("Map")))(), i = ne(a.set);
  return Pe(r, function(o, s) {
    Z(i, a, o, n(s, o, t));
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0 }), a;
} }), L({ target: "Map", proto: !0, real: !0, arity: 1, forced: !0 }, { merge: function(e) {
  for (var t = K(this), r = ne(t.set), n = arguments.length, a = 0; a < n; )
    Pe(arguments[a++], r, { that: t, AS_ENTRIES: !0 });
  return t;
} });
var dm = O.TypeError;
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { reduce: function(e) {
  var t = K(this), r = pt(t), n = arguments.length < 2, a = n ? void 0 : arguments[1];
  if (ne(e), Pe(r, function(i, o) {
    n ? (n = !1, a = o) : a = e(a, o, i, t);
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0 }), n)
    throw dm("Reduce of empty map with no initial value");
  return a;
} }), L({ target: "Map", proto: !0, real: !0, forced: !0 }, { some: function(e) {
  var t = K(this), r = pt(t), n = Ze(e, arguments.length > 1 ? arguments[1] : void 0);
  return Pe(r, function(a, i, o) {
    if (n(i, a, t))
      return o();
  }, { AS_ENTRIES: !0, IS_ITERATOR: !0, INTERRUPTED: !0 }).stopped;
} });
var pm = O.TypeError;
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { update: function(e, t) {
  var r = K(this), n = ne(r.get), a = ne(r.has), i = ne(r.set), o = arguments.length;
  ne(t);
  var s = Z(a, r, e);
  if (!s && o < 3)
    throw pm("Updating absent value");
  var c = s ? Z(n, r, e) : ne(o > 2 ? arguments[2] : void 0)(e, r);
  return Z(i, r, e, t(c, e, r)), r;
} });
var gm = O.TypeError, Nc = function(e, t) {
  var r, n = K(this), a = ne(n.get), i = ne(n.has), o = ne(n.set), s = arguments.length > 2 ? arguments[2] : void 0;
  if (!oe(t) && !oe(s))
    throw gm("At least one callback required");
  return Z(i, n, e) ? (r = Z(a, n, e), oe(t) && (r = t(r), Z(o, n, e, r))) : oe(s) && (r = s(), Z(o, n, e, r)), r;
};
L({ target: "Map", proto: !0, real: !0, forced: !0 }, { upsert: Nc }), L({ target: "Map", proto: !0, real: !0, name: "upsert", forced: !0 }, { updateOrInsert: Nc });
var Mc = lm, hm = Xu, Lr = au, mm = Uu;
function Ui(e, t, r) {
  return Ui = function() {
    if (typeof Reflect > "u" || !Lr || Lr.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(Lr(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }() ? Lr : function(n, a, i) {
    var o = [null];
    o.push.apply(o, a);
    var s = new (mm(Function).apply(n, o))();
    return i && rr(s, i.prototype), s;
  }, Ui.apply(null, arguments);
}
function Wi(e) {
  var t = typeof Mc == "function" ? new Mc() : void 0;
  return Wi = function(r) {
    if (r === null || !function(a) {
      var i;
      return hm(i = Function.toString.call(a)).call(i, "[native code]") !== -1;
    }(r))
      return r;
    if (typeof r != "function")
      throw new TypeError("Super expression must either be null or a function");
    if (t !== void 0) {
      if (t.has(r))
        return t.get(r);
      t.set(r, n);
    }
    function n() {
      return Ui(r, arguments, T(this).constructor);
    }
    return n.prototype = su(r.prototype, { constructor: { value: n, enumerable: !1, writable: !0, configurable: !0 } }), rr(n, r);
  }, Wi(e);
}
function bm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var ef = function(e, t) {
  if (!cr(e) && He(e) !== t.name.toLowerCase() || !cr(e) && t.name.toLowerCase() === "array")
    throw new TypeError("parameter given must be ".concat(t.name));
  return !0;
}, jc = function(e, t) {
  if (!(e instanceof t))
    throw new Error("the hook does not correctly inherit");
  return !0;
}, vm = function(e) {
  if (He(e) !== "object")
    throw new Error("the hook must be a instance, not a class");
  return !0;
}, ym = function(e) {
  U(r, Wi(Error));
  var t = bm(r);
  function r(n, a) {
    var i;
    return j(this, r), (i = t.call(this, n)).name = "Error", i.stack = i.buildStackTrace(a), i;
  }
  return M(r, [{ key: "buildStackTrace", value: function(n) {
    var a, i = n && n.stack ? n.stack : "";
    return b(a = "".concat(this.stack, `
Caused By: `)).call(a, i);
  } }]), r;
}(), rn = new Proxy({}, { get: function(e, t, r) {
  return function() {
  };
} });
function Dc(e, t, r) {
  var n, a;
  if (e === -1)
    rn.warn(b(n = b(a = "Duplicate hook name [".concat(t.HOOK_NAME, "] found, hook [")).call(a, t.toString(), "] ")).call(n, isNaN(r) ? "" : "at index [".concat(r, "] "), "will not take effect."));
  else if (e === -2) {
    var i;
    rn.warn(b(i = "Hook [".concat(t.toString(), "] ")).call(i, isNaN(r) ? "" : "at index [".concat(r, "] "), "is not a valid hook, and will not take effect."));
  }
}
function qr(e) {
  return qi(e) || tf(e);
}
function qi(e) {
  return Object.prototype.isPrototypeOf.call(pe, e);
}
function tf(e) {
  return Object.prototype.isPrototypeOf.call(se, e);
}
function fi(e) {
  return qr(e) && (e == null ? void 0 : e.Cherry$$CUSTOM) === !0;
}
var _m = function() {
  function e(t, r, n) {
    j(this, e), this.$locale = n.locale, this.hookList = {}, this.hookNameList = {}, ef(t, Array), this.registerInternalHooks(t, r), this.registerCustomHooks(r.engine.customSyntax, r);
  }
  return M(e, [{ key: "registerInternalHooks", value: function(t, r) {
    var n = this;
    q(t).call(t, function(a, i) {
      Dc(n.register(a, r), a, i);
    });
  } }, { key: "registerCustomHooks", value: function(t, r) {
    var n = this;
    if (t) {
      var a = ue(t);
      q(a).call(a, function(i) {
        var o, s, c, u, l = {}, f = t[i];
        if (qi(f))
          s = f;
        else {
          if (!qi(u = (c = f) == null ? void 0 : c.syntaxClass) && !tf(u))
            return;
          s = f.syntaxClass, l.force = !!f.force, f.before ? l.before = f.before : f.after && (l.after = f.after);
        }
        qr(s) ? (nt(s, "Cherry$$CUSTOM", { enumerable: !1, configurable: !1, writable: !1, value: !0 }), o = n.register(s, r, l)) : o = -2, Dc(o, s, void 0);
      });
    }
  } }, { key: "getHookList", value: function() {
    return this.hookList;
  } }, { key: "getHookNameList", value: function() {
    return this.hookNameList;
  } }, { key: "register", value: function(t, r, n) {
    var a, i, o = this, s = r.externals, c = r.engine, u = c.syntax;
    if (qr(t)) {
      i = t.HOOK_NAME;
      var l = (u == null ? void 0 : u[i]) || {};
      (a = new t({ externals: s, config: l, globalConfig: c.global })).afterInit(function() {
        a.setLocale(o.$locale);
      });
    } else {
      if (typeof t != "function" || !(a = t(r)) || !qr(a.constructor))
        return -2;
      i = a.getName();
    }
    if (u[i] !== !1 || fi(t)) {
      var f = a.getType();
      if (this.hookNameList[i]) {
        var p;
        if (!fi(t) || !n.force)
          return -1;
        var d = this.hookNameList[i].type;
        this.hookList[d] = Xe(p = this.hookList[d]).call(p, function(E) {
          return E.getName() !== i;
        });
      }
      if (this.hookNameList[i] = { type: f }, this.hookList[f] = this.hookList[f] || [], fi(t)) {
        var g, _, m, h = -1;
        if (n.before)
          (h = Mi(g = this.hookList[f]).call(g, function(E) {
            return E.getName() === n.before;
          })) === -1 && rn.warn(b(_ = "Cannot find hook named [".concat(n.before, `],
            custom hook [`)).call(_, i, "] will append to the end of the hooks."));
        else if (n.after) {
          var v, w;
          (h = Mi(v = this.hookList[f]).call(v, function(E) {
            return E.getName() === n.after;
          })) === -1 ? rn.warn(b(w = "Cannot find hook named [".concat(n.after, `],
              custom hook [`)).call(w, i, "] will append to the end of the hooks.")) : h += 1;
        }
        h < 0 || h >= this.hookList[f].length ? this.hookList[f].push(a) : w1(m = this.hookList[f]).call(m, h, 0, a);
      } else
        this.hookList[f].push(a);
    }
  } }]), e;
}();
function Fc(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Bc(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = Fc(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = Fc(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function Le(e, t, r) {
  var n = e.begin + e.content + e.end;
  return r && (n = n.replace(/\[\\h\]/g, Hc).replace(/\\h/g, Hc)), new RegExp(n, t || "g");
}
function fe() {
  try {
    return new RegExp("(?<=.)"), !0;
  } catch {
  }
  return !1;
}
var Hc = "[ \\t\\u00a0]", nf = "(?:[^\\n]*?\\S[^\\n]*?)", rf = "[\\u0021-\\u002F\\u003a-\\u0040\\u005b-\\u0060\\u007b-\\u007e]", zc = "[\\u0021-\\u002F\\u003a-\\u0040\\u005b\\u005d\\u005e\\u0060\\u007b-\\u007e \\t\\n！“”¥‘’（），。—：；《》？【】「」·～｜]", Gi = new RegExp([/[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+/.source, "@", /[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*/.source].join("")), di = new RegExp("^".concat(Gi.source, "$")), wo = new RegExp('(?:\\S+(?::\\S*)?@)?(?:(?:1\\d\\d|2[01]\\d|22[0-3]|[1-9]\\d?)(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:1\\d\\d|2[0-4]\\d|25[0-4]|[1-9]\\d?))|(?![-_])(?:[-\\w\\xa1-\\xff]{0,63}[^-_]\\.)+(?:[a-zA-Z\\xa1-\\xff]{2,}\\.?))(?::\\d{2,5})?(?:[/?#][^\\s<>\\x00-\\x1f"\\(\\)]*)?'), af = new RegExp("(?:\\/\\/)".concat(wo.source)), Uc = new RegExp("^".concat(wo.source, "$")), Wc = new RegExp("^".concat(af.source, "$"));
function Eo() {
  var e, t = arguments.length > 0 && arguments[0] !== void 0 && arguments[0], r = { begin: "(?:^|\\n)(\\n*)", content: ["(\\h*\\|[^\\n]+\\|?\\h*)", "\\n", "(?:(?:\\h*\\|\\h*:?[-]{1,}:?\\h*)+\\|?\\h*)", "((\\n\\h*\\|[^\\n]+\\|?\\h*)*)"].join(""), end: "(?=$|\\n)" };
  r.reg = Le(r, "g", !0);
  var n = { begin: "(?:^|\\n)(\\n*)", content: ["(\\|?[^\\n|]+(\\|[^\\n|]+)+\\|?)", "\\n", "(?:\\|?\\h*:?[-]{1,}:?[\\h]*(?:\\|[\\h]*:?[-]{1,}:?\\h*)+\\|?)", "((\\n\\|?([^\\n|]+(\\|[^\\n|]*)+)\\|?)*)"].join(""), end: "(?=$|\\n)" };
  return n.reg = Le(n, "g", !0), t === !1 ? { strict: r, loose: n } : Le({ begin: "", content: b(e = "(?:".concat(r.begin + r.content + r.end, "|")).call(e, n.begin + n.content + n.end, ")"), end: "" }, "g", !0);
}
var of = qu;
function Ki(e, t) {
  (t == null || t > e.length) && (t = e.length);
  for (var r = 0, n = new Array(t); r < t; r++)
    n[r] = e[r];
  return n;
}
var So = ma, km = function(e, t, r, n) {
  try {
    return n ? t(K(r)[0], r[1]) : t(r);
  } catch (a) {
    Hi(e, "throw", a);
  }
}, qc = O.Array, sf = _e("iterator"), cf = !1;
try {
  var wm = 0, Gc = { next: function() {
    return { done: !!wm++ };
  }, return: function() {
    cf = !0;
  } };
  Gc[sf] = function() {
    return this;
  }, Array.from(Gc, function() {
    throw 2;
  });
} catch {
}
var Em = !function(e, t) {
  if (!t && !cf)
    return !1;
  var r = !1;
  try {
    var n = {};
    n[sf] = function() {
      return { next: function() {
        return { done: r = !0 };
      } };
    }, e(n);
  } catch {
  }
  return r;
}(function(e) {
  Array.from(e);
});
L({ target: "Array", stat: !0, forced: Em }, { from: function(e) {
  var t = lt(e), r = ia(this), n = arguments.length, a = n > 1 ? arguments[1] : void 0, i = a !== void 0;
  i && (a = Ze(a, n > 2 ? arguments[2] : void 0));
  var o, s, c, u, l, f, p = ma(t), d = 0;
  if (!p || this == qc && Ju(p))
    for (o = bt(t), s = r ? new this(o) : qc(o); o > d; d++)
      f = i ? a(t[d], d) : t[d], Ft(s, d, f);
  else
    for (l = (u = ba(t, p)).next, s = r ? new this() : []; !(c = Z(l, u)).done; d++)
      f = i ? km(u, a, [c.value, d], !0) : c.value, Ft(s, d, f);
  return s.length = d, s;
} });
var lf = X.Array.from, uf = lf;
function ff(e) {
  if (Sn !== void 0 && So(e) != null || e["@@iterator"] != null)
    return uf(e);
}
var Sm = gr("slice"), Am = _e("species"), pi = O.Array, xm = Math.max;
L({ target: "Array", proto: !0, forced: !Sm }, { slice: function(e, t) {
  var r, n, a, i = Ve(this), o = bt(i), s = At(e, o), c = At(t === void 0 ? o : t, o);
  if (en(i) && (r = i.constructor, (ia(r) && (r === pi || en(r.prototype)) || ye(r) && (r = r[Am]) === null) && (r = void 0), r === pi || r === void 0))
    return Qt(i, s, c);
  for (n = new (r === void 0 ? pi : r)(xm(c - s, 0)), a = 0; s < c; s++, a++)
    s in i && Ft(n, a, i[s]);
  return n.length = a, n;
} });
var Cm = Ie("Array").slice, gi = Array.prototype, df = function(e) {
  var t = e.slice;
  return e === gi || de(gi, e) && t === gi.slice ? Cm : t;
}, Tm = df;
function Ao(e, t) {
  var r;
  if (e) {
    if (typeof e == "string")
      return Ki(e, t);
    var n = Tm(r = Object.prototype.toString.call(e)).call(r, 8, -1);
    return n === "Object" && e.constructor && (n = e.constructor.name), n === "Map" || n === "Set" ? uf(e) : n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n) ? Ki(e, t) : void 0;
  }
}
function lr(e) {
  return function(t) {
    if (of(t))
      return Ki(t);
  }(e) || ff(e) || Ao(e) || function() {
    throw new TypeError(`Invalid attempt to spread non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`);
  }();
}
function pf(e) {
  if (of(e))
    return e;
}
function gf() {
  throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`);
}
function hf(e) {
  return pf(e) || ff(e) || Ao(e) || gf();
}
var Ke = df;
function Kc(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Zc(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = Kc(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = Kc(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function ct(e, t, r) {
  var n, a = arguments.length > 3 && arguments[3] !== void 0 && arguments[3], i = arguments.length > 4 && arguments[4] !== void 0 ? arguments[4] : 1;
  if (!t)
    return e;
  t.lastIndex = 0;
  for (var o = 0, s = []; (n = t.exec(e)) !== null; ) {
    var c = { begin: n.index, length: n[0].length };
    if (a && n.index === o - i) {
      var u, l = hf(n), f = l[0], p = Ke(l).call(l, 2);
      s.push({ begin: c.begin + i, length: c.length - i, replacedText: r.apply(void 0, b(u = [Ke(f).call(f, i), ""]).call(u, lr(p))) });
    } else
      s.push(Zc(Zc({}, c), {}, { replacedText: r.apply(void 0, lr(n)) }));
    o = t.lastIndex, t.lastIndex -= i;
  }
  return t.lastIndex = 0, function(d, g) {
    if (!g.length)
      return d;
    var _ = [], m = 0;
    return q(g).call(g, function(h, v) {
      _.push(Ke(d).call(d, m, h.begin)), _.push(h.replacedText), m = h.begin + h.length, v === g.length - 1 && _.push(Ke(d).call(d, m));
    }), _.join("");
  }(e, s);
}
function $m(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var mf = function(e) {
  U(r, pe);
  var t = $m(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i, o) {
    var s, c;
    return b(s = b(c = "".concat(a, '<span style="color:')).call(c, i, '">')).call(s, o, "</span>");
  } }, { key: "makeHtml", value: function(n) {
    return fe() ? n.replace(this.RULE.reg, this.toHtml) : ct(n, this.RULE.reg, this.toHtml, !0, 1);
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))!!" : "(^|[^\\\\])!!", end: "!!", content: "(#[0-9a-zA-Z]{3,6}|[a-z]{3,20})[\\s]([\\w\\W]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Rm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(mf, "HOOK_NAME", "fontColor");
var bf = function(e) {
  U(r, pe);
  var t = Rm(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i, o) {
    var s, c;
    return b(s = b(c = "".concat(a, '<span style="background-color:')).call(c, i, '">')).call(s, o, "</span>");
  } }, { key: "makeHtml", value: function(n) {
    return fe() ? n.replace(this.RULE.reg, this.toHtml) : ct(n, this.RULE.reg, this.toHtml, !0, 1);
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))!!!" : "(^|[^\\\\])!!!", end: "!!!", content: "(#[0-9a-zA-Z]{3,6}|[a-z]{3,10})[\\s]([\\w\\W]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Om(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(bf, "HOOK_NAME", "bgColor");
var vf = function(e) {
  U(r, pe);
  var t = Om(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i, o) {
    var s, c;
    return b(s = b(c = "".concat(a, '<span style="font-size:')).call(c, i, 'px;line-height:1em;">')).call(s, o, "</span>");
  } }, { key: "makeHtml", value: function(n) {
    return this.test(n) ? fe() ? n.replace(this.RULE.reg, this.toHtml) : ct(n, this.RULE.reg, this.toHtml, !0, 1) : n;
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))!" : "(^|[^\\\\])!", end: "!", content: "([0-9]{1,2})[\\s]([\\w\\W]*?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Yc(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Ir(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = Yc(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = Yc(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function Pm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(vf, "HOOK_NAME", "fontSize");
var yf = function(e) {
  U(r, pe);
  var t = Pm(r);
  function r() {
    var n, a = (arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0 }).config;
    return j(this, r), n = t.call(this, { config: a }), a ? (n.needWhitespace = !!a.needWhitespace, n) : B(n);
  }
  return M(r, [{ key: "makeHtml", value: function(n) {
    return this.test(n) ? n.replace(this.RULE.reg, "$1<del>$2</del>") : n;
  } }, { key: "rule", value: function() {
    var n = {};
    return (n = (arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0 }).config.needWhitespace ? Ir(Ir({}, n), {}, { begin: "(^|[\\s])\\~T\\~T", end: "\\~T\\~T(?=\\s|$)", content: "([\\w\\W]+?)" }) : Ir(Ir({}, n), {}, { begin: "(^|[^\\\\])\\~T\\~T", end: "\\~T\\~T", content: "([\\w\\W]+?)" })).reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Lm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(yf, "HOOK_NAME", "strikethrough");
var _f = function(e) {
  U(r, pe);
  var t = Lm(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i) {
    var o;
    return b(o = "".concat(a, "<sup>")).call(o, i, "</sup>");
  } }, { key: "makeHtml", value: function(n) {
    return fe() ? n.replace(this.RULE.reg, this.toHtml) : ct(n, this.RULE.reg, this.toHtml, !0, 1);
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))\\^" : "(^|[^\\\\])\\^", end: "\\^", content: "([\\w\\W]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Im(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(_f, "HOOK_NAME", "sup");
var kf = function(e) {
  U(r, pe);
  var t = Im(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i) {
    var o;
    return b(o = "".concat(a, "<sub>")).call(o, i, "</sub>");
  } }, { key: "makeHtml", value: function(n) {
    return fe() ? n.replace(this.RULE.reg, this.toHtml) : ct(n, this.RULE.reg, this.toHtml, !0, 1);
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))\\^\\^" : "(^|[^\\\\])\\^\\^", end: "\\^\\^", content: "([\\w\\W]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Ht(e, t) {
  return pf(e) || function(r, n) {
    var a = r == null ? null : Sn !== void 0 && So(r) || r["@@iterator"];
    if (a != null) {
      var i, o, s = [], c = !0, u = !1;
      try {
        for (a = a.call(r); !(c = (i = a.next()).done) && (s.push(i.value), !n || s.length !== n); c = !0)
          ;
      } catch (l) {
        u = !0, o = l;
      } finally {
        try {
          c || a.return == null || a.return();
        } finally {
          if (u)
            throw o;
        }
      }
      return s;
    }
  }(e, t) || Ao(e, t) || gf();
}
R(kf, "HOOK_NAME", "sub");
var Gr = rt(function(e) {
  var t = function(r) {
    var n = /(?:^|\s)lang(?:uage)?-([\w-]+)(?=\s|$)/i, a = 0, i = {}, o = { manual: r.Prism && r.Prism.manual, disableWorkerMessageHandler: r.Prism && r.Prism.disableWorkerMessageHandler, util: { encode: function m(h) {
      return h instanceof s ? new s(h.type, m(h.content), h.alias) : Array.isArray(h) ? h.map(m) : h.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/\u00a0/g, " ");
    }, type: function(m) {
      return Object.prototype.toString.call(m).slice(8, -1);
    }, objId: function(m) {
      return m.__id || Object.defineProperty(m, "__id", { value: ++a }), m.__id;
    }, clone: function m(h, v) {
      var w, E;
      switch (v = v || {}, o.util.type(h)) {
        case "Object":
          if (E = o.util.objId(h), v[E])
            return v[E];
          for (var S in w = {}, v[E] = w, h)
            h.hasOwnProperty(S) && (w[S] = m(h[S], v));
          return w;
        case "Array":
          return E = o.util.objId(h), v[E] ? v[E] : (w = [], v[E] = w, h.forEach(function(A, x) {
            w[x] = m(A, v);
          }), w);
        default:
          return h;
      }
    }, getLanguage: function(m) {
      for (; m; ) {
        var h = n.exec(m.className);
        if (h)
          return h[1].toLowerCase();
        m = m.parentElement;
      }
      return "none";
    }, setLanguage: function(m, h) {
      m.className = m.className.replace(RegExp(n, "gi"), ""), m.classList.add("language-" + h);
    }, currentScript: function() {
      if (typeof document > "u")
        return null;
      if ("currentScript" in document)
        return document.currentScript;
      try {
        throw new Error();
      } catch (w) {
        var m = (/at [^(\r\n]*\((.*):[^:]+:[^:]+\)$/i.exec(w.stack) || [])[1];
        if (m) {
          var h = document.getElementsByTagName("script");
          for (var v in h)
            if (h[v].src == m)
              return h[v];
        }
        return null;
      }
    }, isActive: function(m, h, v) {
      for (var w = "no-" + h; m; ) {
        var E = m.classList;
        if (E.contains(h))
          return !0;
        if (E.contains(w))
          return !1;
        m = m.parentElement;
      }
      return !!v;
    } }, languages: { plain: i, plaintext: i, text: i, txt: i, extend: function(m, h) {
      var v = o.util.clone(o.languages[m]);
      for (var w in h)
        v[w] = h[w];
      return v;
    }, insertBefore: function(m, h, v, w) {
      var E = (w = w || o.languages)[m], S = {};
      for (var A in E)
        if (E.hasOwnProperty(A)) {
          if (A == h)
            for (var x in v)
              v.hasOwnProperty(x) && (S[x] = v[x]);
          v.hasOwnProperty(A) || (S[A] = E[A]);
        }
      var D = w[m];
      return w[m] = S, o.languages.DFS(o.languages, function(C, W) {
        W === D && C != m && (this[C] = S);
      }), S;
    }, DFS: function m(h, v, w, E) {
      E = E || {};
      var S = o.util.objId;
      for (var A in h)
        if (h.hasOwnProperty(A)) {
          v.call(h, A, h[A], w || A);
          var x = h[A], D = o.util.type(x);
          D !== "Object" || E[S(x)] ? D !== "Array" || E[S(x)] || (E[S(x)] = !0, m(x, v, A, E)) : (E[S(x)] = !0, m(x, v, null, E));
        }
    } }, plugins: {}, highlightAll: function(m, h) {
      o.highlightAllUnder(document, m, h);
    }, highlightAllUnder: function(m, h, v) {
      var w = { callback: v, container: m, selector: 'code[class*="language-"], [class*="language-"] code, code[class*="lang-"], [class*="lang-"] code' };
      o.hooks.run("before-highlightall", w), w.elements = Array.prototype.slice.apply(w.container.querySelectorAll(w.selector)), o.hooks.run("before-all-elements-highlight", w);
      for (var E, S = 0; E = w.elements[S++]; )
        o.highlightElement(E, h === !0, w.callback);
    }, highlightElement: function(m, h, v) {
      var w = o.util.getLanguage(m), E = o.languages[w];
      o.util.setLanguage(m, w);
      var S = m.parentElement;
      S && S.nodeName.toLowerCase() === "pre" && o.util.setLanguage(S, w);
      var A = { element: m, language: w, grammar: E, code: m.textContent };
      function x(C) {
        A.highlightedCode = C, o.hooks.run("before-insert", A), A.element.innerHTML = A.highlightedCode, o.hooks.run("after-highlight", A), o.hooks.run("complete", A), v && v.call(A.element);
      }
      if (o.hooks.run("before-sanity-check", A), (S = A.element.parentElement) && S.nodeName.toLowerCase() === "pre" && !S.hasAttribute("tabindex") && S.setAttribute("tabindex", "0"), !A.code)
        return o.hooks.run("complete", A), void (v && v.call(A.element));
      if (o.hooks.run("before-highlight", A), A.grammar)
        if (h && r.Worker) {
          var D = new Worker(o.filename);
          D.onmessage = function(C) {
            x(C.data);
          }, D.postMessage(JSON.stringify({ language: A.language, code: A.code, immediateClose: !0 }));
        } else
          x(o.highlight(A.code, A.grammar, A.language));
      else
        x(o.util.encode(A.code));
    }, highlight: function(m, h, v) {
      var w = { code: m, grammar: h, language: v };
      if (o.hooks.run("before-tokenize", w), !w.grammar)
        throw new Error('The language "' + w.language + '" has no grammar.');
      return w.tokens = o.tokenize(w.code, w.grammar), o.hooks.run("after-tokenize", w), s.stringify(o.util.encode(w.tokens), w.language);
    }, tokenize: function(m, h) {
      var v = h.rest;
      if (v) {
        for (var w in v)
          h[w] = v[w];
        delete h.rest;
      }
      var E = new l();
      return f(E, E.head, m), u(m, E, h, E.head, 0), function(S) {
        for (var A = [], x = S.head.next; x !== S.tail; )
          A.push(x.value), x = x.next;
        return A;
      }(E);
    }, hooks: { all: {}, add: function(m, h) {
      var v = o.hooks.all;
      v[m] = v[m] || [], v[m].push(h);
    }, run: function(m, h) {
      var v = o.hooks.all[m];
      if (v && v.length)
        for (var w, E = 0; w = v[E++]; )
          w(h);
    } }, Token: s };
    function s(m, h, v, w) {
      this.type = m, this.content = h, this.alias = v, this.length = 0 | (w || "").length;
    }
    function c(m, h, v, w) {
      m.lastIndex = h;
      var E = m.exec(v);
      if (E && w && E[1]) {
        var S = E[1].length;
        E.index += S, E[0] = E[0].slice(S);
      }
      return E;
    }
    function u(m, h, v, w, E, S) {
      for (var A in v)
        if (v.hasOwnProperty(A) && v[A]) {
          var x = v[A];
          x = Array.isArray(x) ? x : [x];
          for (var D = 0; D < x.length; ++D) {
            if (S && S.cause == A + "," + D)
              return;
            var C = x[D], W = C.inside, V = !!C.lookbehind, G = !!C.greedy, we = C.alias;
            if (G && !C.pattern.global) {
              var ge = C.pattern.toString().match(/[imsuy]*$/)[0];
              C.pattern = RegExp(C.pattern.source, ge + "g");
            }
            for (var Ne = C.pattern || C, Q = w.next, ce = E; Q !== h.tail && !(S && ce >= S.reach); ce += Q.value.length, Q = Q.next) {
              var $e = Q.value;
              if (h.length > m.length)
                return;
              if (!($e instanceof s)) {
                var he, Qe = 1;
                if (G) {
                  if (!(he = c(Ne, ce, m, V)) || he.index >= m.length)
                    break;
                  var at = he.index, Re = he.index + he[0].length, ke = ce;
                  for (ke += Q.value.length; at >= ke; )
                    ke += (Q = Q.next).value.length;
                  if (ce = ke -= Q.value.length, Q.value instanceof s)
                    continue;
                  for (var Ge = Q; Ge !== h.tail && (ke < Re || typeof Ge.value == "string"); Ge = Ge.next)
                    Qe++, ke += Ge.value.length;
                  Qe--, $e = m.slice(ce, ke), he.index -= ce;
                } else if (!(he = c(Ne, 0, $e, V)))
                  continue;
                at = he.index;
                var Tt = he[0], Bn = $e.slice(0, at), yr = $e.slice(at + Tt.length), Hn = ce + $e.length;
                S && Hn > S.reach && (S.reach = Hn);
                var qt = Q.prev;
                if (Bn && (qt = f(h, qt, Bn), ce += Bn.length), p(h, qt, Qe), Q = f(h, qt, new s(A, W ? o.tokenize(Tt, W) : Tt, we, Tt)), yr && f(h, Q, yr), Qe > 1) {
                  var zn = { cause: A + "," + D, reach: Hn };
                  u(m, h, v, Q.prev, ce, zn), S && zn.reach > S.reach && (S.reach = zn.reach);
                }
              }
            }
          }
        }
    }
    function l() {
      var m = { value: null, prev: null, next: null }, h = { value: null, prev: m, next: null };
      m.next = h, this.head = m, this.tail = h, this.length = 0;
    }
    function f(m, h, v) {
      var w = h.next, E = { value: v, prev: h, next: w };
      return h.next = E, w.prev = E, m.length++, E;
    }
    function p(m, h, v) {
      for (var w = h.next, E = 0; E < v && w !== m.tail; E++)
        w = w.next;
      h.next = w, w.prev = h, m.length -= E;
    }
    if (r.Prism = o, s.stringify = function m(h, v) {
      if (typeof h == "string")
        return h;
      if (Array.isArray(h)) {
        var w = "";
        return h.forEach(function(D) {
          w += m(D, v);
        }), w;
      }
      var E = { type: h.type, content: m(h.content, v), tag: "span", classes: ["token", h.type], attributes: {}, language: v }, S = h.alias;
      S && (Array.isArray(S) ? Array.prototype.push.apply(E.classes, S) : E.classes.push(S)), o.hooks.run("wrap", E);
      var A = "";
      for (var x in E.attributes)
        A += " " + x + '="' + (E.attributes[x] || "").replace(/"/g, "&quot;") + '"';
      return "<" + E.tag + ' class="' + E.classes.join(" ") + '"' + A + ">" + E.content + "</" + E.tag + ">";
    }, !r.document)
      return r.addEventListener && (o.disableWorkerMessageHandler || r.addEventListener("message", function(m) {
        var h = JSON.parse(m.data), v = h.language, w = h.code, E = h.immediateClose;
        r.postMessage(o.highlight(w, o.languages[v], v)), E && r.close();
      }, !1)), o;
    var d = o.util.currentScript();
    function g() {
      o.manual || o.highlightAll();
    }
    if (d && (o.filename = d.src, d.hasAttribute("data-manual") && (o.manual = !0)), !o.manual) {
      var _ = document.readyState;
      _ === "loading" || _ === "interactive" && d && d.defer ? document.addEventListener("DOMContentLoaded", g) : window.requestAnimationFrame ? window.requestAnimationFrame(g) : window.setTimeout(g, 16);
    }
    return o;
  }(typeof window < "u" ? window : typeof WorkerGlobalScope < "u" && self instanceof WorkerGlobalScope ? self : {});
  e.exports && (e.exports = t), Lt !== void 0 && (Lt.Prism = t);
});
function Xc(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Nm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
Prism.languages.clike = { comment: [{ pattern: /(^|[^\\])\/\*[\s\S]*?(?:\*\/|$)/, lookbehind: !0, greedy: !0 }, { pattern: /(^|[^\\:])\/\/.*/, lookbehind: !0, greedy: !0 }], string: { pattern: /(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1/, greedy: !0 }, "class-name": { pattern: /(\b(?:class|extends|implements|instanceof|interface|new|trait)\s+|\bcatch\s+\()[\w.\\]+/i, lookbehind: !0, inside: { punctuation: /[.\\]/ } }, keyword: /\b(?:break|catch|continue|do|else|finally|for|function|if|in|instanceof|new|null|return|throw|try|while)\b/, boolean: /\b(?:false|true)\b/, function: /\b\w+(?=\()/, number: /\b0x[\da-f]+\b|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[+-]?\d+)?/i, operator: /[<>]=?|[!=]=?=?|--?|\+\+?|&&?|\|\|?|[?*/~^%]/, punctuation: /[{}[\];(),.:]/ }, Prism.languages.c = Prism.languages.extend("clike", { comment: { pattern: /\/\/(?:[^\r\n\\]|\\(?:\r\n?|\n|(?![\r\n])))*|\/\*[\s\S]*?(?:\*\/|$)/, greedy: !0 }, string: { pattern: /"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"/, greedy: !0 }, "class-name": { pattern: /(\b(?:enum|struct)\s+(?:__attribute__\s*\(\([\s\S]*?\)\)\s*)?)\w+|\b[a-z]\w*_t\b/, lookbehind: !0 }, keyword: /\b(?:_Alignas|_Alignof|_Atomic|_Bool|_Complex|_Generic|_Imaginary|_Noreturn|_Static_assert|_Thread_local|__attribute__|asm|auto|break|case|char|const|continue|default|do|double|else|enum|extern|float|for|goto|if|inline|int|long|register|return|short|signed|sizeof|static|struct|switch|typedef|typeof|union|unsigned|void|volatile|while)\b/, function: /\b[a-z_]\w*(?=\s*\()/i, number: /(?:\b0x(?:[\da-f]+(?:\.[\da-f]*)?|\.[\da-f]+)(?:p[+-]?\d+)?|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[+-]?\d+)?)[ful]{0,4}/i, operator: />>=?|<<=?|->|([-+&|:])\1|[?:~]|[-+*/%&|^!=<>]=?/ }), Prism.languages.insertBefore("c", "string", { char: { pattern: /'(?:\\(?:\r\n|[\s\S])|[^'\\\r\n]){0,32}'/, greedy: !0 } }), Prism.languages.insertBefore("c", "string", { macro: { pattern: /(^[\t ]*)#\s*[a-z](?:[^\r\n\\/]|\/(?!\*)|\/\*(?:[^*]|\*(?!\/))*\*\/|\\(?:\r\n|[\s\S]))*/im, lookbehind: !0, greedy: !0, alias: "property", inside: { string: [{ pattern: /^(#\s*include\s*)<[^>]+>/, lookbehind: !0 }, Prism.languages.c.string], char: Prism.languages.c.char, comment: Prism.languages.c.comment, "macro-name": [{ pattern: /(^#\s*define\s+)\w+\b(?!\()/i, lookbehind: !0 }, { pattern: /(^#\s*define\s+)\w+\b(?=\()/i, lookbehind: !0, alias: "function" }], directive: { pattern: /^(#\s*)[a-z]+/, lookbehind: !0, alias: "keyword" }, "directive-hash": /^#/, punctuation: /##|\\(?=[\r\n])/, expression: { pattern: /\S[\s\S]*/, inside: Prism.languages.c } } } }), Prism.languages.insertBefore("c", "function", { constant: /\b(?:EOF|NULL|SEEK_CUR|SEEK_END|SEEK_SET|__DATE__|__FILE__|__LINE__|__TIMESTAMP__|__TIME__|__func__|stderr|stdin|stdout)\b/ }), delete Prism.languages.c.boolean, function(e) {
  function t(Re, ke) {
    return Re.replace(/<<(\d+)>>/g, function(Ge, Tt) {
      return "(?:" + ke[+Tt] + ")";
    });
  }
  function r(Re, ke, Ge) {
    return RegExp(t(Re, ke), Ge || "");
  }
  function n(Re, ke) {
    for (var Ge = 0; Ge < ke; Ge++)
      Re = Re.replace(/<<self>>/g, function() {
        return "(?:" + Re + ")";
      });
    return Re.replace(/<<self>>/g, "[^\\s\\S]");
  }
  var a = "bool byte char decimal double dynamic float int long object sbyte short string uint ulong ushort var void", i = "class enum interface record struct", o = "add alias and ascending async await by descending from(?=\\s*(?:\\w|$)) get global group into init(?=\\s*;) join let nameof not notnull on or orderby partial remove select set unmanaged value when where with(?=\\s*{)", s = "abstract as base break case catch checked const continue default delegate do else event explicit extern finally fixed for foreach goto if implicit in internal is lock namespace new null operator out override params private protected public readonly ref return sealed sizeof stackalloc static switch this throw try typeof unchecked unsafe using virtual volatile while yield";
  function c(Re) {
    return "\\b(?:" + Re.trim().replace(/ /g, "|") + ")\\b";
  }
  var u = c(i), l = RegExp(c(a + " " + i + " " + o + " " + s)), f = c(i + " " + o + " " + s), p = c(a + " " + i + " " + s), d = n(/<(?:[^<>;=+\-*/%&|^]|<<self>>)*>/.source, 2), g = n(/\((?:[^()]|<<self>>)*\)/.source, 2), _ = /@?\b[A-Za-z_]\w*\b/.source, m = t(/<<0>>(?:\s*<<1>>)?/.source, [_, d]), h = t(/(?!<<0>>)<<1>>(?:\s*\.\s*<<1>>)*/.source, [f, m]), v = /\[\s*(?:,\s*)*\]/.source, w = t(/<<0>>(?:\s*(?:\?\s*)?<<1>>)*(?:\s*\?)?/.source, [h, v]), E = t(/[^,()<>[\];=+\-*/%&|^]|<<0>>|<<1>>|<<2>>/.source, [d, g, v]), S = t(/\(<<0>>+(?:,<<0>>+)+\)/.source, [E]), A = t(/(?:<<0>>|<<1>>)(?:\s*(?:\?\s*)?<<2>>)*(?:\s*\?)?/.source, [S, h, v]), x = { keyword: l, punctuation: /[<>()?,.:[\]]/ }, D = /'(?:[^\r\n'\\]|\\.|\\[Uux][\da-fA-F]{1,8})'/.source, C = /"(?:\\.|[^\\"\r\n])*"/.source, W = /@"(?:""|\\[\s\S]|[^\\"])*"(?!")/.source;
  e.languages.csharp = e.languages.extend("clike", { string: [{ pattern: r(/(^|[^$\\])<<0>>/.source, [W]), lookbehind: !0, greedy: !0 }, { pattern: r(/(^|[^@$\\])<<0>>/.source, [C]), lookbehind: !0, greedy: !0 }], "class-name": [{ pattern: r(/(\busing\s+static\s+)<<0>>(?=\s*;)/.source, [h]), lookbehind: !0, inside: x }, { pattern: r(/(\busing\s+<<0>>\s*=\s*)<<1>>(?=\s*;)/.source, [_, A]), lookbehind: !0, inside: x }, { pattern: r(/(\busing\s+)<<0>>(?=\s*=)/.source, [_]), lookbehind: !0 }, { pattern: r(/(\b<<0>>\s+)<<1>>/.source, [u, m]), lookbehind: !0, inside: x }, { pattern: r(/(\bcatch\s*\(\s*)<<0>>/.source, [h]), lookbehind: !0, inside: x }, { pattern: r(/(\bwhere\s+)<<0>>/.source, [_]), lookbehind: !0 }, { pattern: r(/(\b(?:is(?:\s+not)?|as)\s+)<<0>>/.source, [w]), lookbehind: !0, inside: x }, { pattern: r(/\b<<0>>(?=\s+(?!<<1>>|with\s*\{)<<2>>(?:\s*[=,;:{)\]]|\s+(?:in|when)\b))/.source, [A, p, _]), inside: x }], keyword: l, number: /(?:\b0(?:x[\da-f_]*[\da-f]|b[01_]*[01])|(?:\B\.\d+(?:_+\d+)*|\b\d+(?:_+\d+)*(?:\.\d+(?:_+\d+)*)?)(?:e[-+]?\d+(?:_+\d+)*)?)(?:[dflmu]|lu|ul)?\b/i, operator: />>=?|<<=?|[-=]>|([-+&|])\1|~|\?\?=?|[-+*/%&|^!=<>]=?/, punctuation: /\?\.?|::|[{}[\];(),.:]/ }), e.languages.insertBefore("csharp", "number", { range: { pattern: /\.\./, alias: "operator" } }), e.languages.insertBefore("csharp", "punctuation", { "named-parameter": { pattern: r(/([(,]\s*)<<0>>(?=\s*:)/.source, [_]), lookbehind: !0, alias: "punctuation" } }), e.languages.insertBefore("csharp", "class-name", { namespace: { pattern: r(/(\b(?:namespace|using)\s+)<<0>>(?:\s*\.\s*<<0>>)*(?=\s*[;{])/.source, [_]), lookbehind: !0, inside: { punctuation: /\./ } }, "type-expression": { pattern: r(/(\b(?:default|sizeof|typeof)\s*\(\s*(?!\s))(?:[^()\s]|\s(?!\s)|<<0>>)*(?=\s*\))/.source, [g]), lookbehind: !0, alias: "class-name", inside: x }, "return-type": { pattern: r(/<<0>>(?=\s+(?:<<1>>\s*(?:=>|[({]|\.\s*this\s*\[)|this\s*\[))/.source, [A, h]), inside: x, alias: "class-name" }, "constructor-invocation": { pattern: r(/(\bnew\s+)<<0>>(?=\s*[[({])/.source, [A]), lookbehind: !0, inside: x, alias: "class-name" }, "generic-method": { pattern: r(/<<0>>\s*<<1>>(?=\s*\()/.source, [_, d]), inside: { function: r(/^<<0>>/.source, [_]), generic: { pattern: RegExp(d), alias: "class-name", inside: x } } }, "type-list": { pattern: r(/\b((?:<<0>>\s+<<1>>|record\s+<<1>>\s*<<5>>|where\s+<<2>>)\s*:\s*)(?:<<3>>|<<4>>|<<1>>\s*<<5>>|<<6>>)(?:\s*,\s*(?:<<3>>|<<4>>|<<6>>))*(?=\s*(?:where|[{;]|=>|$))/.source, [u, m, _, A, l.source, g, /\bnew\s*\(\s*\)/.source]), lookbehind: !0, inside: { "record-arguments": { pattern: r(/(^(?!new\s*\()<<0>>\s*)<<1>>/.source, [m, g]), lookbehind: !0, greedy: !0, inside: e.languages.csharp }, keyword: l, "class-name": { pattern: RegExp(A), greedy: !0, inside: x }, punctuation: /[,()]/ } }, preprocessor: { pattern: /(^[\t ]*)#.*/m, lookbehind: !0, alias: "property", inside: { directive: { pattern: /(#)\b(?:define|elif|else|endif|endregion|error|if|line|nullable|pragma|region|undef|warning)\b/, lookbehind: !0, alias: "keyword" } } } });
  var V = C + "|" + D, G = t(/\/(?![*/])|\/\/[^\r\n]*[\r\n]|\/\*(?:[^*]|\*(?!\/))*\*\/|<<0>>/.source, [V]), we = n(t(/[^"'/()]|<<0>>|\(<<self>>*\)/.source, [G]), 2), ge = /\b(?:assembly|event|field|method|module|param|property|return|type)\b/.source, Ne = t(/<<0>>(?:\s*\(<<1>>*\))?/.source, [h, we]);
  e.languages.insertBefore("csharp", "class-name", { attribute: { pattern: r(/((?:^|[^\s\w>)?])\s*\[\s*)(?:<<0>>\s*:\s*)?<<1>>(?:\s*,\s*<<1>>)*(?=\s*\])/.source, [ge, Ne]), lookbehind: !0, greedy: !0, inside: { target: { pattern: r(/^<<0>>(?=\s*:)/.source, [ge]), alias: "keyword" }, "attribute-arguments": { pattern: r(/\(<<0>>*\)/.source, [we]), inside: e.languages.csharp }, "class-name": { pattern: RegExp(h), inside: { punctuation: /\./ } }, punctuation: /[:,]/ } } });
  var Q = /:[^}\r\n]+/.source, ce = n(t(/[^"'/()]|<<0>>|\(<<self>>*\)/.source, [G]), 2), $e = t(/\{(?!\{)(?:(?![}:])<<0>>)*<<1>>?\}/.source, [ce, Q]), he = n(t(/[^"'/()]|\/(?!\*)|\/\*(?:[^*]|\*(?!\/))*\*\/|<<0>>|\(<<self>>*\)/.source, [V]), 2), Qe = t(/\{(?!\{)(?:(?![}:])<<0>>)*<<1>>?\}/.source, [he, Q]);
  function at(Re, ke) {
    return { interpolation: { pattern: r(/((?:^|[^{])(?:\{\{)*)<<0>>/.source, [Re]), lookbehind: !0, inside: { "format-string": { pattern: r(/(^\{(?:(?![}:])<<0>>)*)<<1>>(?=\}$)/.source, [ke, Q]), lookbehind: !0, inside: { punctuation: /^:/ } }, punctuation: /^\{|\}$/, expression: { pattern: /[\s\S]+/, alias: "language-csharp", inside: e.languages.csharp } } }, string: /[\s\S]+/ };
  }
  e.languages.insertBefore("csharp", "string", { "interpolation-string": [{ pattern: r(/(^|[^\\])(?:\$@|@\$)"(?:""|\\[\s\S]|\{\{|<<0>>|[^\\{"])*"/.source, [$e]), lookbehind: !0, greedy: !0, inside: at($e, ce) }, { pattern: r(/(^|[^@\\])\$"(?:\\.|\{\{|<<0>>|[^\\"{])*"/.source, [Qe]), lookbehind: !0, greedy: !0, inside: at(Qe, he) }], char: { pattern: RegExp(D), greedy: !0 } }), e.languages.dotnet = e.languages.cs = e.languages.csharp;
}(Prism), function(e) {
  var t = /\b(?:alignas|alignof|asm|auto|bool|break|case|catch|char|char16_t|char32_t|char8_t|class|co_await|co_return|co_yield|compl|concept|const|const_cast|consteval|constexpr|constinit|continue|decltype|default|delete|do|double|dynamic_cast|else|enum|explicit|export|extern|final|float|for|friend|goto|if|import|inline|int|int16_t|int32_t|int64_t|int8_t|long|module|mutable|namespace|new|noexcept|nullptr|operator|override|private|protected|public|register|reinterpret_cast|requires|return|short|signed|sizeof|static|static_assert|static_cast|struct|switch|template|this|thread_local|throw|try|typedef|typeid|typename|uint16_t|uint32_t|uint64_t|uint8_t|union|unsigned|using|virtual|void|volatile|wchar_t|while)\b/, r = /\b(?!<keyword>)\w+(?:\s*\.\s*\w+)*\b/.source.replace(/<keyword>/g, function() {
    return t.source;
  });
  e.languages.cpp = e.languages.extend("c", { "class-name": [{ pattern: RegExp(/(\b(?:class|concept|enum|struct|typename)\s+)(?!<keyword>)\w+/.source.replace(/<keyword>/g, function() {
    return t.source;
  })), lookbehind: !0 }, /\b[A-Z]\w*(?=\s*::\s*\w+\s*\()/, /\b[A-Z_]\w*(?=\s*::\s*~\w+\s*\()/i, /\b\w+(?=\s*<(?:[^<>]|<(?:[^<>]|<[^<>]*>)*>)*>\s*::\s*\w+\s*\()/], keyword: t, number: { pattern: /(?:\b0b[01']+|\b0x(?:[\da-f']+(?:\.[\da-f']*)?|\.[\da-f']+)(?:p[+-]?[\d']+)?|(?:\b[\d']+(?:\.[\d']*)?|\B\.[\d']+)(?:e[+-]?[\d']+)?)[ful]{0,4}/i, greedy: !0 }, operator: />>=?|<<=?|->|--|\+\+|&&|\|\||[?:~]|<=>|[-+*/%&|^!=<>]=?|\b(?:and|and_eq|bitand|bitor|not|not_eq|or|or_eq|xor|xor_eq)\b/, boolean: /\b(?:false|true)\b/ }), e.languages.insertBefore("cpp", "string", { module: { pattern: RegExp(/(\b(?:import|module)\s+)/.source + "(?:" + /"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"|<[^<>\r\n]*>/.source + "|" + /<mod-name>(?:\s*:\s*<mod-name>)?|:\s*<mod-name>/.source.replace(/<mod-name>/g, function() {
    return r;
  }) + ")"), lookbehind: !0, greedy: !0, inside: { string: /^[<"][\s\S]+/, operator: /:/, punctuation: /\./ } }, "raw-string": { pattern: /R"([^()\\ ]{0,16})\([\s\S]*?\)\1"/, alias: "string", greedy: !0 } }), e.languages.insertBefore("cpp", "keyword", { "generic-function": { pattern: /\b(?!operator\b)[a-z_]\w*\s*<(?:[^<>]|<[^<>]*>)*>(?=\s*\()/i, inside: { function: /^\w+/, generic: { pattern: /<[\s\S]+/, alias: "class-name", inside: e.languages.cpp } } } }), e.languages.insertBefore("cpp", "operator", { "double-colon": { pattern: /::/, alias: "punctuation" } }), e.languages.insertBefore("cpp", "class-name", { "base-clause": { pattern: /(\b(?:class|struct)\s+\w+\s*:\s*)[^;{}"'\s]+(?:\s+[^;{}"'\s]+)*(?=\s*[;{])/, lookbehind: !0, greedy: !0, inside: e.languages.extend("cpp", {}) } }), e.languages.insertBefore("inside", "double-colon", { "class-name": /\b[a-z_]\w*\b(?!\s*::)/i }, e.languages.cpp["base-clause"]);
}(Prism), Prism.languages.markup = { comment: { pattern: /<!--(?:(?!<!--)[\s\S])*?-->/, greedy: !0 }, prolog: { pattern: /<\?[\s\S]+?\?>/, greedy: !0 }, doctype: { pattern: /<!DOCTYPE(?:[^>"'[\]]|"[^"]*"|'[^']*')+(?:\[(?:[^<"'\]]|"[^"]*"|'[^']*'|<(?!!--)|<!--(?:[^-]|-(?!->))*-->)*\]\s*)?>/i, greedy: !0, inside: { "internal-subset": { pattern: /(^[^\[]*\[)[\s\S]+(?=\]>$)/, lookbehind: !0, greedy: !0, inside: null }, string: { pattern: /"[^"]*"|'[^']*'/, greedy: !0 }, punctuation: /^<!|>$|[[\]]/, "doctype-tag": /^DOCTYPE/i, name: /[^\s<>'"]+/ } }, cdata: { pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i, greedy: !0 }, tag: { pattern: /<\/?(?!\d)[^\s>\/=$<%]+(?:\s(?:\s*[^\s>\/=]+(?:\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))|(?=[\s/>])))+)?\s*\/?>/, greedy: !0, inside: { tag: { pattern: /^<\/?[^\s>\/]+/, inside: { punctuation: /^<\/?/, namespace: /^[^\s>\/:]+:/ } }, "special-attr": [], "attr-value": { pattern: /=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+)/, inside: { punctuation: [{ pattern: /^=/, alias: "attr-equals" }, /"|'/] } }, punctuation: /\/?>/, "attr-name": { pattern: /[^\s>\/]+/, inside: { namespace: /^[^\s>\/:]+:/ } } } }, entity: [{ pattern: /&[\da-z]{1,8};/i, alias: "named-entity" }, /&#x?[\da-f]{1,8};/i] }, Prism.languages.markup.tag.inside["attr-value"].inside.entity = Prism.languages.markup.entity, Prism.languages.markup.doctype.inside["internal-subset"].inside = Prism.languages.markup, Prism.hooks.add("wrap", function(e) {
  e.type === "entity" && (e.attributes.title = e.content.replace(/&amp;/, "&"));
}), Object.defineProperty(Prism.languages.markup.tag, "addInlined", { value: function(e, t) {
  var r = {};
  r["language-" + t] = { pattern: /(^<!\[CDATA\[)[\s\S]+?(?=\]\]>$)/i, lookbehind: !0, inside: Prism.languages[t] }, r.cdata = /^<!\[CDATA\[|\]\]>$/i;
  var n = { "included-cdata": { pattern: /<!\[CDATA\[[\s\S]*?\]\]>/i, inside: r } };
  n["language-" + t] = { pattern: /[\s\S]+/, inside: Prism.languages[t] };
  var a = {};
  a[e] = { pattern: RegExp(/(<__[^>]*>)(?:<!\[CDATA\[(?:[^\]]|\](?!\]>))*\]\]>|(?!<!\[CDATA\[)[\s\S])*?(?=<\/__>)/.source.replace(/__/g, function() {
    return e;
  }), "i"), lookbehind: !0, greedy: !0, inside: n }, Prism.languages.insertBefore("markup", "cdata", a);
} }), Object.defineProperty(Prism.languages.markup.tag, "addAttribute", { value: function(e, t) {
  Prism.languages.markup.tag.inside["special-attr"].push({ pattern: RegExp(/(^|["'\s])/.source + "(?:" + e + ")" + /\s*=\s*(?:"[^"]*"|'[^']*'|[^\s'">=]+(?=[\s>]))/.source, "i"), lookbehind: !0, inside: { "attr-name": /^[^\s=]+/, "attr-value": { pattern: /=[\s\S]+/, inside: { value: { pattern: /(^=\s*(["']|(?!["'])))\S[\s\S]*(?=\2$)/, lookbehind: !0, alias: [t, "language-" + t], inside: Prism.languages[t] }, punctuation: [{ pattern: /^=/, alias: "attr-equals" }, /"|'/] } } } });
} }), Prism.languages.html = Prism.languages.markup, Prism.languages.mathml = Prism.languages.markup, Prism.languages.svg = Prism.languages.markup, Prism.languages.xml = Prism.languages.extend("markup", {}), Prism.languages.ssml = Prism.languages.xml, Prism.languages.atom = Prism.languages.xml, Prism.languages.rss = Prism.languages.xml, function(e) {
  var t = /(?:"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"|'(?:\\(?:\r\n|[\s\S])|[^'\\\r\n])*')/;
  e.languages.css = { comment: /\/\*[\s\S]*?\*\//, atrule: { pattern: /@[\w-](?:[^;{\s]|\s+(?![\s{]))*(?:;|(?=\s*\{))/, inside: { rule: /^@[\w-]+/, "selector-function-argument": { pattern: /(\bselector\s*\(\s*(?![\s)]))(?:[^()\s]|\s+(?![\s)])|\((?:[^()]|\([^()]*\))*\))+(?=\s*\))/, lookbehind: !0, alias: "selector" }, keyword: { pattern: /(^|[^\w-])(?:and|not|only|or)(?![\w-])/, lookbehind: !0 } } }, url: { pattern: RegExp("\\burl\\((?:" + t.source + "|" + /(?:[^\\\r\n()"']|\\[\s\S])*/.source + ")\\)", "i"), greedy: !0, inside: { function: /^url/i, punctuation: /^\(|\)$/, string: { pattern: RegExp("^" + t.source + "$"), alias: "url" } } }, selector: { pattern: RegExp(`(^|[{}\\s])[^{}\\s](?:[^{};"'\\s]|\\s+(?![\\s{])|` + t.source + ")*(?=\\s*\\{)"), lookbehind: !0 }, string: { pattern: t, greedy: !0 }, property: { pattern: /(^|[^-\w\xA0-\uFFFF])(?!\s)[-_a-z\xA0-\uFFFF](?:(?!\s)[-\w\xA0-\uFFFF])*(?=\s*:)/i, lookbehind: !0 }, important: /!important\b/i, function: { pattern: /(^|[^-a-z0-9])[-a-z0-9]+(?=\()/i, lookbehind: !0 }, punctuation: /[(){};:,]/ }, e.languages.css.atrule.inside.rest = e.languages.css;
  var r = e.languages.markup;
  r && (r.tag.addInlined("style", "css"), r.tag.addAttribute("style", "css"));
}(Prism), function(e) {
  var t = [/\b(?:async|sync|yield)\*/, /\b(?:abstract|assert|async|await|break|case|catch|class|const|continue|covariant|default|deferred|do|dynamic|else|enum|export|extends|extension|external|factory|final|finally|for|get|hide|if|implements|import|in|interface|library|mixin|new|null|on|operator|part|rethrow|return|set|show|static|super|switch|sync|this|throw|try|typedef|var|void|while|with|yield)\b/], r = /(^|[^\w.])(?:[a-z]\w*\s*\.\s*)*(?:[A-Z]\w*\s*\.\s*)*/.source, n = { pattern: RegExp(r + /[A-Z](?:[\d_A-Z]*[a-z]\w*)?\b/.source), lookbehind: !0, inside: { namespace: { pattern: /^[a-z]\w*(?:\s*\.\s*[a-z]\w*)*(?:\s*\.)?/, inside: { punctuation: /\./ } } } };
  e.languages.dart = e.languages.extend("clike", { "class-name": [n, { pattern: RegExp(r + /[A-Z]\w*(?=\s+\w+\s*[;,=()])/.source), lookbehind: !0, inside: n.inside }], keyword: t, operator: /\bis!|\b(?:as|is)\b|\+\+|--|&&|\|\||<<=?|>>=?|~(?:\/=?)?|[+\-*\/%&^|=!<>]=?|\?/ }), e.languages.insertBefore("dart", "string", { "string-literal": { pattern: /r?(?:("""|''')[\s\S]*?\1|(["'])(?:\\.|(?!\2)[^\\\r\n])*\2(?!\2))/, greedy: !0, inside: { interpolation: { pattern: /((?:^|[^\\])(?:\\{2})*)\$(?:\w+|\{(?:[^{}]|\{[^{}]*\})*\})/, lookbehind: !0, inside: { punctuation: /^\$\{?|\}$/, expression: { pattern: /[\s\S]+/, inside: e.languages.dart } } }, string: /[\s\S]+/ } }, string: void 0 }), e.languages.insertBefore("dart", "class-name", { metadata: { pattern: /@\w+/, alias: "function" } }), e.languages.insertBefore("dart", "class-name", { generics: { pattern: /<(?:[\w\s,.&?]|<(?:[\w\s,.&?]|<(?:[\w\s,.&?]|<[\w\s,.&?]*>)*>)*>)*>/, inside: { "class-name": n, keyword: t, punctuation: /[<>(),.:]/, operator: /[?&|]/ } } });
}(Prism), function(e) {
  e.languages.diff = { coord: [/^(?:\*{3}|-{3}|\+{3}).*$/m, /^@@.*@@$/m, /^\d.*$/m] };
  var t = { "deleted-sign": "-", "deleted-arrow": "<", "inserted-sign": "+", "inserted-arrow": ">", unchanged: " ", diff: "!" };
  Object.keys(t).forEach(function(r) {
    var n = t[r], a = [];
    /^\w+$/.test(r) || a.push(/\w+/.exec(r)[0]), r === "diff" && a.push("bold"), e.languages.diff[r] = { pattern: RegExp("^(?:[" + n + `].*(?:\r
?|
|(?![\\s\\S])))+`, "m"), alias: a, inside: { line: { pattern: /(.)(?=[\s\S]).*(?:\r\n?|\n)?/, lookbehind: !0 }, prefix: { pattern: /[\s\S]/, alias: /\w+/.exec(r)[0] } } };
  }), Object.defineProperty(e.languages.diff, "PREFIXES", { value: t });
}(Prism), function(e) {
  var t = /\\[\r\n](?:\s|\\[\r\n]|#.*(?!.))*(?![\s#]|\\[\r\n])/.source, r = /(?:[ \t]+(?![ \t])(?:<SP_BS>)?|<SP_BS>)/.source.replace(/<SP_BS>/g, function() {
    return t;
  }), n = /"(?:[^"\\\r\n]|\\(?:\r\n|[\s\S]))*"|'(?:[^'\\\r\n]|\\(?:\r\n|[\s\S]))*'/.source, a = /--[\w-]+=(?:<STR>|(?!["'])(?:[^\s\\]|\\.)+)/.source.replace(/<STR>/g, function() {
    return n;
  }), i = { pattern: RegExp(n), greedy: !0 }, o = { pattern: /(^[ \t]*)#.*/m, lookbehind: !0, greedy: !0 };
  function s(c, u) {
    return c = c.replace(/<OPT>/g, function() {
      return a;
    }).replace(/<SP>/g, function() {
      return r;
    }), RegExp(c, u);
  }
  e.languages.docker = { instruction: { pattern: /(^[ \t]*)(?:ADD|ARG|CMD|COPY|ENTRYPOINT|ENV|EXPOSE|FROM|HEALTHCHECK|LABEL|MAINTAINER|ONBUILD|RUN|SHELL|STOPSIGNAL|USER|VOLUME|WORKDIR)(?=\s)(?:\\.|[^\r\n\\])*(?:\\$(?:\s|#.*$)*(?![\s#])(?:\\.|[^\r\n\\])*)*/im, lookbehind: !0, greedy: !0, inside: { options: { pattern: s(/(^(?:ONBUILD<SP>)?\w+<SP>)<OPT>(?:<SP><OPT>)*/.source, "i"), lookbehind: !0, greedy: !0, inside: { property: { pattern: /(^|\s)--[\w-]+/, lookbehind: !0 }, string: [i, { pattern: /(=)(?!["'])(?:[^\s\\]|\\.)+/, lookbehind: !0 }], operator: /\\$/m, punctuation: /=/ } }, keyword: [{ pattern: s(/(^(?:ONBUILD<SP>)?HEALTHCHECK<SP>(?:<OPT><SP>)*)(?:CMD|NONE)\b/.source, "i"), lookbehind: !0, greedy: !0 }, { pattern: s(/(^(?:ONBUILD<SP>)?FROM<SP>(?:<OPT><SP>)*(?!--)[^ \t\\]+<SP>)AS/.source, "i"), lookbehind: !0, greedy: !0 }, { pattern: s(/(^ONBUILD<SP>)\w+/.source, "i"), lookbehind: !0, greedy: !0 }, { pattern: /^\w+/, greedy: !0 }], comment: o, string: i, variable: /\$(?:\w+|\{[^{}"'\\]*\})/, operator: /\\$/m } }, comment: o }, e.languages.dockerfile = e.languages.docker;
}(Prism), Prism.languages.git = { comment: /^#.*/m, deleted: /^[-–].*/m, inserted: /^\+.*/m, string: /("|')(?:\\.|(?!\1)[^\\\r\n])*\1/, command: { pattern: /^.*\$ git .*$/m, inside: { parameter: /\s--?\w+/ } }, coord: /^@@.*@@$/m, "commit-sha1": /^commit \w{40}$/m }, Prism.languages.glsl = Prism.languages.extend("c", { keyword: /\b(?:active|asm|atomic_uint|attribute|[ibdu]?vec[234]|bool|break|buffer|case|cast|centroid|class|coherent|common|const|continue|d?mat[234](?:x[234])?|default|discard|do|double|else|enum|extern|external|false|filter|fixed|flat|float|for|fvec[234]|goto|half|highp|hvec[234]|[iu]?sampler2DMS(?:Array)?|[iu]?sampler2DRect|[iu]?samplerBuffer|[iu]?samplerCube|[iu]?samplerCubeArray|[iu]?sampler[123]D|[iu]?sampler[12]DArray|[iu]?image2DMS(?:Array)?|[iu]?image2DRect|[iu]?imageBuffer|[iu]?imageCube|[iu]?imageCubeArray|[iu]?image[123]D|[iu]?image[12]DArray|if|in|inline|inout|input|int|interface|invariant|layout|long|lowp|mediump|namespace|noinline|noperspective|out|output|partition|patch|precise|precision|public|readonly|resource|restrict|return|sample|sampler[12]DArrayShadow|sampler[12]DShadow|sampler2DRectShadow|sampler3DRect|samplerCubeArrayShadow|samplerCubeShadow|shared|short|sizeof|smooth|static|struct|subroutine|superp|switch|template|this|true|typedef|uint|uniform|union|unsigned|using|varying|void|volatile|while|writeonly)\b/ }), Prism.languages.go = Prism.languages.extend("clike", { string: { pattern: /(^|[^\\])"(?:\\.|[^"\\\r\n])*"|`[^`]*`/, lookbehind: !0, greedy: !0 }, keyword: /\b(?:break|case|chan|const|continue|default|defer|else|fallthrough|for|func|go(?:to)?|if|import|interface|map|package|range|return|select|struct|switch|type|var)\b/, boolean: /\b(?:_|false|iota|nil|true)\b/, number: [/\b0(?:b[01_]+|o[0-7_]+)i?\b/i, /\b0x(?:[a-f\d_]+(?:\.[a-f\d_]*)?|\.[a-f\d_]+)(?:p[+-]?\d+(?:_\d+)*)?i?(?!\w)/i, /(?:\b\d[\d_]*(?:\.[\d_]*)?|\B\.\d[\d_]*)(?:e[+-]?[\d_]+)?i?(?!\w)/i], operator: /[*\/%^!=]=?|\+[=+]?|-[=-]?|\|[=|]?|&(?:=|&|\^=?)?|>(?:>=?|=)?|<(?:<=?|=|-)?|:=|\.\.\./, builtin: /\b(?:append|bool|byte|cap|close|complex|complex(?:64|128)|copy|delete|error|float(?:32|64)|u?int(?:8|16|32|64)?|imag|len|make|new|panic|print(?:ln)?|real|recover|rune|string|uintptr)\b/ }), Prism.languages.insertBefore("go", "string", { char: { pattern: /'(?:\\.|[^'\\\r\n]){0,10}'/, greedy: !0 } }), delete Prism.languages.go["class-name"], Prism.languages["go-mod"] = Prism.languages["go-module"] = { comment: { pattern: /\/\/.*/, greedy: !0 }, version: { pattern: /(^|[\s()[\],])v\d+\.\d+\.\d+(?:[+-][-+.\w]*)?(?![^\s()[\],])/, lookbehind: !0, alias: "number" }, "go-version": { pattern: /((?:^|\s)go\s+)\d+(?:\.\d+){1,2}/, lookbehind: !0, alias: "number" }, keyword: { pattern: /^([ \t]*)(?:exclude|go|module|replace|require|retract)\b/m, lookbehind: !0 }, operator: /=>/, punctuation: /[()[\],]/ }, function(e) {
  var t = /[*&][^\s[\]{},]+/, r = /!(?:<[\w\-%#;/?:@&=+$,.!~*'()[\]]+>|(?:[a-zA-Z\d-]*!)?[\w\-%#;/?:@&=+$.~*'()]+)?/, n = "(?:" + r.source + "(?:[ 	]+" + t.source + ")?|" + t.source + "(?:[ 	]+" + r.source + ")?)", a = /(?:[^\s\x00-\x08\x0e-\x1f!"#%&'*,\-:>?@[\]`{|}\x7f-\x84\x86-\x9f\ud800-\udfff\ufffe\uffff]|[?:-]<PLAIN>)(?:[ \t]*(?:(?![#:])<PLAIN>|:<PLAIN>))*/.source.replace(/<PLAIN>/g, function() {
    return /[^\s\x00-\x08\x0e-\x1f,[\]{}\x7f-\x84\x86-\x9f\ud800-\udfff\ufffe\uffff]/.source;
  }), i = /"(?:[^"\\\r\n]|\\.)*"|'(?:[^'\\\r\n]|\\.)*'/.source;
  function o(s, c) {
    c = (c || "").replace(/m/g, "") + "m";
    var u = /([:\-,[{]\s*(?:\s<<prop>>[ \t]+)?)(?:<<value>>)(?=[ \t]*(?:$|,|\]|\}|(?:[\r\n]\s*)?#))/.source.replace(/<<prop>>/g, function() {
      return n;
    }).replace(/<<value>>/g, function() {
      return s;
    });
    return RegExp(u, c);
  }
  e.languages.yaml = { scalar: { pattern: RegExp(/([\-:]\s*(?:\s<<prop>>[ \t]+)?[|>])[ \t]*(?:((?:\r?\n|\r)[ \t]+)\S[^\r\n]*(?:\2[^\r\n]+)*)/.source.replace(/<<prop>>/g, function() {
    return n;
  })), lookbehind: !0, alias: "string" }, comment: /#.*/, key: { pattern: RegExp(/((?:^|[:\-,[{\r\n?])[ \t]*(?:<<prop>>[ \t]+)?)<<key>>(?=\s*:\s)/.source.replace(/<<prop>>/g, function() {
    return n;
  }).replace(/<<key>>/g, function() {
    return "(?:" + a + "|" + i + ")";
  })), lookbehind: !0, greedy: !0, alias: "atrule" }, directive: { pattern: /(^[ \t]*)%.+/m, lookbehind: !0, alias: "important" }, datetime: { pattern: o(/\d{4}-\d\d?-\d\d?(?:[tT]|[ \t]+)\d\d?:\d{2}:\d{2}(?:\.\d*)?(?:[ \t]*(?:Z|[-+]\d\d?(?::\d{2})?))?|\d{4}-\d{2}-\d{2}|\d\d?:\d{2}(?::\d{2}(?:\.\d*)?)?/.source), lookbehind: !0, alias: "number" }, boolean: { pattern: o(/false|true/.source, "i"), lookbehind: !0, alias: "important" }, null: { pattern: o(/null|~/.source, "i"), lookbehind: !0, alias: "important" }, string: { pattern: o(i), lookbehind: !0, greedy: !0 }, number: { pattern: o(/[+-]?(?:0x[\da-f]+|0o[0-7]+|(?:\d+(?:\.\d*)?|\.\d+)(?:e[+-]?\d+)?|\.inf|\.nan)/.source, "i"), lookbehind: !0 }, tag: r, important: t, punctuation: /---|[:[\]{}\-,|>?]|\.\.\./ }, e.languages.yml = e.languages.yaml;
}(Prism), function(e) {
  var t = /(?:\\.|[^\\\n\r]|(?:\n|\r\n?)(?![\r\n]))/.source;
  function r(u) {
    return u = u.replace(/<inner>/g, function() {
      return t;
    }), RegExp(/((?:^|[^\\])(?:\\{2})*)/.source + "(?:" + u + ")");
  }
  var n = /(?:\\.|``(?:[^`\r\n]|`(?!`))+``|`[^`\r\n]+`|[^\\|\r\n`])+/.source, a = /\|?__(?:\|__)+\|?(?:(?:\n|\r\n?)|(?![\s\S]))/.source.replace(/__/g, function() {
    return n;
  }), i = /\|?[ \t]*:?-{3,}:?[ \t]*(?:\|[ \t]*:?-{3,}:?[ \t]*)+\|?(?:\n|\r\n?)/.source;
  e.languages.markdown = e.languages.extend("markup", {}), e.languages.insertBefore("markdown", "prolog", { "front-matter-block": { pattern: /(^(?:\s*[\r\n])?)---(?!.)[\s\S]*?[\r\n]---(?!.)/, lookbehind: !0, greedy: !0, inside: { punctuation: /^---|---$/, "front-matter": { pattern: /\S+(?:\s+\S+)*/, alias: ["yaml", "language-yaml"], inside: e.languages.yaml } } }, blockquote: { pattern: /^>(?:[\t ]*>)*/m, alias: "punctuation" }, table: { pattern: RegExp("^" + a + i + "(?:" + a + ")*", "m"), inside: { "table-data-rows": { pattern: RegExp("^(" + a + i + ")(?:" + a + ")*$"), lookbehind: !0, inside: { "table-data": { pattern: RegExp(n), inside: e.languages.markdown }, punctuation: /\|/ } }, "table-line": { pattern: RegExp("^(" + a + ")" + i + "$"), lookbehind: !0, inside: { punctuation: /\||:?-{3,}:?/ } }, "table-header-row": { pattern: RegExp("^" + a + "$"), inside: { "table-header": { pattern: RegExp(n), alias: "important", inside: e.languages.markdown }, punctuation: /\|/ } } } }, code: [{ pattern: /((?:^|\n)[ \t]*\n|(?:^|\r\n?)[ \t]*\r\n?)(?: {4}|\t).+(?:(?:\n|\r\n?)(?: {4}|\t).+)*/, lookbehind: !0, alias: "keyword" }, { pattern: /^```[\s\S]*?^```$/m, greedy: !0, inside: { "code-block": { pattern: /^(```.*(?:\n|\r\n?))[\s\S]+?(?=(?:\n|\r\n?)^```$)/m, lookbehind: !0 }, "code-language": { pattern: /^(```).+/, lookbehind: !0 }, punctuation: /```/ } }], title: [{ pattern: /\S.*(?:\n|\r\n?)(?:==+|--+)(?=[ \t]*$)/m, alias: "important", inside: { punctuation: /==+$|--+$/ } }, { pattern: /(^\s*)#.+/m, lookbehind: !0, alias: "important", inside: { punctuation: /^#+|#+$/ } }], hr: { pattern: /(^\s*)([*-])(?:[\t ]*\2){2,}(?=\s*$)/m, lookbehind: !0, alias: "punctuation" }, list: { pattern: /(^\s*)(?:[*+-]|\d+\.)(?=[\t ].)/m, lookbehind: !0, alias: "punctuation" }, "url-reference": { pattern: /!?\[[^\]]+\]:[\t ]+(?:\S+|<(?:\\.|[^>\\])+>)(?:[\t ]+(?:"(?:\\.|[^"\\])*"|'(?:\\.|[^'\\])*'|\((?:\\.|[^)\\])*\)))?/, inside: { variable: { pattern: /^(!?\[)[^\]]+/, lookbehind: !0 }, string: /(?:"(?:\\.|[^"\\])*"|'(?:\\.|[^'\\])*'|\((?:\\.|[^)\\])*\))$/, punctuation: /^[\[\]!:]|[<>]/ }, alias: "url" }, bold: { pattern: r(/\b__(?:(?!_)<inner>|_(?:(?!_)<inner>)+_)+__\b|\*\*(?:(?!\*)<inner>|\*(?:(?!\*)<inner>)+\*)+\*\*/.source), lookbehind: !0, greedy: !0, inside: { content: { pattern: /(^..)[\s\S]+(?=..$)/, lookbehind: !0, inside: {} }, punctuation: /\*\*|__/ } }, italic: { pattern: r(/\b_(?:(?!_)<inner>|__(?:(?!_)<inner>)+__)+_\b|\*(?:(?!\*)<inner>|\*\*(?:(?!\*)<inner>)+\*\*)+\*/.source), lookbehind: !0, greedy: !0, inside: { content: { pattern: /(^.)[\s\S]+(?=.$)/, lookbehind: !0, inside: {} }, punctuation: /[*_]/ } }, strike: { pattern: r(/(~~?)(?:(?!~)<inner>)+\2/.source), lookbehind: !0, greedy: !0, inside: { content: { pattern: /(^~~?)[\s\S]+(?=\1$)/, lookbehind: !0, inside: {} }, punctuation: /~~?/ } }, "code-snippet": { pattern: /(^|[^\\`])(?:``[^`\r\n]+(?:`[^`\r\n]+)*``(?!`)|`[^`\r\n]+`(?!`))/, lookbehind: !0, greedy: !0, alias: ["code", "keyword"] }, url: { pattern: r(/!?\[(?:(?!\])<inner>)+\](?:\([^\s)]+(?:[\t ]+"(?:\\.|[^"\\])*")?\)|[ \t]?\[(?:(?!\])<inner>)+\])/.source), lookbehind: !0, greedy: !0, inside: { operator: /^!/, content: { pattern: /(^\[)[^\]]+(?=\])/, lookbehind: !0, inside: {} }, variable: { pattern: /(^\][ \t]?\[)[^\]]+(?=\]$)/, lookbehind: !0 }, url: { pattern: /(^\]\()[^\s)]+/, lookbehind: !0 }, string: { pattern: /(^[ \t]+)"(?:\\.|[^"\\])*"(?=\)$)/, lookbehind: !0 } } } }), ["url", "bold", "italic", "strike"].forEach(function(u) {
    ["url", "bold", "italic", "strike", "code-snippet"].forEach(function(l) {
      u !== l && (e.languages.markdown[u].inside.content.inside[l] = e.languages.markdown[l]);
    });
  }), e.hooks.add("after-tokenize", function(u) {
    u.language !== "markdown" && u.language !== "md" || function l(f) {
      if (f && typeof f != "string")
        for (var p = 0, d = f.length; p < d; p++) {
          var g = f[p];
          if (g.type === "code") {
            var _ = g.content[1], m = g.content[3];
            if (_ && m && _.type === "code-language" && m.type === "code-block" && typeof _.content == "string") {
              var h = _.content.replace(/\b#/g, "sharp").replace(/\b\+\+/g, "pp"), v = "language-" + (h = (/[a-z][\w-]*/i.exec(h) || [""])[0].toLowerCase());
              m.alias ? typeof m.alias == "string" ? m.alias = [m.alias, v] : m.alias.push(v) : m.alias = [v];
            }
          } else
            l(g.content);
        }
    }(u.tokens);
  }), e.hooks.add("wrap", function(u) {
    if (u.type === "code-block") {
      for (var l = "", f = 0, p = u.classes.length; f < p; f++) {
        var d = u.classes[f], g = /language-(.+)/.exec(d);
        if (g) {
          l = g[1];
          break;
        }
      }
      var _ = e.languages[l];
      if (_)
        u.content = e.highlight(function(h) {
          var v = h.replace(o, "");
          return v = v.replace(/&(\w{1,8}|#x?[\da-f]{1,8});/gi, function(w, E) {
            var S;
            if ((E = E.toLowerCase())[0] === "#")
              return S = E[1] === "x" ? parseInt(E.slice(2), 16) : Number(E.slice(1)), c(S);
            var A = s[E];
            return A || w;
          });
        }(u.content), _, l);
      else if (l && l !== "none" && e.plugins.autoloader) {
        var m = "md-" + (/* @__PURE__ */ new Date()).valueOf() + "-" + Math.floor(1e16 * Math.random());
        u.attributes.id = m, e.plugins.autoloader.loadLanguages(l, function() {
          var h = document.getElementById(m);
          h && (h.innerHTML = e.highlight(h.textContent, e.languages[l], l));
        });
      }
    }
  });
  var o = RegExp(e.languages.markup.tag.pattern.source, "gi"), s = { amp: "&", lt: "<", gt: ">", quot: '"' }, c = String.fromCodePoint || String.fromCharCode;
  e.languages.md = e.languages.markdown;
}(Prism), Prism.languages.graphql = { comment: /#.*/, description: { pattern: /(?:"""(?:[^"]|(?!""")")*"""|"(?:\\.|[^\\"\r\n])*")(?=\s*[a-z_])/i, greedy: !0, alias: "string", inside: { "language-markdown": { pattern: /(^"(?:"")?)(?!\1)[\s\S]+(?=\1$)/, lookbehind: !0, inside: Prism.languages.markdown } } }, string: { pattern: /"""(?:[^"]|(?!""")")*"""|"(?:\\.|[^\\"\r\n])*"/, greedy: !0 }, number: /(?:\B-|\b)\d+(?:\.\d+)?(?:e[+-]?\d+)?\b/i, boolean: /\b(?:false|true)\b/, variable: /\$[a-z_]\w*/i, directive: { pattern: /@[a-z_]\w*/i, alias: "function" }, "attr-name": { pattern: /\b[a-z_]\w*(?=\s*(?:\((?:[^()"]|"(?:\\.|[^\\"\r\n])*")*\))?:)/i, greedy: !0 }, "atom-input": { pattern: /\b[A-Z]\w*Input\b/, alias: "class-name" }, scalar: /\b(?:Boolean|Float|ID|Int|String)\b/, constant: /\b[A-Z][A-Z_\d]*\b/, "class-name": { pattern: /(\b(?:enum|implements|interface|on|scalar|type|union)\s+|&\s*|:\s*|\[)[A-Z_]\w*/, lookbehind: !0 }, fragment: { pattern: /(\bfragment\s+|\.{3}\s*(?!on\b))[a-zA-Z_]\w*/, lookbehind: !0, alias: "function" }, "definition-mutation": { pattern: /(\bmutation\s+)[a-zA-Z_]\w*/, lookbehind: !0, alias: "function" }, "definition-query": { pattern: /(\bquery\s+)[a-zA-Z_]\w*/, lookbehind: !0, alias: "function" }, keyword: /\b(?:directive|enum|extend|fragment|implements|input|interface|mutation|on|query|repeatable|scalar|schema|subscription|type|union)\b/, operator: /[!=|&]|\.{3}/, "property-query": /\w+(?=\s*\()/, object: /\w+(?=\s*\{)/, punctuation: /[!(){}\[\]:=,]/, property: /\w+/ }, Prism.hooks.add("after-tokenize", function(e) {
  if (e.language === "graphql")
    for (var t = e.tokens.filter(function(g) {
      return typeof g != "string" && g.type !== "comment" && g.type !== "scalar";
    }), r = 0; r < t.length; ) {
      var n = t[r++];
      if (n.type === "keyword" && n.content === "mutation") {
        var a = [];
        if (f(["definition-mutation", "punctuation"]) && l(1).content === "(") {
          r += 2;
          var i = p(/^\($/, /^\)$/);
          if (i === -1)
            continue;
          for (; r < i; r++) {
            var o = l(0);
            o.type === "variable" && (d(o, "variable-input"), a.push(o.content));
          }
          r = i + 1;
        }
        if (f(["punctuation", "property-query"]) && l(0).content === "{" && (r++, d(l(0), "property-mutation"), a.length > 0)) {
          var s = p(/^\{$/, /^\}$/);
          if (s === -1)
            continue;
          for (var c = r; c < s; c++) {
            var u = t[c];
            u.type === "variable" && a.indexOf(u.content) >= 0 && d(u, "variable-input");
          }
        }
      }
    }
  function l(g) {
    return t[r + g];
  }
  function f(g, _) {
    _ = _ || 0;
    for (var m = 0; m < g.length; m++) {
      var h = l(m + _);
      if (!h || h.type !== g[m])
        return !1;
    }
    return !0;
  }
  function p(g, _) {
    for (var m = 1, h = r; h < t.length; h++) {
      var v = t[h], w = v.content;
      if (v.type === "punctuation" && typeof w == "string") {
        if (g.test(w))
          m++;
        else if (_.test(w) && --m === 0)
          return h;
      }
    }
    return -1;
  }
  function d(g, _) {
    var m = g.alias;
    m ? Array.isArray(m) || (g.alias = m = [m]) : g.alias = m = [], m.push(_);
  }
}), function(e) {
  e.languages.ruby = e.languages.extend("clike", { comment: { pattern: /#.*|^=begin\s[\s\S]*?^=end/m, greedy: !0 }, "class-name": { pattern: /(\b(?:class|module)\s+|\bcatch\s+\()[\w.\\]+|\b[A-Z_]\w*(?=\s*\.\s*new\b)/, lookbehind: !0, inside: { punctuation: /[.\\]/ } }, keyword: /\b(?:BEGIN|END|alias|and|begin|break|case|class|def|define_method|defined|do|each|else|elsif|end|ensure|extend|for|if|in|include|module|new|next|nil|not|or|prepend|private|protected|public|raise|redo|require|rescue|retry|return|self|super|then|throw|undef|unless|until|when|while|yield)\b/, operator: /\.{2,3}|&\.|===|<?=>|[!=]?~|(?:&&|\|\||<<|>>|\*\*|[+\-*/%<>!^&|=])=?|[?:]/, punctuation: /[(){}[\].,;]/ }), e.languages.insertBefore("ruby", "operator", { "double-colon": { pattern: /::/, alias: "punctuation" } });
  var t = { pattern: /((?:^|[^\\])(?:\\{2})*)#\{(?:[^{}]|\{[^{}]*\})*\}/, lookbehind: !0, inside: { content: { pattern: /^(#\{)[\s\S]+(?=\}$)/, lookbehind: !0, inside: e.languages.ruby }, delimiter: { pattern: /^#\{|\}$/, alias: "punctuation" } } };
  delete e.languages.ruby.function;
  var r = "(?:" + [/([^a-zA-Z0-9\s{(\[<=])(?:(?!\1)[^\\]|\\[\s\S])*\1/.source, /\((?:[^()\\]|\\[\s\S]|\((?:[^()\\]|\\[\s\S])*\))*\)/.source, /\{(?:[^{}\\]|\\[\s\S]|\{(?:[^{}\\]|\\[\s\S])*\})*\}/.source, /\[(?:[^\[\]\\]|\\[\s\S]|\[(?:[^\[\]\\]|\\[\s\S])*\])*\]/.source, /<(?:[^<>\\]|\\[\s\S]|<(?:[^<>\\]|\\[\s\S])*>)*>/.source].join("|") + ")", n = /(?:"(?:\\.|[^"\\\r\n])*"|(?:\b[a-zA-Z_]\w*|[^\s\0-\x7F]+)[?!]?|\$.)/.source;
  e.languages.insertBefore("ruby", "keyword", { "regex-literal": [{ pattern: RegExp(/%r/.source + r + /[egimnosux]{0,6}/.source), greedy: !0, inside: { interpolation: t, regex: /[\s\S]+/ } }, { pattern: /(^|[^/])\/(?!\/)(?:\[[^\r\n\]]+\]|\\.|[^[/\\\r\n])+\/[egimnosux]{0,6}(?=\s*(?:$|[\r\n,.;})#]))/, lookbehind: !0, greedy: !0, inside: { interpolation: t, regex: /[\s\S]+/ } }], variable: /[@$]+[a-zA-Z_]\w*(?:[?!]|\b)/, symbol: [{ pattern: RegExp(/(^|[^:]):/.source + n), lookbehind: !0, greedy: !0 }, { pattern: RegExp(/([\r\n{(,][ \t]*)/.source + n + /(?=:(?!:))/.source), lookbehind: !0, greedy: !0 }], "method-definition": { pattern: /(\bdef\s+)\w+(?:\s*\.\s*\w+)?/, lookbehind: !0, inside: { function: /\b\w+$/, keyword: /^self\b/, "class-name": /^\w+/, punctuation: /\./ } } }), e.languages.insertBefore("ruby", "string", { "string-literal": [{ pattern: RegExp(/%[qQiIwWs]?/.source + r), greedy: !0, inside: { interpolation: t, string: /[\s\S]+/ } }, { pattern: /("|')(?:#\{[^}]+\}|#(?!\{)|\\(?:\r\n|[\s\S])|(?!\1)[^\\#\r\n])*\1/, greedy: !0, inside: { interpolation: t, string: /[\s\S]+/ } }, { pattern: /<<[-~]?([a-z_]\w*)[\r\n](?:.*[\r\n])*?[\t ]*\1/i, alias: "heredoc-string", greedy: !0, inside: { delimiter: { pattern: /^<<[-~]?[a-z_]\w*|\b[a-z_]\w*$/i, inside: { symbol: /\b\w+/, punctuation: /^<<[-~]?/ } }, interpolation: t, string: /[\s\S]+/ } }, { pattern: /<<[-~]?'([a-z_]\w*)'[\r\n](?:.*[\r\n])*?[\t ]*\1/i, alias: "heredoc-string", greedy: !0, inside: { delimiter: { pattern: /^<<[-~]?'[a-z_]\w*'|\b[a-z_]\w*$/i, inside: { symbol: /\b\w+/, punctuation: /^<<[-~]?'|'$/ } }, string: /[\s\S]+/ } }], "command-literal": [{ pattern: RegExp(/%x/.source + r), greedy: !0, inside: { interpolation: t, command: { pattern: /[\s\S]+/, alias: "string" } } }, { pattern: /`(?:#\{[^}]+\}|#(?!\{)|\\(?:\r\n|[\s\S])|[^\\`#\r\n])*`/, greedy: !0, inside: { interpolation: t, command: { pattern: /[\s\S]+/, alias: "string" } } }] }), delete e.languages.ruby.string, e.languages.insertBefore("ruby", "number", { builtin: /\b(?:Array|Bignum|Binding|Class|Continuation|Dir|Exception|FalseClass|File|Fixnum|Float|Hash|IO|Integer|MatchData|Method|Module|NilClass|Numeric|Object|Proc|Range|Regexp|Stat|String|Struct|Symbol|TMS|Thread|ThreadGroup|Time|TrueClass)\b/, constant: /\b[A-Z][A-Z0-9_]*(?:[?!]|\b)/ }), e.languages.rb = e.languages.ruby;
}(Prism), function(e) {
  var t = { pattern: /\\[\\(){}[\]^$+*?|.]/, alias: "escape" }, r = /\\(?:x[\da-fA-F]{2}|u[\da-fA-F]{4}|u\{[\da-fA-F]+\}|0[0-7]{0,2}|[123][0-7]{2}|c[a-zA-Z]|.)/, n = "(?:[^\\\\-]|" + r.source + ")", a = RegExp(n + "-" + n), i = { pattern: /(<|')[^<>']+(?=[>']$)/, lookbehind: !0, alias: "variable" };
  e.languages.regex = { "char-class": { pattern: /((?:^|[^\\])(?:\\\\)*)\[(?:[^\\\]]|\\[\s\S])*\]/, lookbehind: !0, inside: { "char-class-negation": { pattern: /(^\[)\^/, lookbehind: !0, alias: "operator" }, "char-class-punctuation": { pattern: /^\[|\]$/, alias: "punctuation" }, range: { pattern: a, inside: { escape: r, "range-punctuation": { pattern: /-/, alias: "operator" } } }, "special-escape": t, "char-set": { pattern: /\\[wsd]|\\p\{[^{}]+\}/i, alias: "class-name" }, escape: r } }, "special-escape": t, "char-set": { pattern: /\.|\\[wsd]|\\p\{[^{}]+\}/i, alias: "class-name" }, backreference: [{ pattern: /\\(?![123][0-7]{2})[1-9]/, alias: "keyword" }, { pattern: /\\k<[^<>']+>/, alias: "keyword", inside: { "group-name": i } }], anchor: { pattern: /[$^]|\\[ABbGZz]/, alias: "function" }, escape: r, group: [{ pattern: /\((?:\?(?:<[^<>']+>|'[^<>']+'|[>:]|<?[=!]|[idmnsuxU]+(?:-[idmnsuxU]+)?:?))?/, alias: "punctuation", inside: { "group-name": i } }, { pattern: /\)/, alias: "punctuation" }], quantifier: { pattern: /(?:[+*?]|\{\d+(?:,\d*)?\})[?+]?/, alias: "number" }, alternation: { pattern: /\|/, alias: "keyword" } };
}(Prism), Prism.languages.javascript = Prism.languages.extend("clike", { "class-name": [Prism.languages.clike["class-name"], { pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$A-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\.(?:constructor|prototype))/, lookbehind: !0 }], keyword: [{ pattern: /((?:^|\})\s*)catch\b/, lookbehind: !0 }, { pattern: /(^|[^.]|\.\.\.\s*)\b(?:as|assert(?=\s*\{)|async(?=\s*(?:function\b|\(|[$\w\xA0-\uFFFF]|$))|await|break|case|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally(?=\s*(?:\{|$))|for|from(?=\s*(?:['"]|$))|function|(?:get|set)(?=\s*(?:[#\[$\w\xA0-\uFFFF]|$))|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)\b/, lookbehind: !0 }], function: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*(?:\.\s*(?:apply|bind|call)\s*)?\()/, number: { pattern: RegExp(/(^|[^\w$])/.source + "(?:" + /NaN|Infinity/.source + "|" + /0[bB][01]+(?:_[01]+)*n?/.source + "|" + /0[oO][0-7]+(?:_[0-7]+)*n?/.source + "|" + /0[xX][\dA-Fa-f]+(?:_[\dA-Fa-f]+)*n?/.source + "|" + /\d+(?:_\d+)*n/.source + "|" + /(?:\d+(?:_\d+)*(?:\.(?:\d+(?:_\d+)*)?)?|\.\d+(?:_\d+)*)(?:[Ee][+-]?\d+(?:_\d+)*)?/.source + ")" + /(?![\w$])/.source), lookbehind: !0 }, operator: /--|\+\+|\*\*=?|=>|&&=?|\|\|=?|[!=]==|<<=?|>>>?=?|[-+*/%&|^!=<>]=?|\.{3}|\?\?=?|\?\.?|[~:]/ }), Prism.languages.javascript["class-name"][0].pattern = /(\b(?:class|extends|implements|instanceof|interface|new)\s+)[\w.\\]+/, Prism.languages.insertBefore("javascript", "keyword", { regex: { pattern: RegExp(/((?:^|[^$\w\xA0-\uFFFF."'\])\s]|\b(?:return|yield))\s*)/.source + /\//.source + "(?:" + /(?:\[(?:[^\]\\\r\n]|\\.)*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}/.source + "|" + /(?:\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.|\[(?:[^[\]\\\r\n]|\\.)*\])*\])*\]|\\.|[^/\\\[\r\n])+\/[dgimyus]{0,7}v[dgimyus]{0,7}/.source + ")" + /(?=(?:\s|\/\*(?:[^*]|\*(?!\/))*\*\/)*(?:$|[\r\n,.;:})\]]|\/\/))/.source), lookbehind: !0, greedy: !0, inside: { "regex-source": { pattern: /^(\/)[\s\S]+(?=\/[a-z]*$)/, lookbehind: !0, alias: "language-regex", inside: Prism.languages.regex }, "regex-delimiter": /^\/|\/$/, "regex-flags": /^[a-z]+$/ } }, "function-variable": { pattern: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*[=:]\s*(?:async\s*)?(?:\bfunction\b|(?:\((?:[^()]|\([^()]*\))*\)|(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)\s*=>))/, alias: "function" }, parameter: [{ pattern: /(function(?:\s+(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)?\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\))/, lookbehind: !0, inside: Prism.languages.javascript }, { pattern: /(^|[^$\w\xA0-\uFFFF])(?!\s)[_$a-z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*=>)/i, lookbehind: !0, inside: Prism.languages.javascript }, { pattern: /(\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*=>)/, lookbehind: !0, inside: Prism.languages.javascript }, { pattern: /((?:\b|\s|^)(?!(?:as|async|await|break|case|catch|class|const|continue|debugger|default|delete|do|else|enum|export|extends|finally|for|from|function|get|if|implements|import|in|instanceof|interface|let|new|null|of|package|private|protected|public|return|set|static|super|switch|this|throw|try|typeof|undefined|var|void|while|with|yield)(?![$\w\xA0-\uFFFF]))(?:(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*\s*)\(\s*|\]\s*\(\s*)(?!\s)(?:[^()\s]|\s+(?![\s)])|\([^()]*\))+(?=\s*\)\s*\{)/, lookbehind: !0, inside: Prism.languages.javascript }], constant: /\b[A-Z](?:[A-Z_]|\dx?)*\b/ }), Prism.languages.insertBefore("javascript", "string", { hashbang: { pattern: /^#!.*/, greedy: !0, alias: "comment" }, "template-string": { pattern: /`(?:\\[\s\S]|\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}|(?!\$\{)[^\\`])*`/, greedy: !0, inside: { "template-punctuation": { pattern: /^`|`$/, alias: "string" }, interpolation: { pattern: /((?:^|[^\\])(?:\\{2})*)\$\{(?:[^{}]|\{(?:[^{}]|\{[^}]*\})*\})+\}/, lookbehind: !0, inside: { "interpolation-punctuation": { pattern: /^\$\{|\}$/, alias: "punctuation" }, rest: Prism.languages.javascript } }, string: /[\s\S]+/ } }, "string-property": { pattern: /((?:^|[,{])[ \t]*)(["'])(?:\\(?:\r\n|[\s\S])|(?!\2)[^\\\r\n])*\2(?=\s*:)/m, lookbehind: !0, greedy: !0, alias: "property" } }), Prism.languages.insertBefore("javascript", "operator", { "literal-property": { pattern: /((?:^|[,{])[ \t]*)(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*:)/m, lookbehind: !0, alias: "property" } }), Prism.languages.markup && (Prism.languages.markup.tag.addInlined("script", "javascript"), Prism.languages.markup.tag.addAttribute(/on(?:abort|blur|change|click|composition(?:end|start|update)|dblclick|error|focus(?:in|out)?|key(?:down|up)|load|mouse(?:down|enter|leave|move|out|over|up)|reset|resize|scroll|select|slotchange|submit|unload|wheel)/.source, "javascript")), Prism.languages.js = Prism.languages.javascript, function(e) {
  function t(r, n) {
    return "___" + r.toUpperCase() + n + "___";
  }
  Object.defineProperties(e.languages["markup-templating"] = {}, { buildPlaceholders: { value: function(r, n, a, i) {
    if (r.language === n) {
      var o = r.tokenStack = [];
      r.code = r.code.replace(a, function(s) {
        if (typeof i == "function" && !i(s))
          return s;
        for (var c, u = o.length; r.code.indexOf(c = t(n, u)) !== -1; )
          ++u;
        return o[u] = s, c;
      }), r.grammar = e.languages.markup;
    }
  } }, tokenizePlaceholders: { value: function(r, n) {
    if (r.language === n && r.tokenStack) {
      r.grammar = e.languages[n];
      var a = 0, i = Object.keys(r.tokenStack);
      (function o(s) {
        for (var c = 0; c < s.length && !(a >= i.length); c++) {
          var u = s[c];
          if (typeof u == "string" || u.content && typeof u.content == "string") {
            var l = i[a], f = r.tokenStack[l], p = typeof u == "string" ? u : u.content, d = t(n, l), g = p.indexOf(d);
            if (g > -1) {
              ++a;
              var _ = p.substring(0, g), m = new e.Token(n, e.tokenize(f, r.grammar), "language-" + n, f), h = p.substring(g + d.length), v = [];
              _ && v.push.apply(v, o([_])), v.push(m), h && v.push.apply(v, o([h])), typeof u == "string" ? s.splice.apply(s, [c, 1].concat(v)) : u.content = v;
            }
          } else
            u.content && o(u.content);
        }
        return s;
      })(r.tokens);
    }
  } } });
}(Prism), Prism.languages.less = Prism.languages.extend("css", { comment: [/\/\*[\s\S]*?\*\//, { pattern: /(^|[^\\])\/\/.*/, lookbehind: !0 }], atrule: { pattern: /@[\w-](?:\((?:[^(){}]|\([^(){}]*\))*\)|[^(){};\s]|\s+(?!\s))*?(?=\s*\{)/, inside: { punctuation: /[:()]/ } }, selector: { pattern: /(?:@\{[\w-]+\}|[^{};\s@])(?:@\{[\w-]+\}|\((?:[^(){}]|\([^(){}]*\))*\)|[^(){};@\s]|\s+(?!\s))*?(?=\s*\{)/, inside: { variable: /@+[\w-]+/ } }, property: /(?:@\{[\w-]+\}|[\w-])+(?:\+_?)?(?=\s*:)/, operator: /[+\-*\/]/ }), Prism.languages.insertBefore("less", "property", { variable: [{ pattern: /@[\w-]+\s*:/, inside: { punctuation: /:/ } }, /@@?[\w-]+/], "mixin-usage": { pattern: /([{;]\s*)[.#](?!\d)[\w-].*?(?=[(;])/, lookbehind: !0, alias: "function" } }), Prism.languages.scss = Prism.languages.extend("css", { comment: { pattern: /(^|[^\\])(?:\/\*[\s\S]*?\*\/|\/\/.*)/, lookbehind: !0 }, atrule: { pattern: /@[\w-](?:\([^()]+\)|[^()\s]|\s+(?!\s))*?(?=\s+[{;])/, inside: { rule: /@[\w-]+/ } }, url: /(?:[-a-z]+-)?url(?=\()/i, selector: { pattern: /(?=\S)[^@;{}()]?(?:[^@;{}()\s]|\s+(?!\s)|#\{\$[-\w]+\})+(?=\s*\{(?:\}|\s|[^}][^:{}]*[:{][^}]))/, inside: { parent: { pattern: /&/, alias: "important" }, placeholder: /%[-\w]+/, variable: /\$[-\w]+|#\{\$[-\w]+\}/ } }, property: { pattern: /(?:[-\w]|\$[-\w]|#\{\$[-\w]+\})+(?=\s*:)/, inside: { variable: /\$[-\w]+|#\{\$[-\w]+\}/ } } }), Prism.languages.insertBefore("scss", "atrule", { keyword: [/@(?:content|debug|each|else(?: if)?|extend|for|forward|function|if|import|include|mixin|return|use|warn|while)\b/i, { pattern: /( )(?:from|through)(?= )/, lookbehind: !0 }] }), Prism.languages.insertBefore("scss", "important", { variable: /\$[-\w]+|#\{\$[-\w]+\}/ }), Prism.languages.insertBefore("scss", "function", { "module-modifier": { pattern: /\b(?:as|hide|show|with)\b/i, alias: "keyword" }, placeholder: { pattern: /%[-\w]+/, alias: "selector" }, statement: { pattern: /\B!(?:default|optional)\b/i, alias: "keyword" }, boolean: /\b(?:false|true)\b/, null: { pattern: /\bnull\b/, alias: "keyword" }, operator: { pattern: /(\s)(?:[-+*\/%]|[=!]=|<=?|>=?|and|not|or)(?=\s)/, lookbehind: !0 } }), Prism.languages.scss.atrule.inside.rest = Prism.languages.scss, function(e) {
  e.languages.haml = { "multiline-comment": { pattern: /((?:^|\r?\n|\r)([\t ]*))(?:\/|-#).*(?:(?:\r?\n|\r)\2[\t ].+)*/, lookbehind: !0, alias: "comment" }, "multiline-code": [{ pattern: /((?:^|\r?\n|\r)([\t ]*)(?:[~-]|[&!]?=)).*,[\t ]*(?:(?:\r?\n|\r)\2[\t ].*,[\t ]*)*(?:(?:\r?\n|\r)\2[\t ].+)/, lookbehind: !0, inside: e.languages.ruby }, { pattern: /((?:^|\r?\n|\r)([\t ]*)(?:[~-]|[&!]?=)).*\|[\t ]*(?:(?:\r?\n|\r)\2[\t ].*\|[\t ]*)*/, lookbehind: !0, inside: e.languages.ruby }], filter: { pattern: /((?:^|\r?\n|\r)([\t ]*)):[\w-]+(?:(?:\r?\n|\r)(?:\2[\t ].+|\s*?(?=\r?\n|\r)))+/, lookbehind: !0, inside: { "filter-name": { pattern: /^:[\w-]+/, alias: "symbol" } } }, markup: { pattern: /((?:^|\r?\n|\r)[\t ]*)<.+/, lookbehind: !0, inside: e.languages.markup }, doctype: { pattern: /((?:^|\r?\n|\r)[\t ]*)!!!(?: .+)?/, lookbehind: !0 }, tag: { pattern: /((?:^|\r?\n|\r)[\t ]*)[%.#][\w\-#.]*[\w\-](?:\([^)]+\)|\{(?:\{[^}]+\}|[^{}])+\}|\[[^\]]+\])*[\/<>]*/, lookbehind: !0, inside: { attributes: [{ pattern: /(^|[^#])\{(?:\{[^}]+\}|[^{}])+\}/, lookbehind: !0, inside: e.languages.ruby }, { pattern: /\([^)]+\)/, inside: { "attr-value": { pattern: /(=\s*)(?:"(?:\\.|[^\\"\r\n])*"|[^)\s]+)/, lookbehind: !0 }, "attr-name": /[\w:-]+(?=\s*!?=|\s*[,)])/, punctuation: /[=(),]/ } }, { pattern: /\[[^\]]+\]/, inside: e.languages.ruby }], punctuation: /[<>]/ } }, code: { pattern: /((?:^|\r?\n|\r)[\t ]*(?:[~-]|[&!]?=)).+/, lookbehind: !0, inside: e.languages.ruby }, interpolation: { pattern: /#\{[^}]+\}/, inside: { delimiter: { pattern: /^#\{|\}$/, alias: "punctuation" }, ruby: { pattern: /[\s\S]+/, inside: e.languages.ruby } } }, punctuation: { pattern: /((?:^|\r?\n|\r)[\t ]*)[~=\-&!]+/, lookbehind: !0 } };
  for (var t = ["css", { filter: "coffee", language: "coffeescript" }, "erb", "javascript", "less", "markdown", "ruby", "scss", "textile"], r = {}, n = 0, a = t.length; n < a; n++) {
    var i = t[n];
    i = typeof i == "string" ? { filter: i, language: i } : i, e.languages[i.language] && (r["filter-" + i.filter] = { pattern: RegExp("((?:^|\\r?\\n|\\r)([\\t ]*)):{{filter_name}}(?:(?:\\r?\\n|\\r)(?:\\2[\\t ].+|\\s*?(?=\\r?\\n|\\r)))+".replace("{{filter_name}}", function() {
      return i.filter;
    })), lookbehind: !0, inside: { "filter-name": { pattern: /^:[\w-]+/, alias: "symbol" }, text: { pattern: /[\s\S]+/, alias: [i.language, "language-" + i.language], inside: e.languages[i.language] } } });
  }
  e.languages.insertBefore("haml", "filter", r);
}(Prism), Prism.languages.ini = { comment: { pattern: /(^[ \f\t\v]*)[#;][^\n\r]*/m, lookbehind: !0 }, section: { pattern: /(^[ \f\t\v]*)\[[^\n\r\]]*\]?/m, lookbehind: !0, inside: { "section-name": { pattern: /(^\[[ \f\t\v]*)[^ \f\t\v\]]+(?:[ \f\t\v]+[^ \f\t\v\]]+)*/, lookbehind: !0, alias: "selector" }, punctuation: /\[|\]/ } }, key: { pattern: /(^[ \f\t\v]*)[^ \f\n\r\t\v=]+(?:[ \f\t\v]+[^ \f\n\r\t\v=]+)*(?=[ \f\t\v]*=)/m, lookbehind: !0, alias: "attr-name" }, value: { pattern: /(=[ \f\t\v]*)[^ \f\n\r\t\v]+(?:[ \f\t\v]+[^ \f\n\r\t\v]+)*/, lookbehind: !0, alias: "attr-value", inside: { "inner-value": { pattern: /^("|').+(?=\1$)/, lookbehind: !0 } } }, punctuation: /=/ }, function(e) {
  var t = /\b(?:abstract|assert|boolean|break|byte|case|catch|char|class|const|continue|default|do|double|else|enum|exports|extends|final|finally|float|for|goto|if|implements|import|instanceof|int|interface|long|module|native|new|non-sealed|null|open|opens|package|permits|private|protected|provides|public|record(?!\s*[(){}[\]<>=%~.:,;?+\-*/&|^])|requires|return|sealed|short|static|strictfp|super|switch|synchronized|this|throw|throws|to|transient|transitive|try|uses|var|void|volatile|while|with|yield)\b/, r = /(?:[a-z]\w*\s*\.\s*)*(?:[A-Z]\w*\s*\.\s*)*/.source, n = { pattern: RegExp(/(^|[^\w.])/.source + r + /[A-Z](?:[\d_A-Z]*[a-z]\w*)?\b/.source), lookbehind: !0, inside: { namespace: { pattern: /^[a-z]\w*(?:\s*\.\s*[a-z]\w*)*(?:\s*\.)?/, inside: { punctuation: /\./ } }, punctuation: /\./ } };
  e.languages.java = e.languages.extend("clike", { string: { pattern: /(^|[^\\])"(?:\\.|[^"\\\r\n])*"/, lookbehind: !0, greedy: !0 }, "class-name": [n, { pattern: RegExp(/(^|[^\w.])/.source + r + /[A-Z]\w*(?=\s+\w+\s*[;,=()]|\s*(?:\[[\s,]*\]\s*)?::\s*new\b)/.source), lookbehind: !0, inside: n.inside }, { pattern: RegExp(/(\b(?:class|enum|extends|implements|instanceof|interface|new|record|throws)\s+)/.source + r + /[A-Z]\w*\b/.source), lookbehind: !0, inside: n.inside }], keyword: t, function: [e.languages.clike.function, { pattern: /(::\s*)[a-z_]\w*/, lookbehind: !0 }], number: /\b0b[01][01_]*L?\b|\b0x(?:\.[\da-f_p+-]+|[\da-f_]+(?:\.[\da-f_p+-]+)?)\b|(?:\b\d[\d_]*(?:\.[\d_]*)?|\B\.\d[\d_]*)(?:e[+-]?\d[\d_]*)?[dfl]?/i, operator: { pattern: /(^|[^.])(?:<<=?|>>>?=?|->|--|\+\+|&&|\|\||::|[?:~]|[-+*/%&|^!=<>]=?)/m, lookbehind: !0 } }), e.languages.insertBefore("java", "string", { "triple-quoted-string": { pattern: /"""[ \t]*[\r\n](?:(?:"|"")?(?:\\.|[^"\\]))*"""/, greedy: !0, alias: "string" }, char: { pattern: /'(?:\\.|[^'\\\r\n]){1,6}'/, greedy: !0 } }), e.languages.insertBefore("java", "class-name", { annotation: { pattern: /(^|[^.])@\w+(?:\s*\.\s*\w+)*/, lookbehind: !0, alias: "punctuation" }, generics: { pattern: /<(?:[\w\s,.?]|&(?!&)|<(?:[\w\s,.?]|&(?!&)|<(?:[\w\s,.?]|&(?!&)|<(?:[\w\s,.?]|&(?!&))*>)*>)*>)*>/, inside: { "class-name": n, keyword: t, punctuation: /[<>(),.:]/, operator: /[?&|]/ } }, import: [{ pattern: RegExp(/(\bimport\s+)/.source + r + /(?:[A-Z]\w*|\*)(?=\s*;)/.source), lookbehind: !0, inside: { namespace: n.inside.namespace, punctuation: /\./, operator: /\*/, "class-name": /\w+/ } }, { pattern: RegExp(/(\bimport\s+static\s+)/.source + r + /(?:\w+|\*)(?=\s*;)/.source), lookbehind: !0, alias: "static", inside: { namespace: n.inside.namespace, static: /\b\w+$/, punctuation: /\./, operator: /\*/, "class-name": /\w+/ } }], namespace: { pattern: RegExp(/(\b(?:exports|import(?:\s+static)?|module|open|opens|package|provides|requires|to|transitive|uses|with)\s+)(?!<keyword>)[a-z]\w*(?:\.[a-z]\w*)*\.?/.source.replace(/<keyword>/g, function() {
    return t.source;
  })), lookbehind: !0, inside: { punctuation: /\./ } } });
}(Prism), Prism.languages.json = { property: { pattern: /(^|[^\\])"(?:\\.|[^\\"\r\n])*"(?=\s*:)/, lookbehind: !0, greedy: !0 }, string: { pattern: /(^|[^\\])"(?:\\.|[^\\"\r\n])*"(?!\s*:)/, lookbehind: !0, greedy: !0 }, comment: { pattern: /\/\/.*|\/\*[\s\S]*?(?:\*\/|$)/, greedy: !0 }, number: /-?\b\d+(?:\.\d+)?(?:e[+-]?\d+)?\b/i, punctuation: /[{}[\],]/, operator: /:/, boolean: /\b(?:false|true)\b/, null: { pattern: /\bnull\b/, alias: "keyword" } }, Prism.languages.webmanifest = Prism.languages.json, function(e) {
  var t = /("|')(?:\\(?:\r\n?|\n|.)|(?!\1)[^\\\r\n])*\1/;
  e.languages.json5 = e.languages.extend("json", { property: [{ pattern: RegExp(t.source + "(?=\\s*:)"), greedy: !0 }, { pattern: /(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?=\s*:)/, alias: "unquoted" }], string: { pattern: t, greedy: !0 }, number: /[+-]?\b(?:NaN|Infinity|0x[a-fA-F\d]+)\b|[+-]?(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:[eE][+-]?\d+\b)?/ });
}(Prism), Prism.languages.lua = { comment: /^#!.+|--(?:\[(=*)\[[\s\S]*?\]\1\]|.*)/m, string: { pattern: /(["'])(?:(?!\1)[^\\\r\n]|\\z(?:\r\n|\s)|\\(?:\r\n|[^z]))*\1|\[(=*)\[[\s\S]*?\]\2\]/, greedy: !0 }, number: /\b0x[a-f\d]+(?:\.[a-f\d]*)?(?:p[+-]?\d+)?\b|\b\d+(?:\.\B|(?:\.\d*)?(?:e[+-]?\d+)?\b)|\B\.\d+(?:e[+-]?\d+)?\b/i, keyword: /\b(?:and|break|do|else|elseif|end|false|for|function|goto|if|in|local|nil|not|or|repeat|return|then|true|until|while)\b/, function: /(?!\d)\w+(?=\s*(?:[({]))/, operator: [/[-+*%^&|#]|\/\/?|<[<=]?|>[>=]?|[=~]=?/, { pattern: /(^|[^.])\.\.(?!\.)/, lookbehind: !0 }], punctuation: /[\[\](){},;]|\.+|:+/ }, Prism.languages.matlab = { comment: [/%\{[\s\S]*?\}%/, /%.+/], string: { pattern: /\B'(?:''|[^'\r\n])*'/, greedy: !0 }, number: /(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:[eE][+-]?\d+)?(?:[ij])?|\b[ij]\b/, keyword: /\b(?:NaN|break|case|catch|continue|else|elseif|end|for|function|if|inf|otherwise|parfor|pause|pi|return|switch|try|while)\b/, function: /\b(?!\d)\w+(?=\s*\()/, operator: /\.?[*^\/\\']|[+\-:@]|[<>=~]=?|&&?|\|\|?/, punctuation: /\.{3}|[.,;\[\](){}!]/ }, function(e) {
  var t = ["$eq", "$gt", "$gte", "$in", "$lt", "$lte", "$ne", "$nin", "$and", "$not", "$nor", "$or", "$exists", "$type", "$expr", "$jsonSchema", "$mod", "$regex", "$text", "$where", "$geoIntersects", "$geoWithin", "$near", "$nearSphere", "$all", "$elemMatch", "$size", "$bitsAllClear", "$bitsAllSet", "$bitsAnyClear", "$bitsAnySet", "$comment", "$elemMatch", "$meta", "$slice", "$currentDate", "$inc", "$min", "$max", "$mul", "$rename", "$set", "$setOnInsert", "$unset", "$addToSet", "$pop", "$pull", "$push", "$pullAll", "$each", "$position", "$slice", "$sort", "$bit", "$addFields", "$bucket", "$bucketAuto", "$collStats", "$count", "$currentOp", "$facet", "$geoNear", "$graphLookup", "$group", "$indexStats", "$limit", "$listLocalSessions", "$listSessions", "$lookup", "$match", "$merge", "$out", "$planCacheStats", "$project", "$redact", "$replaceRoot", "$replaceWith", "$sample", "$set", "$skip", "$sort", "$sortByCount", "$unionWith", "$unset", "$unwind", "$setWindowFields", "$abs", "$accumulator", "$acos", "$acosh", "$add", "$addToSet", "$allElementsTrue", "$and", "$anyElementTrue", "$arrayElemAt", "$arrayToObject", "$asin", "$asinh", "$atan", "$atan2", "$atanh", "$avg", "$binarySize", "$bsonSize", "$ceil", "$cmp", "$concat", "$concatArrays", "$cond", "$convert", "$cos", "$dateFromParts", "$dateToParts", "$dateFromString", "$dateToString", "$dayOfMonth", "$dayOfWeek", "$dayOfYear", "$degreesToRadians", "$divide", "$eq", "$exp", "$filter", "$first", "$floor", "$function", "$gt", "$gte", "$hour", "$ifNull", "$in", "$indexOfArray", "$indexOfBytes", "$indexOfCP", "$isArray", "$isNumber", "$isoDayOfWeek", "$isoWeek", "$isoWeekYear", "$last", "$last", "$let", "$literal", "$ln", "$log", "$log10", "$lt", "$lte", "$ltrim", "$map", "$max", "$mergeObjects", "$meta", "$min", "$millisecond", "$minute", "$mod", "$month", "$multiply", "$ne", "$not", "$objectToArray", "$or", "$pow", "$push", "$radiansToDegrees", "$range", "$reduce", "$regexFind", "$regexFindAll", "$regexMatch", "$replaceOne", "$replaceAll", "$reverseArray", "$round", "$rtrim", "$second", "$setDifference", "$setEquals", "$setIntersection", "$setIsSubset", "$setUnion", "$size", "$sin", "$slice", "$split", "$sqrt", "$stdDevPop", "$stdDevSamp", "$strcasecmp", "$strLenBytes", "$strLenCP", "$substr", "$substrBytes", "$substrCP", "$subtract", "$sum", "$switch", "$tan", "$toBool", "$toDate", "$toDecimal", "$toDouble", "$toInt", "$toLong", "$toObjectId", "$toString", "$toLower", "$toUpper", "$trim", "$trunc", "$type", "$week", "$year", "$zip", "$count", "$dateAdd", "$dateDiff", "$dateSubtract", "$dateTrunc", "$getField", "$rand", "$sampleRate", "$setField", "$unsetField", "$comment", "$explain", "$hint", "$max", "$maxTimeMS", "$min", "$orderby", "$query", "$returnKey", "$showDiskLoc", "$natural"], r = "(?:" + (t = t.map(function(n) {
    return n.replace("$", "\\$");
  })).join("|") + ")\\b";
  e.languages.mongodb = e.languages.extend("javascript", {}), e.languages.insertBefore("mongodb", "string", { property: { pattern: /(?:(["'])(?:\\(?:\r\n|[\s\S])|(?!\1)[^\\\r\n])*\1|(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*)(?=\s*:)/, greedy: !0, inside: { keyword: RegExp(`^(['"])?` + r + "(?:\\1)?$") } } }), e.languages.mongodb.string.inside = { url: { pattern: /https?:\/\/[-\w@:%.+~#=]{1,256}\.[a-z0-9()]{1,6}\b[-\w()@:%+.~#?&/=]*/i, greedy: !0 }, entity: { pattern: /\b(?:(?:[01]?\d\d?|2[0-4]\d|25[0-5])\.){3}(?:[01]?\d\d?|2[0-4]\d|25[0-5])\b/, greedy: !0 } }, e.languages.insertBefore("mongodb", "constant", { builtin: { pattern: RegExp("\\b(?:" + ["ObjectId", "Code", "BinData", "DBRef", "Timestamp", "NumberLong", "NumberDecimal", "MaxKey", "MinKey", "RegExp", "ISODate", "UUID"].join("|") + ")\\b"), alias: "keyword" } });
}(Prism), function(e) {
  var t = /\$(?:\w[a-z\d]*(?:_[^\x00-\x1F\s"'\\()$]*)?|\{[^}\s"'\\]+\})/i;
  e.languages.nginx = { comment: { pattern: /(^|[\s{};])#.*/, lookbehind: !0, greedy: !0 }, directive: { pattern: /(^|\s)\w(?:[^;{}"'\\\s]|\\.|"(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*'|\s+(?:#.*(?!.)|(?![#\s])))*?(?=\s*[;{])/, lookbehind: !0, greedy: !0, inside: { string: { pattern: /((?:^|[^\\])(?:\\\\)*)(?:"(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*')/, lookbehind: !0, greedy: !0, inside: { escape: { pattern: /\\["'\\nrt]/, alias: "entity" }, variable: t } }, comment: { pattern: /(\s)#.*/, lookbehind: !0, greedy: !0 }, keyword: { pattern: /^\S+/, greedy: !0 }, boolean: { pattern: /(\s)(?:off|on)(?!\S)/, lookbehind: !0 }, number: { pattern: /(\s)\d+[a-z]*(?!\S)/i, lookbehind: !0 }, variable: t } }, punctuation: /[{};]/ };
}(Prism), Prism.languages.objectivec = Prism.languages.extend("c", { string: { pattern: /@?"(?:\\(?:\r\n|[\s\S])|[^"\\\r\n])*"/, greedy: !0 }, keyword: /\b(?:asm|auto|break|case|char|const|continue|default|do|double|else|enum|extern|float|for|goto|if|in|inline|int|long|register|return|self|short|signed|sizeof|static|struct|super|switch|typedef|typeof|union|unsigned|void|volatile|while)\b|(?:@interface|@end|@implementation|@protocol|@class|@public|@protected|@private|@property|@try|@catch|@finally|@throw|@synthesize|@dynamic|@selector)\b/, operator: /-[->]?|\+\+?|!=?|<<?=?|>>?=?|==?|&&?|\|\|?|[~^%?*\/@]/ }), delete Prism.languages.objectivec["class-name"], Prism.languages.objc = Prism.languages.objectivec, Prism.languages.pascal = { directive: { pattern: /\{\$[\s\S]*?\}/, greedy: !0, alias: ["marco", "property"] }, comment: { pattern: /\(\*[\s\S]*?\*\)|\{[\s\S]*?\}|\/\/.*/, greedy: !0 }, string: { pattern: /(?:'(?:''|[^'\r\n])*'(?!')|#[&$%]?[a-f\d]+)+|\^[a-z]/i, greedy: !0 }, asm: { pattern: /(\basm\b)[\s\S]+?(?=\bend\s*[;[])/i, lookbehind: !0, greedy: !0, inside: null }, keyword: [{ pattern: /(^|[^&])\b(?:absolute|array|asm|begin|case|const|constructor|destructor|do|downto|else|end|file|for|function|goto|if|implementation|inherited|inline|interface|label|nil|object|of|operator|packed|procedure|program|record|reintroduce|repeat|self|set|string|then|to|type|unit|until|uses|var|while|with)\b/i, lookbehind: !0 }, { pattern: /(^|[^&])\b(?:dispose|exit|false|new|true)\b/i, lookbehind: !0 }, { pattern: /(^|[^&])\b(?:class|dispinterface|except|exports|finalization|finally|initialization|inline|library|on|out|packed|property|raise|resourcestring|threadvar|try)\b/i, lookbehind: !0 }, { pattern: /(^|[^&])\b(?:absolute|abstract|alias|assembler|bitpacked|break|cdecl|continue|cppdecl|cvar|default|deprecated|dynamic|enumerator|experimental|export|external|far|far16|forward|generic|helper|implements|index|interrupt|iochecks|local|message|name|near|nodefault|noreturn|nostackframe|oldfpccall|otherwise|overload|override|pascal|platform|private|protected|public|published|read|register|reintroduce|result|safecall|saveregisters|softfloat|specialize|static|stdcall|stored|strict|unaligned|unimplemented|varargs|virtual|write)\b/i, lookbehind: !0 }], number: [/(?:[&%]\d+|\$[a-f\d]+)/i, /\b\d+(?:\.\d+)?(?:e[+-]?\d+)?/i], operator: [/\.\.|\*\*|:=|<[<=>]?|>[>=]?|[+\-*\/]=?|[@^=]/, { pattern: /(^|[^&])\b(?:and|as|div|exclude|in|include|is|mod|not|or|shl|shr|xor)\b/, lookbehind: !0 }], punctuation: /\(\.|\.\)|[()\[\]:;,.]/ }, Prism.languages.pascal.asm.inside = Prism.languages.extend("pascal", { asm: void 0, keyword: void 0, operator: void 0 }), Prism.languages.objectpascal = Prism.languages.pascal, function(e) {
  var t = /\/\*[\s\S]*?\*\/|\/\/.*|#(?!\[).*/, r = [{ pattern: /\b(?:false|true)\b/i, alias: "boolean" }, { pattern: /(::\s*)\b[a-z_]\w*\b(?!\s*\()/i, greedy: !0, lookbehind: !0 }, { pattern: /(\b(?:case|const)\s+)\b[a-z_]\w*(?=\s*[;=])/i, greedy: !0, lookbehind: !0 }, /\b(?:null)\b/i, /\b[A-Z_][A-Z0-9_]*\b(?!\s*\()/], n = /\b0b[01]+(?:_[01]+)*\b|\b0o[0-7]+(?:_[0-7]+)*\b|\b0x[\da-f]+(?:_[\da-f]+)*\b|(?:\b\d+(?:_\d+)*\.?(?:\d+(?:_\d+)*)?|\B\.\d+)(?:e[+-]?\d+)?/i, a = /<?=>|\?\?=?|\.{3}|\??->|[!=]=?=?|::|\*\*=?|--|\+\+|&&|\|\||<<|>>|[?~]|[/^|%*&<>.+-]=?/, i = /[{}\[\](),:;]/;
  e.languages.php = { delimiter: { pattern: /\?>$|^<\?(?:php(?=\s)|=)?/i, alias: "important" }, comment: t, variable: /\$+(?:\w+\b|(?=\{))/, package: { pattern: /(namespace\s+|use\s+(?:function\s+)?)(?:\\?\b[a-z_]\w*)+\b(?!\\)/i, lookbehind: !0, inside: { punctuation: /\\/ } }, "class-name-definition": { pattern: /(\b(?:class|enum|interface|trait)\s+)\b[a-z_]\w*(?!\\)\b/i, lookbehind: !0, alias: "class-name" }, "function-definition": { pattern: /(\bfunction\s+)[a-z_]\w*(?=\s*\()/i, lookbehind: !0, alias: "function" }, keyword: [{ pattern: /(\(\s*)\b(?:array|bool|boolean|float|int|integer|object|string)\b(?=\s*\))/i, alias: "type-casting", greedy: !0, lookbehind: !0 }, { pattern: /([(,?]\s*)\b(?:array(?!\s*\()|bool|callable|(?:false|null)(?=\s*\|)|float|int|iterable|mixed|object|self|static|string)\b(?=\s*\$)/i, alias: "type-hint", greedy: !0, lookbehind: !0 }, { pattern: /(\)\s*:\s*(?:\?\s*)?)\b(?:array(?!\s*\()|bool|callable|(?:false|null)(?=\s*\|)|float|int|iterable|mixed|never|object|self|static|string|void)\b/i, alias: "return-type", greedy: !0, lookbehind: !0 }, { pattern: /\b(?:array(?!\s*\()|bool|float|int|iterable|mixed|object|string|void)\b/i, alias: "type-declaration", greedy: !0 }, { pattern: /(\|\s*)(?:false|null)\b|\b(?:false|null)(?=\s*\|)/i, alias: "type-declaration", greedy: !0, lookbehind: !0 }, { pattern: /\b(?:parent|self|static)(?=\s*::)/i, alias: "static-context", greedy: !0 }, { pattern: /(\byield\s+)from\b/i, lookbehind: !0 }, /\bclass\b/i, { pattern: /((?:^|[^\s>:]|(?:^|[^-])>|(?:^|[^:]):)\s*)\b(?:abstract|and|array|as|break|callable|case|catch|clone|const|continue|declare|default|die|do|echo|else|elseif|empty|enddeclare|endfor|endforeach|endif|endswitch|endwhile|enum|eval|exit|extends|final|finally|fn|for|foreach|function|global|goto|if|implements|include|include_once|instanceof|insteadof|interface|isset|list|match|namespace|never|new|or|parent|print|private|protected|public|readonly|require|require_once|return|self|static|switch|throw|trait|try|unset|use|var|while|xor|yield|__halt_compiler)\b/i, lookbehind: !0 }], "argument-name": { pattern: /([(,]\s*)\b[a-z_]\w*(?=\s*:(?!:))/i, lookbehind: !0 }, "class-name": [{ pattern: /(\b(?:extends|implements|instanceof|new(?!\s+self|\s+static))\s+|\bcatch\s*\()\b[a-z_]\w*(?!\\)\b/i, greedy: !0, lookbehind: !0 }, { pattern: /(\|\s*)\b[a-z_]\w*(?!\\)\b/i, greedy: !0, lookbehind: !0 }, { pattern: /\b[a-z_]\w*(?!\\)\b(?=\s*\|)/i, greedy: !0 }, { pattern: /(\|\s*)(?:\\?\b[a-z_]\w*)+\b/i, alias: "class-name-fully-qualified", greedy: !0, lookbehind: !0, inside: { punctuation: /\\/ } }, { pattern: /(?:\\?\b[a-z_]\w*)+\b(?=\s*\|)/i, alias: "class-name-fully-qualified", greedy: !0, inside: { punctuation: /\\/ } }, { pattern: /(\b(?:extends|implements|instanceof|new(?!\s+self\b|\s+static\b))\s+|\bcatch\s*\()(?:\\?\b[a-z_]\w*)+\b(?!\\)/i, alias: "class-name-fully-qualified", greedy: !0, lookbehind: !0, inside: { punctuation: /\\/ } }, { pattern: /\b[a-z_]\w*(?=\s*\$)/i, alias: "type-declaration", greedy: !0 }, { pattern: /(?:\\?\b[a-z_]\w*)+(?=\s*\$)/i, alias: ["class-name-fully-qualified", "type-declaration"], greedy: !0, inside: { punctuation: /\\/ } }, { pattern: /\b[a-z_]\w*(?=\s*::)/i, alias: "static-context", greedy: !0 }, { pattern: /(?:\\?\b[a-z_]\w*)+(?=\s*::)/i, alias: ["class-name-fully-qualified", "static-context"], greedy: !0, inside: { punctuation: /\\/ } }, { pattern: /([(,?]\s*)[a-z_]\w*(?=\s*\$)/i, alias: "type-hint", greedy: !0, lookbehind: !0 }, { pattern: /([(,?]\s*)(?:\\?\b[a-z_]\w*)+(?=\s*\$)/i, alias: ["class-name-fully-qualified", "type-hint"], greedy: !0, lookbehind: !0, inside: { punctuation: /\\/ } }, { pattern: /(\)\s*:\s*(?:\?\s*)?)\b[a-z_]\w*(?!\\)\b/i, alias: "return-type", greedy: !0, lookbehind: !0 }, { pattern: /(\)\s*:\s*(?:\?\s*)?)(?:\\?\b[a-z_]\w*)+\b(?!\\)/i, alias: ["class-name-fully-qualified", "return-type"], greedy: !0, lookbehind: !0, inside: { punctuation: /\\/ } }], constant: r, function: { pattern: /(^|[^\\\w])\\?[a-z_](?:[\w\\]*\w)?(?=\s*\()/i, lookbehind: !0, inside: { punctuation: /\\/ } }, property: { pattern: /(->\s*)\w+/, lookbehind: !0 }, number: n, operator: a, punctuation: i };
  var o = { pattern: /\{\$(?:\{(?:\{[^{}]+\}|[^{}]+)\}|[^{}])+\}|(^|[^\\{])\$+(?:\w+(?:\[[^\r\n\[\]]+\]|->\w+)?)/, lookbehind: !0, inside: e.languages.php }, s = [{ pattern: /<<<'([^']+)'[\r\n](?:.*[\r\n])*?\1;/, alias: "nowdoc-string", greedy: !0, inside: { delimiter: { pattern: /^<<<'[^']+'|[a-z_]\w*;$/i, alias: "symbol", inside: { punctuation: /^<<<'?|[';]$/ } } } }, { pattern: /<<<(?:"([^"]+)"[\r\n](?:.*[\r\n])*?\1;|([a-z_]\w*)[\r\n](?:.*[\r\n])*?\2;)/i, alias: "heredoc-string", greedy: !0, inside: { delimiter: { pattern: /^<<<(?:"[^"]+"|[a-z_]\w*)|[a-z_]\w*;$/i, alias: "symbol", inside: { punctuation: /^<<<"?|[";]$/ } }, interpolation: o } }, { pattern: /`(?:\\[\s\S]|[^\\`])*`/, alias: "backtick-quoted-string", greedy: !0 }, { pattern: /'(?:\\[\s\S]|[^\\'])*'/, alias: "single-quoted-string", greedy: !0 }, { pattern: /"(?:\\[\s\S]|[^\\"])*"/, alias: "double-quoted-string", greedy: !0, inside: { interpolation: o } }];
  e.languages.insertBefore("php", "variable", { string: s, attribute: { pattern: /#\[(?:[^"'\/#]|\/(?![*/])|\/\/.*$|#(?!\[).*$|\/\*(?:[^*]|\*(?!\/))*\*\/|"(?:\\[\s\S]|[^\\"])*"|'(?:\\[\s\S]|[^\\'])*')+\](?=\s*[a-z$#])/im, greedy: !0, inside: { "attribute-content": { pattern: /^(#\[)[\s\S]+(?=\]$)/, lookbehind: !0, inside: { comment: t, string: s, "attribute-class-name": [{ pattern: /([^:]|^)\b[a-z_]\w*(?!\\)\b/i, alias: "class-name", greedy: !0, lookbehind: !0 }, { pattern: /([^:]|^)(?:\\?\b[a-z_]\w*)+/i, alias: ["class-name", "class-name-fully-qualified"], greedy: !0, lookbehind: !0, inside: { punctuation: /\\/ } }], constant: r, number: n, operator: a, punctuation: i } }, delimiter: { pattern: /^#\[|\]$/, alias: "punctuation" } } } }), e.hooks.add("before-tokenize", function(c) {
    /<\?/.test(c.code) && e.languages["markup-templating"].buildPlaceholders(c, "php", /<\?(?:[^"'/#]|\/(?![*/])|("|')(?:\\[\s\S]|(?!\1)[^\\])*\1|(?:\/\/|#(?!\[))(?:[^?\n\r]|\?(?!>))*(?=$|\?>|[\r\n])|#\[|\/\*(?:[^*]|\*(?!\/))*(?:\*\/|$))*?(?:\?>|$)/g);
  }), e.hooks.add("after-tokenize", function(c) {
    e.languages["markup-templating"].tokenizePlaceholders(c, "php");
  });
}(Prism), function(e) {
  var t = /\b(?:bool|bytes|double|s?fixed(?:32|64)|float|[su]?int(?:32|64)|string)\b/;
  e.languages.protobuf = e.languages.extend("clike", { "class-name": [{ pattern: /(\b(?:enum|extend|message|service)\s+)[A-Za-z_]\w*(?=\s*\{)/, lookbehind: !0 }, { pattern: /(\b(?:rpc\s+\w+|returns)\s*\(\s*(?:stream\s+)?)\.?[A-Za-z_]\w*(?:\.[A-Za-z_]\w*)*(?=\s*\))/, lookbehind: !0 }], keyword: /\b(?:enum|extend|extensions|import|message|oneof|option|optional|package|public|repeated|required|reserved|returns|rpc(?=\s+\w)|service|stream|syntax|to)\b(?!\s*=\s*\d)/, function: /\b[a-z_]\w*(?=\s*\()/i }), e.languages.insertBefore("protobuf", "operator", { map: { pattern: /\bmap<\s*[\w.]+\s*,\s*[\w.]+\s*>(?=\s+[a-z_]\w*\s*[=;])/i, alias: "class-name", inside: { punctuation: /[<>.,]/, builtin: t } }, builtin: t, "positional-class-name": { pattern: /(?:\b|\B\.)[a-z_]\w*(?:\.[a-z_]\w*)*(?=\s+[a-z_]\w*\s*[=;])/i, alias: "class-name", inside: { punctuation: /\./ } }, annotation: { pattern: /(\[\s*)[a-z_]\w*(?=\s*=)/i, lookbehind: !0 } });
}(Prism), Prism.languages.python = { comment: { pattern: /(^|[^\\])#.*/, lookbehind: !0, greedy: !0 }, "string-interpolation": { pattern: /(?:f|fr|rf)(?:("""|''')[\s\S]*?\1|("|')(?:\\.|(?!\2)[^\\\r\n])*\2)/i, greedy: !0, inside: { interpolation: { pattern: /((?:^|[^{])(?:\{\{)*)\{(?!\{)(?:[^{}]|\{(?!\{)(?:[^{}]|\{(?!\{)(?:[^{}])+\})+\})+\}/, lookbehind: !0, inside: { "format-spec": { pattern: /(:)[^:(){}]+(?=\}$)/, lookbehind: !0 }, "conversion-option": { pattern: /![sra](?=[:}]$)/, alias: "punctuation" }, rest: null } }, string: /[\s\S]+/ } }, "triple-quoted-string": { pattern: /(?:[rub]|br|rb)?("""|''')[\s\S]*?\1/i, greedy: !0, alias: "string" }, string: { pattern: /(?:[rub]|br|rb)?("|')(?:\\.|(?!\1)[^\\\r\n])*\1/i, greedy: !0 }, function: { pattern: /((?:^|\s)def[ \t]+)[a-zA-Z_]\w*(?=\s*\()/g, lookbehind: !0 }, "class-name": { pattern: /(\bclass\s+)\w+/i, lookbehind: !0 }, decorator: { pattern: /(^[\t ]*)@\w+(?:\.\w+)*/m, lookbehind: !0, alias: ["annotation", "punctuation"], inside: { punctuation: /\./ } }, keyword: /\b(?:_(?=\s*:)|and|as|assert|async|await|break|case|class|continue|def|del|elif|else|except|exec|finally|for|from|global|if|import|in|is|lambda|match|nonlocal|not|or|pass|print|raise|return|try|while|with|yield)\b/, builtin: /\b(?:__import__|abs|all|any|apply|ascii|basestring|bin|bool|buffer|bytearray|bytes|callable|chr|classmethod|cmp|coerce|compile|complex|delattr|dict|dir|divmod|enumerate|eval|execfile|file|filter|float|format|frozenset|getattr|globals|hasattr|hash|help|hex|id|input|int|intern|isinstance|issubclass|iter|len|list|locals|long|map|max|memoryview|min|next|object|oct|open|ord|pow|property|range|raw_input|reduce|reload|repr|reversed|round|set|setattr|slice|sorted|staticmethod|str|sum|super|tuple|type|unichr|unicode|vars|xrange|zip)\b/, boolean: /\b(?:False|None|True)\b/, number: /\b0(?:b(?:_?[01])+|o(?:_?[0-7])+|x(?:_?[a-f0-9])+)\b|(?:\b\d+(?:_\d+)*(?:\.(?:\d+(?:_\d+)*)?)?|\B\.\d+(?:_\d+)*)(?:e[+-]?\d+(?:_\d+)*)?j?(?!\w)/i, operator: /[-+%=]=?|!=|:=|\*\*?=?|\/\/?=?|<[<=>]?|>[=>]?|[&|^~]/, punctuation: /[{}[\];(),.:]/ }, Prism.languages.python["string-interpolation"].inside.interpolation.inside.rest = Prism.languages.python, Prism.languages.py = Prism.languages.python, Prism.languages.r = { comment: /#.*/, string: { pattern: /(['"])(?:\\.|(?!\1)[^\\\r\n])*\1/, greedy: !0 }, "percent-operator": { pattern: /%[^%\s]*%/, alias: "operator" }, boolean: /\b(?:FALSE|TRUE)\b/, ellipsis: /\.\.(?:\.|\d+)/, number: [/\b(?:Inf|NaN)\b/, /(?:\b0x[\dA-Fa-f]+(?:\.\d*)?|\b\d+(?:\.\d*)?|\B\.\d+)(?:[EePp][+-]?\d+)?[iL]?/], keyword: /\b(?:NA|NA_character_|NA_complex_|NA_integer_|NA_real_|NULL|break|else|for|function|if|in|next|repeat|while)\b/, operator: /->?>?|<(?:=|<?-)?|[>=!]=?|::?|&&?|\|\|?|[+*\/^$@~]/, punctuation: /[(){}\[\],;]/ }, function(e) {
  for (var t = /\/\*(?:[^*/]|\*(?!\/)|\/(?!\*)|<self>)*\*\//.source, r = 0; r < 2; r++)
    t = t.replace(/<self>/g, function() {
      return t;
    });
  t = t.replace(/<self>/g, function() {
    return /[^\s\S]/.source;
  }), e.languages.rust = { comment: [{ pattern: RegExp(/(^|[^\\])/.source + t), lookbehind: !0, greedy: !0 }, { pattern: /(^|[^\\:])\/\/.*/, lookbehind: !0, greedy: !0 }], string: { pattern: /b?"(?:\\[\s\S]|[^\\"])*"|b?r(#*)"(?:[^"]|"(?!\1))*"\1/, greedy: !0 }, char: { pattern: /b?'(?:\\(?:x[0-7][\da-fA-F]|u\{(?:[\da-fA-F]_*){1,6}\}|.)|[^\\\r\n\t'])'/, greedy: !0 }, attribute: { pattern: /#!?\[(?:[^\[\]"]|"(?:\\[\s\S]|[^\\"])*")*\]/, greedy: !0, alias: "attr-name", inside: { string: null } }, "closure-params": { pattern: /([=(,:]\s*|\bmove\s*)\|[^|]*\||\|[^|]*\|(?=\s*(?:\{|->))/, lookbehind: !0, greedy: !0, inside: { "closure-punctuation": { pattern: /^\||\|$/, alias: "punctuation" }, rest: null } }, "lifetime-annotation": { pattern: /'\w+/, alias: "symbol" }, "fragment-specifier": { pattern: /(\$\w+:)[a-z]+/, lookbehind: !0, alias: "punctuation" }, variable: /\$\w+/, "function-definition": { pattern: /(\bfn\s+)\w+/, lookbehind: !0, alias: "function" }, "type-definition": { pattern: /(\b(?:enum|struct|trait|type|union)\s+)\w+/, lookbehind: !0, alias: "class-name" }, "module-declaration": [{ pattern: /(\b(?:crate|mod)\s+)[a-z][a-z_\d]*/, lookbehind: !0, alias: "namespace" }, { pattern: /(\b(?:crate|self|super)\s*)::\s*[a-z][a-z_\d]*\b(?:\s*::(?:\s*[a-z][a-z_\d]*\s*::)*)?/, lookbehind: !0, alias: "namespace", inside: { punctuation: /::/ } }], keyword: [/\b(?:Self|abstract|as|async|await|become|box|break|const|continue|crate|do|dyn|else|enum|extern|final|fn|for|if|impl|in|let|loop|macro|match|mod|move|mut|override|priv|pub|ref|return|self|static|struct|super|trait|try|type|typeof|union|unsafe|unsized|use|virtual|where|while|yield)\b/, /\b(?:bool|char|f(?:32|64)|[ui](?:8|16|32|64|128|size)|str)\b/], function: /\b[a-z_]\w*(?=\s*(?:::\s*<|\())/, macro: { pattern: /\b\w+!/, alias: "property" }, constant: /\b[A-Z_][A-Z_\d]+\b/, "class-name": /\b[A-Z]\w*\b/, namespace: { pattern: /(?:\b[a-z][a-z_\d]*\s*::\s*)*\b[a-z][a-z_\d]*\s*::(?!\s*<)/, inside: { punctuation: /::/ } }, number: /\b(?:0x[\dA-Fa-f](?:_?[\dA-Fa-f])*|0o[0-7](?:_?[0-7])*|0b[01](?:_?[01])*|(?:(?:\d(?:_?\d)*)?\.)?\d(?:_?\d)*(?:[Ee][+-]?\d+)?)(?:_?(?:f32|f64|[iu](?:8|16|32|64|size)?))?\b/, boolean: /\b(?:false|true)\b/, punctuation: /->|\.\.=|\.{1,3}|::|[{}[\];(),:]/, operator: /[-+*\/%!^]=?|=[=>]?|&[&=]?|\|[|=]?|<<?=?|>>?=?|[@?]/ }, e.languages.rust["closure-params"].inside.rest = e.languages.rust, e.languages.rust.attribute.inside.string = e.languages.rust.string;
}(Prism), Prism.languages.sql = { comment: { pattern: /(^|[^\\])(?:\/\*[\s\S]*?\*\/|(?:--|\/\/|#).*)/, lookbehind: !0 }, variable: [{ pattern: /@(["'`])(?:\\[\s\S]|(?!\1)[^\\])+\1/, greedy: !0 }, /@[\w.$]+/], string: { pattern: /(^|[^@\\])("|')(?:\\[\s\S]|(?!\2)[^\\]|\2\2)*\2/, greedy: !0, lookbehind: !0 }, identifier: { pattern: /(^|[^@\\])`(?:\\[\s\S]|[^`\\]|``)*`/, greedy: !0, lookbehind: !0, inside: { punctuation: /^`|`$/ } }, function: /\b(?:AVG|COUNT|FIRST|FORMAT|LAST|LCASE|LEN|MAX|MID|MIN|MOD|NOW|ROUND|SUM|UCASE)(?=\s*\()/i, keyword: /\b(?:ACTION|ADD|AFTER|ALGORITHM|ALL|ALTER|ANALYZE|ANY|APPLY|AS|ASC|AUTHORIZATION|AUTO_INCREMENT|BACKUP|BDB|BEGIN|BERKELEYDB|BIGINT|BINARY|BIT|BLOB|BOOL|BOOLEAN|BREAK|BROWSE|BTREE|BULK|BY|CALL|CASCADED?|CASE|CHAIN|CHAR(?:ACTER|SET)?|CHECK(?:POINT)?|CLOSE|CLUSTERED|COALESCE|COLLATE|COLUMNS?|COMMENT|COMMIT(?:TED)?|COMPUTE|CONNECT|CONSISTENT|CONSTRAINT|CONTAINS(?:TABLE)?|CONTINUE|CONVERT|CREATE|CROSS|CURRENT(?:_DATE|_TIME|_TIMESTAMP|_USER)?|CURSOR|CYCLE|DATA(?:BASES?)?|DATE(?:TIME)?|DAY|DBCC|DEALLOCATE|DEC|DECIMAL|DECLARE|DEFAULT|DEFINER|DELAYED|DELETE|DELIMITERS?|DENY|DESC|DESCRIBE|DETERMINISTIC|DISABLE|DISCARD|DISK|DISTINCT|DISTINCTROW|DISTRIBUTED|DO|DOUBLE|DROP|DUMMY|DUMP(?:FILE)?|DUPLICATE|ELSE(?:IF)?|ENABLE|ENCLOSED|END|ENGINE|ENUM|ERRLVL|ERRORS|ESCAPED?|EXCEPT|EXEC(?:UTE)?|EXISTS|EXIT|EXPLAIN|EXTENDED|FETCH|FIELDS|FILE|FILLFACTOR|FIRST|FIXED|FLOAT|FOLLOWING|FOR(?: EACH ROW)?|FORCE|FOREIGN|FREETEXT(?:TABLE)?|FROM|FULL|FUNCTION|GEOMETRY(?:COLLECTION)?|GLOBAL|GOTO|GRANT|GROUP|HANDLER|HASH|HAVING|HOLDLOCK|HOUR|IDENTITY(?:COL|_INSERT)?|IF|IGNORE|IMPORT|INDEX|INFILE|INNER|INNODB|INOUT|INSERT|INT|INTEGER|INTERSECT|INTERVAL|INTO|INVOKER|ISOLATION|ITERATE|JOIN|KEYS?|KILL|LANGUAGE|LAST|LEAVE|LEFT|LEVEL|LIMIT|LINENO|LINES|LINESTRING|LOAD|LOCAL|LOCK|LONG(?:BLOB|TEXT)|LOOP|MATCH(?:ED)?|MEDIUM(?:BLOB|INT|TEXT)|MERGE|MIDDLEINT|MINUTE|MODE|MODIFIES|MODIFY|MONTH|MULTI(?:LINESTRING|POINT|POLYGON)|NATIONAL|NATURAL|NCHAR|NEXT|NO|NONCLUSTERED|NULLIF|NUMERIC|OFF?|OFFSETS?|ON|OPEN(?:DATASOURCE|QUERY|ROWSET)?|OPTIMIZE|OPTION(?:ALLY)?|ORDER|OUT(?:ER|FILE)?|OVER|PARTIAL|PARTITION|PERCENT|PIVOT|PLAN|POINT|POLYGON|PRECEDING|PRECISION|PREPARE|PREV|PRIMARY|PRINT|PRIVILEGES|PROC(?:EDURE)?|PUBLIC|PURGE|QUICK|RAISERROR|READS?|REAL|RECONFIGURE|REFERENCES|RELEASE|RENAME|REPEAT(?:ABLE)?|REPLACE|REPLICATION|REQUIRE|RESIGNAL|RESTORE|RESTRICT|RETURN(?:ING|S)?|REVOKE|RIGHT|ROLLBACK|ROUTINE|ROW(?:COUNT|GUIDCOL|S)?|RTREE|RULE|SAVE(?:POINT)?|SCHEMA|SECOND|SELECT|SERIAL(?:IZABLE)?|SESSION(?:_USER)?|SET(?:USER)?|SHARE|SHOW|SHUTDOWN|SIMPLE|SMALLINT|SNAPSHOT|SOME|SONAME|SQL|START(?:ING)?|STATISTICS|STATUS|STRIPED|SYSTEM_USER|TABLES?|TABLESPACE|TEMP(?:ORARY|TABLE)?|TERMINATED|TEXT(?:SIZE)?|THEN|TIME(?:STAMP)?|TINY(?:BLOB|INT|TEXT)|TOP?|TRAN(?:SACTIONS?)?|TRIGGER|TRUNCATE|TSEQUAL|TYPES?|UNBOUNDED|UNCOMMITTED|UNDEFINED|UNION|UNIQUE|UNLOCK|UNPIVOT|UNSIGNED|UPDATE(?:TEXT)?|USAGE|USE|USER|USING|VALUES?|VAR(?:BINARY|CHAR|CHARACTER|YING)|VIEW|WAITFOR|WARNINGS|WHEN|WHERE|WHILE|WITH(?: ROLLUP|IN)?|WORK|WRITE(?:TEXT)?|YEAR)\b/i, boolean: /\b(?:FALSE|NULL|TRUE)\b/i, number: /\b0x[\da-f]+\b|\b\d+(?:\.\d*)?|\B\.\d+\b/i, operator: /[-+*\/=%^~]|&&?|\|\|?|!=?|<(?:=>?|<|>)?|>[>=]?|\b(?:AND|BETWEEN|DIV|ILIKE|IN|IS|LIKE|NOT|OR|REGEXP|RLIKE|SOUNDS LIKE|XOR)\b/i, punctuation: /[;[\]()`,.]/ }, function(e) {
  e.languages.typescript = e.languages.extend("javascript", { "class-name": { pattern: /(\b(?:class|extends|implements|instanceof|interface|new|type)\s+)(?!keyof\b)(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*(?:\s*<(?:[^<>]|<(?:[^<>]|<[^<>]*>)*>)*>)?/, lookbehind: !0, greedy: !0, inside: null }, builtin: /\b(?:Array|Function|Promise|any|boolean|console|never|number|string|symbol|unknown)\b/ }), e.languages.typescript.keyword.push(/\b(?:abstract|declare|is|keyof|readonly|require)\b/, /\b(?:asserts|infer|interface|module|namespace|type)\b(?=\s*(?:[{_$a-zA-Z\xA0-\uFFFF]|$))/, /\btype\b(?=\s*(?:[\{*]|$))/), delete e.languages.typescript.parameter, delete e.languages.typescript["literal-property"];
  var t = e.languages.extend("typescript", {});
  delete t["class-name"], e.languages.typescript["class-name"].inside = t, e.languages.insertBefore("typescript", "function", { decorator: { pattern: /@[$\w\xA0-\uFFFF]+/, inside: { at: { pattern: /^@/, alias: "operator" }, function: /^[\s\S]+/ } }, "generic-function": { pattern: /#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*\s*<(?:[^<>]|<(?:[^<>]|<[^<>]*>)*>)*>(?=\s*\()/, greedy: !0, inside: { function: /^#?(?!\s)[_$a-zA-Z\xA0-\uFFFF](?:(?!\s)[$\w\xA0-\uFFFF])*/, generic: { pattern: /<[\s\S]+/, alias: "class-name", inside: t } } } }), e.languages.ts = e.languages.typescript;
}(Prism), function(e) {
  var t = e.util.clone(e.languages.javascript), r = /(?:\s|\/\/.*(?!.)|\/\*(?:[^*]|\*(?!\/))\*\/)/.source, n = /(?:\{(?:\{(?:\{[^{}]*\}|[^{}])*\}|[^{}])*\})/.source, a = /(?:\{<S>*\.{3}(?:[^{}]|<BRACES>)*\})/.source;
  function i(c, u) {
    return c = c.replace(/<S>/g, function() {
      return r;
    }).replace(/<BRACES>/g, function() {
      return n;
    }).replace(/<SPREAD>/g, function() {
      return a;
    }), RegExp(c, u);
  }
  a = i(a).source, e.languages.jsx = e.languages.extend("markup", t), e.languages.jsx.tag.pattern = i(/<\/?(?:[\w.:-]+(?:<S>+(?:[\w.:$-]+(?:=(?:"(?:\\[\s\S]|[^\\"])*"|'(?:\\[\s\S]|[^\\'])*'|[^\s{'"/>=]+|<BRACES>))?|<SPREAD>))*<S>*\/?)?>/.source), e.languages.jsx.tag.inside.tag.pattern = /^<\/?[^\s>\/]*/, e.languages.jsx.tag.inside["attr-value"].pattern = /=(?!\{)(?:"(?:\\[\s\S]|[^\\"])*"|'(?:\\[\s\S]|[^\\'])*'|[^\s'">]+)/, e.languages.jsx.tag.inside.tag.inside["class-name"] = /^[A-Z]\w*(?:\.[A-Z]\w*)*$/, e.languages.jsx.tag.inside.comment = t.comment, e.languages.insertBefore("inside", "attr-name", { spread: { pattern: i(/<SPREAD>/.source), inside: e.languages.jsx } }, e.languages.jsx.tag), e.languages.insertBefore("inside", "special-attr", { script: { pattern: i(/=<BRACES>/.source), alias: "language-javascript", inside: { "script-punctuation": { pattern: /^=(?=\{)/, alias: "punctuation" }, rest: e.languages.jsx } } }, e.languages.jsx.tag);
  var o = function(c) {
    return c ? typeof c == "string" ? c : typeof c.content == "string" ? c.content : c.content.map(o).join("") : "";
  }, s = function(c) {
    for (var u = [], l = 0; l < c.length; l++) {
      var f = c[l], p = !1;
      if (typeof f != "string" && (f.type === "tag" && f.content[0] && f.content[0].type === "tag" ? f.content[0].content[0].content === "</" ? u.length > 0 && u[u.length - 1].tagName === o(f.content[0].content[1]) && u.pop() : f.content[f.content.length - 1].content === "/>" || u.push({ tagName: o(f.content[0].content[1]), openedBraces: 0 }) : u.length > 0 && f.type === "punctuation" && f.content === "{" ? u[u.length - 1].openedBraces++ : u.length > 0 && u[u.length - 1].openedBraces > 0 && f.type === "punctuation" && f.content === "}" ? u[u.length - 1].openedBraces-- : p = !0), (p || typeof f == "string") && u.length > 0 && u[u.length - 1].openedBraces === 0) {
        var d = o(f);
        l < c.length - 1 && (typeof c[l + 1] == "string" || c[l + 1].type === "plain-text") && (d += o(c[l + 1]), c.splice(l + 1, 1)), l > 0 && (typeof c[l - 1] == "string" || c[l - 1].type === "plain-text") && (d = o(c[l - 1]) + d, c.splice(l - 1, 1), l--), c[l] = new e.Token("plain-text", d, null, d);
      }
      f.content && typeof f.content != "string" && s(f.content);
    }
  };
  e.hooks.add("after-tokenize", function(c) {
    c.language !== "jsx" && c.language !== "tsx" || s(c.tokens);
  });
}(Prism), function(e) {
  var t = e.util.clone(e.languages.typescript);
  e.languages.tsx = e.languages.extend("jsx", t), delete e.languages.tsx.parameter, delete e.languages.tsx["literal-property"];
  var r = e.languages.tsx.tag;
  r.pattern = RegExp(/(^|[^\w$]|(?=<\/))/.source + "(?:" + r.pattern.source + ")", r.pattern.flags), r.lookbehind = !0;
}(Prism), function(e) {
  e.languages.sass = e.languages.extend("css", { comment: { pattern: /^([ \t]*)\/[\/*].*(?:(?:\r?\n|\r)\1[ \t].+)*/m, lookbehind: !0, greedy: !0 } }), e.languages.insertBefore("sass", "atrule", { "atrule-line": { pattern: /^(?:[ \t]*)[@+=].+/m, greedy: !0, inside: { atrule: /(?:@[\w-]+|[+=])/ } } }), delete e.languages.sass.atrule;
  var t = /\$[-\w]+|#\{\$[-\w]+\}/, r = [/[+*\/%]|[=!]=|<=?|>=?|\b(?:and|not|or)\b/, { pattern: /(\s)-(?=\s)/, lookbehind: !0 }];
  e.languages.insertBefore("sass", "property", { "variable-line": { pattern: /^[ \t]*\$.+/m, greedy: !0, inside: { punctuation: /:/, variable: t, operator: r } }, "property-line": { pattern: /^[ \t]*(?:[^:\s]+ *:.*|:[^:\s].*)/m, greedy: !0, inside: { property: [/[^:\s]+(?=\s*:)/, { pattern: /(:)[^:\s]+/, lookbehind: !0 }], punctuation: /:/, variable: t, operator: r, important: e.languages.sass.important } } }), delete e.languages.sass.property, delete e.languages.sass.important, e.languages.insertBefore("sass", "punctuation", { selector: { pattern: /^([ \t]*)\S(?:,[^,\r\n]+|[^,\r\n]*)(?:,[^,\r\n]+)*(?:,(?:\r?\n|\r)\1[ \t]+\S(?:,[^,\r\n]+|[^,\r\n]*)(?:,[^,\r\n]+)*)*/m, lookbehind: !0, greedy: !0 } });
}(Prism), function(e) {
  var t = "\\b(?:BASH|BASHOPTS|BASH_ALIASES|BASH_ARGC|BASH_ARGV|BASH_CMDS|BASH_COMPLETION_COMPAT_DIR|BASH_LINENO|BASH_REMATCH|BASH_SOURCE|BASH_VERSINFO|BASH_VERSION|COLORTERM|COLUMNS|COMP_WORDBREAKS|DBUS_SESSION_BUS_ADDRESS|DEFAULTS_PATH|DESKTOP_SESSION|DIRSTACK|DISPLAY|EUID|GDMSESSION|GDM_LANG|GNOME_KEYRING_CONTROL|GNOME_KEYRING_PID|GPG_AGENT_INFO|GROUPS|HISTCONTROL|HISTFILE|HISTFILESIZE|HISTSIZE|HOME|HOSTNAME|HOSTTYPE|IFS|INSTANCE|JOB|LANG|LANGUAGE|LC_ADDRESS|LC_ALL|LC_IDENTIFICATION|LC_MEASUREMENT|LC_MONETARY|LC_NAME|LC_NUMERIC|LC_PAPER|LC_TELEPHONE|LC_TIME|LESSCLOSE|LESSOPEN|LINES|LOGNAME|LS_COLORS|MACHTYPE|MAILCHECK|MANDATORY_PATH|NO_AT_BRIDGE|OLDPWD|OPTERR|OPTIND|ORBIT_SOCKETDIR|OSTYPE|PAPERSIZE|PATH|PIPESTATUS|PPID|PS1|PS2|PS3|PS4|PWD|RANDOM|REPLY|SECONDS|SELINUX_INIT|SESSION|SESSIONTYPE|SESSION_MANAGER|SHELL|SHELLOPTS|SHLVL|SSH_AUTH_SOCK|TERM|UID|UPSTART_EVENTS|UPSTART_INSTANCE|UPSTART_JOB|UPSTART_SESSION|USER|WINDOWID|XAUTHORITY|XDG_CONFIG_DIRS|XDG_CURRENT_DESKTOP|XDG_DATA_DIRS|XDG_GREETER_DATA_DIR|XDG_MENU_PREFIX|XDG_RUNTIME_DIR|XDG_SEAT|XDG_SEAT_PATH|XDG_SESSION_DESKTOP|XDG_SESSION_ID|XDG_SESSION_PATH|XDG_SESSION_TYPE|XDG_VTNR|XMODIFIERS)\\b", r = { pattern: /(^(["']?)\w+\2)[ \t]+\S.*/, lookbehind: !0, alias: "punctuation", inside: null }, n = { bash: r, environment: { pattern: RegExp("\\$" + t), alias: "constant" }, variable: [{ pattern: /\$?\(\([\s\S]+?\)\)/, greedy: !0, inside: { variable: [{ pattern: /(^\$\(\([\s\S]+)\)\)/, lookbehind: !0 }, /^\$\(\(/], number: /\b0x[\dA-Fa-f]+\b|(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:[Ee]-?\d+)?/, operator: /--|\+\+|\*\*=?|<<=?|>>=?|&&|\|\||[=!+\-*/%<>^&|]=?|[?~:]/, punctuation: /\(\(?|\)\)?|,|;/ } }, { pattern: /\$\((?:\([^)]+\)|[^()])+\)|`[^`]+`/, greedy: !0, inside: { variable: /^\$\(|^`|\)$|`$/ } }, { pattern: /\$\{[^}]+\}/, greedy: !0, inside: { operator: /:[-=?+]?|[!\/]|##?|%%?|\^\^?|,,?/, punctuation: /[\[\]]/, environment: { pattern: RegExp("(\\{)" + t), lookbehind: !0, alias: "constant" } } }, /\$(?:\w+|[#?*!@$])/], entity: /\\(?:[abceEfnrtv\\"]|O?[0-7]{1,3}|U[0-9a-fA-F]{8}|u[0-9a-fA-F]{4}|x[0-9a-fA-F]{1,2})/ };
  e.languages.bash = { shebang: { pattern: /^#!\s*\/.*/, alias: "important" }, comment: { pattern: /(^|[^"{\\$])#.*/, lookbehind: !0 }, "function-name": [{ pattern: /(\bfunction\s+)[\w-]+(?=(?:\s*\(?:\s*\))?\s*\{)/, lookbehind: !0, alias: "function" }, { pattern: /\b[\w-]+(?=\s*\(\s*\)\s*\{)/, alias: "function" }], "for-or-select": { pattern: /(\b(?:for|select)\s+)\w+(?=\s+in\s)/, alias: "variable", lookbehind: !0 }, "assign-left": { pattern: /(^|[\s;|&]|[<>]\()\w+(?=\+?=)/, inside: { environment: { pattern: RegExp("(^|[\\s;|&]|[<>]\\()" + t), lookbehind: !0, alias: "constant" } }, alias: "variable", lookbehind: !0 }, string: [{ pattern: /((?:^|[^<])<<-?\s*)(\w+)\s[\s\S]*?(?:\r?\n|\r)\2/, lookbehind: !0, greedy: !0, inside: n }, { pattern: /((?:^|[^<])<<-?\s*)(["'])(\w+)\2\s[\s\S]*?(?:\r?\n|\r)\3/, lookbehind: !0, greedy: !0, inside: { bash: r } }, { pattern: /(^|[^\\](?:\\\\)*)"(?:\\[\s\S]|\$\([^)]+\)|\$(?!\()|`[^`]+`|[^"\\`$])*"/, lookbehind: !0, greedy: !0, inside: n }, { pattern: /(^|[^$\\])'[^']*'/, lookbehind: !0, greedy: !0 }, { pattern: /\$'(?:[^'\\]|\\[\s\S])*'/, greedy: !0, inside: { entity: n.entity } }], environment: { pattern: RegExp("\\$?" + t), alias: "constant" }, variable: n.variable, function: { pattern: /(^|[\s;|&]|[<>]\()(?:add|apropos|apt|apt-cache|apt-get|aptitude|aspell|automysqlbackup|awk|basename|bash|bc|bconsole|bg|bzip2|cal|cat|cfdisk|chgrp|chkconfig|chmod|chown|chroot|cksum|clear|cmp|column|comm|composer|cp|cron|crontab|csplit|curl|cut|date|dc|dd|ddrescue|debootstrap|df|diff|diff3|dig|dir|dircolors|dirname|dirs|dmesg|docker|docker-compose|du|egrep|eject|env|ethtool|expand|expect|expr|fdformat|fdisk|fg|fgrep|file|find|fmt|fold|format|free|fsck|ftp|fuser|gawk|git|gparted|grep|groupadd|groupdel|groupmod|groups|grub-mkconfig|gzip|halt|head|hg|history|host|hostname|htop|iconv|id|ifconfig|ifdown|ifup|import|install|ip|jobs|join|kill|killall|less|link|ln|locate|logname|logrotate|look|lpc|lpr|lprint|lprintd|lprintq|lprm|ls|lsof|lynx|make|man|mc|mdadm|mkconfig|mkdir|mke2fs|mkfifo|mkfs|mkisofs|mknod|mkswap|mmv|more|most|mount|mtools|mtr|mutt|mv|nano|nc|netstat|nice|nl|node|nohup|notify-send|npm|nslookup|op|open|parted|passwd|paste|pathchk|ping|pkill|pnpm|podman|podman-compose|popd|pr|printcap|printenv|ps|pushd|pv|quota|quotacheck|quotactl|ram|rar|rcp|reboot|remsync|rename|renice|rev|rm|rmdir|rpm|rsync|scp|screen|sdiff|sed|sendmail|seq|service|sftp|sh|shellcheck|shuf|shutdown|sleep|slocate|sort|split|ssh|stat|strace|su|sudo|sum|suspend|swapon|sync|tac|tail|tar|tee|time|timeout|top|touch|tr|traceroute|tsort|tty|umount|uname|unexpand|uniq|units|unrar|unshar|unzip|update-grub|uptime|useradd|userdel|usermod|users|uudecode|uuencode|v|vcpkg|vdir|vi|vim|virsh|vmstat|wait|watch|wc|wget|whereis|which|who|whoami|write|xargs|xdg-open|yarn|yes|zenity|zip|zsh|zypper)(?=$|[)\s;|&])/, lookbehind: !0 }, keyword: { pattern: /(^|[\s;|&]|[<>]\()(?:case|do|done|elif|else|esac|fi|for|function|if|in|select|then|until|while)(?=$|[)\s;|&])/, lookbehind: !0 }, builtin: { pattern: /(^|[\s;|&]|[<>]\()(?:\.|:|alias|bind|break|builtin|caller|cd|command|continue|declare|echo|enable|eval|exec|exit|export|getopts|hash|help|let|local|logout|mapfile|printf|pwd|read|readarray|readonly|return|set|shift|shopt|source|test|times|trap|type|typeset|ulimit|umask|unalias|unset)(?=$|[)\s;|&])/, lookbehind: !0, alias: "class-name" }, boolean: { pattern: /(^|[\s;|&]|[<>]\()(?:false|true)(?=$|[)\s;|&])/, lookbehind: !0 }, "file-descriptor": { pattern: /\B&\d\b/, alias: "important" }, operator: { pattern: /\d?<>|>\||\+=|=[=~]?|!=?|<<[<-]?|[&\d]?>>|\d[<>]&?|[<>][&=]?|&[>&]?|\|[&|]?/, inside: { "file-descriptor": { pattern: /^\d/, alias: "important" } } }, punctuation: /\$?\(\(?|\)\)?|\.\.|[{}[\];\\]/, number: { pattern: /(^|\s)(?:[1-9]\d*|0)(?:[.,]\d+)?\b/, lookbehind: !0 } }, r.inside = e.languages.bash;
  for (var a = ["comment", "function-name", "for-or-select", "assign-left", "string", "environment", "function", "keyword", "builtin", "boolean", "file-descriptor", "operator", "punctuation", "number"], i = n.variable[1].inside, o = 0; o < a.length; o++)
    i[a[o]] = e.languages.bash[a[o]];
  e.languages.shell = e.languages.bash;
}(Prism), Prism.languages.swift = { comment: { pattern: /(^|[^\\:])(?:\/\/.*|\/\*(?:[^/*]|\/(?!\*)|\*(?!\/)|\/\*(?:[^*]|\*(?!\/))*\*\/)*\*\/)/, lookbehind: !0, greedy: !0 }, "string-literal": [{ pattern: RegExp(/(^|[^"#])/.source + "(?:" + /"(?:\\(?:\((?:[^()]|\([^()]*\))*\)|\r\n|[^(])|[^\\\r\n"])*"/.source + "|" + /"""(?:\\(?:\((?:[^()]|\([^()]*\))*\)|[^(])|[^\\"]|"(?!""))*"""/.source + ")" + /(?!["#])/.source), lookbehind: !0, greedy: !0, inside: { interpolation: { pattern: /(\\\()(?:[^()]|\([^()]*\))*(?=\))/, lookbehind: !0, inside: null }, "interpolation-punctuation": { pattern: /^\)|\\\($/, alias: "punctuation" }, punctuation: /\\(?=[\r\n])/, string: /[\s\S]+/ } }, { pattern: RegExp(/(^|[^"#])(#+)/.source + "(?:" + /"(?:\\(?:#+\((?:[^()]|\([^()]*\))*\)|\r\n|[^#])|[^\\\r\n])*?"/.source + "|" + /"""(?:\\(?:#+\((?:[^()]|\([^()]*\))*\)|[^#])|[^\\])*?"""/.source + ")\\2"), lookbehind: !0, greedy: !0, inside: { interpolation: { pattern: /(\\#+\()(?:[^()]|\([^()]*\))*(?=\))/, lookbehind: !0, inside: null }, "interpolation-punctuation": { pattern: /^\)|\\#+\($/, alias: "punctuation" }, string: /[\s\S]+/ } }], directive: { pattern: RegExp(/#/.source + "(?:" + /(?:elseif|if)\b/.source + "(?:[ 	]*" + /(?:![ \t]*)?(?:\b\w+\b(?:[ \t]*\((?:[^()]|\([^()]*\))*\))?|\((?:[^()]|\([^()]*\))*\))(?:[ \t]*(?:&&|\|\|))?/.source + ")+|" + /(?:else|endif)\b/.source + ")"), alias: "property", inside: { "directive-name": /^#\w+/, boolean: /\b(?:false|true)\b/, number: /\b\d+(?:\.\d+)*\b/, operator: /!|&&|\|\||[<>]=?/, punctuation: /[(),]/ } }, literal: { pattern: /#(?:colorLiteral|column|dsohandle|file(?:ID|Literal|Path)?|function|imageLiteral|line)\b/, alias: "constant" }, "other-directive": { pattern: /#\w+\b/, alias: "property" }, attribute: { pattern: /@\w+/, alias: "atrule" }, "function-definition": { pattern: /(\bfunc\s+)\w+/, lookbehind: !0, alias: "function" }, label: { pattern: /\b(break|continue)\s+\w+|\b[a-zA-Z_]\w*(?=\s*:\s*(?:for|repeat|while)\b)/, lookbehind: !0, alias: "important" }, keyword: /\b(?:Any|Protocol|Self|Type|actor|as|assignment|associatedtype|associativity|async|await|break|case|catch|class|continue|convenience|default|defer|deinit|didSet|do|dynamic|else|enum|extension|fallthrough|fileprivate|final|for|func|get|guard|higherThan|if|import|in|indirect|infix|init|inout|internal|is|isolated|lazy|left|let|lowerThan|mutating|none|nonisolated|nonmutating|open|operator|optional|override|postfix|precedencegroup|prefix|private|protocol|public|repeat|required|rethrows|return|right|safe|self|set|some|static|struct|subscript|super|switch|throw|throws|try|typealias|unowned|unsafe|var|weak|where|while|willSet)\b/, boolean: /\b(?:false|true)\b/, nil: { pattern: /\bnil\b/, alias: "constant" }, "short-argument": /\$\d+\b/, omit: { pattern: /\b_\b/, alias: "keyword" }, number: /\b(?:[\d_]+(?:\.[\de_]+)?|0x[a-f0-9_]+(?:\.[a-f0-9p_]+)?|0b[01_]+|0o[0-7_]+)\b/i, "class-name": /\b[A-Z](?:[A-Z_\d]*[a-z]\w*)?\b/, function: /\b[a-z_]\w*(?=\s*\()/i, constant: /\b(?:[A-Z_]{2,}|k[A-Z][A-Za-z_]+)\b/, operator: /[-+*/%=!<>&|^~?]+|\.[.\-+*/%=!<>&|^~?]+/, punctuation: /[{}[\]();,.:\\]/ }, Prism.languages.swift["string-literal"].forEach(function(e) {
  e.inside.interpolation.inside = Prism.languages.swift;
}), Prism.languages["visual-basic"] = { comment: { pattern: /(?:['‘’]|REM\b)(?:[^\r\n_]|_(?:\r\n?|\n)?)*/i, inside: { keyword: /^REM/i } }, directive: { pattern: /#(?:Const|Else|ElseIf|End|ExternalChecksum|ExternalSource|If|Region)(?:\b_[ \t]*(?:\r\n?|\n)|.)+/i, alias: "property", greedy: !0 }, string: { pattern: /\$?["“”](?:["“”]{2}|[^"“”])*["“”]C?/i, greedy: !0 }, date: { pattern: /#[ \t]*(?:\d+([/-])\d+\1\d+(?:[ \t]+(?:\d+[ \t]*(?:AM|PM)|\d+:\d+(?::\d+)?(?:[ \t]*(?:AM|PM))?))?|\d+[ \t]*(?:AM|PM)|\d+:\d+(?::\d+)?(?:[ \t]*(?:AM|PM))?)[ \t]*#/i, alias: "number" }, number: /(?:(?:\b\d+(?:\.\d+)?|\.\d+)(?:E[+-]?\d+)?|&[HO][\dA-F]+)(?:[FRD]|U?[ILS])?/i, boolean: /\b(?:False|Nothing|True)\b/i, keyword: /\b(?:AddHandler|AddressOf|Alias|And(?:Also)?|As|Boolean|ByRef|Byte|ByVal|Call|Case|Catch|C(?:Bool|Byte|Char|Date|Dbl|Dec|Int|Lng|Obj|SByte|Short|Sng|Str|Type|UInt|ULng|UShort)|Char|Class|Const|Continue|Currency|Date|Decimal|Declare|Default|Delegate|Dim|DirectCast|Do|Double|Each|Else(?:If)?|End(?:If)?|Enum|Erase|Error|Event|Exit|Finally|For|Friend|Function|Get(?:Type|XMLNamespace)?|Global|GoSub|GoTo|Handles|If|Implements|Imports|In|Inherits|Integer|Interface|Is|IsNot|Let|Lib|Like|Long|Loop|Me|Mod|Module|Must(?:Inherit|Override)|My(?:Base|Class)|Namespace|Narrowing|New|Next|Not(?:Inheritable|Overridable)?|Object|Of|On|Operator|Option(?:al)?|Or(?:Else)?|Out|Overloads|Overridable|Overrides|ParamArray|Partial|Private|Property|Protected|Public|RaiseEvent|ReadOnly|ReDim|RemoveHandler|Resume|Return|SByte|Select|Set|Shadows|Shared|short|Single|Static|Step|Stop|String|Structure|Sub|SyncLock|Then|Throw|To|Try|TryCast|Type|TypeOf|U(?:Integer|Long|Short)|Until|Using|Variant|Wend|When|While|Widening|With(?:Events)?|WriteOnly|Xor)\b/i, operator: /[+\-*/\\^<=>&#@$%!]|\b_(?=[ \t]*[\r\n])/, punctuation: /[{}().,:?]/ }, Prism.languages.vb = Prism.languages["visual-basic"], Prism.languages.vba = Prism.languages["visual-basic"], Prism.languages.wasm = { comment: [/\(;[\s\S]*?;\)/, { pattern: /;;.*/, greedy: !0 }], string: { pattern: /"(?:\\[\s\S]|[^"\\])*"/, greedy: !0 }, keyword: [{ pattern: /\b(?:align|offset)=/, inside: { operator: /=/ } }, { pattern: /\b(?:(?:f32|f64|i32|i64)(?:\.(?:abs|add|and|ceil|clz|const|convert_[su]\/i(?:32|64)|copysign|ctz|demote\/f64|div(?:_[su])?|eqz?|extend_[su]\/i32|floor|ge(?:_[su])?|gt(?:_[su])?|le(?:_[su])?|load(?:(?:8|16|32)_[su])?|lt(?:_[su])?|max|min|mul|neg?|nearest|or|popcnt|promote\/f32|reinterpret\/[fi](?:32|64)|rem_[su]|rot[lr]|shl|shr_[su]|sqrt|store(?:8|16|32)?|sub|trunc(?:_[su]\/f(?:32|64))?|wrap\/i64|xor))?|memory\.(?:grow|size))\b/, inside: { punctuation: /\./ } }, /\b(?:anyfunc|block|br(?:_if|_table)?|call(?:_indirect)?|data|drop|elem|else|end|export|func|get_(?:global|local)|global|if|import|local|loop|memory|module|mut|nop|offset|param|result|return|select|set_(?:global|local)|start|table|tee_local|then|type|unreachable)\b/], variable: /\$[\w!#$%&'*+\-./:<=>?@\\^`|~]+/, number: /[+-]?\b(?:\d(?:_?\d)*(?:\.\d(?:_?\d)*)?(?:[eE][+-]?\d(?:_?\d)*)?|0x[\da-fA-F](?:_?[\da-fA-F])*(?:\.[\da-fA-F](?:_?[\da-fA-D])*)?(?:[pP][+-]?\d(?:_?\d)*)?)\b|\binf\b|\bnan(?::0x[\da-fA-F](?:_?[\da-fA-D])*)?\b/, punctuation: /[()]/ }, Gr.manual = !0;
var Mm = { figure: "figure" }, Cn = function(e) {
  U(r, se);
  var t = Nm(r);
  function r(n) {
    var a;
    n.externals;
    var i, o = n.config;
    return j(this, r), a = t.call(this, { needCache: !0 }), r.inlineCodeCache = {}, a.codeCache = {}, a.customLang = [], a.customParser = {}, a.wrap = o.wrap, a.lineNumber = o.lineNumber, a.copyCode = o.copyCode, a.editCode = o.editCode, a.changeLang = o.changeLang, a.mermaid = o.mermaid, a.indentedCodeBlock = o.indentedCodeBlock === void 0 || o.indentedCodeBlock, a.INLINE_CODE_REGEX = /(`+)(.+?(?:\n.+?)*?)\1/g, o && o.customRenderer && (a.customLang = ve(i = ue(o.customRenderer)).call(i, function(s) {
      return s.toLowerCase();
    }), a.customParser = function(s) {
      for (var c = 1; c < arguments.length; c++) {
        var u, l, f = arguments[c] != null ? arguments[c] : {};
        c % 2 ? q(u = Xc(Object(f), !0)).call(u, function(p) {
          R(s, p, f[p]);
        }) : Ce ? Ct(s, Ce(f)) : q(l = Xc(Object(f))).call(l, function(p) {
          nt(s, p, Te(f, p));
        });
      }
      return s;
    }({}, o.customRenderer)), a.customHighlighter = o.highlighter, a;
  }
  return M(r, [{ key: "$codeCache", value: function(n, a) {
    return n && a && (this.codeCache[n] = a), this.codeCache[n] ? this.codeCache[n] : (this.codeCache.length > 40 && (this.codeCache.length = 0), !1);
  } }, { key: "parseCustomLanguage", value: function(n, a, i) {
    var o, s, c, u, l, f = this.customParser[n];
    if (!f || typeof f.render != "function")
      return !1;
    var p = f.render(a, i.sign, this.$engine, this.mermaid);
    if (!p)
      return !1;
    var d = Mm[f.constructor.TYPE] || "div";
    return b(o = b(s = b(c = b(u = b(l = "<".concat(d, ' data-sign="')).call(l, i.sign, '" data-type="')).call(u, n, '" data-lines="')).call(c, i.lines, '">')).call(s, p, "</")).call(o, d, ">");
  } }, { key: "fillTag", value: function(n) {
    var a = [];
    return ve(n).call(n, function(i) {
      if (!i)
        return "";
      for (var o = i; a.length; ) {
        var s, c = a.pop();
        o = b(s = "".concat(c)).call(s, o);
      }
      var u = o.match(/<span class="(.+?)">|<\/span>/g), l = 0;
      if (!u)
        return o;
      for (; u.length; ) {
        var f = u.pop();
        /<\/span>/.test(f) ? l += 1 : l ? l -= 1 : a.unshift(f.match(/<span class="(.+?)">/)[0]);
      }
      for (var p = 0; p < a.length; p++)
        o = "".concat(o, "</span>");
      return o;
    });
  } }, { key: "renderLineNumber", value: function(n) {
    if (!this.lineNumber)
      return n;
    var a = n.split(`
`);
    return a.pop(), a = this.fillTag(a), '<span class="code-line">'.concat(a.join(`</span>
<span class="code-line">`), "</span>");
  } }, { key: "isInternalCustomLangCovered", value: function(n) {
    var a;
    return Me(a = this.customLang).call(a, n) !== -1;
  } }, { key: "computeLines", value: function(n, a, i) {
    var o = a, s = this.getLineCount(n, o);
    return { sign: this.$engine.md5(n.replace(/^\n+/, "") + s), lines: s };
  } }, { key: "appendMermaid", value: function(n, a) {
    var i = n, o = a;
    if (/^flow([ ](TD|LR))?$/i.test(o) && !this.isInternalCustomLangCovered(o)) {
      var s, c = o.match(/^flow(?:[ ](TD|LR))?$/i) || [];
      i = b(s = "graph ".concat(c[1] || "TD", `
`)).call(s, i), o = "mermaid";
    }
    return /^seq$/i.test(o) && !this.isInternalCustomLangCovered(o) && (i = `sequenceDiagram
`.concat(i), o = "mermaid"), o === "mermaid" && (i = i.replace(/(^[\s]*)stateDiagram-v2\n/, `$1stateDiagram
`)), [i, o];
  } }, { key: "wrapCode", value: function(n, a) {
    var i, o;
    return b(i = b(o = '<code class="language-'.concat(a)).call(o, this.wrap ? " wrap" : "", '">')).call(i, n, "</code>");
  } }, { key: "renderCodeBlock", value: function(n, a, i, o) {
    var s, c, u, l, f, p, d, g = n, _ = a;
    return this.customHighlighter ? g = this.customHighlighter(g, _) : (_ && Gr.languages[_] || (_ = "javascript"), g = Gr.highlight(g, Gr.languages[_], _), g = this.renderLineNumber(g)), g = b(s = b(c = b(u = b(l = b(f = b(p = b(d = `<div
        data-sign="`.concat(i, `"
        data-type="codeBlock"
        data-lines="`)).call(d, o, `" 
        data-edit-code="`)).call(p, this.editCode, `" 
        data-copy-code="`)).call(f, this.copyCode, `"
        data-change-lang="`)).call(l, this.changeLang, `"
        data-lang="`)).call(u, a, `"
      >
      <pre class="language-`)).call(c, _, '">')).call(s, this.wrapCode(g, _), `</pre>
    </div>`);
  } }, { key: "$getIndentedCodeReg", value: function() {
    return new RegExp("(?:^|\\n\\s*\\n)(?: {4}|\\t)([\\s\\S]+?)(?=$|\\n( {0,3}[^ \\t\\n]|\\n[^ \\t\\n]))", "g");
  } }, { key: "$getIndentCodeBlock", value: function(n) {
    var a = this;
    return this.indentedCodeBlock ? this.$recoverCodeInIndent(n).replace(this.$getIndentedCodeReg(), function(i, o) {
      var s, c, u = (i.match(/\n/g) || []).length, l = a.$engine.md5(i), f = b(s = b(c = '<pre data-sign="'.concat(l, '" data-lines="')).call(c, u, '"><code>')).call(s, Mn(o.replace(/\n( {4}|\t)/g, `
`)), "</code></pre>");
      return nn(i, a.pushCache(f, l, u));
    }) : n;
  } }, { key: "$replaceCodeInIndent", value: function(n) {
    return this.indentedCodeBlock ? n.replace(this.$getIndentedCodeReg(), function(a) {
      return a.replace(/`/g, "~~~IndentCode");
    }) : n;
  } }, { key: "$recoverCodeInIndent", value: function(n) {
    return this.indentedCodeBlock ? n.replace(this.$getIndentedCodeReg(), function(a) {
      return a.replace(/~~~IndentCode/g, "`");
    }) : n;
  } }, { key: "beforeMakeHtml", value: function(n, a, i) {
    var o = this, s = n;
    return s = (s = this.$replaceCodeInIndent(s)).replace(this.RULE.reg, function(c, u, l, f, p, d) {
      var g, _, m;
      function h(ge) {
        if (l) {
          var Ne = new RegExp(`^
*`, ""), Q = ge.match(Ne)[0];
          ge = Q + l + ge.replace(Ne, function(ce) {
            return "";
          });
        }
        return ge;
      }
      var v = d, w = o.computeLines(c, u, d), E = w.sign, S = w.lines, A = o.$codeCache(E);
      if (A && A !== "")
        return h(o.getCacheWithSpace(o.pushCache(A, E, S), c));
      v = (v = (v = o.$recoverCodeInIndent(v)).replace(/~D/g, "$")).replace(/~T/g, "~");
      var x = (g = u == null || (_ = u.match(/[ ]/g)) === null || _ === void 0 ? void 0 : _.length) !== null && g !== void 0 ? g : 0;
      if (x > 0) {
        var D = new RegExp("(^|\\n)[ ]{1,".concat(x, "}"), "g");
        v = v.replace(D, "$1");
      }
      if (l) {
        var C = new RegExp("(^|\\n)".concat(l), "g");
        v = v.replace(C, "$1");
      }
      var W = N(p).call(p);
      if (/^(math|katex|latex)$/i.test(W) && !o.isInternalCustomLangCovered(W)) {
        var V, G = c.match(/^\s*/g);
        return b(V = "".concat(G, `~D~D
`)).call(V, v, "~D~D");
      }
      var we = Ht(o.appendMermaid(v, W), 2);
      return v = we[0], W = we[1], Me(m = o.customLang).call(m, W.toLowerCase()) !== -1 && (A = o.parseCustomLanguage(W, v, { lines: S, sign: E })) && A !== "" ? (o.$codeCache(E, A), o.getCacheWithSpace(o.pushCache(A, E, S), c)) : (v = v.replace(/~X/g, "\\`"), A = (A = o.renderCodeBlock(v, W, E, S)).replace(/\\/g, "\\\\"), A = o.$codeCache(E, A), h(o.getCacheWithSpace(o.pushCache(A, E, S), c)));
    }), s = s.replace(Eo(!0), function(c) {
      var u;
      return ve(u = c.split("|")).call(u, function(l) {
        return o.makeInlineCode(l);
      }).join("|").replace(/`/g, "\\`");
    }), s = this.makeInlineCode(s), s = this.$getIndentCodeBlock(s);
  } }, { key: "makeInlineCode", value: function(n) {
    var a = this, i = n;
    return this.INLINE_CODE_REGEX.test(i) && (i = (i = i.replace(/\\`/g, "~~not~inlineCode")).replace(this.INLINE_CODE_REGEX, function(o, s, c) {
      if (N(c).call(c) === "`")
        return o;
      var u = c.replace(/~~not~inlineCode/g, "\\`");
      u = (u = a.$replaceSpecialChar(u)).replace(/\\/g, "\\\\");
      var l = "<code>".concat(Mn(u), "</code>"), f = a.$engine.md5(l);
      return r.inlineCodeCache[f] = l, "~~CODE".concat(f, "$");
    }), i = i.replace(/~~not~inlineCode/g, "\\`")), i;
  } }, { key: "makeHtml", value: function(n) {
    return n;
  } }, { key: "$replaceSpecialChar", value: function(n) {
    var a = n.replace(/~Q/g, "\\~");
    return a = (a = (a = (a = a.replace(/~Y/g, "\\!")).replace(/~Z/g, "\\#")).replace(/~&/g, "\\&")).replace(/~K/g, "\\/");
  } }, { key: "rule", value: function() {
    return (n = { begin: /(?:^|\n)(\n*((?:>[\t ]*)*)(?:[^\S\n]*))(`{3,})([^`]*?)\n/, content: /([\w\W]*?)/, end: /[^\S\n]*\3[ \t]*(?=$|\n+)/, reg: new RegExp("") }).reg = new RegExp(n.begin.source + n.content.source + n.end.source, "g"), Bc(Bc({}, n), {}, { begin: n.begin.source, content: n.content.source, end: n.end.source });
    var n;
  } }, { key: "mounted", value: function(n) {
  } }]), r;
}();
function jm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Cn, "HOOK_NAME", "codeBlock"), R(Cn, "inlineCodeCache", {});
var wf = function(e) {
  U(r, se);
  var t = jm(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "makeHtml", value: function(n) {
    return n;
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = n;
    return ue(Cn.inlineCodeCache).length > 0 && (a = a.replace(/~~CODE([0-9a-zA-Z]+)\$/g, function(i, o) {
      return Cn.inlineCodeCache[o];
    }), Cn.inlineCodeCache = {}), a;
  } }, { key: "rule", value: function() {
    var n = { begin: "(`+)[ ]*", end: "[ ]*\\1", content: "(.+?(?:\\n.+?)*?)" };
    return n.reg = Le(n, "g"), n;
  } }]), r;
}();
R(wf, "HOOK_NAME", "inlineCode");
var Dm = rt(function(e) {
  (function() {
    var t = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", r = { rotl: function(n, a) {
      return n << a | n >>> 32 - a;
    }, rotr: function(n, a) {
      return n << 32 - a | n >>> a;
    }, endian: function(n) {
      if (n.constructor == Number)
        return 16711935 & r.rotl(n, 8) | 4278255360 & r.rotl(n, 24);
      for (var a = 0; a < n.length; a++)
        n[a] = r.endian(n[a]);
      return n;
    }, randomBytes: function(n) {
      for (var a = []; n > 0; n--)
        a.push(Math.floor(256 * Math.random()));
      return a;
    }, bytesToWords: function(n) {
      for (var a = [], i = 0, o = 0; i < n.length; i++, o += 8)
        a[o >>> 5] |= n[i] << 24 - o % 32;
      return a;
    }, wordsToBytes: function(n) {
      for (var a = [], i = 0; i < 32 * n.length; i += 8)
        a.push(n[i >>> 5] >>> 24 - i % 32 & 255);
      return a;
    }, bytesToHex: function(n) {
      for (var a = [], i = 0; i < n.length; i++)
        a.push((n[i] >>> 4).toString(16)), a.push((15 & n[i]).toString(16));
      return a.join("");
    }, hexToBytes: function(n) {
      for (var a = [], i = 0; i < n.length; i += 2)
        a.push(parseInt(n.substr(i, 2), 16));
      return a;
    }, bytesToBase64: function(n) {
      for (var a = [], i = 0; i < n.length; i += 3)
        for (var o = n[i] << 16 | n[i + 1] << 8 | n[i + 2], s = 0; s < 4; s++)
          8 * i + 6 * s <= 8 * n.length ? a.push(t.charAt(o >>> 6 * (3 - s) & 63)) : a.push("=");
      return a.join("");
    }, base64ToBytes: function(n) {
      n = n.replace(/[^A-Z0-9+\/]/gi, "");
      for (var a = [], i = 0, o = 0; i < n.length; o = ++i % 4)
        o != 0 && a.push((t.indexOf(n.charAt(i - 1)) & Math.pow(2, -2 * o + 8) - 1) << 2 * o | t.indexOf(n.charAt(i)) >>> 6 - 2 * o);
      return a;
    } };
    e.exports = r;
  })();
}), Zi = { utf8: { stringToBytes: function(e) {
  return Zi.bin.stringToBytes(unescape(encodeURIComponent(e)));
}, bytesToString: function(e) {
  return decodeURIComponent(escape(Zi.bin.bytesToString(e)));
} }, bin: { stringToBytes: function(e) {
  for (var t = [], r = 0; r < e.length; r++)
    t.push(255 & e.charCodeAt(r));
  return t;
}, bytesToString: function(e) {
  for (var t = [], r = 0; r < e.length; r++)
    t.push(String.fromCharCode(e[r]));
  return t.join("");
} } }, Vc = Zi, Fm = function(e) {
  return e != null && (Jc(e) || function(t) {
    return typeof t.readFloatLE == "function" && typeof t.slice == "function" && Jc(t.slice(0, 0));
  }(e) || !!e._isBuffer);
};
function Jc(e) {
  return !!e.constructor && typeof e.constructor.isBuffer == "function" && e.constructor.isBuffer(e);
}
var Ef = rt(function(e) {
  (function() {
    var t = Dm, r = Vc.utf8, n = Fm, a = Vc.bin, i = function(o, s) {
      o.constructor == String ? o = s && s.encoding === "binary" ? a.stringToBytes(o) : r.stringToBytes(o) : n(o) ? o = Array.prototype.slice.call(o, 0) : Array.isArray(o) || o.constructor === Uint8Array || (o = o.toString());
      for (var c = t.bytesToWords(o), u = 8 * o.length, l = 1732584193, f = -271733879, p = -1732584194, d = 271733878, g = 0; g < c.length; g++)
        c[g] = 16711935 & (c[g] << 8 | c[g] >>> 24) | 4278255360 & (c[g] << 24 | c[g] >>> 8);
      c[u >>> 5] |= 128 << u % 32, c[14 + (u + 64 >>> 9 << 4)] = u;
      var _ = i._ff, m = i._gg, h = i._hh, v = i._ii;
      for (g = 0; g < c.length; g += 16) {
        var w = l, E = f, S = p, A = d;
        l = _(l, f, p, d, c[g + 0], 7, -680876936), d = _(d, l, f, p, c[g + 1], 12, -389564586), p = _(p, d, l, f, c[g + 2], 17, 606105819), f = _(f, p, d, l, c[g + 3], 22, -1044525330), l = _(l, f, p, d, c[g + 4], 7, -176418897), d = _(d, l, f, p, c[g + 5], 12, 1200080426), p = _(p, d, l, f, c[g + 6], 17, -1473231341), f = _(f, p, d, l, c[g + 7], 22, -45705983), l = _(l, f, p, d, c[g + 8], 7, 1770035416), d = _(d, l, f, p, c[g + 9], 12, -1958414417), p = _(p, d, l, f, c[g + 10], 17, -42063), f = _(f, p, d, l, c[g + 11], 22, -1990404162), l = _(l, f, p, d, c[g + 12], 7, 1804603682), d = _(d, l, f, p, c[g + 13], 12, -40341101), p = _(p, d, l, f, c[g + 14], 17, -1502002290), l = m(l, f = _(f, p, d, l, c[g + 15], 22, 1236535329), p, d, c[g + 1], 5, -165796510), d = m(d, l, f, p, c[g + 6], 9, -1069501632), p = m(p, d, l, f, c[g + 11], 14, 643717713), f = m(f, p, d, l, c[g + 0], 20, -373897302), l = m(l, f, p, d, c[g + 5], 5, -701558691), d = m(d, l, f, p, c[g + 10], 9, 38016083), p = m(p, d, l, f, c[g + 15], 14, -660478335), f = m(f, p, d, l, c[g + 4], 20, -405537848), l = m(l, f, p, d, c[g + 9], 5, 568446438), d = m(d, l, f, p, c[g + 14], 9, -1019803690), p = m(p, d, l, f, c[g + 3], 14, -187363961), f = m(f, p, d, l, c[g + 8], 20, 1163531501), l = m(l, f, p, d, c[g + 13], 5, -1444681467), d = m(d, l, f, p, c[g + 2], 9, -51403784), p = m(p, d, l, f, c[g + 7], 14, 1735328473), l = h(l, f = m(f, p, d, l, c[g + 12], 20, -1926607734), p, d, c[g + 5], 4, -378558), d = h(d, l, f, p, c[g + 8], 11, -2022574463), p = h(p, d, l, f, c[g + 11], 16, 1839030562), f = h(f, p, d, l, c[g + 14], 23, -35309556), l = h(l, f, p, d, c[g + 1], 4, -1530992060), d = h(d, l, f, p, c[g + 4], 11, 1272893353), p = h(p, d, l, f, c[g + 7], 16, -155497632), f = h(f, p, d, l, c[g + 10], 23, -1094730640), l = h(l, f, p, d, c[g + 13], 4, 681279174), d = h(d, l, f, p, c[g + 0], 11, -358537222), p = h(p, d, l, f, c[g + 3], 16, -722521979), f = h(f, p, d, l, c[g + 6], 23, 76029189), l = h(l, f, p, d, c[g + 9], 4, -640364487), d = h(d, l, f, p, c[g + 12], 11, -421815835), p = h(p, d, l, f, c[g + 15], 16, 530742520), l = v(l, f = h(f, p, d, l, c[g + 2], 23, -995338651), p, d, c[g + 0], 6, -198630844), d = v(d, l, f, p, c[g + 7], 10, 1126891415), p = v(p, d, l, f, c[g + 14], 15, -1416354905), f = v(f, p, d, l, c[g + 5], 21, -57434055), l = v(l, f, p, d, c[g + 12], 6, 1700485571), d = v(d, l, f, p, c[g + 3], 10, -1894986606), p = v(p, d, l, f, c[g + 10], 15, -1051523), f = v(f, p, d, l, c[g + 1], 21, -2054922799), l = v(l, f, p, d, c[g + 8], 6, 1873313359), d = v(d, l, f, p, c[g + 15], 10, -30611744), p = v(p, d, l, f, c[g + 6], 15, -1560198380), f = v(f, p, d, l, c[g + 13], 21, 1309151649), l = v(l, f, p, d, c[g + 4], 6, -145523070), d = v(d, l, f, p, c[g + 11], 10, -1120210379), p = v(p, d, l, f, c[g + 2], 15, 718787259), f = v(f, p, d, l, c[g + 9], 21, -343485551), l = l + w >>> 0, f = f + E >>> 0, p = p + S >>> 0, d = d + A >>> 0;
      }
      return t.endian([l, f, p, d]);
    };
    i._ff = function(o, s, c, u, l, f, p) {
      var d = o + (s & c | ~s & u) + (l >>> 0) + p;
      return (d << f | d >>> 32 - f) + s;
    }, i._gg = function(o, s, c, u, l, f, p) {
      var d = o + (s & u | c & ~u) + (l >>> 0) + p;
      return (d << f | d >>> 32 - f) + s;
    }, i._hh = function(o, s, c, u, l, f, p) {
      var d = o + (s ^ c ^ u) + (l >>> 0) + p;
      return (d << f | d >>> 32 - f) + s;
    }, i._ii = function(o, s, c, u, l, f, p) {
      var d = o + (c ^ (s | ~u)) + (l >>> 0) + p;
      return (d << f | d >>> 32 - f) + s;
    }, i._blocksize = 16, i._digestsize = 16, e.exports = function(o, s) {
      if (o == null)
        throw new Error("Illegal argument " + o);
      var c = t.wordsToBytes(i(o, s));
      return s && s.asBytes ? c : s && s.asString ? a.bytesToString(c) : t.bytesToHex(c);
    };
  })();
}), Nr = {}, hi = /^cherry-inner:\/\/([0-9a-f]+)$/i, St = function() {
  function e() {
    j(this, e);
  }
  return M(e, null, [{ key: "isInnerLink", value: function(t) {
    return hi.test(t);
  } }, { key: "set", value: function(t) {
    var r = Ef(t);
    return Nr[r] = t, "cherry-inner://".concat(r);
  } }, { key: "get", value: function(t) {
    var r, n = Ht((r = t.match(hi)) !== null && r !== void 0 ? r : [], 2)[1];
    if (n)
      return Nr[n];
  } }, { key: "replace", value: function(t, r) {
    var n, a = Ht((n = t.match(hi)) !== null && n !== void 0 ? n : [], 2)[1];
    if (a)
      return Nr[a] = r, t;
  } }, { key: "restoreAll", value: function(t) {
    var r = t.replace(/cherry-inner:\/\/([0-9a-f]+)/gi, function(n) {
      return e.get(n) || n;
    });
    return r;
  } }, { key: "clear", value: function() {
    Nr = {};
  } }]), e;
}();
function Bm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Sf = function(e) {
  U(r, pe);
  var t = Bm(r);
  function r(n) {
    var a, i = n.config, o = n.globalConfig;
    return j(this, r), (a = t.call(this, { config: i })).urlProcessor = o.urlProcessor, a.target = i.target ? 'target="'.concat(i.target, '"') : i.openNewPage ? 'target="_blank"' : "", a.rel = i.rel ? 'rel="'.concat(i.rel, '"') : "", a;
  }
  return M(r, [{ key: "checkBrackets", value: function(n) {
    for (var a = [], i = "[".concat(n, "]"), o = function(c) {
      return 1 & Ke(i).call(i, 0, c).match(/\\*$/)[0].length;
    }, s = i.length - 1; i[s] && (s !== i.length - 1 || !o(s)); s--)
      if (i[s] !== "]" || o(s) || a.push("]"), i[s] === "[" && !o(s) && (a.pop(), !a.length))
        return { isValid: !0, coreText: Ke(i).call(i, s + 1, i.length - 1), extraLeadingChar: Ke(i).call(i, 0, s) };
    return { isValid: !1, coreText: n, extraLeadingChar: "" };
  } }, { key: "toHtml", value: function(n, a, i, o, s, c, u) {
    var l = o === void 0 ? "ref" : "url", f = "";
    if (l === "ref")
      return n;
    if (l === "url") {
      var p, d = this.checkBrackets(i), g = d.isValid, _ = d.coreText, m = d.extraLeadingChar;
      if (!g)
        return n;
      f = s && N(s).call(s) !== "" ? ' title="'.concat(Mn(s.replace(/["']/g, "")), '"') : "", u ? f += ' target="'.concat(u.replace(/{target\s*=\s*(.*?)}/, "$1"), '"') : this.target && (f += " ".concat(this.target));
      var h, v, w, E, S = N(o).call(o).replace(/~1D/g, "~D"), A = _.replace(/~1D/g, "~D");
      return Z1(S) ? (S = Nt(S = this.urlProcessor(S, "link")), b(h = b(v = b(w = b(E = "".concat(a + m, '<a href="')).call(E, St.set(S), '" ')).call(w, this.rel, " ")).call(v, f, ">")).call(h, A, "</a>")) : b(p = "".concat(a + m, "<span>")).call(p, i, "</span>");
    }
    return n;
  } }, { key: "toStdMarkdown", value: function(n) {
    return n;
  } }, { key: "makeHtml", value: function(n) {
    var a, i, o = n.replace(this.RULE.reg, function(s) {
      return s.replace(/~D/g, "~1D");
    });
    return fe() ? o = o.replace(this.RULE.reg, je(a = this.toHtml).call(a, this)) : o = ct(o, this.RULE.reg, je(i = this.toHtml).call(i, this), !0, 1), o = o.replace(this.RULE.reg, function(s) {
      return s.replace(/~1D/g, "~D");
    }), o;
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))" : "(^|[^\\\\])", content: ["\\[([^\\n]*?)\\]", "[ \\t]*", "".concat(`(?:\\(([^\\s)]+)(?:[ \\t]((?:".*?")|(?:'.*?')))?\\)|\\[(`).concat(nf, ")\\]") + ")", "(\\{target\\s*=\\s*(_blank|_parent|_self|_top)\\})?"].join(""), end: "" };
    return n.reg = Le(n, "g"), n;
  } }]), r;
}();
R(Sf, "HOOK_NAME", "link");
var Hm = O.RangeError;
L({ target: "String", proto: !0 }, { repeat: function(e) {
  var t = tt(an(this)), r = "", n = dr(e);
  if (n < 0 || n == 1 / 0)
    throw Hm("Wrong number of repetitions");
  for (; n > 0; (n >>>= 1) && (t += t))
    1 & n && (r += t);
  return r;
} });
var zm = Ie("String").repeat, mi = String.prototype, gt = function(e) {
  var t = e.repeat;
  return typeof e == "string" || e === mi || de(mi, e) && t === mi.repeat ? zm : t;
};
function Um(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Af = function(e) {
  U(r, pe);
  var t = Um(r);
  function r() {
    var n, a = (arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0 }).config;
    return j(this, r), n = t.call(this, { config: a }), a ? (n.allowWhitespace = !!a.allowWhitespace, n) : B(n);
  }
  return M(r, [{ key: "makeHtml", value: function(n, a) {
    var i = function(s, c, u, l) {
      var f, p, d, g = u.length % 2 == 1 ? "em" : "strong", _ = Math.floor(u.length / 2), m = gt("<strong>").call("<strong>", _), h = gt("</strong>").call("</strong>", _);
      return g === "em" && (m += "<em>", h = "</em>".concat(h)), b(f = b(p = b(d = "".concat(c)).call(d, m)).call(p, a(l).html.replace(/_/g, "~U"))).call(f, h);
    }, o = n;
    return o = (o = this.allowWhitespace ? (o = (o = o.replace(/(^|\n[\s]*)(\*)([^\s*](?:.*?)(?:(?:\n.*?)*?))\*/g, i)).replace(/(^|\n[\s]*)(\*{2,})((?:.*?)(?:(?:\n.*?)*?))\2/g, i)).replace(/([^\n*\\\s][ ]*)(\*+)((?:.*?)(?:(?:\n.*?)*?))\2/g, i) : o.replace(this.RULE.asterisk.reg, i)).replace(this.RULE.underscore.reg, function(s, c, u, l, f, p) {
      var d, g, _;
      if (N(l).call(l) === "")
        return s;
      var m = u.length % 2 == 1 ? "em" : "strong", h = Math.floor(u.length / 2), v = gt("<strong>").call("<strong>", h), w = gt("</strong>").call("</strong>", h), E = a(l).html;
      return m === "em" && (v += "<em>", w = "</em>".concat(w)), b(d = b(g = b(_ = "".concat(c)).call(_, v)).call(g, E)).call(d, w);
    }), o.replace(/~U/g, "_");
  } }, { key: "test", value: function(n, a) {
    return this.RULE[a].reg && this.RULE[a].reg.test(n);
  } }, { key: "rule", value: function() {
    var n = (arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0 }).config, a = !!n && !!n.allowWhitespace, i = function(c, u) {
      var l, f, p, d = "[^".concat(u, "\\s]");
      return c ? "(?:.*?)(?:(?:\\n.*?)*?)" : b(l = b(f = b(p = "(".concat(d, "|")).call(p, d, `(.*?(
`)).call(f, d, ".*)*)")).call(l, d, ")");
    }, o = { begin: "(^|[^\\\\])([*]+)", content: "(".concat(i(a, "*"), ")"), end: "\\2" }, s = { begin: "(^|".concat(zc, ")(_+)"), content: "(".concat(i(a, "_"), ")"), end: "\\2(?=".concat(zc, "|$)") };
    return o.reg = Le(o, "g"), s.reg = Le(s, "g"), { asterisk: o, underscore: s };
  } }]), r;
}();
function Wm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Af, "HOOK_NAME", "fontEmphasis");
var xf = function(e) {
  U(r, se);
  var t = Wm(r);
  function r(n) {
    var a;
    return j(this, r), (a = t.call(this)).initBrReg(n.globalConfig.classicBr), a;
  }
  return M(r, [{ key: "makeHtml", value: function(n, a) {
    var i = this;
    return this.test(n) ? n.replace(this.RULE.reg, function(o, s, c) {
      var u;
      if (i.isContainsCache(o, !0))
        return o;
      var l, f = function(p) {
        var d, g, _, m, h, v;
        if (N(p).call(p) === "")
          return "";
        var w = a(p), E = w.sign, S = w.html, A = "p";
        new RegExp("<(".concat(vr, ")[^>]*>"), "i").test(S) && (A = "div");
        var x = i.getLineCount(p, p);
        return b(d = b(g = b(_ = b(m = b(h = b(v = "<".concat(A, ' data-sign="')).call(v, E)).call(h, x, '" data-type="')).call(m, A, '" data-lines="')).call(_, x, '">')).call(g, i.$cleanParagraph(S), "</")).call(d, A, ">");
      };
      return i.isContainsCache(c) ? i.makeExcludingCached(b(l = "".concat(s)).call(l, c), f) : f(b(u = "".concat(s)).call(u, c));
    }) : n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:^|\\n)(\\n*)", end: "(?=\\s*$|\\n\\n)", content: "([\\s\\S]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
R(xf, "HOOK_NAME", "normalParagraph");
L({ target: "Reflect", stat: !0 }, { get: function e(t, r) {
  var n, a, i = arguments.length < 3 ? t : arguments[2];
  return K(t) === i ? t[r] : (n = on.f(t, r)) ? function(o) {
    return o !== void 0 && (F(o, "value") || F(o, "writable"));
  }(n) ? n.value : n.get === void 0 ? void 0 : Z(n.get, i) : ye(a = On(t)) ? e(a, r, i) : void 0;
} });
var Qc = X.Reflect.get, qm = Yu;
function dt() {
  return dt = typeof Reflect < "u" && Qc ? Qc : function(e, t, r) {
    var n = function(i, o) {
      for (; !Object.prototype.hasOwnProperty.call(i, o) && (i = T(i)) !== null; )
        ;
      return i;
    }(e, t);
    if (n) {
      var a = qm(n, t);
      return a.get ? a.get.call(arguments.length < 3 ? e : r) : a.value;
    }
  }, dt.apply(this, arguments);
}
function Gm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Mr = "atx", jr = "setext", Km = /[\s\-_]/, Zm = /[A-Za-z]/, Ym = /[0-9]/, Cf = function(e) {
  U(r, se);
  var t = Gm(r);
  function r() {
    var n, a = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0, externals: void 0 };
    a.externals;
    var i = a.config;
    return j(this, r), (n = t.call(this, { needCache: !0 })).strict = !i || !!i.strict, n.RULE = n.rule(), n.headerIDCache = [], n.headerIDCounter = {}, n.config = i || {}, n;
  }
  return M(r, [{ key: "$parseTitleText", value: function() {
    var n = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : "";
    return typeof n != "string" ? "" : n.replace(/<.*?>/g, "").replace(/&#60;/g, "<").replace(/&#62;/g, ">");
  } }, { key: "$generateId", value: function(n) {
    for (var a = !(arguments.length > 1 && arguments[1] !== void 0) || arguments[1], i = n.length, o = "", s = 0; s < i; s++) {
      var c = n.charAt(s);
      if (Zm.test(c))
        o += a ? c.toLowerCase() : c;
      else if (Ym.test(c))
        o += c;
      else if (Km.test(c))
        o += o.length < 1 || o.charAt(o.length - 1) !== "-" ? "-" : "";
      else if (c.charCodeAt(0) > 255)
        try {
          o += encodeURIComponent(c);
        } catch {
        }
    }
    return o;
  } }, { key: "generateIDNoDup", value: function(n) {
    var a, i = n.replace(/&#60;/g, "<").replace(/&#62;/g, ">"), o = this.$generateId(i, !0), s = Me(a = this.headerIDCache).call(a, o);
    if (s !== -1)
      this.headerIDCounter[s] += 1, o += "-".concat(this.headerIDCounter[s] + 1);
    else {
      var c = this.headerIDCache.push(o);
      this.headerIDCounter[c - 1] = 1;
    }
    return o;
  } }, { key: "$wrapHeader", value: function(n, a, i, o) {
    var s, c, u, l, f, p, d, g = o(N(n).call(n)), _ = g.html, m = _.match(/\s+\{#([A-Za-z0-9-]+)\}$/);
    m !== null && (_ = _.substring(0, m.index), d = Ht(m, 2)[1]);
    var h = this.$parseTitleText(_);
    d || (d = this.generateIDNoDup(h.replace(/~fn#([0-9]+)#/g, "")));
    var v = "safe_".concat(d), w = this.$engine.md5(b(s = b(c = b(u = "".concat(a, "-")).call(u, g.sign, "-")).call(c, d, "-")).call(s, i));
    return { html: [b(l = b(f = b(p = "<h".concat(a, ' id="')).call(p, v, '" data-sign="')).call(f, w, '" data-lines="')).call(l, i, '">'), this.$getAnchor(d), "".concat(_), "</h".concat(a, ">")].join(""), sign: "".concat(w) };
  } }, { key: "$getAnchor", value: function(n) {
    return (this.config.anchorStyle || "default") === "none" ? "" : '<a class="anchor" href="#'.concat(n, '"></a>');
  } }, { key: "beforeMakeHtml", value: function(n) {
    var a = this, i = n;
    return this.test(i, Mr) && (i = i.replace(this.RULE[Mr].reg, function(o, s, c, u) {
      return N(u).call(u) === "" ? o : a.getCacheWithSpace(a.pushCache(o), o, !0);
    })), this.test(i, jr) && (i = i.replace(this.RULE[jr].reg, function(o, s, c) {
      return N(c).call(c) === "" || a.isContainsCache(c) ? o : a.getCacheWithSpace(a.pushCache(o), o, !0);
    })), i;
  } }, { key: "makeHtml", value: function(n, a) {
    var i = this, o = this.restoreCache(n);
    return this.test(o, Mr) && (o = o.replace(this.RULE[Mr].reg, function(s, c, u, l) {
      var f = ji(c, i.getLineCount(s.replace(/^\n+/, ""))), p = l.replace(/\s+#+\s*$/, ""), d = i.$wrapHeader(p, u.length, f, a), g = d.html, _ = d.sign;
      return i.getCacheWithSpace(i.pushCache(g, _, f), s, !0);
    })), this.test(o, jr) && (o = o.replace(this.RULE[jr].reg, function(s, c, u, l) {
      if (i.isContainsCache(u))
        return s;
      var f = ji(c, i.getLineCount(s.replace(/^\n+/, ""))), p = l[0] === "-" ? 2 : 1, d = i.$wrapHeader(u, p, f, a), g = d.html, _ = d.sign;
      return i.getCacheWithSpace(i.pushCache(g, _, f), s, !0);
    })), o;
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = dt(T(r.prototype), "afterMakeHtml", this).call(this, n);
    return this.headerIDCache = [], this.headerIDCounter = {}, a;
  } }, { key: "test", value: function(n, a) {
    return this.RULE[a].reg && this.RULE[a].reg.test(n);
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:^|\\n)(\\n*)", content: ["(?:\\h*", "(.+)", ")\\n", "(?:\\h*", "([=]+|[-]+)", ")"].join(""), end: "(?=$|\\n)" };
    n.reg = Le(n, "g", !0);
    var a = { begin: "(?:^|\\n)(\\n*)(?:\\h*(#{1,6}))", content: "(.+?)", end: "(?=$|\\n)" };
    return this.strict && (a.begin += "(?=\\h+)"), a.reg = Le(a, "g", !0), { setext: n, atx: a };
  } }]), r;
}();
function Xm(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Cf, "HOOK_NAME", "header");
var Tf = function(e) {
  U(r, pe);
  var t = Xm(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "rule", value: function() {
    return { begin: "", content: "", end: "", reg: new RegExp("") };
  } }, { key: "beforeMakeHtml", value: function(n) {
    return n.replace(/\\\n/g, `\\ 
`);
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = n.replace(/~Q/g, "~");
    return a = (a = (a = (a = (a = a.replace(/~X/g, "`")).replace(/~Y/g, "!")).replace(/~Z/g, "#")).replace(/~&/g, "&")).replace(/~K/g, "/");
  } }]), r;
}();
R(Tf, "HOOK_NAME", "transfer");
var Vm = O.TypeError, el = function(e) {
  return function(t, r, n, a) {
    ne(r);
    var i = lt(t), o = to(i), s = bt(i), c = e ? s - 1 : 0, u = e ? -1 : 1;
    if (n < 2)
      for (; ; ) {
        if (c in o) {
          a = o[c], c += u;
          break;
        }
        if (c += u, e ? c < 0 : s <= c)
          throw Vm("Reduce of empty array with no initial value");
      }
    for (; e ? c >= 0 : s > c; c += u)
      c in o && (a = r(a, o[c], c, i));
    return a;
  };
}, Jm = { left: el(!1), right: el(!0) }, Qm = ht(O.process) == "process", eb = Jm.left, tb = ha("reduce");
L({ target: "Array", proto: !0, forced: !tb || !Qm && Tn > 79 && Tn < 83 }, { reduce: function(e) {
  var t = arguments.length;
  return eb(this, e, t, t > 1 ? arguments[1] : void 0);
} });
var nb = Ie("Array").reduce, bi = Array.prototype, na = function(e) {
  var t = e.reduce;
  return e === bi || de(bi, e) && t === bi.reduce ? nb : t;
};
function tl(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Dr(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = tl(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = tl(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function rb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var nl = "loose", rl = "strict", $f = function(e) {
  U(r, se);
  var t = rb(r);
  function r(n) {
    var a, i = n.externals, o = n.config;
    j(this, r), a = t.call(this, { needCache: !0 });
    var s = o.enableChart, c = o.chartRenderEngine, u = o.externals, l = o.chartEngineOptions, f = l === void 0 ? {} : l;
    if (a.chartRenderEngine = null, s === !0)
      try {
        a.chartRenderEngine = new c(Dr(Dr({}, i && u instanceof Array && na(u).call(u, function(p, d) {
          return delete f[d], Dr(Dr({}, p), {}, R({}, d, i[d]));
        }, {})), {}, { renderer: "svg", width: 500, height: 300 }, f));
      } catch (p) {
        console.warn(p);
      }
    return a;
  }
  return M(r, [{ key: "$extendColumns", value: function(n, a) {
    var i = a - n.length;
    return i < 1 ? n : b(n).call(n, gt("&nbsp;|").call("&nbsp;|", i).split("|", i));
  } }, { key: "$parseChartOptions", value: function(n) {
    if (!this.chartRenderEngine)
      return null;
    var a = /^[ ]*:(\w+):(?:[ ]*{(.*?)}[ ]*)?$/;
    if (!a.test(n))
      return null;
    var i = Ht(n.match(a), 3), o = i[1], s = i[2];
    return { type: o, options: s ? s.split(/\s*,\s*/) : ["x", "y"] };
  } }, { key: "$parseColumnAlignRules", value: function(n) {
    var a = ["U", "L", "R", "C"];
    return { textAlignRules: ve(n).call(n, function(i) {
      var o = N(i).call(i), s = 0;
      return /^:/.test(o) && (s += 1), /:$/.test(o) && (s += 2), a[s];
    }), COLUMN_ALIGN_MAP: { L: "left", R: "right", C: "center" } };
  } }, { key: "$parseTable", value: function(n, a, i) {
    var o, s, c, u, l, f, p = this, d = 0, g = ve(n).call(n, function(W, V) {
      var G = W.replace(/\\\|/g, "~CS").split("|");
      return G[0] === "" && G.shift(), G[G.length - 1] === "" && G.pop(), V !== 1 && (d = Math.max(d, G.length)), G;
    }), _ = this.$parseColumnAlignRules(g[1]), m = _.textAlignRules, h = _.COLUMN_ALIGN_MAP, v = { header: [], rows: [], colLength: d, rowLength: g.length - 2 }, w = this.$parseChartOptions(g[0][0]), E = this.$engine.md5(g[0][0]);
    w && (g[0][0] = "");
    var S = ve(o = this.$extendColumns(g[0], d)).call(o, function(W, V) {
      var G, we;
      v.header.push(W.replace(/~CS/g, "\\|"));
      var ge = a(N(G = W.replace(/~CS/g, "\\|")).call(G)).html;
      return b(we = "~CTH".concat(m[V] || "U", " ")).call(we, ge, " ~CTH$");
    }).join(""), A = na(g).call(g, function(W, V, G) {
      var we;
      if (G <= 1)
        return W;
      var ge = G - 2;
      v.rows[ge] = [];
      var Ne = ve(we = p.$extendColumns(V, d)).call(we, function(Q, ce) {
        var $e, he;
        v.rows[ge].push(Q.replace(/~CS/g, "\\|"));
        var Qe = a(N($e = Q.replace(/~CS/g, "\\|")).call($e)).html;
        return b(he = "~CTD".concat(m[ce] || "U", " ")).call(he, Qe, " ~CTD$");
      });
      return W.push("~CTR".concat(Ne.join(""), "~CTR$")), W;
    }, []).join(""), x = this.$renderTable(h, S, A, i);
    if (!w)
      return x;
    var D = this.chartRenderEngine.render(w.type, w.options, v), C = b(s = b(c = b(u = b(l = '<figure id="table_chart_'.concat(E, "_")).call(l, x.sign, `"
      data-sign="table_chart_`)).call(u, E, "_")).call(c, x.sign, '" data-lines="0">')).call(s, D, "</figure>");
    return { html: b(f = "".concat(C)).call(f, x.html), sign: E + x.sign };
  } }, { key: "$testHeadEmpty", value: function(n) {
    var a = n.replace(/&nbsp;/g, "").replace(/\s/g, "").replace(/(~CTH\$|~CTHU|~CTHL|~CTHR|~CTHC)/g, "");
    return (a == null ? void 0 : a.length) > 0;
  } }, { key: "$renderTable", value: function(n, a, i, o) {
    var s, c, u, l, f = this.$testHeadEmpty(a) ? b(s = "~CTHD".concat(a, "~CTHD$~CTBD")).call(s, i, "~CTBD$") : "~CTBD".concat(i, "~CTBD$"), p = this.$engine.md5(f), d = f.replace(/~CTHD\$/g, "</thead>").replace(/~CTHD/g, "<thead>").replace(/~CTBD\$/g, "</tbody>").replace(/~CTBD/g, "</tbody>").replace(/~CTR\$/g, "</tr>").replace(/~CTR/g, "<tr>").replace(/[ ]?~CTH\$/g, "</th>").replace(/[ ]?~CTD\$/g, "</td>").replace(/~CT(D|H)(L|R|C|U)[ ]?/g, function(g, _, m) {
      var h = "<t".concat(_);
      return h += m === "U" ? ">" : ' align="'.concat(n[m], '">');
    }).replace(/\\\|/g, "|");
    return { html: b(c = b(u = b(l = '<div class="cherry-table-container" data-sign="'.concat(p)).call(l, o, '" data-lines="')).call(u, o, `">
        <table class="cherry-table">`)).call(c, d, "</table></div>"), sign: p };
  } }, { key: "makeHtml", value: function(n, a) {
    var i = this, o = n;
    return this.test(o, rl) && (o = o.replace(this.RULE[rl].reg, function(s, c) {
      var u, l = i.getLineCount(s, c), f = ve(u = N(s).call(s).split(/\n/)).call(u, function(_) {
        var m;
        return N(m = String(_)).call(m);
      }), p = i.$parseTable(f, a, l), d = p.html, g = p.sign;
      return i.getCacheWithSpace(i.pushCache(d, g, l), s);
    })), this.test(o, nl) && (o = o.replace(this.RULE[nl].reg, function(s, c) {
      var u, l = i.getLineCount(s, c), f = ve(u = N(s).call(s).split(/\n/)).call(u, function(_) {
        var m;
        return N(m = String(_)).call(m);
      }), p = i.$parseTable(f, a, l), d = p.html, g = p.sign;
      return i.getCacheWithSpace(i.pushCache(d, g, l), s);
    })), o;
  } }, { key: "test", value: function(n, a) {
    return this.RULE[a].reg && this.RULE[a].reg.test(n);
  } }, { key: "rule", value: function() {
    return Eo();
  } }]), r;
}();
function Ye() {
  return (typeof window > "u" ? "undefined" : He(window)) === "object";
}
function ab(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R($f, "HOOK_NAME", "table");
var Rf = function(e) {
  U(r, se);
  var t = ab(r);
  function r(n) {
    var a;
    return j(this, r), (a = t.call(this, { needCache: !0 })).classicBr = Gu("classicBr") ? Ku() : n.globalConfig.classicBr, a;
  }
  return M(r, [{ key: "beforeMakeHtml", value: function(n) {
    var a = this;
    return this.test(n) ? n.replace(this.RULE.reg, function(i, o, s) {
      var c, u;
      if (s === 0)
        return i;
      var l, f, p = (c = (u = o.match(/\n/g)) === null || u === void 0 ? void 0 : u.length) !== null && c !== void 0 ? c : 0, d = "br".concat(p), g = "";
      Ye() ? g = a.classicBr ? b(l = '<span data-sign="'.concat(d, '" data-type="br" data-lines="')).call(l, p, '"></span>') : b(f = '<p data-sign="'.concat(d, '" data-type="br" data-lines="')).call(f, p, '">&nbsp;</p>') : g = a.classicBr ? "" : "<br/>";
      var _ = a.pushCache(g, d, p);
      return `

`.concat(_, `
`);
    }) : n;
  } }, { key: "makeHtml", value: function(n, a) {
    return n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:\\n)", end: "", content: "((?:\\h*\\n){2,})" };
    return n.reg = Le(n, "g", !0), n;
  } }]), r;
}();
function ib(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Rf, "HOOK_NAME", "br");
var Of = function(e) {
  U(r, se);
  var t = ib(r);
  function r() {
    return j(this, r), t.call(this, { needCache: !0 });
  }
  return M(r, [{ key: "beforeMakeHtml", value: function(n) {
    var a = this;
    return n.replace(this.RULE.reg, function(i, o) {
      var s, c = (o.match(/\n/g) || []).length + 1, u = "hr".concat(c);
      return nn(i, a.pushCache(b(s = '<hr data-sign="'.concat(u, '" data-lines="')).call(s, c, '" />'), u));
    });
  } }, { key: "makeHtml", value: function(n, a) {
    return n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:^|\\n)(\\n*)[ ]*", end: "(?=$|\\n)", content: "((?:-[ \\t]*){3,}|(?:\\*[ \\t]*){3,}|(?:_[ \\t]*){3,})" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
R(Of, "HOOK_NAME", "hr");
var ra = { processExtendAttributesInAlt: function(e) {
  var t = e.match(/#([0-9]+(px|em|pt|pc|in|mm|cm|ex|%)|auto)/g);
  if (!t)
    return "";
  var r = "", n = Ht(t, 2), a = n[0], i = n[1];
  return a && (r = ' width="'.concat(a.replace(/[ #]*/g, ""), '"')), i && (r += ' height="'.concat(i.replace(/[ #]*/g, ""), '"')), r;
}, processExtendStyleInAlt: function(e) {
  var t = this.$getAlignment(e), r = "", n = e.match(/#(border|shadow|radius|B|S|R)/g);
  if (n)
    for (var a = 0; a < n.length; a++)
      switch (n[a]) {
        case "#border":
        case "#B":
          t += "border:1px solid #888888;padding: 2px;box-sizing: border-box;", r += " cherry-img-border";
          break;
        case "#shadow":
        case "#S":
          t += "box-shadow:0 2px 15px -5px rgb(0 0 0 / 50%);", r += " cherry-img-shadow";
          break;
        case "#radius":
        case "#R":
          t += "border-radius: 15px;", r += " cherry-img-radius";
      }
  return { extendStyles: t, extendClasses: r };
}, $getAlignment: function(e) {
  var t = e.match(/#(center|right|left|float-right|float-left)/i);
  if (!t)
    return "";
  switch (Ht(t, 2)[1]) {
    case "center":
      return "transform:translateX(-50%);margin-left:50%;display:block;";
    case "right":
      return "transform:translateX(-100%);margin-left:100%;margin-right:-100%;display:block;";
    case "left":
      return "transform:translateX(0);margin-left:0;display:block;";
    case "float-right":
      return "float:right;transform:translateX(0);margin-left:0;display:block;";
    case "float-left":
      return "float:left;transform:translateX(0);margin-left:0;display:block;";
  }
} };
function al(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function ob(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var il = function(e, t, r, n, a, i, o, s, c) {
  var u = a === void 0 ? "ref" : "url", l = "";
  if (u === "ref")
    return t;
  if (u === "url") {
    var f, p, d, g, _, m, h, v, w = ra.processExtendAttributesInAlt(n), E = ra.processExtendStyleInAlt(n), S = E.extendStyles, A = E.extendClasses;
    S && (S = ' style="'.concat(S, '" ')), A && (A = ' class="'.concat(A, '" ')), l = i && N(i).call(i) !== "" ? ' title="'.concat(qe(i), '"') : "", o && (l += ' poster="'.concat(Nt(o), '"'));
    var x = c.urlProcessor(a, e), D = b(f = b(p = b(d = b(g = b(_ = b(m = b(h = "<".concat(e, ' src="')).call(h, St.set(Nt(x)), '"')).call(m, l, " ")).call(_, w, " ")).call(g, S, " ")).call(d, A, ' controls="controls">')).call(p, qe(n || ""), "</")).call(f, e, ">");
    return b(v = "".concat(r)).call(v, s.videoWrapper ? s.videoWrapper(a) : D);
  }
  return t;
}, Pf = function(e) {
  U(r, pe);
  var t = ob(r);
  function r(n) {
    var a, i = n.config, o = n.globalConfig;
    return j(this, r), (a = t.call(this, null)).urlProcessor = o.urlProcessor, a.extendMedia = { tag: ["video", "audio"], replacer: { video: function(s, c, u, l, f, p) {
      return il("video", s, c, u, l, f, p, i, o);
    }, audio: function(s, c, u, l, f, p) {
      return il("audio", s, c, u, l, f, p, i, o);
    } } }, a.RULE = a.rule(a.extendMedia), a;
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i, o, s, c, u) {
    var l = o === void 0 ? "ref" : "url", f = "";
    if (l === "ref")
      return n;
    if (l === "url") {
      var p, d, g, _, m, h, v, w, E = ra.processExtendAttributesInAlt(i), S = ra.processExtendStyleInAlt(i), A = S.extendStyles, x = S.extendClasses;
      A && (A = ' style="'.concat(A, '" ')), x && (x = ' class="'.concat(x, '" ')), f = s && N(s).call(s) !== "" ? ' title="'.concat(qe(s.replace(/["']/g, "")), '"') : "";
      var D, C = "src", W = this.$engine.$cherry.options;
      if (W.callback && W.callback.beforeImageMounted) {
        var V = W.callback.beforeImageMounted(C, o);
        C = V.srcProp || C, D = V.src || o;
      }
      var G = u ? u.replace(/[{}]/g, "").replace(/([^=\s]+)=([^\s]+)/g, '$1="$2"').replace(/&/g, "&amp;") : "";
      return b(p = b(d = b(g = b(_ = b(m = b(h = b(v = b(w = "".concat(a, "<img ")).call(w, C, '="')).call(v, St.set(Nt(this.urlProcessor(D, "image"))), '" ')).call(h, E, " ")).call(m, A, " ")).call(_, x, ' alt="')).call(g, qe(i || ""), '"')).call(d, f, " ")).call(p, G, "/>");
    }
    return n;
  } }, { key: "toMediaHtml", value: function(n, a, i, o, s, c, u, l, f) {
    var p, d;
    if (!this.extendMedia.replacer[i])
      return n;
    for (var g = arguments.length, _ = new Array(g > 9 ? g - 9 : 0), m = 9; m < g; m++)
      _[m - 9] = arguments[m];
    return (p = this.extendMedia.replacer[i]).call.apply(p, b(d = [this, n, a, o, s, c, f]).call(d, _));
  } }, { key: "makeHtml", value: function(n) {
    var a, i, o, s, c = n;
    return this.test(c) && (c = fe() ? c.replace(this.RULE.reg, je(a = this.toHtml).call(a, this)) : ct(c, this.RULE.reg, je(i = this.toHtml).call(i, this), !0, 1)), this.testMedia(c) && (c = fe() ? c.replace(this.RULE.regExtend, je(o = this.toMediaHtml).call(o, this)) : ct(c, this.RULE.regExtend, je(s = this.toMediaHtml).call(s, this), !0, 1)), c;
  } }, { key: "testMedia", value: function(n) {
    return this.RULE.regExtend && this.RULE.regExtend.test(n);
  } }, { key: "rule", value: function(n) {
    var a = { begin: fe() ? "((?<!\\\\))!" : "(^|[^\\\\])!", content: ["\\[([^\\n]*?)\\]", "[ \\t]*", "".concat(`(?:\\(([^"][^\\s]+?)(?:[ \\t]((?:".*?")|(?:'.*?')))?\\)|\\[(`).concat(nf, ")\\]") + ")"].join(""), end: "({[^{}]+?})?" };
    if (n) {
      var i = function(o) {
        for (var s = 1; s < arguments.length; s++) {
          var c, u, l = arguments[s] != null ? arguments[s] : {};
          s % 2 ? q(c = al(Object(l), !0)).call(c, function(f) {
            R(o, f, l[f]);
          }) : Ce ? Ct(o, Ce(l)) : q(u = al(Object(l))).call(u, function(f) {
            nt(o, f, Te(l, f));
          });
        }
        return o;
      }({}, a);
      i.begin = fe() ? "((?<!\\\\))!(".concat(n.tag.join("|"), ")") : "(^|[^\\\\])!(".concat(n.tag.join("|"), ")"), i.end = "({poster=(.*)})?", a.regExtend = Le(i, "g");
    }
    return a.reg = Le(a, "g"), a;
  } }]), r;
}();
function ol(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function sl(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = ol(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = ol(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function sb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Pf, "HOOK_NAME", "image");
function cl(e) {
  var t;
  if (He(e) !== "object" && ue(e).length < 1)
    return "";
  var r = [""];
  return q(t = ue(e)).call(t, function(n) {
    var a;
    r.push(b(a = "".concat(n, '="')).call(a, e[n], '"'));
  }), r.join(" ");
}
function cb(e, t) {
  for (var r = /^(\t|[ ])/, n = e; r.test(n); )
    t.space += n[0] === "	" ? 4 : 1, n = n.replace(r, "");
  return n;
}
function lb(e, t) {
  var r = /^((([*+-]|\d+[.]|[a-z]\.|[I一二三四五六七八九十]+\.)[ \t]+)([^\r]*?)($|\n{2,}(?=\S)(?![ \t]*(?:[*+-]|\d+[.]|[a-z]\.|[I一二三四五六七八九十]+\.)[ \t]+)))/;
  return r.test(e) ? e.replace(r, function(n, a, i, o, s) {
    return t.type = i.search(/[*+-]/g) > -1 ? "ul" : "ol", t.listStyle = function(c) {
      return /^[a-z]/.test(c) ? "lower-greek" : /^[一二三四五六七八九十]/.test(c) ? "cjk-ideographic" : /^I/.test(c) ? "upper-roman" : /^\+/.test(c) ? "circle" : /^\*/.test(c) ? "square" : "default";
    }(i), t.start = Number(i.replace(".", "")) ? Number(i.replace(".", "")) : 1, s;
  }) : (t.type = "blank", e);
}
var ub = M(function e() {
  j(this, e), this.index = 0, this.space = 0, this.type = "", this.start = 1, this.listStyle = "", this.strs = [], this.children = [], this.lines = 0;
}), Lf = function(e) {
  U(r, se);
  var t = sb(r);
  function r(n) {
    var a, i = n.config;
    return j(this, r), (a = t.call(this, { needCache: !0 })).config = i || {}, a.tree = [], a.emptyLines = 0, a.indentSpace = Math.max(a.config.indentSpace, 2), a;
  }
  return M(r, [{ key: "addNode", value: function(n, a, i, o) {
    n.type === "blank" ? this.tree[o].strs.push(n.strs[0]) : (this.tree[i].children.push(a), this.tree[a] = sl(sl({}, n), {}, { parent: i }));
  } }, { key: "buildTree", value: function(n, a) {
    var i = n.split(`
`);
    this.tree = [], i.unshift("");
    for (var o = n.match(/\n*$/g)[0].length, s = 0; s < i.length - o; s++) {
      var c = new ub();
      if (i[s] = cb(i[s], c), i[s] = lb(i[s], c), c.strs.push(a(i[s]).html), c.index = s, s !== 0) {
        for (var u = s - 1; !this.tree[u]; )
          u -= 1;
        if (c.type === "blank")
          this.addNode(c, s, this.tree[u].parent, u);
        else {
          for (; !this.tree[u] || this.tree[u].space > c.space; )
            u -= 1;
          var l = c.space, f = this.tree[u].space;
          l < f + this.indentSpace ? this.config.listNested && this.tree[u].type !== c.type ? this.addNode(c, s, u) : this.addNode(c, s, this.tree[u].parent) : l < f + this.indentSpace + 4 ? this.addNode(c, s, u) : (c.type = "blank", this.addNode(c, s, this.tree[u].parent, u));
        }
      } else
        c.space = -2, this.tree.push(c);
    }
  } }, { key: "renderSubTree", value: function(n, a, i) {
    var o, s, c, u = this, l = 0, f = {}, p = na(a).call(a, function(d, g) {
      var _, m, h, v = u.tree[g], w = { class: "cherry-list-item" }, E = "<p>".concat(v.strs.join("<br>"), "</p>");
      v.lines += u.getLineCount(v.strs.join(`
`));
      var S = v.children.length ? u.renderTree(g) : "";
      return n.lines += v.lines, l += v.lines, /<span class="ch-icon ch-icon-(square|check)"><\/span>/.test(E) && (w.class += " check-list-item"), b(_ = b(m = b(h = "".concat(d, "<li")).call(h, cl(w), ">")).call(m, E)).call(_, S, "</li>");
    }, "");
    return n.parent === void 0 && (f["data-lines"] = n.index === 0 ? l + this.emptyLines : l, f["data-sign"] = this.sign), a[0] && i === "ol" && (f.start = this.tree[a[0]].start), f.class = "cherry-list__".concat(this.tree[a[0]].listStyle), b(o = b(s = b(c = "<".concat(i)).call(c, cl(f), ">")).call(s, p, "</")).call(o, i, ">");
  } }, { key: "renderTree", value: function(n) {
    var a = this, i = 0, o = this.tree[n], s = o.children;
    return na(s).call(s, function(c, u, l) {
      if (l === 0 || a.tree[s[l]].type === a.tree[s[l - 1]].type)
        return c;
      var f = a.renderSubTree(o, Ke(s).call(s, i, l), a.tree[s[l - 1]].type);
      return i = l, c + f;
    }, "") + (s.length ? this.renderSubTree(o, Ke(s).call(s, i, s.length), this.tree[s[s.length - 1]].type) : "");
  } }, { key: "toHtml", value: function(n, a) {
    var i, o;
    this.emptyLines = (i = (o = n.match(/^\n\n/)) === null || o === void 0 ? void 0 : o.length) !== null && i !== void 0 ? i : 0;
    var s = n.replace(/~0$/g, "").replace(/^\n+/, "");
    this.buildTree(function(u) {
      return u.replace(/^((?:|[\t ]+)[*+-]\s+)\[(\s|x)\]/gm, function(l, f, p) {
        var d, g = /\s/.test(p) ? '<span class="ch-icon ch-icon-square"></span>' : '<span class="ch-icon ch-icon-check"></span>';
        return b(d = "".concat(f)).call(d, g);
      });
    }(s), a);
    var c = this.renderTree(0);
    return this.pushCache(c, this.sign, this.$getLineNum(n));
  } }, { key: "$getLineNum", value: function(n) {
    var a, i, o, s, c = (a = (i = n.match(/^\n\n/)) === null || i === void 0 ? void 0 : i.length) !== null && a !== void 0 ? a : 0;
    return (o = (s = n.replace(/^\n+/, "").replace(/\n+$/, `
`).match(/\n/g)) === null || s === void 0 ? void 0 : s.length) !== null && o !== void 0 ? o : 0 + c;
  } }, { key: "makeHtml", value: function(n, a) {
    var i = this, o = "".concat(n, "~0");
    return this.test(o) && (o = o.replace(this.RULE.reg, function(s) {
      return i.getCacheWithSpace(i.checkCache(s, a, i.$getLineNum(s)), s);
    })), o = o.replace(/~0$/g, "");
  } }, { key: "rule", value: function() {
    var n = { begin: `(?:^|
)(
*)(([ ]{0,3}([*+-]|\\d+[.]|[a-z]\\.|[I一二三四五六七八九十]+\\.)[ \\t]+)`, content: "([^\\r]+?)", end: "(~0|\\n{2,}(?=\\S)(?![ \\t]*(?:[*+-]|\\d+[.]|[a-z]\\.|[I一二三四五六七八九十]+\\.)[ \\t]+)))" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "gm"), n;
  } }]), r;
}();
function fb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
function ll(e) {
  for (var t = /^(\t|[ ]{1,4})/, r = e, n = 0; t.test(r); )
    r = r.replace(/^(\t|[ ]{1,4})/g, ""), n += 1;
  return n;
}
R(Lf, "HOOK_NAME", "list");
var If = function(e) {
  U(r, se);
  var t = fb(r);
  function r() {
    return j(this, r), t.call(this, { needCache: !0 });
  }
  return M(r, [{ key: "handleMatch", value: function(n, a) {
    var i = this;
    return n.replace(this.RULE.reg, function(o, s, c) {
      for (var u, l, f = a(c), p = f.sign, d = f.html, g = i.signWithCache(d) || p, _ = i.getLineCount(o, s), m = /^(([ \t]{0,3}([*+-]|\d+[.]|[a-z]\.|[I一二三四五六七八九十]+\.)[ \t]+)([^\r]+?)($|\n{2,}(?=\S)(?![ \t]*(?:[*+-]|\d+[.]|[a-z]\.|[I一二三四五六七八九十]+\.)[ \t]+)))/, h = ll(s), v = d.split(`
`), w = /^[>\s]+/, E = />/g, S = 1, A = 0, x = b(u = b(l = '<blockquote data-sign="'.concat(g, "_")).call(l, _, '" data-lines="')).call(u, _, '">'), D = 0; v[D]; D++) {
        if (D !== 0) {
          var C = ll(v[D]);
          if (C <= h && m.test(v[D]))
            break;
          h = C;
        }
        var W = v[D].replace(w, function(V) {
          var G = V.match(E);
          return A = G && G.length > S ? G.length : S, "";
        });
        S === A && D !== 0 && (x += "<br>"), S < A && (x += gt("<blockquote>").call("<blockquote>", A - S), S = A), x += W;
      }
      return x += gt("</blockquote>").call("</blockquote>", S), i.getCacheWithSpace(i.pushCache(x, g, _), o);
    });
  } }, { key: "makeHtml", value: function(n, a) {
    return this.test(n) ? this.handleMatch(n, a) : n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:^|\\n)(\\s*)", content: ["(", ">(?:.+?\\n(?![*+-]|\\d+[.]|[a-z]\\.))(?:>*.+?\\n(?![*+-]|\\d+[.]|[a-z]\\.))*(?:>*.+?)", "|", ">(?:.+?)", ")"].join(""), end: "(?=(\\n)|$)" };
    return n.reg = Le(n, "g"), n;
  } }]), r;
}();
function db(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(If, "HOOK_NAME", "blockquote");
var Nf = function(e) {
  U(r, pe);
  var t = db(r);
  function r(n) {
    var a, i = n.config, o = n.globalConfig;
    return j(this, r), (a = t.call(this, { config: i })).urlProcessor = o.urlProcessor, a.enableShortLink = !!i.enableShortLink, a.shortLinkLength = i.shortLinkLength, a.target = i.target ? 'target="'.concat(i.target, '"') : i.openNewPage ? 'target="_blank"' : "", a.rel = i.rel ? 'rel="'.concat(i.rel, '"') : "", a;
  }
  return M(r, [{ key: "isLinkInHtmlAttribute", value: function(n, a, i) {
    for (var o, s = new RegExp(["<", "([a-zA-Z][a-zA-Z0-9-]*)", "(", ["\\s+[a-zA-Z_:][a-zA-Z0-9_.:-]*", "(", ["\\s*=\\s*", "(", ["([^\\s\"'=<>`]+)", "('[^']*')", '("[^"]*")'].join("|"), ")"].join(""), ")?"].join(""), ")*", "\\s*[/]?>"].join(""), "g"); (o = s.exec(n)) !== null && !(o.index > a + i); )
      if (o.index < a && o.index + o[0].length >= a + i)
        return !0;
    return !1;
  } }, { key: "isLinkInATag", value: function(n, a, i) {
    for (var o, s = /<a.*>[^<]*<\/a>/g; (o = s.exec(n)) !== null && !(o.index > a + i); )
      if (o.index < a && o.index + o[0].length >= a + i)
        return !0;
    return !1;
  } }, { key: "makeHtml", value: function(n, a) {
    var i = this;
    return this.test(n) && (Gi.test(n) || wo.test(n)) ? n.replace(this.RULE.reg, function(o, s, c, u, l, f, p) {
      var d, g, _;
      if (i.isLinkInHtmlAttribute(p, f, c.length + u.length) || i.isLinkInATag(p, f, c.length + u.length))
        return o;
      var m = c.toLowerCase(), h = "", v = "", w = !0;
      if ((s !== "<" && s !== "&#60;" || l !== ">" && l !== "&#62;") && (h = s, v = l, w = !1), N(u).call(u) === "" || !w && m === "" && !/www\./.test(u))
        return o;
      switch (m) {
        case "javascript:":
          return o;
        case "mailto:":
          var E, S, A, x, D, C;
          return di.test(u) ? b(E = b(S = b(A = b(x = b(D = "".concat(h, '<a href="')).call(D, Nt(b(C = "".concat(m)).call(C, u)), '" ')).call(x, i.target, " ")).call(A, i.rel, ">")).call(S, qe(u), "</a>")).call(E, v) : o;
        case "":
          var W, V, G, we, ge, Ne, Q, ce, $e, he;
          if (h === v || !w)
            return di.test(u) ? b(W = b(V = b(G = b(we = b(ge = "".concat(h, '<a href="mailto:')).call(ge, Nt(u), '" ')).call(we, i.target, " ")).call(G, i.rel, ">")).call(V, qe(u), "</a>")).call(W, v) : Uc.test(u) ? b(Ne = b(Q = "".concat(h)).call(Q, i.renderLink("//".concat(u), u))).call(Ne, v) : o;
          if (w)
            return di.test(u) ? b(ce = b($e = b(he = '<a href="mailto:'.concat(Nt(u), '" ')).call(he, i.target, " ")).call($e, i.rel, ">")).call(ce, qe(u), "</a>") : Wc.test(u) || Uc.test(u) ? i.renderLink(u) : o;
        default:
          return Wc.test(u) ? b(d = b(g = "".concat(h)).call(g, i.renderLink(b(_ = "".concat(m)).call(_, u)))).call(d, v) : o;
      }
      return o;
    }) : n;
  } }, { key: "rule", value: function() {
    var n, a = { begin: "(<?)", content: ["((?:[a-z][a-z0-9+.-]{1,31}:)?)", b(n = "((?:".concat(af.source, ")|(?:")).call(n, Gi.source, "))")].join(""), end: "(>?)" };
    return a.reg = Le(a, "ig"), a;
  } }, { key: "renderLink", value: function(n, a) {
    var i, o, s, c, u = a;
    if (typeof u != "string")
      if (this.enableShortLink) {
        var l, f = n.replace(/^https?:\/\//i, "");
        u = b(l = "".concat(f.substring(0, this.shortLinkLength))).call(l, f.length > this.shortLinkLength ? "..." : "");
      } else
        u = n;
    var p = this.urlProcessor(n, "autolink");
    return b(i = b(o = b(s = b(c = "<a ".concat(this.target, " ")).call(c, this.rel, ' title="')).call(s, qe(n).replace(/_/g, "\\_"), '"  href="')).call(o, Nt(p).replace(/_/g, "\\_"), '">')).call(i, qe(u).replace(/_/g, "\\_"), "</a>");
  } }]), r;
}();
function aa() {
  var e, t, r, n;
  Ye() && (this.katex = (e = (t = this.externals) === null || t === void 0 ? void 0 : t.katex) !== null && e !== void 0 ? e : window.katex, this.MathJax = (r = (n = this.externals) === null || n === void 0 ? void 0 : n.MathJax) !== null && r !== void 0 ? r : window.MathJax);
}
R(Nf, "HOOK_NAME", "autoLink");
var ul = ["&", "<", ">", '"', "'"], Mf = function(e) {
  var t = e.replace(new RegExp(rf, "g"), function(r) {
    return Me(ul).call(ul, r) !== -1 ? Mn(r) : "\\".concat(r);
  });
  return t;
}, pb = _o.trim, gb = z("".charAt), Fr = O.parseFloat, fl = O.Symbol, dl = fl && fl.iterator, pl = 1 / Fr(Nn + "-0") != -1 / 0 || dl && !J(function() {
  Fr(Object(dl));
}) ? function(e) {
  var t = pb(tt(e)), r = Fr(t);
  return r === 0 && gb(t, 0) == "-" ? -0 : r;
} : Fr;
L({ global: !0, forced: parseFloat != pl }, { parseFloat: pl }), X.parseFloat;
var hb = lf, mb = _e("match"), bb = O.TypeError, jf = function(e) {
  if (function(t) {
    var r;
    return ye(t) && ((r = t[mb]) !== void 0 ? !!r : ht(t) == "RegExp");
  }(e))
    throw bb("The method doesn't accept regular expressions");
  return e;
}, vb = _e("match"), Df = function(e) {
  var t = /./;
  try {
    "/./"[e](t);
  } catch {
    try {
      return t[vb] = !1, "/./"[e](t);
    } catch {
    }
  }
  return !1;
};
on.f;
var gl = z("".startsWith), yb = z("".slice), _b = Math.min, kb = Df("startsWith");
L({ target: "String", proto: !0, forced: !kb }, { startsWith: function(e) {
  var t = tt(an(this));
  jf(e);
  var r = Jl(_b(arguments.length > 1 ? arguments[1] : void 0, t.length)), n = tt(e);
  return gl ? gl(t, n, r) : yb(t, r, r + n.length) === n;
} });
var wb = Ie("String").startsWith, vi = String.prototype, Ff = function(e) {
  var t = e.startsWith;
  return typeof e == "string" || e === vi || de(vi, e) && t === vi.startsWith ? wb : t;
};
function Bf(e, t) {
  if (!e || !e.tagName)
    return "";
  var r, n, a = document.createElement("div");
  return a.appendChild(e.cloneNode(!1)), r = a.innerHTML, t && (n = Me(r).call(r, ">") + 1, r = r.substring(0, n) + e.innerHTML + r.substring(n)), a = null, r;
}
function Br(e) {
  var t, r = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "", n = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : {}, a = document.createElement(e);
  return a.className = r, n !== void 0 && q(t = ue(n)).call(t, function(i) {
    var o = n[i];
    if (Ff(i).call(i, "data-")) {
      var s = i.replace(/^data-/, "");
      a.dataset[s] = o;
    } else
      a.setAttribute(i, o);
  }), a;
}
var Eb = rt(function(e, t) {
  e.exports = function() {
    const { entries: r, setPrototypeOf: n, isFrozen: a, getPrototypeOf: i, getOwnPropertyDescriptor: o } = Object;
    let { freeze: s, seal: c, create: u } = Object, { apply: l, construct: f } = typeof Reflect < "u" && Reflect;
    s || (s = function(H) {
      return H;
    }), c || (c = function(H) {
      return H;
    }), l || (l = function(H, I, ee) {
      return H.apply(I, ee);
    }), f || (f = function(H, I) {
      return new H(...I);
    });
    const p = x(Array.prototype.forEach), d = x(Array.prototype.pop), g = x(Array.prototype.push), _ = x(String.prototype.toLowerCase), m = x(String.prototype.toString), h = x(String.prototype.match), v = x(String.prototype.replace), w = x(String.prototype.indexOf), E = x(String.prototype.trim), S = x(RegExp.prototype.test), A = D(TypeError);
    function x(H) {
      return function(I) {
        for (var ee = arguments.length, Y = new Array(ee > 1 ? ee - 1 : 0), De = 1; De < ee; De++)
          Y[De - 1] = arguments[De];
        return l(H, I, Y);
      };
    }
    function D(H) {
      return function() {
        for (var I = arguments.length, ee = new Array(I), Y = 0; Y < I; Y++)
          ee[Y] = arguments[Y];
        return f(H, ee);
      };
    }
    function C(H, I) {
      let ee = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : _;
      n && n(H, null);
      let Y = I.length;
      for (; Y--; ) {
        let De = I[Y];
        if (typeof De == "string") {
          const it = ee(De);
          it !== De && (a(I) || (I[Y] = it), De = it);
        }
        H[De] = !0;
      }
      return H;
    }
    function W(H) {
      for (let I = 0; I < H.length; I++)
        o(H, I) === void 0 && (H[I] = null);
      return H;
    }
    function V(H) {
      const I = u(null);
      for (const [ee, Y] of r(H))
        o(H, ee) !== void 0 && (Array.isArray(Y) ? I[ee] = W(Y) : Y && typeof Y == "object" && Y.constructor === Object ? I[ee] = V(Y) : I[ee] = Y);
      return I;
    }
    function G(H, I) {
      for (; H !== null; ) {
        const Y = o(H, I);
        if (Y) {
          if (Y.get)
            return x(Y.get);
          if (typeof Y.value == "function")
            return x(Y.value);
        }
        H = i(H);
      }
      function ee(Y) {
        return console.warn("fallback value for", Y), null;
      }
      return ee;
    }
    const we = s(["a", "abbr", "acronym", "address", "area", "article", "aside", "audio", "b", "bdi", "bdo", "big", "blink", "blockquote", "body", "br", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "content", "data", "datalist", "dd", "decorator", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "element", "em", "fieldset", "figcaption", "figure", "font", "footer", "form", "h1", "h2", "h3", "h4", "h5", "h6", "head", "header", "hgroup", "hr", "html", "i", "img", "input", "ins", "kbd", "label", "legend", "li", "main", "map", "mark", "marquee", "menu", "menuitem", "meter", "nav", "nobr", "ol", "optgroup", "option", "output", "p", "picture", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "section", "select", "shadow", "small", "source", "spacer", "span", "strike", "strong", "style", "sub", "summary", "sup", "table", "tbody", "td", "template", "textarea", "tfoot", "th", "thead", "time", "tr", "track", "tt", "u", "ul", "var", "video", "wbr"]), ge = s(["svg", "a", "altglyph", "altglyphdef", "altglyphitem", "animatecolor", "animatemotion", "animatetransform", "circle", "clippath", "defs", "desc", "ellipse", "filter", "font", "g", "glyph", "glyphref", "hkern", "image", "line", "lineargradient", "marker", "mask", "metadata", "mpath", "path", "pattern", "polygon", "polyline", "radialgradient", "rect", "stop", "style", "switch", "symbol", "text", "textpath", "title", "tref", "tspan", "view", "vkern"]), Ne = s(["feBlend", "feColorMatrix", "feComponentTransfer", "feComposite", "feConvolveMatrix", "feDiffuseLighting", "feDisplacementMap", "feDistantLight", "feDropShadow", "feFlood", "feFuncA", "feFuncB", "feFuncG", "feFuncR", "feGaussianBlur", "feImage", "feMerge", "feMergeNode", "feMorphology", "feOffset", "fePointLight", "feSpecularLighting", "feSpotLight", "feTile", "feTurbulence"]), Q = s(["animate", "color-profile", "cursor", "discard", "font-face", "font-face-format", "font-face-name", "font-face-src", "font-face-uri", "foreignobject", "hatch", "hatchpath", "mesh", "meshgradient", "meshpatch", "meshrow", "missing-glyph", "script", "set", "solidcolor", "unknown", "use"]), ce = s(["math", "menclose", "merror", "mfenced", "mfrac", "mglyph", "mi", "mlabeledtr", "mmultiscripts", "mn", "mo", "mover", "mpadded", "mphantom", "mroot", "mrow", "ms", "mspace", "msqrt", "mstyle", "msub", "msup", "msubsup", "mtable", "mtd", "mtext", "mtr", "munder", "munderover", "mprescripts"]), $e = s(["maction", "maligngroup", "malignmark", "mlongdiv", "mscarries", "mscarry", "msgroup", "mstack", "msline", "msrow", "semantics", "annotation", "annotation-xml", "mprescripts", "none"]), he = s(["#text"]), Qe = s(["accept", "action", "align", "alt", "autocapitalize", "autocomplete", "autopictureinpicture", "autoplay", "background", "bgcolor", "border", "capture", "cellpadding", "cellspacing", "checked", "cite", "class", "clear", "color", "cols", "colspan", "controls", "controlslist", "coords", "crossorigin", "datetime", "decoding", "default", "dir", "disabled", "disablepictureinpicture", "disableremoteplayback", "download", "draggable", "enctype", "enterkeyhint", "face", "for", "headers", "height", "hidden", "high", "href", "hreflang", "id", "inputmode", "integrity", "ismap", "kind", "label", "lang", "list", "loading", "loop", "low", "max", "maxlength", "media", "method", "min", "minlength", "multiple", "muted", "name", "nonce", "noshade", "novalidate", "nowrap", "open", "optimum", "pattern", "placeholder", "playsinline", "poster", "preload", "pubdate", "radiogroup", "readonly", "rel", "required", "rev", "reversed", "role", "rows", "rowspan", "spellcheck", "scope", "selected", "shape", "size", "sizes", "span", "srclang", "start", "src", "srcset", "step", "style", "summary", "tabindex", "title", "translate", "type", "usemap", "valign", "value", "width", "xmlns", "slot"]), at = s(["accent-height", "accumulate", "additive", "alignment-baseline", "ascent", "attributename", "attributetype", "azimuth", "basefrequency", "baseline-shift", "begin", "bias", "by", "class", "clip", "clippathunits", "clip-path", "clip-rule", "color", "color-interpolation", "color-interpolation-filters", "color-profile", "color-rendering", "cx", "cy", "d", "dx", "dy", "diffuseconstant", "direction", "display", "divisor", "dur", "edgemode", "elevation", "end", "fill", "fill-opacity", "fill-rule", "filter", "filterunits", "flood-color", "flood-opacity", "font-family", "font-size", "font-size-adjust", "font-stretch", "font-style", "font-variant", "font-weight", "fx", "fy", "g1", "g2", "glyph-name", "glyphref", "gradientunits", "gradienttransform", "height", "href", "id", "image-rendering", "in", "in2", "k", "k1", "k2", "k3", "k4", "kerning", "keypoints", "keysplines", "keytimes", "lang", "lengthadjust", "letter-spacing", "kernelmatrix", "kernelunitlength", "lighting-color", "local", "marker-end", "marker-mid", "marker-start", "markerheight", "markerunits", "markerwidth", "maskcontentunits", "maskunits", "max", "mask", "media", "method", "mode", "min", "name", "numoctaves", "offset", "operator", "opacity", "order", "orient", "orientation", "origin", "overflow", "paint-order", "path", "pathlength", "patterncontentunits", "patterntransform", "patternunits", "points", "preservealpha", "preserveaspectratio", "primitiveunits", "r", "rx", "ry", "radius", "refx", "refy", "repeatcount", "repeatdur", "restart", "result", "rotate", "scale", "seed", "shape-rendering", "specularconstant", "specularexponent", "spreadmethod", "startoffset", "stddeviation", "stitchtiles", "stop-color", "stop-opacity", "stroke-dasharray", "stroke-dashoffset", "stroke-linecap", "stroke-linejoin", "stroke-miterlimit", "stroke-opacity", "stroke", "stroke-width", "style", "surfacescale", "systemlanguage", "tabindex", "targetx", "targety", "transform", "transform-origin", "text-anchor", "text-decoration", "text-rendering", "textlength", "type", "u1", "u2", "unicode", "values", "viewbox", "visibility", "version", "vert-adv-y", "vert-origin-x", "vert-origin-y", "width", "word-spacing", "wrap", "writing-mode", "xchannelselector", "ychannelselector", "x", "x1", "x2", "xmlns", "y", "y1", "y2", "z", "zoomandpan"]), Re = s(["accent", "accentunder", "align", "bevelled", "close", "columnsalign", "columnlines", "columnspan", "denomalign", "depth", "dir", "display", "displaystyle", "encoding", "fence", "frame", "height", "href", "id", "largeop", "length", "linethickness", "lspace", "lquote", "mathbackground", "mathcolor", "mathsize", "mathvariant", "maxsize", "minsize", "movablelimits", "notation", "numalign", "open", "rowalign", "rowlines", "rowspacing", "rowspan", "rspace", "rquote", "scriptlevel", "scriptminsize", "scriptsizemultiplier", "selection", "separator", "separators", "stretchy", "subscriptshift", "supscriptshift", "symmetric", "voffset", "width", "xmlns"]), ke = s(["xlink:href", "xml:id", "xlink:title", "xml:space", "xmlns:xlink"]), Ge = c(/\{\{[\w\W]*|[\w\W]*\}\}/gm), Tt = c(/<%[\w\W]*|[\w\W]*%>/gm), Bn = c(/\${[\w\W]*}/gm), yr = c(/^data-[\-\w.\u00B7-\uFFFF]/), Hn = c(/^aria-[\-\w]+$/), qt = c(/^(?:(?:(?:f|ht)tps?|mailto|tel|callto|sms|cid|xmpp):|[^a-z]|[a-z+.\-]+(?:[^a-z+.\-:]|$))/i), zn = c(/^(?:\w+script|data):/i), pd = c(/[\u0000-\u0020\u00A0\u1680\u180E\u2000-\u2029\u205F\u3000]/g), $o = c(/^html$/i);
    var Ro = Object.freeze({ __proto__: null, MUSTACHE_EXPR: Ge, ERB_EXPR: Tt, TMPLIT_EXPR: Bn, DATA_ATTR: yr, ARIA_ATTR: Hn, IS_ALLOWED_URI: qt, IS_SCRIPT_OR_DATA: zn, ATTR_WHITESPACE: pd, DOCTYPE_NAME: $o });
    const gd = function() {
      return typeof window > "u" ? null : window;
    }, hd = function(H, I) {
      if (typeof H != "object" || typeof H.createPolicy != "function")
        return null;
      let ee = null;
      const Y = "data-tt-policy-suffix";
      I && I.hasAttribute(Y) && (ee = I.getAttribute(Y));
      const De = "dompurify" + (ee ? "#" + ee : "");
      try {
        return H.createPolicy(De, { createHTML: (it) => it, createScriptURL: (it) => it });
      } catch {
        return console.warn("TrustedTypes policy " + De + " could not be created."), null;
      }
    };
    function Oo() {
      let H = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : gd();
      const I = (y) => Oo(y);
      if (I.version = "3.0.8", I.removed = [], !H || !H.document || H.document.nodeType !== 9)
        return I.isSupported = !1, I;
      let { document: ee } = H;
      const Y = ee, De = Y.currentScript, { DocumentFragment: it, HTMLTemplateElement: bd, Node: ya, Element: Po, NodeFilter: _a, NamedNodeMap: vd = H.NamedNodeMap || H.MozNamedAttrMap, HTMLFormElement: yd, DOMParser: _d, trustedTypes: _r } = H, kr = Po.prototype, kd = G(kr, "cloneNode"), wd = G(kr, "nextSibling"), Ed = G(kr, "childNodes"), ka = G(kr, "parentNode");
      if (typeof bd == "function") {
        const y = ee.createElement("template");
        y.content && y.content.ownerDocument && (ee = y.content.ownerDocument);
      }
      let Fe, Un = "";
      const { implementation: wa, createNodeIterator: Sd, createDocumentFragment: Ad, getElementsByTagName: xd } = ee, { importNode: Cd } = Y;
      let ut = {};
      I.isSupported = typeof r == "function" && typeof ka == "function" && wa && wa.createHTMLDocument !== void 0;
      const { MUSTACHE_EXPR: Ea, ERB_EXPR: Sa, TMPLIT_EXPR: Aa, DATA_ATTR: Td, ARIA_ATTR: $d, IS_SCRIPT_OR_DATA: Rd, ATTR_WHITESPACE: Lo } = Ro;
      let { IS_ALLOWED_URI: Io } = Ro, Ee = null;
      const No = C({}, [...we, ...ge, ...Ne, ...ce, ...he]);
      let Se = null;
      const Mo = C({}, [...Qe, ...at, ...Re, ...ke]);
      let me = Object.seal(u(null, { tagNameCheck: { writable: !0, configurable: !1, enumerable: !0, value: null }, attributeNameCheck: { writable: !0, configurable: !1, enumerable: !0, value: null }, allowCustomizedBuiltInElements: { writable: !0, configurable: !1, enumerable: !0, value: !1 } })), Wn = null, xa = null, jo = !0, Ca = !0, Do = !1, Fo = !0, fn = !1, Gt = !1, Ta = !1, $a = !1, dn = !1, wr = !1, Er = !1, Bo = !0, Ho = !1;
      const Od = "user-content-";
      let Ra = !0, qn = !1, pn = {}, gn = null;
      const zo = C({}, ["annotation-xml", "audio", "colgroup", "desc", "foreignobject", "head", "iframe", "math", "mi", "mn", "mo", "ms", "mtext", "noembed", "noframes", "noscript", "plaintext", "script", "style", "svg", "template", "thead", "title", "video", "xmp"]);
      let Uo = null;
      const Wo = C({}, ["audio", "video", "img", "source", "image", "track"]);
      let Oa = null;
      const qo = C({}, ["alt", "class", "for", "id", "label", "name", "pattern", "placeholder", "role", "summary", "title", "value", "style", "xmlns"]), Sr = "http://www.w3.org/1998/Math/MathML", Ar = "http://www.w3.org/2000/svg", _t = "http://www.w3.org/1999/xhtml";
      let hn = _t, Pa = !1, La = null;
      const Pd = C({}, [Sr, Ar, _t], m);
      let Gn = null;
      const Ld = ["application/xhtml+xml", "text/html"], Id = "text/html";
      let Ae = null, mn = null;
      const Nd = ee.createElement("form"), Go = function(y) {
        return y instanceof RegExp || y instanceof Function;
      }, Ia = function() {
        let y = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : {};
        if (!mn || mn !== y) {
          if (y && typeof y == "object" || (y = {}), y = V(y), Gn = Ld.indexOf(y.PARSER_MEDIA_TYPE) === -1 ? Id : y.PARSER_MEDIA_TYPE, Ae = Gn === "application/xhtml+xml" ? m : _, Ee = "ALLOWED_TAGS" in y ? C({}, y.ALLOWED_TAGS, Ae) : No, Se = "ALLOWED_ATTR" in y ? C({}, y.ALLOWED_ATTR, Ae) : Mo, La = "ALLOWED_NAMESPACES" in y ? C({}, y.ALLOWED_NAMESPACES, m) : Pd, Oa = "ADD_URI_SAFE_ATTR" in y ? C(V(qo), y.ADD_URI_SAFE_ATTR, Ae) : qo, Uo = "ADD_DATA_URI_TAGS" in y ? C(V(Wo), y.ADD_DATA_URI_TAGS, Ae) : Wo, gn = "FORBID_CONTENTS" in y ? C({}, y.FORBID_CONTENTS, Ae) : zo, Wn = "FORBID_TAGS" in y ? C({}, y.FORBID_TAGS, Ae) : {}, xa = "FORBID_ATTR" in y ? C({}, y.FORBID_ATTR, Ae) : {}, pn = "USE_PROFILES" in y && y.USE_PROFILES, jo = y.ALLOW_ARIA_ATTR !== !1, Ca = y.ALLOW_DATA_ATTR !== !1, Do = y.ALLOW_UNKNOWN_PROTOCOLS || !1, Fo = y.ALLOW_SELF_CLOSE_IN_ATTR !== !1, fn = y.SAFE_FOR_TEMPLATES || !1, Gt = y.WHOLE_DOCUMENT || !1, dn = y.RETURN_DOM || !1, wr = y.RETURN_DOM_FRAGMENT || !1, Er = y.RETURN_TRUSTED_TYPE || !1, $a = y.FORCE_BODY || !1, Bo = y.SANITIZE_DOM !== !1, Ho = y.SANITIZE_NAMED_PROPS || !1, Ra = y.KEEP_CONTENT !== !1, qn = y.IN_PLACE || !1, Io = y.ALLOWED_URI_REGEXP || qt, hn = y.NAMESPACE || _t, me = y.CUSTOM_ELEMENT_HANDLING || {}, y.CUSTOM_ELEMENT_HANDLING && Go(y.CUSTOM_ELEMENT_HANDLING.tagNameCheck) && (me.tagNameCheck = y.CUSTOM_ELEMENT_HANDLING.tagNameCheck), y.CUSTOM_ELEMENT_HANDLING && Go(y.CUSTOM_ELEMENT_HANDLING.attributeNameCheck) && (me.attributeNameCheck = y.CUSTOM_ELEMENT_HANDLING.attributeNameCheck), y.CUSTOM_ELEMENT_HANDLING && typeof y.CUSTOM_ELEMENT_HANDLING.allowCustomizedBuiltInElements == "boolean" && (me.allowCustomizedBuiltInElements = y.CUSTOM_ELEMENT_HANDLING.allowCustomizedBuiltInElements), fn && (Ca = !1), wr && (dn = !0), pn && (Ee = C({}, he), Se = [], pn.html === !0 && (C(Ee, we), C(Se, Qe)), pn.svg === !0 && (C(Ee, ge), C(Se, at), C(Se, ke)), pn.svgFilters === !0 && (C(Ee, Ne), C(Se, at), C(Se, ke)), pn.mathMl === !0 && (C(Ee, ce), C(Se, Re), C(Se, ke))), y.ADD_TAGS && (Ee === No && (Ee = V(Ee)), C(Ee, y.ADD_TAGS, Ae)), y.ADD_ATTR && (Se === Mo && (Se = V(Se)), C(Se, y.ADD_ATTR, Ae)), y.ADD_URI_SAFE_ATTR && C(Oa, y.ADD_URI_SAFE_ATTR, Ae), y.FORBID_CONTENTS && (gn === zo && (gn = V(gn)), C(gn, y.FORBID_CONTENTS, Ae)), Ra && (Ee["#text"] = !0), Gt && C(Ee, ["html", "head", "body"]), Ee.table && (C(Ee, ["tbody"]), delete Wn.tbody), y.TRUSTED_TYPES_POLICY) {
            if (typeof y.TRUSTED_TYPES_POLICY.createHTML != "function")
              throw A('TRUSTED_TYPES_POLICY configuration option must provide a "createHTML" hook.');
            if (typeof y.TRUSTED_TYPES_POLICY.createScriptURL != "function")
              throw A('TRUSTED_TYPES_POLICY configuration option must provide a "createScriptURL" hook.');
            Fe = y.TRUSTED_TYPES_POLICY, Un = Fe.createHTML("");
          } else
            Fe === void 0 && (Fe = hd(_r, De)), Fe !== null && typeof Un == "string" && (Un = Fe.createHTML(""));
          s && s(y), mn = y;
        }
      }, Ko = C({}, ["mi", "mo", "mn", "ms", "mtext"]), Zo = C({}, ["foreignobject", "desc", "title", "annotation-xml"]), Md = C({}, ["title", "style", "font", "a", "script"]), Yo = C({}, [...ge, ...Ne, ...Q]), Xo = C({}, [...ce, ...$e]), jd = function(y) {
        let P = ka(y);
        P && P.tagName || (P = { namespaceURI: hn, tagName: "template" });
        const $ = _(y.tagName), te = _(P.tagName);
        return !!La[y.namespaceURI] && (y.namespaceURI === Ar ? P.namespaceURI === _t ? $ === "svg" : P.namespaceURI === Sr ? $ === "svg" && (te === "annotation-xml" || Ko[te]) : !!Yo[$] : y.namespaceURI === Sr ? P.namespaceURI === _t ? $ === "math" : P.namespaceURI === Ar ? $ === "math" && Zo[te] : !!Xo[$] : y.namespaceURI === _t ? !(P.namespaceURI === Ar && !Zo[te]) && !(P.namespaceURI === Sr && !Ko[te]) && !Xo[$] && (Md[$] || !Yo[$]) : !(Gn !== "application/xhtml+xml" || !La[y.namespaceURI]));
      }, Kt = function(y) {
        g(I.removed, { element: y });
        try {
          y.parentNode.removeChild(y);
        } catch {
          y.remove();
        }
      }, Na = function(y, P) {
        try {
          g(I.removed, { attribute: P.getAttributeNode(y), from: P });
        } catch {
          g(I.removed, { attribute: null, from: P });
        }
        if (P.removeAttribute(y), y === "is" && !Se[y])
          if (dn || wr)
            try {
              Kt(P);
            } catch {
            }
          else
            try {
              P.setAttribute(y, "");
            } catch {
            }
      }, Vo = function(y) {
        let P = null, $ = null;
        if ($a)
          y = "<remove></remove>" + y;
        else {
          const le = h(y, /^[\r\n\t ]+/);
          $ = le && le[0];
        }
        Gn === "application/xhtml+xml" && hn === _t && (y = '<html xmlns="http://www.w3.org/1999/xhtml"><head></head><body>' + y + "</body></html>");
        const te = Fe ? Fe.createHTML(y) : y;
        if (hn === _t)
          try {
            P = new _d().parseFromString(te, Gn);
          } catch {
          }
        if (!P || !P.documentElement) {
          P = wa.createDocument(hn, "template", null);
          try {
            P.documentElement.innerHTML = Pa ? Un : te;
          } catch {
          }
        }
        const Oe = P.body || P.documentElement;
        return y && $ && Oe.insertBefore(ee.createTextNode($), Oe.childNodes[0] || null), hn === _t ? xd.call(P, Gt ? "html" : "body")[0] : Gt ? P.documentElement : Oe;
      }, Jo = function(y) {
        return Sd.call(y.ownerDocument || y, y, _a.SHOW_ELEMENT | _a.SHOW_COMMENT | _a.SHOW_TEXT, null);
      }, Dd = function(y) {
        return y instanceof yd && (typeof y.nodeName != "string" || typeof y.textContent != "string" || typeof y.removeChild != "function" || !(y.attributes instanceof vd) || typeof y.removeAttribute != "function" || typeof y.setAttribute != "function" || typeof y.namespaceURI != "string" || typeof y.insertBefore != "function" || typeof y.hasChildNodes != "function");
      }, Qo = function(y) {
        return typeof ya == "function" && y instanceof ya;
      }, kt = function(y, P, $) {
        ut[y] && p(ut[y], (te) => {
          te.call(I, P, $, mn);
        });
      }, es = function(y) {
        let P = null;
        if (kt("beforeSanitizeElements", y, null), Dd(y))
          return Kt(y), !0;
        const $ = Ae(y.nodeName);
        if (kt("uponSanitizeElement", y, { tagName: $, allowedTags: Ee }), y.hasChildNodes() && !Qo(y.firstElementChild) && S(/<[/\w]/g, y.innerHTML) && S(/<[/\w]/g, y.textContent))
          return Kt(y), !0;
        if (!Ee[$] || Wn[$]) {
          if (!Wn[$] && ns($) && (me.tagNameCheck instanceof RegExp && S(me.tagNameCheck, $) || me.tagNameCheck instanceof Function && me.tagNameCheck($)))
            return !1;
          if (Ra && !gn[$]) {
            const te = ka(y) || y.parentNode, Oe = Ed(y) || y.childNodes;
            if (Oe && te)
              for (let le = Oe.length - 1; le >= 0; --le)
                te.insertBefore(kd(Oe[le], !0), wd(y));
          }
          return Kt(y), !0;
        }
        return y instanceof Po && !jd(y) ? (Kt(y), !0) : $ !== "noscript" && $ !== "noembed" && $ !== "noframes" || !S(/<\/no(script|embed|frames)/i, y.innerHTML) ? (fn && y.nodeType === 3 && (P = y.textContent, p([Ea, Sa, Aa], (te) => {
          P = v(P, te, " ");
        }), y.textContent !== P && (g(I.removed, { element: y.cloneNode() }), y.textContent = P)), kt("afterSanitizeElements", y, null), !1) : (Kt(y), !0);
      }, ts = function(y, P, $) {
        if (Bo && (P === "id" || P === "name") && ($ in ee || $ in Nd))
          return !1;
        if (!(Ca && !xa[P] && S(Td, P))) {
          if (!(jo && S($d, P))) {
            if (!Se[P] || xa[P]) {
              if (!(ns(y) && (me.tagNameCheck instanceof RegExp && S(me.tagNameCheck, y) || me.tagNameCheck instanceof Function && me.tagNameCheck(y)) && (me.attributeNameCheck instanceof RegExp && S(me.attributeNameCheck, P) || me.attributeNameCheck instanceof Function && me.attributeNameCheck(P)) || P === "is" && me.allowCustomizedBuiltInElements && (me.tagNameCheck instanceof RegExp && S(me.tagNameCheck, $) || me.tagNameCheck instanceof Function && me.tagNameCheck($))))
                return !1;
            } else if (!Oa[P]) {
              if (!S(Io, v($, Lo, ""))) {
                if ((P !== "src" && P !== "xlink:href" && P !== "href" || y === "script" || w($, "data:") !== 0 || !Uo[y]) && !(Do && !S(Rd, v($, Lo, "")))) {
                  if ($)
                    return !1;
                }
              }
            }
          }
        }
        return !0;
      }, ns = function(y) {
        return y.indexOf("-") > 0;
      }, rs = function(y) {
        kt("beforeSanitizeAttributes", y, null);
        const { attributes: P } = y;
        if (!P)
          return;
        const $ = { attrName: "", attrValue: "", keepAttr: !0, allowedAttributes: Se };
        let te = P.length;
        for (; te--; ) {
          const Oe = P[te], { name: le, namespaceURI: Kn, value: wt } = Oe, ft = Ae(le);
          let Ue = le === "value" ? wt : E(wt);
          if ($.attrName = ft, $.attrValue = Ue, $.keepAttr = !0, $.forceKeepAttr = void 0, kt("uponSanitizeAttribute", y, $), Ue = $.attrValue, $.forceKeepAttr || (Na(le, y), !$.keepAttr))
            continue;
          if (!Fo && S(/\/>/i, Ue)) {
            Na(le, y);
            continue;
          }
          fn && p([Ea, Sa, Aa], (is) => {
            Ue = v(Ue, is, " ");
          });
          const as = Ae(y.nodeName);
          if (ts(as, ft, Ue)) {
            if (!Ho || ft !== "id" && ft !== "name" || (Na(le, y), Ue = Od + Ue), Fe && typeof _r == "object" && typeof _r.getAttributeType == "function" && !Kn)
              switch (_r.getAttributeType(as, ft)) {
                case "TrustedHTML":
                  Ue = Fe.createHTML(Ue);
                  break;
                case "TrustedScriptURL":
                  Ue = Fe.createScriptURL(Ue);
              }
            try {
              Kn ? y.setAttributeNS(Kn, le, Ue) : y.setAttribute(le, Ue), d(I.removed);
            } catch {
            }
          }
        }
        kt("afterSanitizeAttributes", y, null);
      }, Fd = function y(P) {
        let $ = null;
        const te = Jo(P);
        for (kt("beforeSanitizeShadowDOM", P, null); $ = te.nextNode(); )
          kt("uponSanitizeShadowNode", $, null), es($) || ($.content instanceof it && y($.content), rs($));
        kt("afterSanitizeShadowDOM", P, null);
      };
      return I.sanitize = function(y) {
        let P = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {}, $ = null, te = null, Oe = null, le = null;
        if (Pa = !y, Pa && (y = "<!-->"), typeof y != "string" && !Qo(y)) {
          if (typeof y.toString != "function")
            throw A("toString is not a function");
          if (typeof (y = y.toString()) != "string")
            throw A("dirty is not a string, aborting");
        }
        if (!I.isSupported)
          return y;
        if (Ta || Ia(P), I.removed = [], typeof y == "string" && (qn = !1), qn) {
          if (y.nodeName) {
            const ft = Ae(y.nodeName);
            if (!Ee[ft] || Wn[ft])
              throw A("root node is forbidden and cannot be sanitized in-place");
          }
        } else if (y instanceof ya)
          $ = Vo("<!---->"), te = $.ownerDocument.importNode(y, !0), te.nodeType === 1 && te.nodeName === "BODY" || te.nodeName === "HTML" ? $ = te : $.appendChild(te);
        else {
          if (!dn && !fn && !Gt && y.indexOf("<") === -1)
            return Fe && Er ? Fe.createHTML(y) : y;
          if ($ = Vo(y), !$)
            return dn ? null : Er ? Un : "";
        }
        $ && $a && Kt($.firstChild);
        const Kn = Jo(qn ? y : $);
        for (; Oe = Kn.nextNode(); )
          es(Oe) || (Oe.content instanceof it && Fd(Oe.content), rs(Oe));
        if (qn)
          return y;
        if (dn) {
          if (wr)
            for (le = Ad.call($.ownerDocument); $.firstChild; )
              le.appendChild($.firstChild);
          else
            le = $;
          return (Se.shadowroot || Se.shadowrootmode) && (le = Cd.call(Y, le, !0)), le;
        }
        let wt = Gt ? $.outerHTML : $.innerHTML;
        return Gt && Ee["!doctype"] && $.ownerDocument && $.ownerDocument.doctype && $.ownerDocument.doctype.name && S($o, $.ownerDocument.doctype.name) && (wt = "<!DOCTYPE " + $.ownerDocument.doctype.name + `>
` + wt), fn && p([Ea, Sa, Aa], (ft) => {
          wt = v(wt, ft, " ");
        }), Fe && Er ? Fe.createHTML(wt) : wt;
      }, I.setConfig = function() {
        Ia(arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : {}), Ta = !0;
      }, I.clearConfig = function() {
        mn = null, Ta = !1;
      }, I.isValidAttribute = function(y, P, $) {
        mn || Ia({});
        const te = Ae(y), Oe = Ae(P);
        return ts(te, Oe, $);
      }, I.addHook = function(y, P) {
        typeof P == "function" && (ut[y] = ut[y] || [], g(ut[y], P));
      }, I.removeHook = function(y) {
        if (ut[y])
          return d(ut[y]);
      }, I.removeHooks = function(y) {
        ut[y] && (ut[y] = []);
      }, I.removeAllHooks = function() {
        ut = {};
      }, I;
    }
    var md = Oo();
    return md;
  }();
}), va = Eb(window);
function Sb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Hf = function(e) {
  U(r, se);
  var t = Sb(r);
  function r(n) {
    var a, i, o = n.config;
    return j(this, r), R(Be(i = t.call(this, { needCache: !0 })), "engine", "MathJax"), R(Be(i), "katex", void 0), R(Be(i), "MathJax", void 0), i.engine = Ye() ? (a = o.engine) !== null && a !== void 0 ? a : "MathJax" : "node", i;
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i, o) {
    var s, c, u;
    je(aa).call(aa, this)("engine");
    var l = n.replace(/^[ \f\r\t\v]*/, "").replace(/\s*$/, ""), f = a.replace(/^[ \f\r\t\v]*\n/, ""), p = this.$engine.md5(n), d = this.getLineCount(l, f);
    /\n/.test(a) || (d -= 1), /\n\s*$/.test(n) || (d -= 1), d = d > 0 ? d : 0;
    var g = b(s = b(c = '<div data-sign="'.concat(p, `" class="Cherry-Math" data-type="mathBlock"
          data-lines="`)).call(c, d, '">$$')).call(s, Mf(o), "$$</div>");
    if (this.engine === "katex") {
      var _, m, h = this.katex.renderToString(o, { throwOnError: !1, displayMode: !0 });
      g = b(_ = b(m = '<div data-sign="'.concat(p, `" class="Cherry-Math" data-type="mathBlock"
            data-lines="`)).call(m, d, '">')).call(_, h, "</div>");
    }
    if ((u = this.MathJax) !== null && u !== void 0 && u.tex2svg) {
      var v, w, E = Bf(this.MathJax.tex2svg(o), !0);
      g = b(v = b(w = '<div data-sign="'.concat(p, `" class="Cherry-Math" data-type="mathBlock"
            data-lines="`)).call(w, d, '">')).call(v, E, "</div>");
    }
    return g = va.sanitize(g), i + this.getCacheWithSpace(this.pushCache(g, p, d), n);
  } }, { key: "beforeMakeHtml", value: function(n) {
    var a, i;
    return fe() ? n.replace(this.RULE.reg, je(i = this.toHtml).call(i, this)) : ct(n, this.RULE.reg, je(a = this.toHtml).call(a, this), !0, 1);
  } }, { key: "makeHtml", value: function(n) {
    return n;
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "(\\s*)((?<!\\\\))~D~D\\s*" : "(\\s*)(^|[^\\\\])~D~D\\s*", content: "([\\w\\W]*?)", end: "(\\s*)~D~D(?:\\s{0,1})" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Ab(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Hf, "HOOK_NAME", "mathBlock");
var zf = function(e) {
  U(r, se);
  var t = Ab(r);
  function r(n) {
    var a, i, o = n.config;
    return j(this, r), R(Be(i = t.call(this, { needCache: !0 })), "engine", "MathJax"), R(Be(i), "katex", void 0), R(Be(i), "MathJax", void 0), i.engine = Ye() ? (a = o.engine) !== null && a !== void 0 ? a : "MathJax" : "node", i;
  }
  return M(r, [{ key: "toHtml", value: function(n, a, i) {
    var o, s, c, u;
    if (!i)
      return n;
    je(aa).call(aa, this)("engine");
    var l = i.match(/\n/g), f = l ? l.length + 2 : 2, p = this.$engine.md5(n), d = b(o = b(s = "".concat(a, `<span class="Cherry-InlineMath" data-type="mathBlock"
        data-lines="`)).call(s, f, '">$')).call(o, Mf(i), "$</span>");
    if (this.engine === "katex" && (c = this.katex) !== null && c !== void 0 && c.renderToString) {
      var g, _, m = this.katex.renderToString(i, { throwOnError: !1 });
      d = b(g = b(_ = "".concat(a, '<span class="Cherry-InlineMath" data-type="mathBlock" data-lines="')).call(_, f, '">')).call(g, m, "</span>");
    }
    if ((u = this.MathJax) !== null && u !== void 0 && u.tex2svg) {
      var h, v, w = Bf(this.MathJax.tex2svg(i, { em: 12, ex: 6, display: !1 }), !0);
      d = b(h = b(v = "".concat(a, '<span class="Cherry-InlineMath" data-type="mathBlock" data-lines="')).call(v, f, '">')).call(h, w, "</span>");
    }
    return d = va.sanitize(d), this.pushCache(d, se.IN_PARAGRAPH_CACHE_KEY_PREFIX + p);
  } }, { key: "beforeMakeHtml", value: function(n) {
    var a = this, i = n;
    return i = i.replace(Eo(!0), function(o) {
      var s;
      return ve(s = o.split("|")).call(s, function(c) {
        return a.makeInlineMath(c);
      }).join("|").replace(/\\~D/g, "~D").replace(/~D/g, "\\~D");
    }), this.makeInlineMath(i);
  } }, { key: "makeInlineMath", value: function(n) {
    var a, i;
    return this.test(n) ? fe() ? n.replace(this.RULE.reg, je(i = this.toHtml).call(i, this)) : ct(n, this.RULE.reg, je(a = this.toHtml).call(a, this), !0, 1) : n;
  } }, { key: "makeHtml", value: function(n) {
    return n;
  } }, { key: "rule", value: function() {
    var n = { begin: fe() ? "((?<!\\\\))~D\\n?" : "(^|[^\\\\])~D\\n?", content: "(.*?)\\n?", end: "~D" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
R(zf, "HOOK_NAME", "inlineMath");
L({ target: "Array", proto: !0 }, { fill: function(e) {
  for (var t = lt(this), r = bt(t), n = arguments.length, a = At(n > 1 ? arguments[1] : void 0, r), i = n > 2 ? arguments[2] : void 0, o = i === void 0 ? r : At(i, r); o > a; )
    t[a++] = e;
  return t;
} });
var xb = Ie("Array").fill, yi = Array.prototype, hl = function(e) {
  var t = e.fill;
  return e === yi || de(yi, e) && t === yi.fill ? xb : t;
};
function Cb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
function Uf(e) {
  return e;
}
var ml = { tocStyle: "plain", tocNodeClass: "toc-li", tocContainerClass: "toc", tocTitleClass: "toc-title", linkProcessor: Uf }, bl = '<p data-sign="empty-toc" data-lines="1">&nbsp;</p>', Wf = function(e) {
  U(r, se);
  var t = Cb(r);
  function r(n) {
    var a, i;
    n.externals;
    var o = n.config;
    return j(this, r), R(Be(i = t.call(this, { needCache: !0 })), "tocStyle", "nested"), R(Be(i), "tocNodeClass", "toc-li"), R(Be(i), "tocContainerClass", "toc"), R(Be(i), "tocTitleClass", "toc-title"), R(Be(i), "linkProcessor", Uf), R(Be(i), "baseLevel", 1), R(Be(i), "isFirstTocToken", !0), R(Be(i), "allowMultiToc", !1), q(a = ue(ml)).call(a, function(s) {
      i[s] = o[s] || ml[s];
    }), i;
  }
  return M(r, [{ key: "beforeMakeHtml", value: function(n) {
    var a = this, i = n;
    return this.test(i, "extend") && (i = i.replace(this.RULE.extend.reg, function(o, s, c) {
      var u;
      if (!a.allowMultiToc && !a.isFirstTocToken)
        return b(u = `
`.concat(s)).call(u, bl);
      var l = a.pushCache(o);
      return a.isFirstTocToken = !1, nn(o, l);
    })), this.test(i, "standard") && (i = i.replace(this.RULE.standard.reg, function(o, s, c) {
      var u;
      return a.allowMultiToc || a.isFirstTocToken ? (a.isFirstTocToken = !1, nn(o, a.pushCache(o))) : b(u = `
`.concat(s)).call(u, bl);
    })), i;
  } }, { key: "makeHtml", value: function(n) {
    return n;
  } }, { key: "$makeLevel", value: function(n) {
    for (var a = "", i = this.baseLevel; i < n; i++)
      a += "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;";
    return a;
  } }, { key: "$makeTocItem", value: function(n, a) {
    var i, o, s, c, u, l = !(arguments.length > 2 && arguments[2] !== void 0) || arguments[2], f = "";
    a && (f = this.$makeLevel(n.level));
    var p = this.linkProcessor("#".concat(n.id).replace(/safe_/g, ""));
    return b(i = b(o = b(s = b(c = b(u = '<li class="'.concat(this.tocNodeClass, '">')).call(u, f, '<a href="')).call(c, p, '" class="level-')).call(s, n.level, '">')).call(o, n.text, "</a>")).call(i, l ? "</li>" : "");
  } }, { key: "$makePlainToc", value: function(n) {
    var a = this;
    return ve(n).call(n, function(i) {
      return a.$makeTocItem(i, !0);
    }).join("");
  } }, { key: "$makeNestedToc", value: function(n) {
    var a, i, o = this, s = 0, c = hl(a = new Array(7)).call(a, !1), u = hl(i = new Array(7)).call(i, !1), l = "";
    q(n).call(n, function(p) {
      var d = p.level;
      if (s === 0) {
        for (var g = d; g >= o.baseLevel; g--)
          l += "<ul>", u[g] = !0;
        return l += o.$makeTocItem(p, !1, !1), c[d] = !0, void (s = d);
      }
      if (d < s) {
        for (var _ = s; _ >= d; _--)
          c[_] && (l += "</li>", c[_] = !1), u[_] && _ > d && (l += "</ul>", u[_] = !1);
        c[d] = !0, l += o.$makeTocItem(p, !1, !1), s = d;
      } else if (d === s)
        c[s] && (l += "</li>"), l += o.$makeTocItem(p, !1, !1), c[d] = !0, u[d] = !0;
      else {
        for (var m = s + 1; m <= d; m++)
          l += "<ul>", u[m] = !0;
        c[d] = !0, l += o.$makeTocItem(p, !1, !1), s = d;
      }
    });
    for (var f = s; f >= this.baseLevel; f--)
      c[f] && (l += "</li>", c[f] = !1), u[f] && (l += "</ul>", u[f] = !1);
    return l;
  } }, { key: "$makeToc", value: function(n, a, i) {
    var o, s, c, u = ji(i, 1), l = b(o = b(s = b(c = '<dir class="'.concat(this.tocContainerClass, '" data-sign="')).call(c, a, "-")).call(s, u, '" data-lines="')).call(o, u, '">');
    return l += '<p class="'.concat(this.tocTitleClass, '">目录</p>'), n.length <= 0 ? "" : (this.baseLevel = Math.min.apply(Math, lr(ve(n).call(n, function(f) {
      return f.level;
    }))), this.tocStyle === "nested" ? l += this.$makeNestedToc(n) : l += this.$makePlainToc(n), l += "</dir>");
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = this, i = dt(T(r.prototype), "afterMakeHtml", this).call(this, n), o = [], s = "";
    return i.replace(/<h([1-6])[^>]*? id="([^"]+?)"[^>]*?>(?:<a[^/]+?\/a>|)(.+?)<\/h\1>/g, function(c, u, l, f) {
      var p, d = f.replace(/~fn#[0-9]+#/g, "");
      o.push({ level: +u, id: l, text: d }), s += b(p = "".concat(u)).call(p, l);
    }), s = this.$engine.md5(s), i = i.replace(/(?:^|\n)(\[\[|\[|【【)(toc|TOC)(\]\]|\]|】】)([<~])/, function(c) {
      return c.replace(/(\]\]|\]|】】)([<~])/, `$1
$2`);
    }), i = (i = i.replace(this.RULE.extend.reg, function(c, u) {
      return a.$makeToc(o, s, u);
    })).replace(this.RULE.standard.reg, function(c, u) {
      return a.$makeToc(o, s, u);
    }), this.isFirstTocToken = !0, i;
  } }, { key: "test", value: function(n, a) {
    return !!this.RULE[a].reg && this.RULE[a].reg.test(n);
  } }, { key: "rule", value: function() {
    var n = { begin: "(?:^|\\n)(\\n*)", end: "(?=$|\\n)", content: "[ ]*((?:【【|\\[\\[)(?:toc|TOC)(?:\\]\\]|】】))[ ]*" };
    n.reg = new RegExp(n.begin + n.content + n.end, "g");
    var a = { begin: "(?:^|\\n)(\\n*)", end: "(?=$|\\n)", content: "[ ]*(\\[(?:toc|TOC)\\])[ ]*" };
    return a.reg = new RegExp(a.begin + a.content + a.end, "g"), { extend: n, standard: a };
  } }]), r;
}();
function Tb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Wf, "HOOK_NAME", "toc");
var qf = function(e) {
  U(r, se);
  var t = Tb(r);
  function r(n) {
    var a;
    return n.externals, n.config, j(this, r), (a = t.call(this)).footnoteCache = {}, a.footnoteMap = {}, a.footnote = [], a;
  }
  return M(r, [{ key: "$cleanCache", value: function() {
    this.footnoteCache = {}, this.footnoteMap = {}, this.footnote = [];
  } }, { key: "pushFootnoteCache", value: function(n, a) {
    this.footnoteCache[n] = a;
  } }, { key: "getFootnoteCache", value: function(n) {
    return this.footnoteCache[n] || null;
  } }, { key: "pushFootNote", value: function(n, a) {
    var i, o, s, c, u, l;
    if (this.footnoteMap[n])
      return this.footnoteMap[n];
    var f = this.footnote.length + 1, p = {};
    p.fn = b(i = b(o = b(s = '<sup><a href="#fn:'.concat(f, '" id="fnref:')).call(s, f, '" title="')).call(o, n, '" class="footnote">[')).call(i, f, "]</a></sup>"), p.fnref = b(c = b(u = b(l = '<a href="#fnref:'.concat(f, '" id="fn:')).call(l, f, '" title="')).call(u, n, '" class="footnote-ref">[')).call(c, f, "]</a>"), p.num = f, p.note = N(a).call(a), this.footnote.push(p);
    var d = "\0~fn#".concat(f - 1, "#\0");
    return this.footnoteMap[n] = d, d;
  } }, { key: "getFootNote", value: function() {
    return this.footnote;
  } }, { key: "formatFootNote", value: function() {
    var n, a = this.getFootNote();
    if (a.length <= 0)
      return "";
    var i = ve(a).call(a, function(s) {
      var c;
      return b(c = `<div class="one-footnote">
`.concat(s.fnref)).call(c, s.note, `
</div>`);
    }).join(""), o = this.$engine.md5(i);
    return i = b(n = '<div class="footnote" data-sign="'.concat(o, '" data-lines="0"><div class="footnote-title">脚注</div>')).call(n, i, "</div>");
  } }, { key: "beforeMakeHtml", value: function(n) {
    var a = this, i = n;
    return this.test(i) && (i = i.replace(this.RULE.reg, function(o, s, c, u) {
      return a.pushFootnoteCache(c, u), (o.match(/\n/g) || []).join("");
    }), i = i.replace(/\[\^([^\]]+?)\](?!:)/g, function(o, s) {
      var c = a.getFootnoteCache(s);
      return c ? a.pushFootNote(s, c) : o;
    }), i += this.formatFootNote()), i;
  } }, { key: "makeHtml", value: function(n, a) {
    return n;
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = this.getFootNote(), i = n.replace(/\0~fn#([0-9]+)#\0/g, function(o, s) {
      return a[s].fn;
    });
    return this.$cleanCache(), i;
  } }, { key: "rule", value: function() {
    var n = { begin: "(^|\\n)[ 	]*", content: ["\\[\\^([^\\]]+?)\\]:\\h*", "([\\s\\S]+?)"].join(""), end: "(?=\\s*$|\\n\\n)" };
    return n.reg = Le(n, "g", !0), n;
  } }]), r;
}();
function $b(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(qf, "HOOK_NAME", "footnote");
var Gf = function(e) {
  U(r, se);
  var t = $b(r);
  function r(n) {
    var a;
    return n.externals, n.config, j(this, r), (a = t.call(this)).commentCache = {}, a;
  }
  return M(r, [{ key: "$cleanCache", value: function() {
    this.commentCache = {};
  } }, { key: "pushCommentReferenceCache", value: function(n, a) {
    var i, o = hf(a.split(/[ ]+/g)), s = o[0], c = Ke(o).call(o, 1), u = St.set(s);
    this.commentCache["".concat(n).toLowerCase()] = b(i = [u]).call(i, lr(c)).join(" ");
  } }, { key: "getCommentReferenceCache", value: function(n) {
    return this.commentCache["".concat(n).toLowerCase()] || null;
  } }, { key: "beforeMakeHtml", value: function(n) {
    var a = this, i = n;
    return this.test(i) && (i = i.replace(this.RULE.reg, function(o, s, c, u) {
      var l;
      return a.pushCommentReferenceCache(c, u), ((l = o.match(/\n/g)) !== null && l !== void 0 ? l : []).join("");
    }), i = i.replace(/(\[[^\]\n]+?\])?(?:\[([^\]\n]+?)\])/g, function(o, s, c) {
      var u, l, f = a.getCommentReferenceCache(c);
      return f ? s ? b(l = "".concat(s, "(")).call(l, f, ")") : b(u = "[".concat(c, "](")).call(u, f, ")") : o;
    }), this.$cleanCache()), i;
  } }, { key: "makeHtml", value: function(n, a) {
    return n;
  } }, { key: "afterMakeHtml", value: function(n) {
    return St.restoreAll(n);
  } }, { key: "rule", value: function() {
    var n = { begin: "(^|\\n)[ 	]*", content: ["\\[([^^][^\\]]*?)\\]:\\h*", "([^\\n]+?)"].join(""), end: "(?=$|\\n)" };
    return n.reg = Le(n, "g", !0), n;
  } }]), r;
}();
R(Gf, "HOOK_NAME", "commentReference");
var Rb = sn.some, Ob = ha("some");
L({ target: "Array", proto: !0, forced: !Ob }, { some: function(e) {
  return Rb(this, e, arguments.length > 1 ? arguments[1] : void 0);
} });
var Pb = Ie("Array").some, _i = Array.prototype, Lb = function(e) {
  var t = e.some;
  return e === _i || de(_i, e) && t === _i.some ? Pb : t;
};
function Ib(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var vl = ["href", "src"];
va.addHook("afterSanitizeAttributes", function(e) {
  q(vl).call(vl, function(t) {
    if (e.hasAttribute(t)) {
      var r = e.getAttribute(t);
      e.setAttribute(t, r.replace(/\\/g, "%5c"));
    }
  });
});
var Kf = function(e) {
  U(r, se);
  var t = Ib(r);
  function r() {
    return j(this, r), t.call(this, { needCache: !0 });
  }
  return M(r, [{ key: "isAutoLinkTag", value: function(n) {
    var a = [/^<([a-z][a-z0-9+.-]{1,31}:\/\/[^<> `]+)>$/i, /^<(mailto:[^<> `]+)>$/i, /^<([^()<>[\]:'@\\,"\s`]+@[^()<>[\]:'@\\,"\s`.]+\.[^()<>[\]:'@\\,"\s`]+)>$/i];
    return Lb(a).call(a, function(i) {
      return i.test(n);
    });
  } }, { key: "isHtmlComment", value: function(n) {
    return /^<!--.*?-->$/.test(n);
  } }, { key: "beforeMakeHtml", value: function(n, a) {
    var i = this;
    this.$engine.htmlWhiteListAppend ? (this.htmlWhiteListAppend = new RegExp("^(".concat(this.$engine.htmlWhiteListAppend, ")( |$|/)"), "i"), this.htmlWhiteList = this.$engine.htmlWhiteListAppend.split("|")) : (this.htmlWhiteListAppend = !1, this.htmlWhiteList = []);
    var o = n;
    return o = function(s) {
      if (typeof s != "string")
        return "";
      var c = s.replace(/&(\w+);?/g, function(u, l) {
        return Me(u).call(u, ";") === -1 || Me(Ac).call(Ac, l.toLowerCase()) === -1 ? u.replace(/&/g, "&amp;") : u;
      });
      return c = c.replace(/&#(?!x)(\d*);?/gi, function(u, l) {
        return Bi(l) || Me(u).call(u, ";") === -1 || l.lenth > 7 || !xc(l) ? u.replace(/&/g, "&amp;") : u;
      }), c = c.replace(/&#x([0-9a-f]*);?/gi, function(u, l) {
        if (Bi(l))
          return u.replace(/&/g, "&amp;");
        var f = "0x".concat(l), p = xn(f, 16);
        return isNaN(p) || Me(u).call(u, ";") === -1 || l.lenth > 6 || !xc(f) ? u.replace(/&/g, "&amp;") : u;
      }), c;
    }(o = Vu(o)), o = (o = (o = o.replace(/<[/]?(.*?)>/g, function(s, c) {
      return K1.test(c) || i.isAutoLinkTag(s) || i.isHtmlComment(s) || i.htmlWhiteListAppend !== !1 && i.htmlWhiteListAppend.test(c) ? s.replace(/</g, "$#60;").replace(/>/g, "$#62;") : s.replace(/</g, "&#60;").replace(/>/g, "&#62;");
    })).replace(/<(?=\/?(\w|\n|$))/g, "&#60;")).replace(/\$#60;/g, "<").replace(/\$#62;/g, ">");
  } }, { key: "makeHtml", value: function(n, a) {
    return n;
  } }, { key: "afterMakeHtml", value: function(n) {
    var a = n, i = { ALLOW_UNKNOWN_PROTOCOLS: !0, ADD_ATTR: ["target"] };
    if (this.htmlWhiteListAppend !== !1) {
      var o;
      if (i.ADD_TAGS = this.htmlWhiteList, (this.htmlWhiteListAppend.test("style") || this.htmlWhiteListAppend.test("ALL")) && (a = a.replace(/<style(>| [^>]*>).*?<\/style>/gi, function(s) {
        return s.replace(/<br>/gi, "");
      })), (this.htmlWhiteListAppend.test("iframe") || this.htmlWhiteListAppend.test("ALL")) && (i.ADD_ATTR = b(o = i.ADD_ATTR).call(o, ["align", "frameborder", "height", "longdesc", "marginheight", "marginwidth", "name", "sandbox", "scrolling", "seamless", "src", "srcdoc", "width"]), i.SANITIZE_DOM = !1, a = a.replace(/<iframe(>| [^>]*>).*?<\/iframe>/gi, function(s) {
        return s.replace(/<br>/gi, "").replace(/\n/g, "");
      })), this.htmlWhiteListAppend.test("script") || this.htmlWhiteListAppend.test("ALL"))
        return a = a.replace(/<script(>| [^>]*>).*?<\/script>/gi, function(s) {
          return s.replace(/<br>/gi, "");
        }), a;
    }
    return Ye() || (i.FORBID_ATTR = ["data-sign", "data-lines"]), va.sanitize(a, i);
  } }]), r;
}();
R(Kf, "HOOK_NAME", "htmlBlock");
var Nb = { "+1": "1f44d", "-1": "1f44e", 100: "1f4af", 1234: "1f522", "1st_place_medal": "1f947", "2nd_place_medal": "1f948", "3rd_place_medal": "1f949", "8ball": "1f3b1", a: "1f170", ab: "1f18e", abacus: "1f9ee", abc: "1f524", abcd: "1f521", accept: "1f251", adhesive_bandage: "1fa79", adult: "1f9d1", aerial_tramway: "1f6a1", afghanistan: "1f1e6-1f1eb", airplane: "2708", aland_islands: "1f1e6-1f1fd", alarm_clock: "23f0", albania: "1f1e6-1f1f1", alembic: "2697", algeria: "1f1e9-1f1ff", alien: "1f47d", ambulance: "1f691", american_samoa: "1f1e6-1f1f8", amphora: "1f3fa", anchor: "2693", andorra: "1f1e6-1f1e9", angel: "1f47c", anger: "1f4a2", angola: "1f1e6-1f1f4", angry: "1f620", anguilla: "1f1e6-1f1ee", anguished: "1f627", ant: "1f41c", antarctica: "1f1e6-1f1f6", antigua_barbuda: "1f1e6-1f1ec", apple: "1f34e", aquarius: "2652", argentina: "1f1e6-1f1f7", aries: "2648", armenia: "1f1e6-1f1f2", arrow_backward: "25c0", arrow_double_down: "23ec", arrow_double_up: "23eb", arrow_down: "2b07", arrow_down_small: "1f53d", arrow_forward: "25b6", arrow_heading_down: "2935", arrow_heading_up: "2934", arrow_left: "2b05", arrow_lower_left: "2199", arrow_lower_right: "2198", arrow_right: "27a1", arrow_right_hook: "21aa", arrow_up: "2b06", arrow_up_down: "2195", arrow_up_small: "1f53c", arrow_upper_left: "2196", arrow_upper_right: "2197", arrows_clockwise: "1f503", arrows_counterclockwise: "1f504", art: "1f3a8", articulated_lorry: "1f69b", artificial_satellite: "1f6f0", artist: "1f9d1-1f3a8", aruba: "1f1e6-1f1fc", ascension_island: "1f1e6-1f1e8", asterisk: "002a-20e3", astonished: "1f632", astronaut: "1f9d1-1f680", athletic_shoe: "1f45f", atm: "1f3e7", atom_symbol: "269b", australia: "1f1e6-1f1fa", austria: "1f1e6-1f1f9", auto_rickshaw: "1f6fa", avocado: "1f951", axe: "1fa93", azerbaijan: "1f1e6-1f1ff", b: "1f171", baby: "1f476", baby_bottle: "1f37c", baby_chick: "1f424", baby_symbol: "1f6bc", back: "1f519", bacon: "1f953", badger: "1f9a1", badminton: "1f3f8", bagel: "1f96f", baggage_claim: "1f6c4", baguette_bread: "1f956", bahamas: "1f1e7-1f1f8", bahrain: "1f1e7-1f1ed", balance_scale: "2696", bald_man: "1f468-1f9b2", bald_woman: "1f469-1f9b2", ballet_shoes: "1fa70", balloon: "1f388", ballot_box: "1f5f3", ballot_box_with_check: "2611", bamboo: "1f38d", banana: "1f34c", bangbang: "203c", bangladesh: "1f1e7-1f1e9", banjo: "1fa95", bank: "1f3e6", bar_chart: "1f4ca", barbados: "1f1e7-1f1e7", barber: "1f488", baseball: "26be", basket: "1f9fa", basketball: "1f3c0", basketball_man: "26f9-2642", basketball_woman: "26f9-2640", bat: "1f987", bath: "1f6c0", bathtub: "1f6c1", battery: "1f50b", beach_umbrella: "1f3d6", bear: "1f43b", bearded_person: "1f9d4", bed: "1f6cf", bee: "1f41d", beer: "1f37a", beers: "1f37b", beetle: "1f41e", beginner: "1f530", belarus: "1f1e7-1f1fe", belgium: "1f1e7-1f1ea", belize: "1f1e7-1f1ff", bell: "1f514", bellhop_bell: "1f6ce", benin: "1f1e7-1f1ef", bento: "1f371", bermuda: "1f1e7-1f1f2", beverage_box: "1f9c3", bhutan: "1f1e7-1f1f9", bicyclist: "1f6b4", bike: "1f6b2", biking_man: "1f6b4-2642", biking_woman: "1f6b4-2640", bikini: "1f459", billed_cap: "1f9e2", biohazard: "2623", bird: "1f426", birthday: "1f382", black_circle: "26ab", black_flag: "1f3f4", black_heart: "1f5a4", black_joker: "1f0cf", black_large_square: "2b1b", black_medium_small_square: "25fe", black_medium_square: "25fc", black_nib: "2712", black_small_square: "25aa", black_square_button: "1f532", blond_haired_man: "1f471-2642", blond_haired_person: "1f471", blond_haired_woman: "1f471-2640", blonde_woman: "1f471-2640", blossom: "1f33c", blowfish: "1f421", blue_book: "1f4d8", blue_car: "1f699", blue_heart: "1f499", blue_square: "1f7e6", blush: "1f60a", boar: "1f417", boat: "26f5", bolivia: "1f1e7-1f1f4", bomb: "1f4a3", bone: "1f9b4", book: "1f4d6", bookmark: "1f516", bookmark_tabs: "1f4d1", books: "1f4da", boom: "1f4a5", boot: "1f462", bosnia_herzegovina: "1f1e7-1f1e6", botswana: "1f1e7-1f1fc", bouncing_ball_man: "26f9-2642", bouncing_ball_person: "26f9", bouncing_ball_woman: "26f9-2640", bouquet: "1f490", bouvet_island: "1f1e7-1f1fb", bow: "1f647", bow_and_arrow: "1f3f9", bowing_man: "1f647-2642", bowing_woman: "1f647-2640", bowl_with_spoon: "1f963", bowling: "1f3b3", boxing_glove: "1f94a", boy: "1f466", brain: "1f9e0", brazil: "1f1e7-1f1f7", bread: "1f35e", breast_feeding: "1f931", bricks: "1f9f1", bride_with_veil: "1f470", bridge_at_night: "1f309", briefcase: "1f4bc", british_indian_ocean_territory: "1f1ee-1f1f4", british_virgin_islands: "1f1fb-1f1ec", broccoli: "1f966", broken_heart: "1f494", broom: "1f9f9", brown_circle: "1f7e4", brown_heart: "1f90e", brown_square: "1f7eb", brunei: "1f1e7-1f1f3", bug: "1f41b", building_construction: "1f3d7", bulb: "1f4a1", bulgaria: "1f1e7-1f1ec", bullettrain_front: "1f685", bullettrain_side: "1f684", burkina_faso: "1f1e7-1f1eb", burrito: "1f32f", burundi: "1f1e7-1f1ee", bus: "1f68c", business_suit_levitating: "1f574", busstop: "1f68f", bust_in_silhouette: "1f464", busts_in_silhouette: "1f465", butter: "1f9c8", butterfly: "1f98b", cactus: "1f335", cake: "1f370", calendar: "1f4c6", call_me_hand: "1f919", calling: "1f4f2", cambodia: "1f1f0-1f1ed", camel: "1f42b", camera: "1f4f7", camera_flash: "1f4f8", cameroon: "1f1e8-1f1f2", camping: "1f3d5", canada: "1f1e8-1f1e6", canary_islands: "1f1ee-1f1e8", cancer: "264b", candle: "1f56f", candy: "1f36c", canned_food: "1f96b", canoe: "1f6f6", cape_verde: "1f1e8-1f1fb", capital_abcd: "1f520", capricorn: "2651", car: "1f697", card_file_box: "1f5c3", card_index: "1f4c7", card_index_dividers: "1f5c2", caribbean_netherlands: "1f1e7-1f1f6", carousel_horse: "1f3a0", carrot: "1f955", cartwheeling: "1f938", cat: "1f431", cat2: "1f408", cayman_islands: "1f1f0-1f1fe", cd: "1f4bf", central_african_republic: "1f1e8-1f1eb", ceuta_melilla: "1f1ea-1f1e6", chad: "1f1f9-1f1e9", chains: "26d3", chair: "1fa91", champagne: "1f37e", chart: "1f4b9", chart_with_downwards_trend: "1f4c9", chart_with_upwards_trend: "1f4c8", checkered_flag: "1f3c1", cheese: "1f9c0", cherries: "1f352", cherry_blossom: "1f338", chess_pawn: "265f", chestnut: "1f330", chicken: "1f414", child: "1f9d2", children_crossing: "1f6b8", chile: "1f1e8-1f1f1", chipmunk: "1f43f", chocolate_bar: "1f36b", chopsticks: "1f962", christmas_island: "1f1e8-1f1fd", christmas_tree: "1f384", church: "26ea", cinema: "1f3a6", circus_tent: "1f3aa", city_sunrise: "1f307", city_sunset: "1f306", cityscape: "1f3d9", cl: "1f191", clamp: "1f5dc", clap: "1f44f", clapper: "1f3ac", classical_building: "1f3db", climbing: "1f9d7", climbing_man: "1f9d7-2642", climbing_woman: "1f9d7-2640", clinking_glasses: "1f942", clipboard: "1f4cb", clipperton_island: "1f1e8-1f1f5", clock1: "1f550", clock10: "1f559", clock1030: "1f565", clock11: "1f55a", clock1130: "1f566", clock12: "1f55b", clock1230: "1f567", clock130: "1f55c", clock2: "1f551", clock230: "1f55d", clock3: "1f552", clock330: "1f55e", clock4: "1f553", clock430: "1f55f", clock5: "1f554", clock530: "1f560", clock6: "1f555", clock630: "1f561", clock7: "1f556", clock730: "1f562", clock8: "1f557", clock830: "1f563", clock9: "1f558", clock930: "1f564", closed_book: "1f4d5", closed_lock_with_key: "1f510", closed_umbrella: "1f302", cloud: "2601", cloud_with_lightning: "1f329", cloud_with_lightning_and_rain: "26c8", cloud_with_rain: "1f327", cloud_with_snow: "1f328", clown_face: "1f921", clubs: "2663", cn: "1f1e8-1f1f3", coat: "1f9e5", cocktail: "1f378", coconut: "1f965", cocos_islands: "1f1e8-1f1e8", coffee: "2615", coffin: "26b0", cold_face: "1f976", cold_sweat: "1f630", collision: "1f4a5", colombia: "1f1e8-1f1f4", comet: "2604", comoros: "1f1f0-1f1f2", compass: "1f9ed", computer: "1f4bb", computer_mouse: "1f5b1", confetti_ball: "1f38a", confounded: "1f616", confused: "1f615", congo_brazzaville: "1f1e8-1f1ec", congo_kinshasa: "1f1e8-1f1e9", congratulations: "3297", construction: "1f6a7", construction_worker: "1f477", construction_worker_man: "1f477-2642", construction_worker_woman: "1f477-2640", control_knobs: "1f39b", convenience_store: "1f3ea", cook: "1f9d1-1f373", cook_islands: "1f1e8-1f1f0", cookie: "1f36a", cool: "1f192", cop: "1f46e", copyright: "00a9", corn: "1f33d", costa_rica: "1f1e8-1f1f7", cote_divoire: "1f1e8-1f1ee", couch_and_lamp: "1f6cb", couple: "1f46b", couple_with_heart: "1f491", couple_with_heart_man_man: "1f468-2764-1f468", couple_with_heart_woman_man: "1f469-2764-1f468", couple_with_heart_woman_woman: "1f469-2764-1f469", couplekiss: "1f48f", couplekiss_man_man: "1f468-2764-1f48b-1f468", couplekiss_man_woman: "1f469-2764-1f48b-1f468", couplekiss_woman_woman: "1f469-2764-1f48b-1f469", cow: "1f42e", cow2: "1f404", cowboy_hat_face: "1f920", crab: "1f980", crayon: "1f58d", credit_card: "1f4b3", crescent_moon: "1f319", cricket: "1f997", cricket_game: "1f3cf", croatia: "1f1ed-1f1f7", crocodile: "1f40a", croissant: "1f950", crossed_fingers: "1f91e", crossed_flags: "1f38c", crossed_swords: "2694", crown: "1f451", cry: "1f622", crying_cat_face: "1f63f", crystal_ball: "1f52e", cuba: "1f1e8-1f1fa", cucumber: "1f952", cup_with_straw: "1f964", cupcake: "1f9c1", cupid: "1f498", curacao: "1f1e8-1f1fc", curling_stone: "1f94c", curly_haired_man: "1f468-1f9b1", curly_haired_woman: "1f469-1f9b1", curly_loop: "27b0", currency_exchange: "1f4b1", curry: "1f35b", cursing_face: "1f92c", custard: "1f36e", customs: "1f6c3", cut_of_meat: "1f969", cyclone: "1f300", cyprus: "1f1e8-1f1fe", czech_republic: "1f1e8-1f1ff", dagger: "1f5e1", dancer: "1f483", dancers: "1f46f", dancing_men: "1f46f-2642", dancing_women: "1f46f-2640", dango: "1f361", dark_sunglasses: "1f576", dart: "1f3af", dash: "1f4a8", date: "1f4c5", de: "1f1e9-1f1ea", deaf_man: "1f9cf-2642", deaf_person: "1f9cf", deaf_woman: "1f9cf-2640", deciduous_tree: "1f333", deer: "1f98c", denmark: "1f1e9-1f1f0", department_store: "1f3ec", derelict_house: "1f3da", desert: "1f3dc", desert_island: "1f3dd", desktop_computer: "1f5a5", detective: "1f575", diamond_shape_with_a_dot_inside: "1f4a0", diamonds: "2666", diego_garcia: "1f1e9-1f1ec", disappointed: "1f61e", disappointed_relieved: "1f625", diving_mask: "1f93f", diya_lamp: "1fa94", dizzy: "1f4ab", dizzy_face: "1f635", djibouti: "1f1e9-1f1ef", dna: "1f9ec", do_not_litter: "1f6af", dog: "1f436", dog2: "1f415", dollar: "1f4b5", dolls: "1f38e", dolphin: "1f42c", dominica: "1f1e9-1f1f2", dominican_republic: "1f1e9-1f1f4", door: "1f6aa", doughnut: "1f369", dove: "1f54a", dragon: "1f409", dragon_face: "1f432", dress: "1f457", dromedary_camel: "1f42a", drooling_face: "1f924", drop_of_blood: "1fa78", droplet: "1f4a7", drum: "1f941", duck: "1f986", dumpling: "1f95f", dvd: "1f4c0", "e-mail": "1f4e7", eagle: "1f985", ear: "1f442", ear_of_rice: "1f33e", ear_with_hearing_aid: "1f9bb", earth_africa: "1f30d", earth_americas: "1f30e", earth_asia: "1f30f", ecuador: "1f1ea-1f1e8", egg: "1f95a", eggplant: "1f346", egypt: "1f1ea-1f1ec", eight: "0038-20e3", eight_pointed_black_star: "2734", eight_spoked_asterisk: "2733", eject_button: "23cf", el_salvador: "1f1f8-1f1fb", electric_plug: "1f50c", elephant: "1f418", elf: "1f9dd", elf_man: "1f9dd-2642", elf_woman: "1f9dd-2640", email: "2709", end: "1f51a", england: "1f3f4-e0067-e0062-e0065-e006e-e0067-e007f", envelope: "2709", envelope_with_arrow: "1f4e9", equatorial_guinea: "1f1ec-1f1f6", eritrea: "1f1ea-1f1f7", es: "1f1ea-1f1f8", estonia: "1f1ea-1f1ea", ethiopia: "1f1ea-1f1f9", eu: "1f1ea-1f1fa", euro: "1f4b6", european_castle: "1f3f0", european_post_office: "1f3e4", european_union: "1f1ea-1f1fa", evergreen_tree: "1f332", exclamation: "2757", exploding_head: "1f92f", expressionless: "1f611", eye: "1f441", eye_speech_bubble: "1f441-1f5e8", eyeglasses: "1f453", eyes: "1f440", face_with_head_bandage: "1f915", face_with_thermometer: "1f912", facepalm: "1f926", facepunch: "1f44a", factory: "1f3ed", factory_worker: "1f9d1-1f3ed", fairy: "1f9da", fairy_man: "1f9da-2642", fairy_woman: "1f9da-2640", falafel: "1f9c6", falkland_islands: "1f1eb-1f1f0", fallen_leaf: "1f342", family: "1f46a", family_man_boy: "1f468-1f466", family_man_boy_boy: "1f468-1f466-1f466", family_man_girl: "1f468-1f467", family_man_girl_boy: "1f468-1f467-1f466", family_man_girl_girl: "1f468-1f467-1f467", family_man_man_boy: "1f468-1f468-1f466", family_man_man_boy_boy: "1f468-1f468-1f466-1f466", family_man_man_girl: "1f468-1f468-1f467", family_man_man_girl_boy: "1f468-1f468-1f467-1f466", family_man_man_girl_girl: "1f468-1f468-1f467-1f467", family_man_woman_boy: "1f468-1f469-1f466", family_man_woman_boy_boy: "1f468-1f469-1f466-1f466", family_man_woman_girl: "1f468-1f469-1f467", family_man_woman_girl_boy: "1f468-1f469-1f467-1f466", family_man_woman_girl_girl: "1f468-1f469-1f467-1f467", family_woman_boy: "1f469-1f466", family_woman_boy_boy: "1f469-1f466-1f466", family_woman_girl: "1f469-1f467", family_woman_girl_boy: "1f469-1f467-1f466", family_woman_girl_girl: "1f469-1f467-1f467", family_woman_woman_boy: "1f469-1f469-1f466", family_woman_woman_boy_boy: "1f469-1f469-1f466-1f466", family_woman_woman_girl: "1f469-1f469-1f467", family_woman_woman_girl_boy: "1f469-1f469-1f467-1f466", family_woman_woman_girl_girl: "1f469-1f469-1f467-1f467", farmer: "1f9d1-1f33e", faroe_islands: "1f1eb-1f1f4", fast_forward: "23e9", fax: "1f4e0", fearful: "1f628", feet: "1f43e", female_detective: "1f575-2640", female_sign: "2640", ferris_wheel: "1f3a1", ferry: "26f4", field_hockey: "1f3d1", fiji: "1f1eb-1f1ef", file_cabinet: "1f5c4", file_folder: "1f4c1", film_projector: "1f4fd", film_strip: "1f39e", finland: "1f1eb-1f1ee", fire: "1f525", fire_engine: "1f692", fire_extinguisher: "1f9ef", firecracker: "1f9e8", firefighter: "1f9d1-1f692", fireworks: "1f386", first_quarter_moon: "1f313", first_quarter_moon_with_face: "1f31b", fish: "1f41f", fish_cake: "1f365", fishing_pole_and_fish: "1f3a3", fist: "270a", fist_left: "1f91b", fist_oncoming: "1f44a", fist_raised: "270a", fist_right: "1f91c", five: "0035-20e3", flags: "1f38f", flamingo: "1f9a9", flashlight: "1f526", flat_shoe: "1f97f", fleur_de_lis: "269c", flight_arrival: "1f6ec", flight_departure: "1f6eb", flipper: "1f42c", floppy_disk: "1f4be", flower_playing_cards: "1f3b4", flushed: "1f633", flying_disc: "1f94f", flying_saucer: "1f6f8", fog: "1f32b", foggy: "1f301", foot: "1f9b6", football: "1f3c8", footprints: "1f463", fork_and_knife: "1f374", fortune_cookie: "1f960", fountain: "26f2", fountain_pen: "1f58b", four: "0034-20e3", four_leaf_clover: "1f340", fox_face: "1f98a", fr: "1f1eb-1f1f7", framed_picture: "1f5bc", free: "1f193", french_guiana: "1f1ec-1f1eb", french_polynesia: "1f1f5-1f1eb", french_southern_territories: "1f1f9-1f1eb", fried_egg: "1f373", fried_shrimp: "1f364", fries: "1f35f", frog: "1f438", frowning: "1f626", frowning_face: "2639", frowning_man: "1f64d-2642", frowning_person: "1f64d", frowning_woman: "1f64d-2640", fu: "1f595", fuelpump: "26fd", full_moon: "1f315", full_moon_with_face: "1f31d", funeral_urn: "26b1", gabon: "1f1ec-1f1e6", gambia: "1f1ec-1f1f2", game_die: "1f3b2", garlic: "1f9c4", gb: "1f1ec-1f1e7", gear: "2699", gem: "1f48e", gemini: "264a", genie: "1f9de", genie_man: "1f9de-2642", genie_woman: "1f9de-2640", georgia: "1f1ec-1f1ea", ghana: "1f1ec-1f1ed", ghost: "1f47b", gibraltar: "1f1ec-1f1ee", gift: "1f381", gift_heart: "1f49d", giraffe: "1f992", girl: "1f467", globe_with_meridians: "1f310", gloves: "1f9e4", goal_net: "1f945", goat: "1f410", goggles: "1f97d", golf: "26f3", golfing: "1f3cc", golfing_man: "1f3cc-2642", golfing_woman: "1f3cc-2640", gorilla: "1f98d", grapes: "1f347", greece: "1f1ec-1f1f7", green_apple: "1f34f", green_book: "1f4d7", green_circle: "1f7e2", green_heart: "1f49a", green_salad: "1f957", green_square: "1f7e9", greenland: "1f1ec-1f1f1", grenada: "1f1ec-1f1e9", grey_exclamation: "2755", grey_question: "2754", grimacing: "1f62c", grin: "1f601", grinning: "1f600", guadeloupe: "1f1ec-1f1f5", guam: "1f1ec-1f1fa", guard: "1f482", guardsman: "1f482-2642", guardswoman: "1f482-2640", guatemala: "1f1ec-1f1f9", guernsey: "1f1ec-1f1ec", guide_dog: "1f9ae", guinea: "1f1ec-1f1f3", guinea_bissau: "1f1ec-1f1fc", guitar: "1f3b8", gun: "1f52b", guyana: "1f1ec-1f1fe", haircut: "1f487", haircut_man: "1f487-2642", haircut_woman: "1f487-2640", haiti: "1f1ed-1f1f9", hamburger: "1f354", hammer: "1f528", hammer_and_pick: "2692", hammer_and_wrench: "1f6e0", hamster: "1f439", hand: "270b", hand_over_mouth: "1f92d", handbag: "1f45c", handball_person: "1f93e", handshake: "1f91d", hankey: "1f4a9", hash: "0023-20e3", hatched_chick: "1f425", hatching_chick: "1f423", headphones: "1f3a7", health_worker: "1f9d1-2695", hear_no_evil: "1f649", heard_mcdonald_islands: "1f1ed-1f1f2", heart: "2764", heart_decoration: "1f49f", heart_eyes: "1f60d", heart_eyes_cat: "1f63b", heartbeat: "1f493", heartpulse: "1f497", hearts: "2665", heavy_check_mark: "2714", heavy_division_sign: "2797", heavy_dollar_sign: "1f4b2", heavy_exclamation_mark: "2757", heavy_heart_exclamation: "2763", heavy_minus_sign: "2796", heavy_multiplication_x: "2716", heavy_plus_sign: "2795", hedgehog: "1f994", helicopter: "1f681", herb: "1f33f", hibiscus: "1f33a", high_brightness: "1f506", high_heel: "1f460", hiking_boot: "1f97e", hindu_temple: "1f6d5", hippopotamus: "1f99b", hocho: "1f52a", hole: "1f573", honduras: "1f1ed-1f1f3", honey_pot: "1f36f", honeybee: "1f41d", hong_kong: "1f1ed-1f1f0", horse: "1f434", horse_racing: "1f3c7", hospital: "1f3e5", hot_face: "1f975", hot_pepper: "1f336", hotdog: "1f32d", hotel: "1f3e8", hotsprings: "2668", hourglass: "231b", hourglass_flowing_sand: "23f3", house: "1f3e0", house_with_garden: "1f3e1", houses: "1f3d8", hugs: "1f917", hungary: "1f1ed-1f1fa", hushed: "1f62f", ice_cream: "1f368", ice_cube: "1f9ca", ice_hockey: "1f3d2", ice_skate: "26f8", icecream: "1f366", iceland: "1f1ee-1f1f8", id: "1f194", ideograph_advantage: "1f250", imp: "1f47f", inbox_tray: "1f4e5", incoming_envelope: "1f4e8", india: "1f1ee-1f1f3", indonesia: "1f1ee-1f1e9", infinity: "267e", information_desk_person: "1f481", information_source: "2139", innocent: "1f607", interrobang: "2049", iphone: "1f4f1", iran: "1f1ee-1f1f7", iraq: "1f1ee-1f1f6", ireland: "1f1ee-1f1ea", isle_of_man: "1f1ee-1f1f2", israel: "1f1ee-1f1f1", it: "1f1ee-1f1f9", izakaya_lantern: "1f3ee", jack_o_lantern: "1f383", jamaica: "1f1ef-1f1f2", japan: "1f5fe", japanese_castle: "1f3ef", japanese_goblin: "1f47a", japanese_ogre: "1f479", jeans: "1f456", jersey: "1f1ef-1f1ea", jigsaw: "1f9e9", jordan: "1f1ef-1f1f4", joy: "1f602", joy_cat: "1f639", joystick: "1f579", jp: "1f1ef-1f1f5", judge: "1f9d1-2696", juggling_person: "1f939", kaaba: "1f54b", kangaroo: "1f998", kazakhstan: "1f1f0-1f1ff", kenya: "1f1f0-1f1ea", key: "1f511", keyboard: "2328", keycap_ten: "1f51f", kick_scooter: "1f6f4", kimono: "1f458", kiribati: "1f1f0-1f1ee", kiss: "1f48b", kissing: "1f617", kissing_cat: "1f63d", kissing_closed_eyes: "1f61a", kissing_heart: "1f618", kissing_smiling_eyes: "1f619", kite: "1fa81", kiwi_fruit: "1f95d", kneeling_man: "1f9ce-2642", kneeling_person: "1f9ce", kneeling_woman: "1f9ce-2640", knife: "1f52a", koala: "1f428", koko: "1f201", kosovo: "1f1fd-1f1f0", kr: "1f1f0-1f1f7", kuwait: "1f1f0-1f1fc", kyrgyzstan: "1f1f0-1f1ec", lab_coat: "1f97c", label: "1f3f7", lacrosse: "1f94d", lantern: "1f3ee", laos: "1f1f1-1f1e6", large_blue_circle: "1f535", large_blue_diamond: "1f537", large_orange_diamond: "1f536", last_quarter_moon: "1f317", last_quarter_moon_with_face: "1f31c", latin_cross: "271d", latvia: "1f1f1-1f1fb", laughing: "1f606", leafy_green: "1f96c", leaves: "1f343", lebanon: "1f1f1-1f1e7", ledger: "1f4d2", left_luggage: "1f6c5", left_right_arrow: "2194", left_speech_bubble: "1f5e8", leftwards_arrow_with_hook: "21a9", leg: "1f9b5", lemon: "1f34b", leo: "264c", leopard: "1f406", lesotho: "1f1f1-1f1f8", level_slider: "1f39a", liberia: "1f1f1-1f1f7", libra: "264e", libya: "1f1f1-1f1fe", liechtenstein: "1f1f1-1f1ee", light_rail: "1f688", link: "1f517", lion: "1f981", lips: "1f444", lipstick: "1f484", lithuania: "1f1f1-1f1f9", lizard: "1f98e", llama: "1f999", lobster: "1f99e", lock: "1f512", lock_with_ink_pen: "1f50f", lollipop: "1f36d", loop: "27bf", lotion_bottle: "1f9f4", lotus_position: "1f9d8", lotus_position_man: "1f9d8-2642", lotus_position_woman: "1f9d8-2640", loud_sound: "1f50a", loudspeaker: "1f4e2", love_hotel: "1f3e9", love_letter: "1f48c", love_you_gesture: "1f91f", low_brightness: "1f505", luggage: "1f9f3", luxembourg: "1f1f1-1f1fa", lying_face: "1f925", m: "24c2", macau: "1f1f2-1f1f4", macedonia: "1f1f2-1f1f0", madagascar: "1f1f2-1f1ec", mag: "1f50d", mag_right: "1f50e", mage: "1f9d9", mage_man: "1f9d9-2642", mage_woman: "1f9d9-2640", magnet: "1f9f2", mahjong: "1f004", mailbox: "1f4eb", mailbox_closed: "1f4ea", mailbox_with_mail: "1f4ec", mailbox_with_no_mail: "1f4ed", malawi: "1f1f2-1f1fc", malaysia: "1f1f2-1f1fe", maldives: "1f1f2-1f1fb", male_detective: "1f575-2642", male_sign: "2642", mali: "1f1f2-1f1f1", malta: "1f1f2-1f1f9", man: "1f468", man_artist: "1f468-1f3a8", man_astronaut: "1f468-1f680", man_cartwheeling: "1f938-2642", man_cook: "1f468-1f373", man_dancing: "1f57a", man_facepalming: "1f926-2642", man_factory_worker: "1f468-1f3ed", man_farmer: "1f468-1f33e", man_firefighter: "1f468-1f692", man_health_worker: "1f468-2695", man_in_manual_wheelchair: "1f468-1f9bd", man_in_motorized_wheelchair: "1f468-1f9bc", man_in_tuxedo: "1f935", man_judge: "1f468-2696", man_juggling: "1f939-2642", man_mechanic: "1f468-1f527", man_office_worker: "1f468-1f4bc", man_pilot: "1f468-2708", man_playing_handball: "1f93e-2642", man_playing_water_polo: "1f93d-2642", man_scientist: "1f468-1f52c", man_shrugging: "1f937-2642", man_singer: "1f468-1f3a4", man_student: "1f468-1f393", man_teacher: "1f468-1f3eb", man_technologist: "1f468-1f4bb", man_with_gua_pi_mao: "1f472", man_with_probing_cane: "1f468-1f9af", man_with_turban: "1f473-2642", mandarin: "1f34a", mango: "1f96d", mans_shoe: "1f45e", mantelpiece_clock: "1f570", manual_wheelchair: "1f9bd", maple_leaf: "1f341", marshall_islands: "1f1f2-1f1ed", martial_arts_uniform: "1f94b", martinique: "1f1f2-1f1f6", mask: "1f637", massage: "1f486", massage_man: "1f486-2642", massage_woman: "1f486-2640", mate: "1f9c9", mauritania: "1f1f2-1f1f7", mauritius: "1f1f2-1f1fa", mayotte: "1f1fe-1f1f9", meat_on_bone: "1f356", mechanic: "1f9d1-1f527", mechanical_arm: "1f9be", mechanical_leg: "1f9bf", medal_military: "1f396", medal_sports: "1f3c5", medical_symbol: "2695", mega: "1f4e3", melon: "1f348", memo: "1f4dd", men_wrestling: "1f93c-2642", menorah: "1f54e", mens: "1f6b9", mermaid: "1f9dc-2640", merman: "1f9dc-2642", merperson: "1f9dc", metal: "1f918", metro: "1f687", mexico: "1f1f2-1f1fd", microbe: "1f9a0", micronesia: "1f1eb-1f1f2", microphone: "1f3a4", microscope: "1f52c", middle_finger: "1f595", milk_glass: "1f95b", milky_way: "1f30c", minibus: "1f690", minidisc: "1f4bd", mobile_phone_off: "1f4f4", moldova: "1f1f2-1f1e9", monaco: "1f1f2-1f1e8", money_mouth_face: "1f911", money_with_wings: "1f4b8", moneybag: "1f4b0", mongolia: "1f1f2-1f1f3", monkey: "1f412", monkey_face: "1f435", monocle_face: "1f9d0", monorail: "1f69d", montenegro: "1f1f2-1f1ea", montserrat: "1f1f2-1f1f8", moon: "1f314", moon_cake: "1f96e", morocco: "1f1f2-1f1e6", mortar_board: "1f393", mosque: "1f54c", mosquito: "1f99f", motor_boat: "1f6e5", motor_scooter: "1f6f5", motorcycle: "1f3cd", motorized_wheelchair: "1f9bc", motorway: "1f6e3", mount_fuji: "1f5fb", mountain: "26f0", mountain_bicyclist: "1f6b5", mountain_biking_man: "1f6b5-2642", mountain_biking_woman: "1f6b5-2640", mountain_cableway: "1f6a0", mountain_railway: "1f69e", mountain_snow: "1f3d4", mouse: "1f42d", mouse2: "1f401", movie_camera: "1f3a5", moyai: "1f5ff", mozambique: "1f1f2-1f1ff", mrs_claus: "1f936", muscle: "1f4aa", mushroom: "1f344", musical_keyboard: "1f3b9", musical_note: "1f3b5", musical_score: "1f3bc", mute: "1f507", myanmar: "1f1f2-1f1f2", nail_care: "1f485", name_badge: "1f4db", namibia: "1f1f3-1f1e6", national_park: "1f3de", nauru: "1f1f3-1f1f7", nauseated_face: "1f922", nazar_amulet: "1f9ff", necktie: "1f454", negative_squared_cross_mark: "274e", nepal: "1f1f3-1f1f5", nerd_face: "1f913", netherlands: "1f1f3-1f1f1", neutral_face: "1f610", new: "1f195", new_caledonia: "1f1f3-1f1e8", new_moon: "1f311", new_moon_with_face: "1f31a", new_zealand: "1f1f3-1f1ff", newspaper: "1f4f0", newspaper_roll: "1f5de", next_track_button: "23ed", ng: "1f196", ng_man: "1f645-2642", ng_woman: "1f645-2640", nicaragua: "1f1f3-1f1ee", niger: "1f1f3-1f1ea", nigeria: "1f1f3-1f1ec", night_with_stars: "1f303", nine: "0039-20e3", niue: "1f1f3-1f1fa", no_bell: "1f515", no_bicycles: "1f6b3", no_entry: "26d4", no_entry_sign: "1f6ab", no_good: "1f645", no_good_man: "1f645-2642", no_good_woman: "1f645-2640", no_mobile_phones: "1f4f5", no_mouth: "1f636", no_pedestrians: "1f6b7", no_smoking: "1f6ad", "non-potable_water": "1f6b1", norfolk_island: "1f1f3-1f1eb", north_korea: "1f1f0-1f1f5", northern_mariana_islands: "1f1f2-1f1f5", norway: "1f1f3-1f1f4", nose: "1f443", notebook: "1f4d3", notebook_with_decorative_cover: "1f4d4", notes: "1f3b6", nut_and_bolt: "1f529", o: "2b55", o2: "1f17e", ocean: "1f30a", octopus: "1f419", oden: "1f362", office: "1f3e2", office_worker: "1f9d1-1f4bc", oil_drum: "1f6e2", ok: "1f197", ok_hand: "1f44c", ok_man: "1f646-2642", ok_person: "1f646", ok_woman: "1f646-2640", old_key: "1f5dd", older_adult: "1f9d3", older_man: "1f474", older_woman: "1f475", om: "1f549", oman: "1f1f4-1f1f2", on: "1f51b", oncoming_automobile: "1f698", oncoming_bus: "1f68d", oncoming_police_car: "1f694", oncoming_taxi: "1f696", one: "0031-20e3", one_piece_swimsuit: "1fa71", onion: "1f9c5", open_book: "1f4d6", open_file_folder: "1f4c2", open_hands: "1f450", open_mouth: "1f62e", open_umbrella: "2602", ophiuchus: "26ce", orange: "1f34a", orange_book: "1f4d9", orange_circle: "1f7e0", orange_heart: "1f9e1", orange_square: "1f7e7", orangutan: "1f9a7", orthodox_cross: "2626", otter: "1f9a6", outbox_tray: "1f4e4", owl: "1f989", ox: "1f402", oyster: "1f9aa", package: "1f4e6", page_facing_up: "1f4c4", page_with_curl: "1f4c3", pager: "1f4df", paintbrush: "1f58c", pakistan: "1f1f5-1f1f0", palau: "1f1f5-1f1fc", palestinian_territories: "1f1f5-1f1f8", palm_tree: "1f334", palms_up_together: "1f932", panama: "1f1f5-1f1e6", pancakes: "1f95e", panda_face: "1f43c", paperclip: "1f4ce", paperclips: "1f587", papua_new_guinea: "1f1f5-1f1ec", parachute: "1fa82", paraguay: "1f1f5-1f1fe", parasol_on_ground: "26f1", parking: "1f17f", parrot: "1f99c", part_alternation_mark: "303d", partly_sunny: "26c5", partying_face: "1f973", passenger_ship: "1f6f3", passport_control: "1f6c2", pause_button: "23f8", paw_prints: "1f43e", peace_symbol: "262e", peach: "1f351", peacock: "1f99a", peanuts: "1f95c", pear: "1f350", pen: "1f58a", pencil: "1f4dd", pencil2: "270f", penguin: "1f427", pensive: "1f614", people_holding_hands: "1f9d1-1f91d-1f9d1", performing_arts: "1f3ad", persevere: "1f623", person_bald: "1f9d1-1f9b2", person_curly_hair: "1f9d1-1f9b1", person_fencing: "1f93a", person_in_manual_wheelchair: "1f9d1-1f9bd", person_in_motorized_wheelchair: "1f9d1-1f9bc", person_red_hair: "1f9d1-1f9b0", person_white_hair: "1f9d1-1f9b3", person_with_probing_cane: "1f9d1-1f9af", person_with_turban: "1f473", peru: "1f1f5-1f1ea", petri_dish: "1f9eb", philippines: "1f1f5-1f1ed", phone: "260e", pick: "26cf", pie: "1f967", pig: "1f437", pig2: "1f416", pig_nose: "1f43d", pill: "1f48a", pilot: "1f9d1-2708", pinching_hand: "1f90f", pineapple: "1f34d", ping_pong: "1f3d3", pirate_flag: "1f3f4-2620", pisces: "2653", pitcairn_islands: "1f1f5-1f1f3", pizza: "1f355", place_of_worship: "1f6d0", plate_with_cutlery: "1f37d", play_or_pause_button: "23ef", pleading_face: "1f97a", point_down: "1f447", point_left: "1f448", point_right: "1f449", point_up: "261d", point_up_2: "1f446", poland: "1f1f5-1f1f1", police_car: "1f693", police_officer: "1f46e", policeman: "1f46e-2642", policewoman: "1f46e-2640", poodle: "1f429", poop: "1f4a9", popcorn: "1f37f", portugal: "1f1f5-1f1f9", post_office: "1f3e3", postal_horn: "1f4ef", postbox: "1f4ee", potable_water: "1f6b0", potato: "1f954", pouch: "1f45d", poultry_leg: "1f357", pound: "1f4b7", pout: "1f621", pouting_cat: "1f63e", pouting_face: "1f64e", pouting_man: "1f64e-2642", pouting_woman: "1f64e-2640", pray: "1f64f", prayer_beads: "1f4ff", pregnant_woman: "1f930", pretzel: "1f968", previous_track_button: "23ee", prince: "1f934", princess: "1f478", printer: "1f5a8", probing_cane: "1f9af", puerto_rico: "1f1f5-1f1f7", punch: "1f44a", purple_circle: "1f7e3", purple_heart: "1f49c", purple_square: "1f7ea", purse: "1f45b", pushpin: "1f4cc", put_litter_in_its_place: "1f6ae", qatar: "1f1f6-1f1e6", question: "2753", rabbit: "1f430", rabbit2: "1f407", raccoon: "1f99d", racehorse: "1f40e", racing_car: "1f3ce", radio: "1f4fb", radio_button: "1f518", radioactive: "2622", rage: "1f621", railway_car: "1f683", railway_track: "1f6e4", rainbow: "1f308", rainbow_flag: "1f3f3-1f308", raised_back_of_hand: "1f91a", raised_eyebrow: "1f928", raised_hand: "270b", raised_hand_with_fingers_splayed: "1f590", raised_hands: "1f64c", raising_hand: "1f64b", raising_hand_man: "1f64b-2642", raising_hand_woman: "1f64b-2640", ram: "1f40f", ramen: "1f35c", rat: "1f400", razor: "1fa92", receipt: "1f9fe", record_button: "23fa", recycle: "267b", red_car: "1f697", red_circle: "1f534", red_envelope: "1f9e7", red_haired_man: "1f468-1f9b0", red_haired_woman: "1f469-1f9b0", red_square: "1f7e5", registered: "00ae", relaxed: "263a", relieved: "1f60c", reminder_ribbon: "1f397", repeat: "1f501", repeat_one: "1f502", rescue_worker_helmet: "26d1", restroom: "1f6bb", reunion: "1f1f7-1f1ea", revolving_hearts: "1f49e", rewind: "23ea", rhinoceros: "1f98f", ribbon: "1f380", rice: "1f35a", rice_ball: "1f359", rice_cracker: "1f358", rice_scene: "1f391", right_anger_bubble: "1f5ef", ring: "1f48d", ringed_planet: "1fa90", robot: "1f916", rocket: "1f680", rofl: "1f923", roll_eyes: "1f644", roll_of_paper: "1f9fb", roller_coaster: "1f3a2", romania: "1f1f7-1f1f4", rooster: "1f413", rose: "1f339", rosette: "1f3f5", rotating_light: "1f6a8", round_pushpin: "1f4cd", rowboat: "1f6a3", rowing_man: "1f6a3-2642", rowing_woman: "1f6a3-2640", ru: "1f1f7-1f1fa", rugby_football: "1f3c9", runner: "1f3c3", running: "1f3c3", running_man: "1f3c3-2642", running_shirt_with_sash: "1f3bd", running_woman: "1f3c3-2640", rwanda: "1f1f7-1f1fc", sa: "1f202", safety_pin: "1f9f7", safety_vest: "1f9ba", sagittarius: "2650", sailboat: "26f5", sake: "1f376", salt: "1f9c2", samoa: "1f1fc-1f1f8", san_marino: "1f1f8-1f1f2", sandal: "1f461", sandwich: "1f96a", santa: "1f385", sao_tome_principe: "1f1f8-1f1f9", sari: "1f97b", sassy_man: "1f481-2642", sassy_woman: "1f481-2640", satellite: "1f4e1", satisfied: "1f606", saudi_arabia: "1f1f8-1f1e6", sauna_man: "1f9d6-2642", sauna_person: "1f9d6", sauna_woman: "1f9d6-2640", sauropod: "1f995", saxophone: "1f3b7", scarf: "1f9e3", school: "1f3eb", school_satchel: "1f392", scientist: "1f9d1-1f52c", scissors: "2702", scorpion: "1f982", scorpius: "264f", scotland: "1f3f4-e0067-e0062-e0073-e0063-e0074-e007f", scream: "1f631", scream_cat: "1f640", scroll: "1f4dc", seat: "1f4ba", secret: "3299", see_no_evil: "1f648", seedling: "1f331", selfie: "1f933", senegal: "1f1f8-1f1f3", serbia: "1f1f7-1f1f8", service_dog: "1f415-1f9ba", seven: "0037-20e3", seychelles: "1f1f8-1f1e8", shallow_pan_of_food: "1f958", shamrock: "2618", shark: "1f988", shaved_ice: "1f367", sheep: "1f411", shell: "1f41a", shield: "1f6e1", shinto_shrine: "26e9", ship: "1f6a2", shirt: "1f455", poo: "1f4a9", shoe: "1f45e", shopping: "1f6cd", shopping_cart: "1f6d2", shorts: "1fa73", shower: "1f6bf", shrimp: "1f990", shrug: "1f937", shushing_face: "1f92b", sierra_leone: "1f1f8-1f1f1", signal_strength: "1f4f6", singapore: "1f1f8-1f1ec", singer: "1f9d1-1f3a4", sint_maarten: "1f1f8-1f1fd", six: "0036-20e3", six_pointed_star: "1f52f", skateboard: "1f6f9", ski: "1f3bf", skier: "26f7", skull: "1f480", skull_and_crossbones: "2620", skunk: "1f9a8", sled: "1f6f7", sleeping: "1f634", sleeping_bed: "1f6cc", sleepy: "1f62a", slightly_frowning_face: "1f641", slightly_smiling_face: "1f642", slot_machine: "1f3b0", sloth: "1f9a5", slovakia: "1f1f8-1f1f0", slovenia: "1f1f8-1f1ee", small_airplane: "1f6e9", small_blue_diamond: "1f539", small_orange_diamond: "1f538", small_red_triangle: "1f53a", small_red_triangle_down: "1f53b", smile: "1f604", smile_cat: "1f638", smiley: "1f603", smiley_cat: "1f63a", smiling_face_with_three_hearts: "1f970", smiling_imp: "1f608", smirk: "1f60f", smirk_cat: "1f63c", smoking: "1f6ac", snail: "1f40c", snake: "1f40d", sneezing_face: "1f927", snowboarder: "1f3c2", snowflake: "2744", snowman: "26c4", snowman_with_snow: "2603", soap: "1f9fc", sob: "1f62d", soccer: "26bd", socks: "1f9e6", softball: "1f94e", solomon_islands: "1f1f8-1f1e7", somalia: "1f1f8-1f1f4", soon: "1f51c", sos: "1f198", sound: "1f509", south_africa: "1f1ff-1f1e6", south_georgia_south_sandwich_islands: "1f1ec-1f1f8", south_sudan: "1f1f8-1f1f8", space_invader: "1f47e", spades: "2660", spaghetti: "1f35d", sparkle: "2747", sparkler: "1f387", sparkles: "2728", sparkling_heart: "1f496", speak_no_evil: "1f64a", speaker: "1f508", speaking_head: "1f5e3", speech_balloon: "1f4ac", speedboat: "1f6a4", spider: "1f577", spider_web: "1f578", spiral_calendar: "1f5d3", spiral_notepad: "1f5d2", sponge: "1f9fd", spoon: "1f944", squid: "1f991", sri_lanka: "1f1f1-1f1f0", st_barthelemy: "1f1e7-1f1f1", st_helena: "1f1f8-1f1ed", st_kitts_nevis: "1f1f0-1f1f3", st_lucia: "1f1f1-1f1e8", st_martin: "1f1f2-1f1eb", st_pierre_miquelon: "1f1f5-1f1f2", st_vincent_grenadines: "1f1fb-1f1e8", stadium: "1f3df", standing_man: "1f9cd-2642", standing_person: "1f9cd", standing_woman: "1f9cd-2640", star: "2b50", star2: "1f31f", star_and_crescent: "262a", star_of_david: "2721", star_struck: "1f929", stars: "1f320", station: "1f689", statue_of_liberty: "1f5fd", steam_locomotive: "1f682", stethoscope: "1fa7a", stew: "1f372", stop_button: "23f9", stop_sign: "1f6d1", stopwatch: "23f1", straight_ruler: "1f4cf", strawberry: "1f353", stuck_out_tongue: "1f61b", stuck_out_tongue_closed_eyes: "1f61d", stuck_out_tongue_winking_eye: "1f61c", student: "1f9d1-1f393", studio_microphone: "1f399", stuffed_flatbread: "1f959", sudan: "1f1f8-1f1e9", sun_behind_large_cloud: "1f325", sun_behind_rain_cloud: "1f326", sun_behind_small_cloud: "1f324", sun_with_face: "1f31e", sunflower: "1f33b", sunglasses: "1f60e", sunny: "2600", sunrise: "1f305", sunrise_over_mountains: "1f304", superhero: "1f9b8", superhero_man: "1f9b8-2642", superhero_woman: "1f9b8-2640", supervillain: "1f9b9", supervillain_man: "1f9b9-2642", supervillain_woman: "1f9b9-2640", surfer: "1f3c4", surfing_man: "1f3c4-2642", surfing_woman: "1f3c4-2640", suriname: "1f1f8-1f1f7", sushi: "1f363", suspension_railway: "1f69f", svalbard_jan_mayen: "1f1f8-1f1ef", swan: "1f9a2", swaziland: "1f1f8-1f1ff", sweat: "1f613", sweat_drops: "1f4a6", sweat_smile: "1f605", sweden: "1f1f8-1f1ea", sweet_potato: "1f360", swim_brief: "1fa72", swimmer: "1f3ca", swimming_man: "1f3ca-2642", swimming_woman: "1f3ca-2640", switzerland: "1f1e8-1f1ed", symbols: "1f523", synagogue: "1f54d", syria: "1f1f8-1f1fe", syringe: "1f489", "t-rex": "1f996", taco: "1f32e", tada: "1f389", taiwan: "1f1f9-1f1fc", tajikistan: "1f1f9-1f1ef", takeout_box: "1f961", tanabata_tree: "1f38b", tangerine: "1f34a", tanzania: "1f1f9-1f1ff", taurus: "2649", taxi: "1f695", tea: "1f375", teacher: "1f9d1-1f3eb", technologist: "1f9d1-1f4bb", teddy_bear: "1f9f8", telephone: "260e", telephone_receiver: "1f4de", telescope: "1f52d", tennis: "1f3be", tent: "26fa", test_tube: "1f9ea", thailand: "1f1f9-1f1ed", thermometer: "1f321", thinking: "1f914", thought_balloon: "1f4ad", thread: "1f9f5", three: "0033-20e3", thumbsdown: "1f44e", thumbsup: "1f44d", ticket: "1f3ab", tickets: "1f39f", tiger: "1f42f", tiger2: "1f405", timer_clock: "23f2", timor_leste: "1f1f9-1f1f1", tipping_hand_man: "1f481-2642", tipping_hand_person: "1f481", tipping_hand_woman: "1f481-2640", tired_face: "1f62b", tm: "2122", togo: "1f1f9-1f1ec", toilet: "1f6bd", tokelau: "1f1f9-1f1f0", tokyo_tower: "1f5fc", tomato: "1f345", tonga: "1f1f9-1f1f4", tongue: "1f445", toolbox: "1f9f0", tooth: "1f9b7", top: "1f51d", tophat: "1f3a9", tornado: "1f32a", tr: "1f1f9-1f1f7", trackball: "1f5b2", tractor: "1f69c", traffic_light: "1f6a5", train: "1f68b", train2: "1f686", tram: "1f68a", triangular_flag_on_post: "1f6a9", triangular_ruler: "1f4d0", trident: "1f531", trinidad_tobago: "1f1f9-1f1f9", tristan_da_cunha: "1f1f9-1f1e6", triumph: "1f624", trolleybus: "1f68e", trophy: "1f3c6", tropical_drink: "1f379", tropical_fish: "1f420", truck: "1f69a", trumpet: "1f3ba", tshirt: "1f455", tulip: "1f337", tumbler_glass: "1f943", tunisia: "1f1f9-1f1f3", turkey: "1f983", turkmenistan: "1f1f9-1f1f2", turks_caicos_islands: "1f1f9-1f1e8", turtle: "1f422", tuvalu: "1f1f9-1f1fb", tv: "1f4fa", twisted_rightwards_arrows: "1f500", two: "0032-20e3", two_hearts: "1f495", two_men_holding_hands: "1f46c", two_women_holding_hands: "1f46d", u5272: "1f239", u5408: "1f234", u55b6: "1f23a", u6307: "1f22f", u6708: "1f237", u6709: "1f236", u6e80: "1f235", u7121: "1f21a", u7533: "1f238", u7981: "1f232", u7a7a: "1f233", uganda: "1f1fa-1f1ec", uk: "1f1ec-1f1e7", ukraine: "1f1fa-1f1e6", umbrella: "2614", unamused: "1f612", underage: "1f51e", unicorn: "1f984", united_arab_emirates: "1f1e6-1f1ea", united_nations: "1f1fa-1f1f3", unlock: "1f513", up: "1f199", upside_down_face: "1f643", uruguay: "1f1fa-1f1fe", us: "1f1fa-1f1f8", us_outlying_islands: "1f1fa-1f1f2", us_virgin_islands: "1f1fb-1f1ee", uzbekistan: "1f1fa-1f1ff", v: "270c", vampire: "1f9db", vampire_man: "1f9db-2642", vampire_woman: "1f9db-2640", vanuatu: "1f1fb-1f1fa", vatican_city: "1f1fb-1f1e6", venezuela: "1f1fb-1f1ea", vertical_traffic_light: "1f6a6", vhs: "1f4fc", vibration_mode: "1f4f3", video_camera: "1f4f9", video_game: "1f3ae", vietnam: "1f1fb-1f1f3", violin: "1f3bb", virgo: "264d", volcano: "1f30b", volleyball: "1f3d0", vomiting_face: "1f92e", vs: "1f19a", vulcan_salute: "1f596", waffle: "1f9c7", wales: "1f3f4-e0067-e0062-e0077-e006c-e0073-e007f", walking: "1f6b6", walking_man: "1f6b6-2642", walking_woman: "1f6b6-2640", wallis_futuna: "1f1fc-1f1eb", waning_crescent_moon: "1f318", waning_gibbous_moon: "1f316", warning: "26a0", wastebasket: "1f5d1", watch: "231a", water_buffalo: "1f403", water_polo: "1f93d", watermelon: "1f349", wave: "1f44b", wavy_dash: "3030", waxing_crescent_moon: "1f312", waxing_gibbous_moon: "1f314", wc: "1f6be", weary: "1f629", wedding: "1f492", weight_lifting: "1f3cb", weight_lifting_man: "1f3cb-2642", weight_lifting_woman: "1f3cb-2640", western_sahara: "1f1ea-1f1ed", whale: "1f433", whale2: "1f40b", wheel_of_dharma: "2638", wheelchair: "267f", white_check_mark: "2705", white_circle: "26aa", white_flag: "1f3f3", white_flower: "1f4ae", white_haired_man: "1f468-1f9b3", white_haired_woman: "1f469-1f9b3", white_heart: "1f90d", white_large_square: "2b1c", white_medium_small_square: "25fd", white_medium_square: "25fb", white_small_square: "25ab", white_square_button: "1f533", wilted_flower: "1f940", wind_chime: "1f390", wind_face: "1f32c", wine_glass: "1f377", wink: "1f609", wolf: "1f43a", woman: "1f469", woman_artist: "1f469-1f3a8", woman_astronaut: "1f469-1f680", woman_cartwheeling: "1f938-2640", woman_cook: "1f469-1f373", woman_dancing: "1f483", woman_facepalming: "1f926-2640", woman_factory_worker: "1f469-1f3ed", woman_farmer: "1f469-1f33e", woman_firefighter: "1f469-1f692", woman_health_worker: "1f469-2695", woman_in_manual_wheelchair: "1f469-1f9bd", woman_in_motorized_wheelchair: "1f469-1f9bc", woman_judge: "1f469-2696", woman_juggling: "1f939-2640", woman_mechanic: "1f469-1f527", woman_office_worker: "1f469-1f4bc", woman_pilot: "1f469-2708", woman_playing_handball: "1f93e-2640", woman_playing_water_polo: "1f93d-2640", woman_scientist: "1f469-1f52c", woman_shrugging: "1f937-2640", woman_singer: "1f469-1f3a4", woman_student: "1f469-1f393", woman_teacher: "1f469-1f3eb", woman_technologist: "1f469-1f4bb", woman_with_headscarf: "1f9d5", woman_with_probing_cane: "1f469-1f9af", woman_with_turban: "1f473-2640", womans_clothes: "1f45a", womans_hat: "1f452", women_wrestling: "1f93c-2640", womens: "1f6ba", woozy_face: "1f974", world_map: "1f5fa", worried: "1f61f", wrench: "1f527", wrestling: "1f93c", writing_hand: "270d", x: "274c", yarn: "1f9f6", yawning_face: "1f971", yellow_circle: "1f7e1", yellow_heart: "1f49b", yellow_square: "1f7e8", yemen: "1f1fe-1f1ea", yen: "1f4b4", yin_yang: "262f", yo_yo: "1fa80", yum: "1f60b", zambia: "1f1ff-1f1f2", zany_face: "1f92a", zap: "26a1", zebra: "1f993", zero: "0030-20e3", zimbabwe: "1f1ff-1f1fc", zipper_mouth_face: "1f910", zombie: "1f9df", zombie_man: "1f9df-2642", zombie_woman: "1f9df-2640", zzz: "1f4a4" };
function yl(e, t) {
  var r = ue(e);
  if (xe) {
    var n = xe(e);
    t && (n = Xe(n).call(n, function(a) {
      return Te(e, a).enumerable;
    })), r.push.apply(r, n);
  }
  return r;
}
function Mb(e) {
  for (var t = 1; t < arguments.length; t++) {
    var r, n, a = arguments[t] != null ? arguments[t] : {};
    t % 2 ? q(r = yl(Object(a), !0)).call(r, function(i) {
      R(e, i, a[i]);
    }) : Ce ? Ct(e, Ce(a)) : q(n = yl(Object(a))).call(n, function(i) {
      nt(e, i, Te(a, i));
    });
  }
  return e;
}
function jb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
function Db() {
  for (var e = [], t = 0, r = "", n = 0, a = arguments.length; n !== a; ++n) {
    var i = +(n < 0 || arguments.length <= n ? void 0 : arguments[n]);
    if (!(i < 1114111 && i >>> 0 === i))
      throw new RangeError("Invalid code point: ".concat(i));
    i <= 65535 ? t = e.push(i) : (i -= 65536, t = e.push(55296 + (i >> 10), i % 1024 + 56320)), t >= 16383 && (r += String.fromCharCode.apply(null, e), e.length = 0);
  }
  return r + String.fromCharCode.apply(null, e);
}
var Zf = function(e) {
  U(r, pe);
  var t = jb(r);
  function r() {
    var n, a = (arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : { config: void 0 }).config;
    if (j(this, r), (n = t.call(this, { config: a })).options = { useUnicode: !0, upperCase: !1, customHandled: !1, resourceURL: "https://github.githubassets.com/images/icons/emoji/unicode/${code}.png?v8", emojis: Mb({}, Nb) }, He(a) !== "object")
      return B(n);
    var i = a.useUnicode, o = a.customResourceURL, s = a.customRenderer, c = a.upperCase;
    return n.options.useUnicode = typeof i == "boolean" ? i : n.options.useUnicode, n.options.upperCase = typeof c == "boolean" ? c : n.options.upperCase, i === !1 && typeof o == "string" && (n.options.resourceURL = o), typeof s == "function" && (n.options.customHandled = !0, n.options.customRenderer = s), n;
  }
  return M(r, [{ key: "makeHtml", value: function(n, a) {
    var i = this;
    return this.test(n) ? n.replace(this.RULE.reg, function(o, s) {
      var c;
      if (i.options.customHandled && typeof i.options.customRenderer == "function")
        return i.options.customRenderer(s);
      var u = i.options.emojis[s];
      if (typeof u != "string")
        return o;
      if (i.options.useUnicode) {
        var l, f = ve(l = u.split("-")).call(l, function(d) {
          return "0x".concat(d);
        });
        return Db.apply(void 0, lr(f));
      }
      i.options.upperCase && (u = u.toUpperCase());
      var p = i.options.resourceURL.replace(/\$\{code\}/g, u);
      return b(c = '<img class="emoji" src="'.concat(p, '" alt="')).call(c, qe(s), '" />');
    }) : n;
  } }, { key: "rule", value: function() {
    var n = { begin: ":", content: "([a-zA-Z0-9+_]+?)", end: ":" };
    return n.reg = Le(n, "g"), n;
  } }]), r;
}();
function Fb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Zf, "HOOK_NAME", "emoji");
var Yf = function(e) {
  U(r, pe);
  var t = Fb(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "makeHtml", value: function(n) {
    return this.test(n) ? n.replace(this.RULE.reg, '$1<span style="text-decoration: underline;">$2</span>$3') : n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(^| )\\/", end: "\\/( |$)", content: "([^\\n]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function Bb(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(Yf, "HOOK_NAME", "underline");
var Xf = function(e) {
  U(r, pe);
  var t = Bb(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "makeHtml", value: function(n) {
    return this.test(n) ? n.replace(this.RULE.reg, "$1<mark>$2</mark>$3") : n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(^| )==", end: "==( |$|\\n)", content: "([^\\n]+?)" };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
R(Xf, "HOOK_NAME", "highLight");
var Hb = Au;
X.JSON || (X.JSON = { stringify: JSON.stringify });
var zb = function(e, t, r) {
  return Mt(X.JSON.stringify, null, arguments);
}, Ub = io.includes, Wb = J(function() {
  return !Array(1).includes();
});
L({ target: "Array", proto: !0, forced: Wb }, { includes: function(e) {
  return Ub(this, e, arguments.length > 1 ? arguments[1] : void 0);
} });
var qb = Ie("Array").includes, Gb = z("".indexOf);
L({ target: "String", proto: !0, forced: !Df("includes") }, { includes: function(e) {
  return !!~Gb(tt(an(this)), tt(jf(e)), arguments.length > 1 ? arguments[1] : void 0);
} });
var Kb = Ie("String").includes, ki = Array.prototype, wi = String.prototype, Vf = function(e) {
  var t = e.includes;
  return e === ki || de(ki, e) && t === ki.includes ? qb : typeof e == "string" || e === wi || de(wi, e) && t === wi.includes ? Kb : t;
}, Zb = O.TypeError, Yb = /MSIE .\./.test(zr), Xb = O.Function, _l = function(e) {
  return Yb ? function(t, r) {
    var n = function(o, s) {
      if (o < s)
        throw Zb("Not enough arguments");
      return o;
    }(arguments.length, 1) > 2, a = oe(t) ? t : Xb(t), i = n ? Qt(arguments, 2) : void 0;
    return e(n ? function() {
      Mt(a, this, i);
    } : a, r);
  } : e;
}, Jf = { setTimeout: _l(O.setTimeout), setInterval: _l(O.setInterval) }, kl = Jf.setInterval;
L({ global: !0, bind: !0, forced: O.setInterval !== kl }, { setInterval: kl });
var wl = Jf.setTimeout;
L({ global: !0, bind: !0, forced: O.setTimeout !== wl }, { setTimeout: wl });
var Ei = X.setTimeout, Vb = function(e, t) {
  for (var r = -1, n = e == null ? 0 : e.length, a = Array(n); ++r < n; )
    a[r] = t(e[r], r, e);
  return a;
}, Jb = function(e) {
  return typeof e == "symbol" || Wt(e) && cn(e) == "[object Symbol]";
}, El = Bt ? Bt.prototype : void 0, Sl = El ? El.toString : void 0, Qb = function e(t) {
  if (typeof t == "string")
    return t;
  if (Pn(t))
    return Vb(t, e) + "";
  if (Jb(t))
    return Sl ? Sl.call(t) : "";
  var r = t + "";
  return r == "0" && 1 / t == -1 / 0 ? "-0" : r;
}, ev = function(e) {
  return e == null ? "" : Qb(e);
}, Qf = /[\\^$.*+?()[\]{}|]/g, tv = RegExp(Qf.source), nv = function(e) {
  return (e = ev(e)) && tv.test(e) ? e.replace(Qf, "\\$&") : e;
}, rv = "/·￥、：“”【】（）《》".concat("#"), av = [{ icon: "h1", label: "H1 Heading", keyword: "head1", value: "# " }, { icon: "h2", label: "H2  Heading", keyword: "head2", value: "## " }, { icon: "h3", label: "H3  Heading", keyword: "head3", value: "### " }, { icon: "table", label: "Table", keyword: "table", value: `| Header | Header | Header |
| --- | --- | --- |
| Content | Content | Content |
` }, { icon: "code", label: "Code", keyword: "code", value: "```\n\n```\n" }, { icon: "link", label: "Link", keyword: "link", value: "[title](https://url)", selection: { from: 19, to: 14 } }, { icon: "checklist", label: "Checklist", keyword: "checklist", value: `- [ ] item
- [x] item` }, { icon: "tips", label: "Panel", keyword: "panel tips info warning danger success", value: `::: primary title
content
:::
` }, { icon: "insertFlow", label: "Detail", keyword: "detail", value: `+++ 点击展开更多
内容
++- 默认展开
内容
++ 默认收起
内容
+++
` }], Al = [{ icon: "FullWidth", label: "`", keyword: "···", value: "`" }, { icon: "FullWidth", label: "$", keyword: "￥", value: "$" }, { icon: "FullWidth", label: "/", keyword: "、", value: "/" }, { icon: "FullWidth", label: "\\", keyword: "、", value: "\\" }, { icon: "FullWidth", label: ":", keyword: "：", value: ":" }, { icon: "FullWidth", label: '"', keyword: "“", value: '"' }, { icon: "FullWidth", label: '"', keyword: "”", value: '"' }, { icon: "FullWidth", label: "[", keyword: "【", value: "[" }, { icon: "FullWidth", label: "]", keyword: "】", value: "]" }, { icon: "FullWidth", label: "(", keyword: "（", value: "(" }, { icon: "FullWidth", label: ")", keyword: "）", value: ")" }, { icon: "FullWidth", label: "<", keyword: "《", value: "<" }, { icon: "FullWidth", label: ">", keyword: "》", value: ">" }], iv = [{ icon: "FullWidth", label: "[]", keyword: "【】", value: "[]", goLeft: 1 }, { icon: "FullWidth", label: "【】", keyword: "【", value: "【】", goLeft: 1 }, { icon: "link", label: "Link", keyword: "【】", value: "[title](https://url)", selection: { from: 19, to: 14 } }, { icon: "FullWidth", label: "()", keyword: "（", value: "()", goLeft: 1 }, { icon: "FullWidth", label: "（）", keyword: "（", value: "（）", goLeft: 1 }, { icon: "FullWidth", label: "<>", keyword: "《》", value: "<>", goLeft: 1 }, { icon: "FullWidth", label: "《》", keyword: "《》", value: "《》", goLeft: 1 }, { icon: "FullWidth", label: '""', keyword: "“”", value: '""', goLeft: 1 }, { icon: "FullWidth", label: "“”", keyword: "“”", value: "”“", goLeft: 1 }], ov = b(Al).call(Al, iv), xl = ou, sv = function() {
  return "CodeMirror.Pass";
};
function cv(e, t) {
  var r = Hb !== void 0 && So(e) || e["@@iterator"];
  if (!r) {
    if (cr(e) || (r = function(c, u) {
      var l;
      if (c) {
        if (typeof c == "string")
          return Cl(c, u);
        var f = Ke(l = Object.prototype.toString.call(c)).call(l, 8, -1);
        if (f === "Object" && c.constructor && (f = c.constructor.name), f === "Map" || f === "Set")
          return hb(c);
        if (f === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(f))
          return Cl(c, u);
      }
    }(e)) || t && e && typeof e.length == "number") {
      r && (e = r);
      var n = 0, a = function() {
      };
      return { s: a, n: function() {
        return n >= e.length ? { done: !0 } : { done: !1, value: e[n++] };
      }, e: function(c) {
        throw c;
      }, f: a };
    }
    throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`);
  }
  var i, o = !0, s = !1;
  return { s: function() {
    r = r.call(e);
  }, n: function() {
    var c = r.next();
    return o = c.done, c;
  }, e: function(c) {
    s = !0, i = c;
  }, f: function() {
    try {
      o || r.return == null || r.return();
    } finally {
      if (s)
        throw i;
    }
  } };
}
function Cl(e, t) {
  (t == null || t > e.length) && (t = e.length);
  for (var r = 0, n = new Array(t); r < t; r++)
    n[r] = e[r];
  return n;
}
function lv(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var ed = function(e) {
  U(r, pe);
  var t = lv(r);
  function r(n) {
    var a, i = n.config;
    return j(this, r), (a = t.call(this, { needCache: !0 })).config = i, a.RULE = a.rule(), a.suggesterPanel = new uv(), a;
  }
  return M(r, [{ key: "afterInit", value: function(n) {
    Ye() && (typeof n == "function" && n(), this.initConfig(this.config));
  } }, { key: "initConfig", value: function(n) {
    var a = this, i = n.suggester;
    this.suggester = {};
    var o, s = [], c = cv(rv);
    try {
      var u = function() {
        var l = o.value;
        s.push({ keyword: l, suggestList: function(f, p) {
          var d = f.toLowerCase(), g = function(v, w) {
            var E, S, A, x = b(E = []).call(E, av), D = b(S = []).call(S, ov);
            return q(x).call(x, function(C) {
              C.label = w ? w[C.label] : C.label;
            }), q(D).call(D, function(C) {
              C.label = w ? w[C.label] : C.label;
            }), (v[0] === "/" || v[0] === "、" || Vf("#").call("#", v[0])) && q(x).call(x, function(C) {
              C.keyword = "".concat(v[0], C.keyword);
            }), Xe(A = b(D).call(D, x)).call(A, function(C) {
              var W;
              return Ff(W = C.keyword).call(W, v[0]);
            });
          }(l, this.$locale);
          if (/^\s$/.test(d))
            p(!1);
          else {
            var _ = d.replace(/\s+/g, "").replace(new RegExp("^".concat(l), "g"), "").split("").join(".*?"), m = new RegExp("^.*?".concat(_, ".*?$"), "i"), h = Xe(g).call(g, function(v) {
              return !d || m.test(v.keyword);
            });
            p(h.length !== 0 && h);
          }
        } });
      };
      for (c.s(); !(o = c.n()).done; )
        u();
    } catch (l) {
      c.e(l);
    } finally {
      c.f();
    }
    i = i ? b(i).call(i, s) : s, q(i).call(i, function(l) {
      l.suggestList ? (l.keyword || (l.keyword = "@"), a.suggester[l.keyword] = l) : console.warn("[cherry-suggester]: the suggestList of config is missing.");
    }), this.suggesterPanel.hasEditor() && (this.suggesterPanel.editor = null);
  } }, { key: "makeHtml", value: function(n) {
    var a, i;
    if (!this.RULE.reg)
      return n;
    if (!this.suggesterPanel.hasEditor() && Ye()) {
      var o = this.$engine.$cherry.editor;
      this.suggesterPanel.setEditor(o), this.suggesterPanel.setSuggester(this.suggester), this.suggesterPanel.bindEvent();
    }
    return fe() ? n.replace(this.RULE.reg, je(i = this.toHtml).call(i, this)) : ct(n, this.RULE.reg, je(a = this.toHtml).call(a, this), !0, 1);
  } }, { key: "toHtml", value: function(n, a, i, o) {
    var s, c, u, l, f;
    return o ? ((c = this.suggester[i]) === null || c === void 0 || (u = c.echo) === null || u === void 0 ? void 0 : u.call(this, o)) || b(l = b(f = "".concat(a, '<span class="cherry-suggestion">')).call(f, i)).call(l, o, "</span>") : ((s = this.suggester[i]) === null || s === void 0 ? void 0 : s.echo) === !1 ? "".concat(a) : this.suggester[i] ? o ? a + o : "".concat(a) : a + o;
  } }, { key: "rule", value: function() {
    var n, a, i;
    if (!this.suggester || ue(this.suggester).length <= 0)
      return {};
    var o = ve(n = ue(this.suggester)).call(n, function(s) {
      return nv(s);
    }).join("|");
    return { reg: new RegExp(b(a = b(i = "".concat(fe() ? "((?<!\\\\))[ ]" : "(^|[^\\\\])[ ]", "(")).call(i, o, ")(([^")).call(a, o, "\\s])+)"), "g") };
  } }, { key: "mounted", value: function() {
    if (!this.suggesterPanel.hasEditor() && Ye()) {
      var n = this.$engine.$cherry.editor;
      this.suggesterPanel.setEditor(n), this.suggesterPanel.setSuggester(this.suggester), this.suggesterPanel.bindEvent();
    }
  } }]), r;
}();
R(ed, "HOOK_NAME", "suggester");
var uv = function() {
  function e() {
    j(this, e), R(this, "panelWrap", '<div class="cherry-suggester-panel"></div>'), this.searchCache = !1, this.searchKeyCache = [], this.optionList = [], this.cursorMove = !0, this.suggesterConfig = {};
  }
  return M(e, [{ key: "tryCreatePanel", value: function() {
    var t, r, n;
    !this.$suggesterPanel && Ye() && document && ((t = document) === null || t === void 0 || (r = t.body) === null || r === void 0 || r.appendChild(this.createDom(this.panelWrap)), this.$suggesterPanel = (n = document) === null || n === void 0 ? void 0 : n.querySelector(".cherry-suggester-panel"));
  } }, { key: "hasEditor", value: function() {
    return !!this.editor && !!this.editor.editor.display && !!this.editor.editor.display.wrapper;
  } }, { key: "setEditor", value: function(t) {
    this.editor = t;
  } }, { key: "setSuggester", value: function(t) {
    this.suggesterConfig = t;
  } }, { key: "bindEvent", value: function() {
    var t = this, r = !1;
    this.editor.editor.on("change", function(i, o) {
      r = !0, t.onCodeMirrorChange(i, o);
    }), this.editor.editor.on("keydown", function(i, o) {
      r = !0, t.enableRelate() && t.onKeyDown(i, o);
    }), this.editor.editor.on("cursorActivity", function() {
      r || t.stopRelate(), r = !1;
    });
    var n = this.editor.editor.getOption("extraKeys"), a = ["Up", "Down", "Enter"];
    q(a).call(a, function(i) {
      if (typeof n[i] == "function") {
        var o = n[i];
        n[i] = function(c) {
          if (t.cursorMove) {
            var u = o.call(c, c);
            if (u)
              return u;
          }
        };
      } else if (n[i]) {
        if (typeof n[i] == "string") {
          var s = n[i];
          n[i] = function(c) {
            t.cursorMove && t.editor.editor.execCommand(s);
          };
        }
      } else
        n[i] = function() {
          if (t.cursorMove)
            return sv();
        };
    }), this.editor.editor.setOption("extraKeys", n), this.editor.editor.on("scroll", function(i, o) {
      t.searchCache && t.relocatePanel(t.editor.editor);
    }), this.onClickPanelItem();
  } }, { key: "onClickPanelItem", value: function() {
    var t = this;
    this.tryCreatePanel(), this.$suggesterPanel.addEventListener("click", function(r) {
      var n, a, i, o, s = (n = t.$suggesterPanel, a = r.target, o = -1, q(i = n.childNodes).call(i, function(c, u) {
        return c === a ? o = u : "";
      }), o);
      s > -1 && t.pasteSelectResult(s), t.stopRelate();
    }, !1);
  } }, { key: "showSuggesterPanel", value: function(t) {
    var r = t.left, n = t.top, a = t.items;
    this.tryCreatePanel(), !this.$suggesterPanel && Ye() && (document.body.appendChild(this.createDom(this.panelWrap)), this.$suggesterPanel = document.querySelector(".cherry-suggester-panel")), this.updatePanel(a), this.$suggesterPanel.style.left = "".concat(r, "px"), this.$suggesterPanel.style.top = "".concat(n, "px"), this.$suggesterPanel.style.display = "block", this.$suggesterPanel.style.position = "absolute", this.$suggesterPanel.style.zIndex = "100";
  } }, { key: "hideSuggesterPanel", value: function() {
    this.tryCreatePanel(), this.$suggesterPanel && (this.$suggesterPanel.style.display = "none");
  } }, { key: "updatePanel", value: function(t) {
    var r = this;
    this.tryCreatePanel();
    var n = ve(t).call(t, function(i, o) {
      if (He(i) === "object" && i !== null) {
        var s, c = i.label;
        return i != null && i.icon && (c = b(s = '<i class="ch-icon ch-icon-'.concat(i.icon, '"></i>')).call(s, c)), r.renderPanelItem(c, o === 0);
      }
      return r.renderPanelItem(i, o === 0);
    }).join(""), a = this.suggesterConfig[this.keyword];
    a && typeof a.suggestListRender == "function" && (n = a.suggestListRender.call(this, t) || n), this.$suggesterPanel.innerHTML = "", typeof n == "string" ? this.$suggesterPanel.innerHTML = n : cr(n) && n.length > 0 ? q(n).call(n, function(i) {
      r.$suggesterPanel.appendChild(i);
    }) : He(n) === "object" && n.nodeType === 1 && this.$suggesterPanel.appendChild(n);
  } }, { key: "renderPanelItem", value: function(t, r) {
    return r ? '<div class="cherry-suggester-panel__item cherry-suggester-panel__item--selected">'.concat(t, "</div>") : '<div class="cherry-suggester-panel__item">'.concat(t, "</div>");
  } }, { key: "createDom", value: function() {
    var t = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : "";
    this.template || (this.template = document.createElement("div")), this.template.innerHTML = N(t).call(t);
    var r = document.createDocumentFragment();
    return ve(Array.prototype).call(this.template.childNodes, function(n, a) {
      r.appendChild(n);
    }), r;
  } }, { key: "relocatePanel", value: function(t) {
    var r = document.querySelector(".CodeMirror-cursors .CodeMirror-cursor");
    if (r || (r = document.querySelector(".CodeMirror-selected")), !r)
      return !1;
    var n = t.getCursor(), a = t.lineInfo(n.line).handle.height, i = r.getBoundingClientRect(), o = i.top + a, s = i.left;
    this.showSuggesterPanel({ left: s, top: o, items: this.optionList });
  } }, { key: "getCursorPos", value: function(t) {
    var r = document.querySelector(".CodeMirror-cursors .CodeMirror-cursor");
    if (!r)
      return null;
    var n = t.getCursor(), a = t.lineInfo(n.line).handle.height, i = r.getBoundingClientRect(), o = i.top + a;
    return { left: i.left, top: o };
  } }, { key: "startRelate", value: function(t, r, n) {
    this.cursorFrom = n, this.keyword = r, this.searchCache = !0, this.relocatePanel(t);
  } }, { key: "stopRelate", value: function() {
    this.hideSuggesterPanel(), this.cursorFrom = null, this.cursorTo = null, this.keyword = "", this.searchKeyCache = [], this.searchCache = !1, this.cursorMove = !0, this.optionList = [];
  } }, { key: "pasteSelectResult", value: function(t, r) {
    if (this.cursorTo && this.cursorTo !== this.cursorFrom || (this.cursorTo = JSON.parse(zb(this.cursorFrom))), this.cursorTo) {
      this.cursorTo.ch += 1;
      var n = this.cursorFrom, a = this.cursorTo;
      if (this.optionList[t]) {
        var i = "";
        if (He(this.optionList[t]) === "object" && this.optionList[t] !== null && typeof this.optionList[t].value == "string")
          i = this.optionList[t].value;
        else if (He(this.optionList[t]) === "object" && this.optionList[t] !== null && typeof this.optionList[t].value == "function")
          i = this.optionList[t].value();
        else if (typeof this.optionList[t] == "string")
          i = "".concat(this.optionList[t], " ");
        else {
          var o;
          i = b(o = " ".concat(this.keyword)).call(o, this.optionList[t], " ");
        }
        if (i && this.editor.editor.replaceRange(i, n, a), this.optionList[t].goLeft) {
          var s = this.editor.editor.getCursor();
          this.editor.editor.setCursor(s.line, s.ch - this.optionList[t].goLeft);
        }
        if (this.optionList[t].selection) {
          var c = this.editor.editor.getCursor().line, u = this.editor.editor.getCursor().ch;
          this.editor.editor.setSelection({ line: c, ch: u - this.optionList[t].selection.from }, { line: c, ch: u - this.optionList[t].selection.to });
        }
      }
    }
  } }, { key: "findSelectedItemIndex", value: function() {
    return Mi(Array.prototype).call(this.$suggesterPanel.childNodes, function(t) {
      return t.classList.contains("cherry-suggester-panel__item--selected");
    });
  } }, { key: "enableRelate", value: function() {
    return this.searchCache;
  } }, { key: "onCodeMirrorChange", value: function(t, r) {
    var n = this, a = r.text, i = r.from, o = r.to, s = r.origin, c = a.length === 1 ? a[0] : "";
    if (!this.enableRelate() && this.suggesterConfig[c] && this.startRelate(t, c, i), this.enableRelate() && (c || s === "+delete")) {
      var u;
      if (this.cursorTo = o, c)
        this.searchKeyCache.push(c);
      else if (s === "+delete" && (this.searchKeyCache.pop(), this.searchKeyCache.length === 0))
        return void this.stopRelate();
      typeof ((u = this.suggesterConfig[this.keyword]) === null || u === void 0 ? void 0 : u.suggestList) == "function" && this.suggesterConfig[this.keyword].suggestList(this.searchKeyCache.join(""), function(l) {
        l !== !1 ? (n.optionList = l && l.length ? l : [], n.updatePanel(n.optionList)) : n.stopRelate();
      });
    }
  } }, { key: "onKeyDown", value: function(t, r) {
    var n, a = this;
    if (this.tryCreatePanel(), !this.$suggesterPanel)
      return !1;
    var i = r.keyCode;
    if (Vf(n = [38, 40]).call(n, i)) {
      if (this.optionList.length === 0)
        return void Ei(function() {
          a.stopRelate();
        }, 0);
      this.cursorMove = !1;
      var o = this.$suggesterPanel.querySelector(".cherry-suggester-panel__item--selected"), s = null;
      i !== 38 || o.previousElementSibling ? i !== 40 || o.nextElementSibling ? i === 38 ? s = o.previousElementSibling : i === 40 && (s = o.nextElementSibling) : s = this.$suggesterPanel.firstElementChild : s = this.$suggesterPanel.lastElementChild, o.classList.remove("cherry-suggester-panel__item--selected"), s.classList.add("cherry-suggester-panel__item--selected");
    } else
      i === 13 ? (r.stopPropagation(), this.cursorMove = !1, this.pasteSelectResult(this.findSelectedItemIndex(), r), t.focus(), Ei(function() {
        a.stopRelate();
      }, 0)) : i !== 27 && i !== 37 && i !== 39 || (r.stopPropagation(), t.focus(), Ei(function() {
        a.stopRelate();
      }, 0));
  } }]), e;
}();
function fv(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var td = function(e) {
  U(r, pe);
  var t = fv(r);
  function r() {
    return j(this, r), t.apply(this, arguments);
  }
  return M(r, [{ key: "makeHtml", value: function(n) {
    return this.test(n) ? n.replace(this.RULE.reg, "$1<ruby>$2<rt>$3</rt></ruby>$4") : n;
  } }, { key: "rule", value: function() {
    var n = { begin: "(^| )\\{", end: "\\}( |$)", content: `([^
]+?)\\|([^
]+?)` };
    return n.reg = new RegExp(n.begin + n.content + n.end, "g"), n;
  } }]), r;
}();
function dv(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(td, "HOOK_NAME", "ruby");
var nd = function(e) {
  U(r, se);
  var t = dv(r);
  function r(n) {
    var a;
    return j(this, r), (a = t.call(this, { needCache: !0 })).initBrReg(n.globalConfig.classicBr), a;
  }
  return M(r, [{ key: "makeHtml", value: function(n, a) {
    var i = this;
    return n.replace(this.RULE.reg, function(o, s, c, u) {
      var l, f, p, d, g, _ = i.getLineCount(o, s), m = i.$engine.md5(o), h = i.$getPanelInfo(c, u, a), v = h.title, w = h.body, E = h.appendStyle, S = h.className;
      return nn(o, i.pushCache(b(l = b(f = b(p = b(d = b(g = '<div class="'.concat(S, '" data-sign="')).call(g, m, '" data-lines="')).call(d, _, '" ')).call(p, E, ">")).call(f, v)).call(l, w, "</div>"), m, _));
    });
  } }, { key: "$getClassByType", value: function(n) {
    return /(left|right|center)/i.test(n) ? "cherry-text-align cherry-text-align__".concat(n) : "cherry-panel cherry-panel__".concat(n);
  } }, { key: "$getPanelInfo", value: function(n, a, i) {
    var o, s = this, c = { type: this.$getTargetType(n), title: i(this.$getTitle(n)).html, body: a, appendStyle: "", className: "" };
    c.className = this.$getClassByType(c.type), /(left|right|center)/i.test(c.type) && (c.appendStyle = 'style="text-align:'.concat(c.type, ';"')), c.title = b(o = '<div class="cherry-panel--title '.concat(c.title ? "cherry-panel--title__not-empty" : "", '">')).call(o, c.title, "</div>");
    var u = function(f) {
      var p, d;
      if (N(f).call(f) === "")
        return "";
      var g = i(f).html, _ = "p";
      return new RegExp("<(".concat(vr, ")[^>]*>"), "i").test(g) && (_ = "div"), b(p = b(d = "<".concat(_, ">")).call(d, s.$cleanParagraph(g), "</")).call(p, _, ">");
    }, l = "";
    return l = this.isContainsCache(c.body) ? this.makeExcludingCached(c.body, u) : u(c.body), c.body = '<div class="cherry-panel--body">'.concat(l, "</div>"), c;
  } }, { key: "$getTitle", value: function(n) {
    var a = N(n).call(n);
    return /\s/.test(a) ? a.replace(/[^\s]+\s/, "") : "";
  } }, { key: "$getTargetType", value: function(n) {
    var a = /\s/.test(N(n).call(n)) ? N(n).call(n).replace(/\s.*$/, "") : n;
    switch (N(a).call(a).toLowerCase()) {
      case "primary":
      case "p":
      default:
        return "primary";
      case "info":
      case "i":
        return "info";
      case "warning":
      case "w":
        return "warning";
      case "danger":
      case "d":
        return "danger";
      case "success":
      case "s":
        return "success";
      case "right":
      case "r":
        return "right";
      case "center":
      case "c":
        return "center";
      case "left":
      case "l":
        return "left";
    }
  } }, { key: "rule", value: function() {
    return (n = { begin: /(?:^|\n)(\n*(?:[^\S\n]*)):::([^:][^\n]+?)\s*\n/, content: /([\w\W]*?)/, end: /\n[ \t]*:::[ \t]*(?=$|\n+)/ }).reg = new RegExp(n.begin.source + n.content.source + n.end.source, "g"), n;
    var n;
  } }]), r;
}();
function pv(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
R(nd, "HOOK_NAME", "panel");
var rd = function(e) {
  U(r, se);
  var t = pv(r);
  function r() {
    return j(this, r), t.call(this, { needCache: !0 });
  }
  return M(r, [{ key: "makeHtml", value: function(n, a) {
    var i = this;
    return n.replace(this.RULE.reg, function(o, s, c, u, l) {
      var f, p, d, g = i.getLineCount(o, s), _ = i.$engine.md5(o), m = i.$getDetailInfo(c, u, l, a), h = m.type, v = m.html;
      return nn(o, i.pushCache(b(f = b(p = b(d = '<div class="cherry-detail cherry-detail__'.concat(h, '" data-sign="')).call(d, _, '" data-lines="')).call(p, g, '" >')).call(f, v, "</div>"), _, g));
    });
  } }, { key: "$getDetailInfo", value: function(n, a, i, o) {
    var s = this, c = /\n\s*(\+\+|\+\+-)\s*[^\n]+\n/.test(i) ? "multiple" : "single", u = i.split(/\n\s*(\+\+[-]{0,1}\s*[^\n]+)\n/), l = n === "-", f = a, p = "";
    return c === "multiple" ? q(u).call(u, function(d) {
      if (/\+\+/.test(d))
        return l = /\+\+-/.test(d), f = d.replace(/\+\+[-]{0,1}\s*([^\n]+)$/, "$1"), !0;
      p += s.$getDetailHtml(l, f, d, o);
    }) : p = this.$getDetailHtml(l, f, i, o), { type: c, html: p };
  } }, { key: "$getDetailHtml", value: function(n, a, i, o) {
    var s = this, c = "<details ".concat(n ? "open" : "", ">"), u = function(f) {
      var p, d;
      if (N(f).call(f) === "")
        return "";
      var g = o(f).html, _ = "p";
      return new RegExp("<(".concat(vr, ")[^>]*>"), "i").test(g) && (_ = "div"), b(p = b(d = "<".concat(_, ">")).call(d, s.$cleanParagraph(g), "</")).call(p, _, ">");
    };
    c += "<summary>".concat(o(a).html, "</summary>");
    var l = "";
    return l = this.isContainsCache(i) ? this.makeExcludingCached(i, u) : u(i), c += '<div class="cherry-detail-body">'.concat(l, "</div>"), c += "</details>";
  } }, { key: "rule", value: function() {
    return (n = { begin: /(?:^|\n)(\n*(?:[^\S\n]*))\+\+\+([-]{0,1})\s+([^\n]+)\n/, content: /([\w\W]+?)/, end: /\n[ \t]*\+\+\+[ \t]*(?=$|\n+)/ }).reg = new RegExp(n.begin.source + n.content.source + n.end.source, "g"), n;
    var n;
  } }]), r;
}();
R(rd, "HOOK_NAME", "detail");
var Tl = [Cn, wf, Hf, zf, Kf, qf, Gf, Tf, Rf, $f, If, Wf, Cf, Of, Lf, rd, nd, xf, Zf, Pf, Sf, Nf, Af, bf, mf, vf, kf, _f, td, yf, Yf, Xf, ed], gv = { run: function(e) {
  var t, r = "<div>".concat(e, "</div>");
  this.tagParser.formatEngine = this.mdFormatEngine, r = r.replace(/<!--[\s\S]*?-->/g, "");
  var n = this.htmlParser.parseHtml(r);
  return n = this.paragraphStyleClear(n), N(t = this.$dealHtml(n).replace(/\n{3,}/g, `


`).replace(/&gt;/g, ">").replace(/&lt;/g, "<").replace(/&amp;/g, "&")).call(t, `
`);
}, $dealHtml: function(e) {
  for (var t = "", r = 0; r < e.length; r++) {
    var n = e[r];
    n.type === "tag" ? t = this.$handleTagObject(n, t) : n.type === "text" && n.content.length > 0 && (t += n.content.replace(/&nbsp;/g, " ").replace(/[\n]+/g, `
`).replace(/^[ \t\n]+\n\s*$/, `
`));
  }
  return t;
}, $handleTagObject: function(e, t) {
  var r, n = t;
  return e.attrs.class && /(ch-icon-square|ch-icon-check)/.test(e.attrs.class) ? Me(r = e.attrs.class).call(r, "ch-icon-check") >= 0 ? n += "[x]" : n += "[ ]" : e.attrs.class && /cherry-code-preview-lang-select/.test(e.attrs.class) ? n += "" : n += this.$dealTag(e), n;
}, $dealTag: function(e) {
  var t = this, r = "";
  return e.children && (r = t.$dealHtml(e.children)), /(style|meta|link|script)/.test(e.name) ? "" : e.name === "code" || e.name === "pre" ? t.tagParser.codeParser(e, t.$dealCodeTag(e), e.name === "pre") : typeof t.tagParser["".concat(e.name, "Parser")] == "function" ? t.tagParser["".concat(e.name, "Parser")](e, r) : r;
}, $dealCodeTag: function(e) {
  if (e.children.length < 0)
    return "";
  for (var t = "", r = 0; r < e.children.length; r++) {
    var n = e.children[r];
    n.type !== "text" ? (n.name === "li" && (t += `
`), n.name === "br" && (t += `
`), t += this.$dealCodeTag(n)) : t += n.content;
  }
  return t;
}, htmlParser: { attrRE: /([\w-]+)|['"]{1}([^'"]*)['"]{1}/g, lookup: { area: !0, base: !0, br: !0, col: !0, embed: !0, hr: !0, img: !0, video: !0, input: !0, keygen: !0, link: !0, menuitem: !0, meta: !0, param: !0, source: !0, track: !0, wbr: !0 }, tagRE: /<(?:"[^"]*"['"]*|'[^']*'['"]*|[^'">])+>/g, empty: xl ? xl(null) : {}, parseTags: function(e) {
  var t, r = this, n = 0, a = { type: "tag", name: "", voidElement: !1, attrs: {}, children: [] };
  return e.replace(this.attrRE, function(i) {
    n % 2 ? t = i : n === 0 ? ((r.lookup[i] || e.charAt(e.length - 2) === "/") && (a.voidElement = !0), a.name = i) : a.attrs[t] = i.replace(/['"]/g, ""), n += 1;
  }), a;
}, parseHtml: function(e, t) {
  var r = this, n = t || {};
  n.components || (n.components = this.empty);
  var a, i = [], o = -1, s = [], c = {}, u = !1;
  return e.replace(this.tagRE, function(l, f) {
    if (u) {
      if (l !== "</".concat(a.name, ">"))
        return;
      u = !1;
    }
    var p, d = l.charAt(1) !== "/", g = f + l.length, _ = e.charAt(g);
    d && (o += 1, (a = r.parseTags(l)).type === "tag" && n.components[a.name] && (a.type = "component", u = !0), a.voidElement || u || !_ || _ === "<" || a.children.push({ type: "text", content: Ke(e).call(e, g, Me(e).call(e, "<", g)) }), c[a.tagName] = a, o === 0 && i.push(a), (p = s[o - 1]) && p.children.push(a), s[o] = a), d && !a.voidElement || (o -= 1, !u && _ !== "<" && _ && s[o] && s[o].children.push({ type: "text", content: Ke(e).call(e, g, Me(e).call(e, "<", g)) }));
  }), i;
} }, tagParser: { formatEngine: {}, pParser: function(e, t) {
  var r = t;
  return /\n$/.test(r) ? r : "".concat(r, `
`);
}, divParser: function(e, t) {
  var r = t;
  return /\n$/.test(r) ? r : "".concat(r, `
`);
}, spanParser: function(e, t) {
  var r = t.replace(/\t/g, "").replace(/\n/g, " ");
  return e.attrs && e.attrs.style, r;
}, codeParser: function(e, t) {
  var r = arguments.length > 2 && arguments[2] !== void 0 && arguments[2];
  return this.formatEngine.convertCode(t, r);
}, brParser: function(e, t) {
  return this.formatEngine.convertBr(t, `
`);
}, imgParser: function(e, t) {
  return e.attrs && e.attrs["data-control"] === "tapd-graph" ? this.formatEngine.convertGraph(e.attrs.title, e.attrs.src, e.attrs["data-origin-xml"], e) : e.attrs && e.attrs.src ? this.formatEngine.convertImg(e.attrs.alt, e.attrs.src) : void 0;
}, videoParser: function(e, t) {
  if (e.attrs && e.attrs.src)
    return this.formatEngine.convertVideo(t, e.attrs.src, e.attrs.poster, e.attrs.title);
}, bParser: function(e, t) {
  for (var r = t.split(`
`), n = [], a = 0; a < r.length; a++)
    n.push(this.formatEngine.convertB(r[a]));
  return n.join(`
`);
}, iParser: function(e, t) {
  for (var r = t.split(`
`), n = [], a = 0; a < r.length; a++)
    n.push(this.formatEngine.convertI(r[a]));
  return n.join(`
`);
}, strikeParser: function(e, t) {
  for (var r = t.split(`
`), n = [], a = 0; a < r.length; a++)
    n.push(this.formatEngine.convertStrike(r[a]));
  return n.join(`
`);
}, delParser: function(e, t) {
  for (var r = t.split(`
`), n = [], a = 0; a < r.length; a++)
    n.push(this.formatEngine.convertDel(r[a]));
  return n.join(`
`);
}, uParser: function(e, t) {
  for (var r = t.split(`
`), n = [], a = 0; a < r.length; a++)
    n.push(this.formatEngine.convertU(r[a]));
  return n.join(`
`);
}, aParser: function(e, t) {
  return e.attrs && e.attrs.href ? this.formatEngine.convertA(t, e.attrs.href) : "";
}, supParser: function(e, t) {
  return this.formatEngine.convertSup(t);
}, subParser: function(e, t) {
  return this.formatEngine.convertSub(t);
}, tdParser: function(e, t) {
  return this.formatEngine.convertTd(t);
}, trParser: function(e, t) {
  return this.formatEngine.convertTr(t);
}, thParser: function(e, t) {
  return this.formatEngine.convertTh(t);
}, theadParser: function(e, t) {
  return this.formatEngine.convertThead(t);
}, tableParser: function(e, t) {
  return this.formatEngine.convertTable(t);
}, liParser: function(e, t) {
  return this.formatEngine.convertLi(t);
}, ulParser: function(e, t) {
  return this.formatEngine.convertUl(t);
}, olParser: function(e, t) {
  return this.formatEngine.convertOl(t);
}, strongParser: function(e, t) {
  return this.formatEngine.convertStrong(t);
}, hrParser: function(e, t) {
  return this.formatEngine.convertHr(t);
}, h1Parser: function(e, t) {
  return this.formatEngine.convertH1(t);
}, h2Parser: function(e, t) {
  return this.formatEngine.convertH2(t);
}, h3Parser: function(e, t) {
  return this.formatEngine.convertH3(t);
}, h4Parser: function(e, t) {
  return this.formatEngine.convertH4(t);
}, h5Parser: function(e, t) {
  return this.formatEngine.convertH5(t);
}, h6Parser: function(e, t) {
  return this.formatEngine.convertH6(t);
}, blockquoteParser: function(e, t) {
  return this.formatEngine.convertBlockquote(t.replace(/\n+/g, `
`));
}, addressParser: function(e, t) {
  return this.formatEngine.convertAddress(t.replace(/\n+/g, `
`));
}, styleParser: { colorAttrParser: function(e) {
  var t = e.match(/color:\s*(#[a-zA-Z0-9]{3,6});/);
  return t && t[1] ? t[1] : "";
}, sizeAttrParser: function(e) {
  var t = e.match(/font-size:\s*([a-zA-Z0-9-]+?);/);
  if (t && t[1]) {
    var r, n = 0;
    if (/[0-9]+px/.test(t[1]))
      n = N(r = t[1].replace(/px/, "")).call(r);
    else
      switch (t[1]) {
        case "x-small":
          n = 10;
          break;
        case "small":
          n = 12;
          break;
        case "medium":
          n = 16;
          break;
        case "large":
          n = 18;
          break;
        case "x-large":
          n = 24;
          break;
        case "xx-large":
          n = 32;
          break;
        default:
          n = "";
      }
    return n > 0 ? n : "";
  }
  return "";
}, bgColorAttrParser: function(e) {
  var t = e.match(/background-color:\s*([^;]+?);/);
  if (t && t[1]) {
    var r = "";
    if (/rgb\([ 0-9]+,[ 0-9]+,[ 0-9]+\)/.test(t[1])) {
      var n, a, i, o, s, c = t[1].match(/rgb\(([ 0-9]+),([ 0-9]+),([ 0-9]+)\)/);
      c[1] && c[2] && c[3] && (c[1] = xn(N(n = c[1]).call(n), 10), c[2] = xn(N(a = c[2]).call(a), 10), c[3] = xn(N(i = c[3]).call(i), 10), r = b(o = b(s = "#".concat(c[1].toString(16))).call(s, c[2].toString(16))).call(o, c[3].toString(16)));
    } else
      r = Ht(t, 2)[1];
    return r;
  }
  return "";
} } }, mdFormatEngine: { convertColor: function(e, t) {
  var r, n = N(e).call(e);
  return !n || /\n/.test(n) ? n : t ? b(r = "!!".concat(t, " ")).call(r, n, "!!") : n;
}, convertSize: function(e, t) {
  var r, n = N(e).call(e);
  return !n || /\n/.test(n) ? n : t ? b(r = "!".concat(t, " ")).call(r, n, "!") : n;
}, convertBgColor: function(e, t) {
  var r, n = N(e).call(e);
  return !n || /\n/.test(n) ? n : t ? b(r = "!!!".concat(t, " ")).call(r, n, "!!!") : n;
}, convertBr: function(e, t) {
  return e + t;
}, convertCode: function(e) {
  var t = arguments.length > 1 && arguments[1] !== void 0 && arguments[1];
  return /\n/.test(e) || t ? "```\n".concat(e.replace(/\n+$/, ""), "\n```") : "`".concat(e.replace(/`/g, "\\`"), "`");
}, convertB: function(e) {
  return /^\s*$/.test(e) ? "" : "**".concat(e, "**");
}, convertI: function(e) {
  return /^\s*$/.test(e) ? "" : "*".concat(e, "*");
}, convertU: function(e) {
  return /^\s*$/.test(e) ? "" : " /".concat(e, "/ ");
}, convertImg: function(e, t) {
  var r, n = e && e.length > 0 ? e : "image";
  return b(r = "![".concat(n, "](")).call(r, t, ")");
}, convertGraph: function(e, t, r, n) {
  var a, i, o, s = e && e.length > 0 ? e : "graph", c = "";
  if (n)
    try {
      var u, l = n.attrs;
      q(u = ue(l)).call(u, function(f) {
        var p;
        Object.prototype.hasOwnProperty.call(l, f) && Me(f).call(f, "data-graph-") >= 0 && l[f] && (c += b(p = " ".concat(f, "=")).call(p, l[f]));
      });
    } catch {
    }
  return b(a = b(i = b(o = "![".concat(s, "](")).call(o, t, "){data-control=tapd-graph data-origin-xml=")).call(i, r)).call(a, c, "}");
}, convertVideo: function(e, t, r, n) {
  var a, i, o = n && n.length > 0 ? n : "video";
  return b(a = b(i = "!video[".concat(o, "](")).call(i, t, "){poster=")).call(a, r, "}");
}, convertA: function(e, t) {
  var r;
  if (e === t)
    return "".concat(e, " ");
  var n = N(e).call(e);
  return n && b(r = "[".concat(n, "](")).call(r, t, ")");
}, convertSup: function(e) {
  return "^".concat(N(e).call(e).replace(/\^/g, "\\^"), "^");
}, convertSub: function(e) {
  return "^^".concat(N(e).call(e).replace(/\^\^/g, "\\^\\^"), "^^");
}, convertTd: function(e) {
  return "~|".concat(N(e).call(e).replace(/\n{1,}/g, "<br>").replace(/ /g, "~s~"), " ~|");
}, convertTh: function(e) {
  return /^\s*$/.test(e) ? "" : "~|".concat(N(e).call(e).replace(/\n{1,}/g, "<br>"), " ~|");
}, convertTr: function(e) {
  return /^\s*$/.test(e) ? "" : "".concat(N(e).call(e).replace(/\n/g, ""), `
`);
}, convertThead: function(e) {
  var t, r = "".concat(e.replace(/[ \t]+/g, "").replace(/~\|~\|/g, "~|").replace(/~\|/g, "|"), `
`), n = r.match(/\|/g).length - 1;
  return b(t = "".concat(r, "|")).call(t, gt(":-:|").call(":-:|", n), `
`);
}, convertTable: function(e) {
  var t = `
`.concat(e.replace(/[ \t]+/g, "").replace(/~\|~\|/g, "~|").replace(/~\|/g, "|"), `
`).replace(/\n{2,}/g, `
`).replace(/\n[ \t]+\n/g, `
`).replace(/~s~/g, " ");
  if (!/\|:-:\|/.test(t)) {
    var r, n, a = t.match(/^\n[^\n]+\n/)[0].match(/\|/g).length - 1;
    t = b(r = b(n = `
|`.concat(gt(" |").call(" |", a), `
|`)).call(n, gt(":-:|").call(":-:|", a))).call(r, t);
  }
  return t;
}, convertLi: function(e) {
  return "- ".concat(e.replace(/^\n/, "").replace(/\n+$/, "").replace(/\n+/g, `
	`), `
`);
}, convertUl: function(e) {
  return "".concat(e, `
`);
}, convertOl: function(e) {
  for (var t = e.split(`
`), r = 1, n = 0; n < t.length; n++)
    /^- /.test(t[n]) && (t[n] = t[n].replace(/^- /, "".concat(r, ". ")), r += 1);
  var a = t.join(`
`);
  return "".concat(a, `
`);
}, convertStrong: function(e) {
  return /^\s*$/.test(e) ? "" : "**".concat(e, "**");
}, convertStrike: function(e) {
  return /^\s*$/.test(e) ? "" : "~~".concat(e, "~~");
}, convertDel: function(e) {
  return /^\s*$/.test(e) ? "" : "~~".concat(e, "~~");
}, convertHr: function(e) {
  return /^\s*$/.test(e) ? `

----
` : `

----
`.concat(e);
}, convertH1: function(e) {
  return "# ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertH2: function(e) {
  return "## ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertH3: function(e) {
  return "### ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertH4: function(e) {
  return "#### ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertH5: function(e) {
  return "##### ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertH6: function(e) {
  return "###### ".concat(N(e).call(e).replace(/\n+$/, ""), `

`);
}, convertBlockquote: function(e) {
  return ">".concat(N(e).call(e), `

`);
}, convertAddress: function(e) {
  return ">".concat(N(e).call(e), `

`);
} }, paragraphStyleClear: function(e) {
  for (var t = 0; t < e[0].children.length; t++) {
    for (var r = [e[0].children[t]], n = []; r.length; ) {
      var a = r.shift(), i = this.notEmptyTagCount(a);
      if (i === 1)
        n.push(a);
      else if (i > 1)
        for (var o = 0; o < a.children.length; o++)
          r.push(a.children[o]);
      else
        n.length === 1 && this.clearChildColorAttrs(n.pop()), n = [];
    }
    n.length === 1 && this.clearChildColorAttrs(n.pop());
  }
  return e;
}, notEmptyTagCount: function(e) {
  if (!e || e.voidElement || e.type === "tag" && !e.children.length || e.type === "text" && !e.content.replace(/(\r|\n|\s)+/g, ""))
    return 0;
  if (e.children && e.children.length) {
    for (var t = 0, r = 0; r < e.children.length; r++)
      t += this.notEmptyTagCount(e.children[r]);
    return t;
  }
  return 1;
}, clearChildColorAttrs: function(e) {
  var t = this;
  this.forEachHtmlParsedItems(e, function(r) {
    t.clearSelfNodeColorAttrs(r);
  });
}, clearSelfNodeColorAttrs: function(e) {
  if (e.attrs && e.attrs.style) {
    for (var t = e.attrs.style.split(";"), r = [], n = 0; n < t.length; n++) {
      var a;
      t[n] && Me(a = t[n]).call(a, "color") === -1 && r.push(t[n]);
    }
    r.length ? e.attrs.style = "".concat(r.join(";"), ";") : delete e.attrs.style;
  }
}, forEachHtmlParsedItems: function(e, t) {
  if (e && (t(e), e.children && e.children.length))
    for (var r = 0; r < e.children.length; r++)
      this.forEachHtmlParsedItems(e.children[r], t);
} }, hv = gv, mv = function() {
  function e(t, r) {
    j(this, e), this.$cherry = r, nt(this, "_cherry", { get: function() {
      return rn.warn("`_engine._cherry` is deprecated. Use `$engine.$cherry` instead."), this.$cherry;
    } }), this.initMath(t), this.$configInit(t), this.hookCenter = new _m(Tl, t, r), this.hooks = this.hookCenter.getHookList(), this.md5Cache = {}, this.md5StrMap = {}, this.markdownParams = t, this.currentStrMd5 = [], this.htmlWhiteListAppend = t.engine.global.htmlWhiteList;
  }
  return M(e, [{ key: "initMath", value: function(t) {
    var r = t.externals, n = t.engine.syntax, a = n.mathBlock.plugins;
    if (Ye() && (n.mathBlock.src || n.inlineMath.src) && !r.MathJax && !window.MathJax) {
      (function(o) {
        if (Ye()) {
          var s = o ? ["input/asciimath", "[tex]/noerrors", "[tex]/cancel", "[tex]/color", "[tex]/boldsymbol"] : [];
          window.MathJax = { startup: { elements: [".Cherry-Math", ".Cherry-InlineMath"], typeset: !0 }, tex: { inlineMath: [["$", "$"], ["\\(", "\\)"]], displayMath: [["$$", "$$"], ["\\[", "\\]"]], tags: "ams", packages: { "[+]": ["noerrors", "cancel", "color"] }, macros: { bm: ["{\\boldsymbol{#1}}", 1] } }, options: { skipHtmlTags: ["script", "noscript", "style", "textarea", "pre", "code", "a"], ignoreHtmlClass: "tex2jax_ignore", processHtmlClass: "tex2jax_process", enableMenu: !1 }, loader: { load: s } };
        }
      })(a);
      var i = document.createElement("script");
      i.src = n.mathBlock.src ? n.mathBlock.src : n.inlineMath.src, i.async = !0, i.src && document.head.appendChild(i);
    }
  } }, { key: "$configInit", value: function(t) {
    if (t.hooksConfig && ef(t.hooksConfig.hooksList, Array))
      for (var r = 0; r < t.hooksConfig.hooksList.length; r++) {
        var n = t.hooksConfig.hooksList[r];
        try {
          n.getType() === "sentence" && jc(n, pe), n.getType() === "paragraph" && jc(n, se), vm(n), Tl.push(n);
        } catch {
          throw new Error("the hook does not correctly inherit");
        }
      }
  } }, { key: "$beforeMakeHtml", value: function(t) {
    var r = t.replace(/~/g, "~T");
    return (r = (r = (r = r.replace(/\$/g, "~D")).replace(/\r\n/g, `
`)).replace(/\r/g, `
`))[r.length - 1] !== `
` && (r += `
`), r = this.$fireHookAction(r, "sentence", "beforeMakeHtml"), r = this.$fireHookAction(r, "paragraph", "beforeMakeHtml");
  } }, { key: "$afterMakeHtml", value: function(t) {
    var r = this.$fireHookAction(t, "paragraph", "afterMakeHtml");
    return r = (r = (r = (r = r.replace(/~D/g, "$")).replace(/~T/g, "~")).replace(/\\<\//g, "\\ </")).replace(new RegExp("\\\\(".concat(rf, ")"), "g"), function(n, a) {
      return a === "&" ? n : Mn(a);
    }).replace(/\\&(?!(amp|lt|gt|quot|apos);)/, function() {
      return "&amp;";
    }), r = (r = r.replace(/\\ <\//g, "\\</")).replace(/id="safe_(?=.*?")/g, 'id="'), r = St.restoreAll(r);
  } }, { key: "$dealSentenceByCache", value: function(t) {
    var r = this;
    return this.$checkCache(t, function(n) {
      return r.$dealSentence(n);
    });
  } }, { key: "$dealSentence", value: function(t) {
    var r;
    return this.$fireHookAction(t, "sentence", "makeHtml", je(r = this.$dealSentenceByCache).call(r, this));
  } }, { key: "$fireHookAction", value: function(t, r, n, a) {
    var i = this, o = t, s = n === "afterMakeHtml" ? "reduceRight" : "reduce";
    if (!this.hooks && !this.hooks[r] && !this.hooks[r][s])
      return o;
    try {
      o = this.hooks[r][s](function(c, u) {
        return u.$engine || (u.$engine = i, nt(u, "_engine", { get: function() {
          return rn.warn("`this._engine` is deprecated. Use `this.$engine` instead."), this.$engine;
        } })), u[n] ? u[n](c, a, i.markdownParams) : c;
      }, o);
    } catch (c) {
      throw new ym(c);
    }
    return o;
  } }, { key: "md5", value: function(t) {
    return this.md5StrMap[t] || (this.md5StrMap[t] = Ef(t)), this.md5StrMap[t];
  } }, { key: "$checkCache", value: function(t, r) {
    var n = this.md5(t);
    return this.md5Cache[n] === void 0 && (this.md5Cache[n] = r(t)), { sign: n, html: this.md5Cache[n] };
  } }, { key: "$dealParagraph", value: function(t) {
    var r;
    return this.$fireHookAction(t, "paragraph", "makeHtml", je(r = this.$dealSentenceByCache).call(r, this));
  } }, { key: "makeHtml", value: function(t) {
    var r = this.$beforeMakeHtml(t);
    return r = this.$dealParagraph(r), r = this.$afterMakeHtml(r);
  } }, { key: "mounted", value: function() {
    this.$fireHookAction("", "sentence", "mounted"), this.$fireHookAction("", "paragraph", "mounted");
  } }, { key: "makeMarkdown", value: function(t) {
    return hv.run(t);
  } }]), e;
}();
function bv(e) {
  var t = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "absolute", r = e.getBoundingClientRect();
  return t === "fixed" ? r : t === "sidebar" ? { left: ad.getTargetParentByButton(e).offsetLeft - 130 + r.width, top: e.offsetTop + r.height / 2, width: r.width, height: r.height } : { left: e.offsetLeft, top: e.offsetTop, width: r.width, height: r.height };
}
var ad = function() {
  function e(t) {
    j(this, e), R(this, "_onClick", void 0), this.$cherry = t, this.bubbleMenu = !1, this.subMenu = null, this.name = "", this.editor = t.editor, this.locale = t.locale, this.dom = null, this.updateMarkdown = !0, this.subMenuConfig = [], this.noIcon = !1, this.cacheOnce = !1, this.positionModel = "absolute", typeof this._onClick == "function" && (rn.warn("`MenuBase._onClick` is deprecated. Override `fire` instead"), this.fire = this._onClick);
  }
  return M(e, [{ key: "getSubMenuConfig", value: function() {
    return this.subMenuConfig;
  } }, { key: "setName", value: function(t, r) {
    this.name = t, this.iconName = r;
  } }, { key: "setCacheOnce", value: function(t) {
    this.cacheOnce = t;
  } }, { key: "getAndCleanCacheOnce", value: function() {
    this.updateMarkdown = !0;
    var t = this.cacheOnce;
    return this.cacheOnce = !1, t;
  } }, { key: "hasCacheOnce", value: function() {
    return this.cacheOnce !== !1;
  } }, { key: "createBtn", value: function() {
    var t = arguments.length > 0 && arguments[0] !== void 0 && arguments[0], r = Br("span", t ? "cherry-dropdown-item" : "cherry-toolbar-button cherry-toolbar-".concat(this.iconName ? this.iconName : this.name), { title: this.locale[this.name] || qe(this.name) });
    if (this.iconName && !this.noIcon) {
      var n = Br("i", "ch-icon ch-icon-".concat(this.iconName));
      r.appendChild(n);
    }
    return (t || this.noIcon) && (r.innerHTML += this.locale[this.name] || qe(this.name)), t || this.dom || (this.dom = r), r;
  } }, { key: "createSubBtnByConfig", value: function(t) {
    var r = t.name, n = t.iconName, a = t.onclick, i = Br("span", "cherry-dropdown-item", { title: this.locale[r] || qe(r) });
    if (n) {
      var o = Br("i", "ch-icon ch-icon-".concat(n));
      i.appendChild(o);
    }
    return i.innerHTML += this.locale[r] || qe(r), i.addEventListener("click", a, !1), i;
  } }, { key: "fire", value: function(t) {
    var r = this, n = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "";
    if (t == null || t.stopPropagation(), typeof this.onClick == "function") {
      var a = this.editor.editor.getSelections();
      this.isSelections = a.length > 1;
      var i = ve(a).call(a, function(o, s, c) {
        return r.onClick(o, n, t) || c[s];
      });
      !this.bubbleMenu && this.updateMarkdown && (this.editor.editor.replaceSelections(i, "around"), this.editor.editor.focus(), this.$afterClick());
    }
  } }, { key: "$getSelectionRange", value: function() {
    var t = this.editor.editor.listSelections()[0], r = t.anchor, n = t.head;
    return r.line === n.line && r.ch > n.ch || r.line > n.line ? { begin: n, end: r } : { begin: r, end: n };
  } }, { key: "registerAfterClickCb", value: function(t) {
    this.afterClickCb = t;
  } }, { key: "$afterClick", value: function() {
    typeof this.afterClickCb != "function" || this.isSelections || (this.afterClickCb(), this.afterClickCb = null);
  } }, { key: "setLessSelection", value: function(t, r) {
    var n, a, i, o, s = this.editor.editor, c = this.$getSelectionRange(), u = c.begin, l = c.end, f = { line: ((n = t.match(/\n/g)) === null || n === void 0 ? void 0 : n.length) > 0 ? u.line + t.match(/\n/g).length : u.line, ch: ((a = t.match(/\n/g)) === null || a === void 0 ? void 0 : a.length) > 0 ? t.replace(/^[\s\S]*?\n([^\n]*)$/, "$1").length : u.ch + t.length }, p = ((i = r.match(/\n/g)) === null || i === void 0 ? void 0 : i.length) > 0 ? l.line - r.match(/\n/g).length : l.line, d = { line: p, ch: ((o = r.match(/\n/g)) === null || o === void 0 ? void 0 : o.length) > 0 ? s.getLine(p).length : l.ch - r.length };
    s.setSelection(f, d);
  } }, { key: "getMoreSelection", value: function(t, r, n) {
    var a = this.editor.editor, i = this.$getSelectionRange(), o = i.begin, s = i.end, c = /\n/.test(t) ? 0 : o.ch - t.length;
    c = c < 0 ? 0 : c;
    var u, l = /\n/.test(t) ? o.line - t.match(/\n/g).length : o.line, f = { line: l = l < 0 ? 0 : l, ch: c }, p = s.line, d = s.ch;
    /\n/.test(r) ? (p = s.line + r.match(/\n/g).length, d = (u = a.getLine(p)) === null || u === void 0 ? void 0 : u.length) : d = a.getLine(s.line).length < s.ch + r.length ? a.getLine(s.line).length : s.ch + r.length;
    var g = { line: p, ch: d };
    a.setSelection(f, g), n() === !1 && a.setSelection(o, s);
  } }, { key: "getSelection", value: function(t) {
    var r = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : "word", n = arguments.length > 2 && arguments[2] !== void 0 && arguments[2], a = this.editor.editor;
    if (this.isSelections || t && !n)
      return t;
    if (r === "line") {
      var i = this.$getSelectionRange(), o = i.begin, s = i.end;
      return a.setSelection({ line: o.line, ch: 0 }, { line: s.line, ch: a.getLine(s.line).length }), a.getSelection();
    }
    if (r === "word") {
      var c = a.findWordAt(a.getCursor()), u = c.anchor, l = c.head;
      return a.setSelection(u, l), a.getSelection();
    }
  } }, { key: "bindSubClick", value: function(t, r) {
    return this.fire(null, t);
  } }, { key: "onClick", value: function(t, r, n) {
    return t;
  } }, { key: "shortcutKeys", get: function() {
    return [];
  } }, { key: "getMenuPosition", value: function() {
    var t = e.getTargetParentByButton(this.dom), r = /cherry-sidebar/.test(t.className);
    return /cherry-bubble/.test(t.className) || /cherry-floatmenu/.test(t.className) ? this.positionModel = "fixed" : this.positionModel = r ? "sidebar" : "absolute", bv(this.dom, this.positionModel);
  } }], [{ key: "getTargetParentByButton", value: function(t) {
    var r = t.parentElement;
    return /toolbar-(left|right)/.test(r.className) && (r = r.parentElement), r;
  } }]), e;
}();
function $l(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
function Rl(e, t, r) {
  var n, a = {};
  return q(n = ue(e)).call(n, function(i) {
    Me(t).call(t, i) !== -1 && (He(r) === "object" ? typeof r[i] == "string" ? He(e[i]) === r[i] && (a[i] = e[i]) : e[i] instanceof r[i] && (a[i] = e[i]) : typeof r == "string" && He(e[i]) === r && (a[i] = e[i]));
  }), a;
}
var vv = { HOOKS_TYPE_LIST: In }, Ol = [];
Ye() || q(Ol).call(Ol, function(e) {
});
var er = function() {
  function e() {
    j(this, e);
  }
  return M(e, null, [{ key: "usePlugin", value: function(t) {
    var r;
    if (this === e)
      throw new Error("`usePlugin` is not allowed to called through CherryStatic class.");
    if (this.initialized)
      throw new Error("The function `usePlugin` should be called before Cherry is instantiated.");
    if (t.$cherry$mounted !== !0) {
      for (var n = arguments.length, a = new Array(n > 1 ? n - 1 : 0), i = 1; i < n; i++)
        a[i - 1] = arguments[i];
      t.install.apply(t, b(r = [this.config.defaults]).call(r, a)), t.$cherry$mounted = !0;
    }
  } }]), e;
}();
R(er, "createSyntaxHook", function(e, t, r) {
  var n, a = t === In.PAR ? se : pe, i = Rl(r, ["beforeMakeHtml", "makeHtml", "afterMakeHtml", "rule", "test"], "function"), o = { needCache: r.needCache, defaultCache: r.defaultCache };
  return n = function(s) {
    U(u, s);
    var c = $l(u);
    function u() {
      var l, f = arguments.length > 0 && arguments[0] !== void 0 ? arguments[0] : {};
      return j(this, u), (l = t === In.PAR ? c.call(this, { needCache: !!o.needCache, defaultCache: o.defaultCache }) : c.call(this)).config = f.config, B(l);
    }
    return M(u, [{ key: "beforeMakeHtml", value: function() {
      for (var l, f, p = arguments.length, d = new Array(p), g = 0; g < p; g++)
        d[g] = arguments[g];
      return i.beforeMakeHtml ? i.beforeMakeHtml.apply(this, d) : (l = dt(T(u.prototype), "beforeMakeHtml", this)).call.apply(l, b(f = [this]).call(f, d));
    } }, { key: "makeHtml", value: function() {
      for (var l, f, p = arguments.length, d = new Array(p), g = 0; g < p; g++)
        d[g] = arguments[g];
      return i.makeHtml ? i.makeHtml.apply(this, d) : (l = dt(T(u.prototype), "makeHtml", this)).call.apply(l, b(f = [this]).call(f, d));
    } }, { key: "afterMakeHtml", value: function() {
      for (var l, f, p = arguments.length, d = new Array(p), g = 0; g < p; g++)
        d[g] = arguments[g];
      return i.afterMakeHtml ? i.afterMakeHtml.apply(this, d) : (l = dt(T(u.prototype), "afterMakeHtml", this)).call.apply(l, b(f = [this]).call(f, d));
    } }, { key: "test", value: function() {
      for (var l, f, p = arguments.length, d = new Array(p), g = 0; g < p; g++)
        d[g] = arguments[g];
      return i.test ? i.test.apply(this, d) : (l = dt(T(u.prototype), "test", this)).call.apply(l, b(f = [this]).call(f, d));
    } }, { key: "rule", value: function() {
      for (var l, f, p = arguments.length, d = new Array(p), g = 0; g < p; g++)
        d[g] = arguments[g];
      return i.rule ? i.rule.apply(this, d) : (l = dt(T(u.prototype), "rule", this)).call.apply(l, b(f = [this]).call(f, d));
    } }]), u;
  }(a), R(n, "HOOK_NAME", e), n;
}), R(er, "createMenuHook", function(e, t) {
  var r = Rl(t, ["subMenuConfig", "onClick", "shortcutKeys", "iconName"], { subMenuConfig: Array, onClick: "function", shortcutKeys: Array, iconName: "string" });
  return function(n) {
    U(i, ad);
    var a = $l(i);
    function i(o) {
      var s;
      return j(this, i), s = a.call(this, o), r.iconName || (s.noIcon = !0), s.setName(e, r.iconName), s.subMenuConfig = r.subMenuConfig || [], s;
    }
    return M(i, [{ key: "onClick", value: function() {
      for (var o, s, c = arguments.length, u = new Array(c), l = 0; l < c; l++)
        u[l] = arguments[l];
      return r.onClick ? r.onClick.apply(this, u) : (o = dt(T(i.prototype), "onClick", this)).call.apply(o, b(s = [this]).call(s, u));
    } }, { key: "shortcutKeys", get: function() {
      return r.shortcutKeys ? r.shortcutKeys : [];
    } }]), i;
  }();
}), R(er, "constants", vv), R(er, "VERSION", "0.8.34");
var yv = function(e, t) {
  for (var r = -1, n = e == null ? 0 : e.length; ++r < n && t(e[r], r, e) !== !1; )
    ;
  return e;
}, _v = Lu(Object.keys, Object), kv = Object.prototype.hasOwnProperty, wv = function(e) {
  if (!bo(e))
    return _v(e);
  var t = [];
  for (var r in Object(e))
    kv.call(e, r) && r != "constructor" && t.push(r);
  return t;
}, xo = function(e) {
  return ga(e) ? Hu(e) : wv(e);
}, Ev = function(e, t) {
  return e && mr(t, xo(t), e);
}, Sv = function(e, t) {
  return e && mr(t, br(t), e);
}, Av = function(e, t) {
  for (var r = -1, n = e == null ? 0 : e.length, a = 0, i = []; ++r < n; ) {
    var o = e[r];
    t(o, r, e) && (i[a++] = o);
  }
  return i;
}, id = function() {
  return [];
}, xv = Object.prototype.propertyIsEnumerable, Pl = Object.getOwnPropertySymbols, Cv = Pl ? function(e) {
  return e == null ? [] : (e = Object(e), Av(Pl(e), function(t) {
    return xv.call(e, t);
  }));
} : id, Co = Cv, Tv = function(e, t) {
  return mr(e, Co(e), t);
}, od = function(e, t) {
  for (var r = -1, n = t.length, a = e.length; ++r < n; )
    e[a + r] = t[r];
  return e;
}, sd = Object.getOwnPropertySymbols ? function(e) {
  for (var t = []; e; )
    od(t, Co(e)), e = mo(e);
  return t;
} : id, $v = function(e, t) {
  return mr(e, sd(e), t);
}, cd = function(e, t, r) {
  var n = t(e);
  return Pn(e) ? n : od(n, r(e));
}, Rv = function(e) {
  return cd(e, xo, Co);
}, Ov = function(e) {
  return cd(e, br, sd);
}, Yi = un(yt, "DataView"), Xi = un(yt, "Promise"), Vi = un(yt, "Set"), Ji = un(yt, "WeakMap"), Ll = "[object Map]", Il = "[object Promise]", Nl = "[object Set]", Ml = "[object WeakMap]", jl = "[object DataView]", Pv = ln(Yi), Lv = ln(or), Iv = ln(Xi), Nv = ln(Vi), Mv = ln(Ji), Yt = cn;
(Yi && Yt(new Yi(new ArrayBuffer(1))) != jl || or && Yt(new or()) != Ll || Xi && Yt(Xi.resolve()) != Il || Vi && Yt(new Vi()) != Nl || Ji && Yt(new Ji()) != Ml) && (Yt = function(e) {
  var t = cn(e), r = t == "[object Object]" ? e.constructor : void 0, n = r ? ln(r) : "";
  if (n)
    switch (n) {
      case Pv:
        return jl;
      case Lv:
        return Ll;
      case Iv:
        return Il;
      case Nv:
        return Nl;
      case Mv:
        return Ml;
    }
  return t;
});
var To = Yt, jv = Object.prototype.hasOwnProperty, Dv = function(e) {
  var t = e.length, r = new e.constructor(t);
  return t && typeof e[0] == "string" && jv.call(e, "index") && (r.index = e.index, r.input = e.input), r;
}, Fv = function(e, t) {
  var r = t ? ho(e.buffer) : e.buffer;
  return new e.constructor(r, e.byteOffset, e.byteLength);
}, Bv = /\w*$/, Hv = function(e) {
  var t = new e.constructor(e.source, Bv.exec(e));
  return t.lastIndex = e.lastIndex, t;
}, Dl = Bt ? Bt.prototype : void 0, Fl = Dl ? Dl.valueOf : void 0, zv = function(e) {
  return Fl ? Object(Fl.call(e)) : {};
}, Uv = function(e, t, r) {
  var n = e.constructor;
  switch (t) {
    case "[object ArrayBuffer]":
      return ho(e);
    case "[object Boolean]":
    case "[object Date]":
      return new n(+e);
    case "[object DataView]":
      return Fv(e, r);
    case "[object Float32Array]":
    case "[object Float64Array]":
    case "[object Int8Array]":
    case "[object Int16Array]":
    case "[object Int32Array]":
    case "[object Uint8Array]":
    case "[object Uint8ClampedArray]":
    case "[object Uint16Array]":
    case "[object Uint32Array]":
      return Ou(e, r);
    case "[object Map]":
    case "[object Set]":
      return new n();
    case "[object Number]":
    case "[object String]":
      return new n(e);
    case "[object RegExp]":
      return Hv(e);
    case "[object Symbol]":
      return zv(e);
  }
}, Wv = function(e) {
  return Wt(e) && To(e) == "[object Map]";
}, Bl = Ln && Ln.isMap, qv = Bl ? yo(Bl) : Wv, Gv = function(e) {
  return Wt(e) && To(e) == "[object Set]";
}, Hl = Ln && Ln.isSet, Kv = Hl ? yo(Hl) : Gv, ld = "[object Arguments]", ud = "[object Function]", fd = "[object Object]", ae = {};
ae[ld] = ae["[object Array]"] = ae["[object ArrayBuffer]"] = ae["[object DataView]"] = ae["[object Boolean]"] = ae["[object Date]"] = ae["[object Float32Array]"] = ae["[object Float64Array]"] = ae["[object Int8Array]"] = ae["[object Int16Array]"] = ae["[object Int32Array]"] = ae["[object Map]"] = ae["[object Number]"] = ae[fd] = ae["[object RegExp]"] = ae["[object Set]"] = ae["[object String]"] = ae["[object Symbol]"] = ae["[object Uint8Array]"] = ae["[object Uint8ClampedArray]"] = ae["[object Uint16Array]"] = ae["[object Uint32Array]"] = !0, ae["[object Error]"] = ae[ud] = ae["[object WeakMap]"] = !1;
var Zv = function e(t, r, n, a, i, o) {
  var s, c = 1 & r, u = 2 & r, l = 4 & r;
  if (n && (s = i ? n(t, a, i, o) : n(t)), s !== void 0)
    return s;
  if (!Ut(t))
    return t;
  var f = Pn(t);
  if (f) {
    if (s = Dv(t), !c)
      return Pu(t, s);
  } else {
    var p = To(t), d = p == ud || p == "[object GeneratorFunction]";
    if (vo(t))
      return Ru(t, c);
    if (p == fd || p == ld || d && !i) {
      if (s = u || d ? {} : Iu(t), !c)
        return u ? $v(t, Sv(s, t)) : Tv(t, Ev(s, t));
    } else {
      if (!ae[p])
        return i ? t : {};
      s = Uv(t, p, c);
    }
  }
  o || (o = new $u());
  var g = o.get(t);
  if (g)
    return g;
  o.set(t, s), Kv(t) ? t.forEach(function(m) {
    s.add(e(m, r, n, m, t, o));
  }) : qv(t) && t.forEach(function(m, h) {
    s.set(h, e(m, r, n, h, t, o));
  });
  var _ = f ? void 0 : (l ? u ? Ov : Rv : u ? br : xo)(t);
  return yv(_ || t, function(m, h) {
    _ && (m = t[h = m]), Fu(s, h, e(m, r, n, h, t, o));
  }), s;
}, dd = function(e) {
  return Zv(e, 5);
}, Ot = { urlProcessor: function(e, t) {
  return e;
}, fileUpload: function(e, t) {
  if (/video/i.test(e.type))
    t("images/demo-dog.png", { name: "".concat(e.name.replace(/\.[^.]+$/, "")), poster: "images/demo-dog.png?poster=true", isBorder: !0, isShadow: !0, isRadius: !0 });
  else if (/image/i.test(e.type)) {
    var r = new FileReader();
    r.onload = function(n) {
      var a = n.target.result;
      t(a, { name: "".concat(e.name.replace(/\.[^.]+$/, "")), isShadow: !0, width: "60%", height: "auto" });
    }, r.readAsDataURL(e);
  } else
    t("images/demo-dog.png");
}, afterChange: function(e, t) {
}, afterInit: function(e, t) {
}, beforeImageMounted: function(e, t) {
  return { srcProp: e, src: t };
}, onClickPreview: function(e) {
}, onCopyCode: function(e, t) {
  return t;
}, changeString2Pinyin: function(e) {
  return e;
} }, Yv = dd({ externals: {}, openai: { apiKey: "", ignoreError: !1 }, engine: { global: { classicBr: !1, urlProcessor: Ot.urlProcessor, htmlWhiteList: "" }, syntax: { link: { target: "", rel: "" }, autoLink: { target: "", rel: "", enableShortLink: !0, shortLinkLength: 20 }, list: { listNested: !1, indentSpace: 2 }, table: { enableChart: !1 }, inlineCode: { theme: "red" }, codeBlock: { theme: "dark", wrap: !0, lineNumber: !0, copyCode: !0, editCode: !0, changeLang: !0, customRenderer: {}, mermaid: { svg2img: !1 }, indentedCodeBlock: !0 }, emoji: { useUnicode: !0 }, fontEmphasis: { allowWhitespace: !1 }, strikethrough: { needWhitespace: !1 }, mathBlock: { engine: "MathJax", src: "", plugins: !0 }, inlineMath: { engine: "MathJax", src: "" }, toc: { allowMultiToc: !1 }, header: { anchorStyle: "default" } } }, editor: { id: "code", name: "code", autoSave2Textarea: !1, theme: "default", height: "100%", defaultModel: "edit&preview", convertWhenPaste: !0, codemirror: { autofocus: !0 }, writingStyle: "normal", keepDocumentScrollAfterInit: !1 }, toolbars: { theme: "dark", showToolbar: !0, toolbar: ["bold", "italic", "strikethrough", "|", "color", "header", "ruby", "|", "list", "panel", "detail", { insert: ["image", "audio", "video", "link", "hr", "br", "code", "formula", "toc", "table", "line-table", "bar-table", "pdf", "word"] }, "graph", "settings"], toolbarRight: [], sidebar: [], bubble: ["bold", "italic", "underline", "strikethrough", "sub", "sup", "quote", "|", "size", "color"], float: ["h1", "h2", "h3", "|", "checklist", "quote", "table", "code"], toc: !1, shortcutKey: {}, config: { formula: { showLatexLive: !0, templateConfig: !1 } } }, drawioIframeUrl: "", fileUpload: Ot.fileUpload, fileTypeLimitMap: { video: "video/*", audio: "audio/*", image: "image/*", word: ".doc,.docx", pdf: ".pdf", file: "*" }, callback: { afterChange: Ot.afterChange, afterInit: Ot.afterInit, beforeImageMounted: Ot.beforeImageMounted, onClickPreview: Ot.onClickPreview, onCopyCode: Ot.onCopyCode, changeString2Pinyin: Ot.changeString2Pinyin }, previewer: { dom: !1, className: "cherry-markdown", enablePreviewerBubble: !0, lazyLoadImg: { loadingImgPath: "", maxNumPerTime: 2, noLoadImgNum: 5, autoLoadImgNum: 5, maxTryTimesPerSrc: 2, beforeLoadOneImgCallback: function(e) {
  return !0;
}, failLoadOneImgCallback: function(e) {
}, afterLoadOneImgCallback: function(e) {
}, afterLoadAllImgCallback: function() {
} } }, theme: [{ className: "default", label: "默认" }, { className: "dark", label: "暗黑" }, { className: "light", label: "明亮" }, { className: "green", label: "清新" }, { className: "red", label: "热情" }, { className: "violet", label: "淡雅" }, { className: "blue", label: "清幽" }], isPreviewOnly: !1, autoScrollByCursor: !0, forceAppend: !0, locale: "zh_CN", autoScrollByHashAfterInit: !1 });
function Xv(e) {
  var t = function() {
    if (typeof Reflect > "u" || !k || k.sham)
      return !1;
    if (typeof Proxy == "function")
      return !0;
    try {
      return Boolean.prototype.valueOf.call(k(Boolean, [], function() {
      })), !0;
    } catch {
      return !1;
    }
  }();
  return function() {
    var r, n = T(e);
    if (t) {
      var a = T(this).constructor;
      r = k(n, arguments, a);
    } else
      r = n.apply(this, arguments);
    return B(this, r);
  };
}
var Qi = function(e) {
  U(r, er);
  var t = Xv(r);
  function r(n) {
    var a;
    j(this, r), a = t.call(this), r.initialized = !0;
    var i, o = dd(r.config.defaults), s = i1({}, o, n, L1);
    return typeof s.engine.global.urlProcessor == "function" && (s.engine.global.urlProcessor = (i = s.engine.global.urlProcessor, function(c, u) {
      if (St.isInnerLink(c)) {
        var l = i(St.get(c), u);
        return St.replace(c, l);
      }
      return i(c, u);
    })), B(a, new mv(s, { options: s }));
  }
  return M(r);
}();
R(Qi, "initialized", !1), R(Qi, "config", { defaults: Yv });
var Vv = Qi;
const Jv = (e, t) => {
  const r = e.__vccOpts || e;
  for (const [n, a] of t)
    r[n] = a;
  return r;
}, ty = {
  convertMdToHtml: function(e) {
    var t = new Vv({
      // 配置同 cherrymarkdown
    }), r = t.makeHtml(e);
    return r;
  }
}, Qv = {};
function ey(e, t, r, n, a, i) {
  return null;
}
const ny = /* @__PURE__ */ Jv(Qv, [["render", ey]]);
export {
  ty as MarkdownViewer,
  ny as default
};
