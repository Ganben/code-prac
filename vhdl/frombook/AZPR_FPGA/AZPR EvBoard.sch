<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE eagle SYSTEM "eagle.dtd">
<eagle version="6.1">
<drawing>
<settings>
<setting alwaysvectorfont="no"/>
<setting verticaltext="up"/>
</settings>
<grid distance="0.1" unitdist="inch" unit="inch" style="lines" multiple="1" display="no" altdistance="0.01" altunitdist="inch" altunit="inch"/>
<layers>
<layer number="1" name="Top" color="4" fill="1" visible="no" active="no"/>
<layer number="16" name="Bottom" color="1" fill="1" visible="no" active="no"/>
<layer number="17" name="Pads" color="2" fill="1" visible="no" active="no"/>
<layer number="18" name="Vias" color="2" fill="1" visible="no" active="no"/>
<layer number="19" name="Unrouted" color="6" fill="1" visible="no" active="no"/>
<layer number="20" name="Dimension" color="15" fill="1" visible="no" active="no"/>
<layer number="21" name="tPlace" color="7" fill="1" visible="no" active="no"/>
<layer number="22" name="bPlace" color="7" fill="1" visible="no" active="no"/>
<layer number="23" name="tOrigins" color="15" fill="1" visible="no" active="no"/>
<layer number="24" name="bOrigins" color="15" fill="1" visible="no" active="no"/>
<layer number="25" name="tNames" color="7" fill="1" visible="no" active="no"/>
<layer number="26" name="bNames" color="7" fill="1" visible="no" active="no"/>
<layer number="27" name="tValues" color="7" fill="1" visible="no" active="no"/>
<layer number="28" name="bValues" color="7" fill="1" visible="no" active="no"/>
<layer number="29" name="tStop" color="7" fill="3" visible="no" active="no"/>
<layer number="30" name="bStop" color="7" fill="6" visible="no" active="no"/>
<layer number="31" name="tCream" color="7" fill="4" visible="no" active="no"/>
<layer number="32" name="bCream" color="7" fill="5" visible="no" active="no"/>
<layer number="33" name="tFinish" color="6" fill="3" visible="no" active="no"/>
<layer number="34" name="bFinish" color="6" fill="6" visible="no" active="no"/>
<layer number="35" name="tGlue" color="7" fill="4" visible="no" active="no"/>
<layer number="36" name="bGlue" color="7" fill="5" visible="no" active="no"/>
<layer number="37" name="tTest" color="7" fill="1" visible="no" active="no"/>
<layer number="38" name="bTest" color="7" fill="1" visible="no" active="no"/>
<layer number="39" name="tKeepout" color="4" fill="11" visible="no" active="no"/>
<layer number="40" name="bKeepout" color="1" fill="11" visible="no" active="no"/>
<layer number="41" name="tRestrict" color="4" fill="10" visible="no" active="no"/>
<layer number="42" name="bRestrict" color="1" fill="10" visible="no" active="no"/>
<layer number="43" name="vRestrict" color="2" fill="10" visible="no" active="no"/>
<layer number="44" name="Drills" color="7" fill="1" visible="no" active="no"/>
<layer number="45" name="Holes" color="7" fill="1" visible="no" active="no"/>
<layer number="46" name="Milling" color="3" fill="1" visible="no" active="no"/>
<layer number="47" name="Measures" color="7" fill="1" visible="no" active="no"/>
<layer number="48" name="Document" color="7" fill="1" visible="no" active="no"/>
<layer number="49" name="Reference" color="7" fill="1" visible="no" active="no"/>
<layer number="51" name="tDocu" color="7" fill="1" visible="no" active="no"/>
<layer number="52" name="bDocu" color="7" fill="1" visible="no" active="no"/>
<layer number="91" name="Nets" color="2" fill="1" visible="yes" active="yes"/>
<layer number="92" name="Busses" color="1" fill="1" visible="yes" active="yes"/>
<layer number="93" name="Pins" color="2" fill="1" visible="no" active="yes"/>
<layer number="94" name="Symbols" color="4" fill="1" visible="yes" active="yes"/>
<layer number="95" name="Names" color="7" fill="1" visible="yes" active="yes"/>
<layer number="96" name="Values" color="7" fill="1" visible="yes" active="yes"/>
<layer number="97" name="Info" color="7" fill="1" visible="no" active="yes"/>
<layer number="98" name="Guide" color="6" fill="1" visible="yes" active="yes"/>
<layer number="100" name="CATEGORY" color="15" fill="1" visible="yes" active="yes"/>
</layers>
<schematic xreflabel="%F%N/%S.%C%R" xrefpart="/%S.%C%R">
<libraries>
<library name="00_Respon">
<packages>
<package name="SMDPAD_1605_04INCH">
<wire x1="-10.16" y1="0" x2="-3.54" y2="0" width="0.2032" layer="1"/>
<wire x1="-1.54" y1="0" x2="0" y2="0" width="0.2032" layer="1"/>
<wire x1="-3.5" y1="1" x2="-2.5" y2="1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="1" x2="-1.5" y2="1" width="0.254" layer="21"/>
<wire x1="-3.5" y1="-1" x2="-2.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="-1" x2="-1.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="-1" x2="-2.5" y2="1" width="0.254" layer="21"/>
<pad name="P$1" x="-10.16" y="0" drill="0.8" diameter="1.6"/>
<pad name="P$2" x="0" y="0" drill="0.8" diameter="1.6"/>
<smd name="1" x="-1.54" y="0" dx="1" dy="1" layer="1"/>
<smd name="2" x="-3.54" y="0" dx="1" dy="1" layer="1"/>
<text x="-11.43" y="0.635" size="1.27" layer="25" font="vector" ratio="20" rot="R180">&gt;NAME</text>
<text x="1.27" y="-0.635" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-3.04" y1="-0.5" x2="-2.04" y2="0.5" layer="41"/>
<rectangle x1="-1.5" y1="-0.5" x2="-0.5" y2="0.5" layer="29"/>
</package>
<package name="SMDPAD_1605_02INCH">
<wire x1="-5.08" y1="0" x2="-3.54" y2="0" width="0.2032" layer="1"/>
<wire x1="0" y1="0" x2="-1.54" y2="0" width="0.2032" layer="1"/>
<wire x1="-3.5" y1="1" x2="-2.5" y2="1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="1" x2="-1.5" y2="1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="1" x2="-2.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-3.5" y1="-1" x2="-2.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="-1" x2="-1.5" y2="-1" width="0.254" layer="21"/>
<pad name="P$1" x="-5.08" y="0" drill="0.8" diameter="1.6"/>
<pad name="P$2" x="0" y="0" drill="0.8" diameter="1.6"/>
<smd name="1" x="-1.54" y="0" dx="1" dy="1" layer="1"/>
<smd name="2" x="-3.54" y="0" dx="1" dy="1" layer="1"/>
<text x="-6.35" y="0.635" size="1.27" layer="25" font="vector" ratio="20" rot="R180">&gt;NAME</text>
<text x="1.27" y="-0.635" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-3.04" y1="-0.5" x2="-2.04" y2="0.5" layer="41"/>
<rectangle x1="-1" y1="-0.5" x2="-0.5" y2="0.5" layer="29"/>
<rectangle x1="-4.5" y1="-0.5" x2="-4" y2="0.5" layer="29"/>
</package>
<package name="SMDPAD_1605_03INCH">
<wire x1="-7.62" y1="0" x2="-3.54" y2="0" width="0.2032" layer="1"/>
<wire x1="-1.54" y1="0" x2="0" y2="0" width="0.2032" layer="1"/>
<wire x1="-3.5" y1="1" x2="-2.5" y2="1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="1" x2="-1.5" y2="1" width="0.254" layer="21"/>
<wire x1="-3.5" y1="-1" x2="-2.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="-1" x2="-1.5" y2="-1" width="0.254" layer="21"/>
<wire x1="-2.5" y1="-1" x2="-2.5" y2="1" width="0.254" layer="21"/>
<pad name="P$1" x="-7.62" y="0" drill="0.8" diameter="1.6"/>
<pad name="P$2" x="0" y="0" drill="0.8" diameter="1.6"/>
<smd name="1" x="-1.54" y="0" dx="1" dy="1" layer="1"/>
<smd name="2" x="-3.54" y="0" dx="1" dy="1" layer="1"/>
<text x="-8.89" y="0.635" size="1.27" layer="25" font="vector" ratio="20" rot="R180">&gt;NAME</text>
<text x="1.27" y="-0.635" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-3.04" y1="-0.5" x2="-2.04" y2="0.5" layer="41"/>
<rectangle x1="-1.5" y1="-0.5" x2="-0.5" y2="0.5" layer="29"/>
</package>
<package name="SMD_1605_R">
<wire x1="-1" y1="1" x2="0" y2="1" width="0.254" layer="21"/>
<wire x1="0" y1="1" x2="1" y2="1" width="0.254" layer="21"/>
<wire x1="0" y1="1" x2="0" y2="-1" width="0.254" layer="21"/>
<wire x1="-1" y1="-1" x2="1" y2="-1" width="0.254" layer="21"/>
<smd name="1" x="1" y="0" dx="1" dy="1" layer="1"/>
<smd name="2" x="-1" y="0" dx="1" dy="1" layer="1"/>
<text x="1.905" y="-0.635" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="-1.905" y="0.635" size="1.27" layer="27" font="vector" ratio="20" rot="R180">&gt;VALUE</text>
<rectangle x1="-0.5" y1="-0.5" x2="0.5" y2="0.5" layer="41"/>
</package>
<package name="SMD_1605">
<wire x1="-1" y1="1" x2="0" y2="1" width="0.254" layer="21"/>
<wire x1="0" y1="1" x2="1" y2="1" width="0.254" layer="21"/>
<wire x1="0" y1="1" x2="0" y2="-1" width="0.254" layer="21"/>
<wire x1="-1" y1="-1" x2="1" y2="-1" width="0.254" layer="21"/>
<smd name="1" x="1" y="0" dx="1" dy="1" layer="1"/>
<smd name="2" x="-1" y="0" dx="1" dy="1" layer="1"/>
<text x="1.905" y="-0.635" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="-1.905" y="0.635" size="1.27" layer="27" font="vector" ratio="20" rot="R180">&gt;VALUE</text>
<rectangle x1="-0.5" y1="-0.5" x2="0.5" y2="0.5" layer="41"/>
</package>
<package name="SMD_HC-49_S">
<wire x1="-5" y1="-2" x2="-5" y2="-3" width="0.254" layer="21"/>
<wire x1="-5" y1="-3" x2="5" y2="-3" width="0.254" layer="21"/>
<wire x1="5" y1="-3" x2="5" y2="-2" width="0.254" layer="21"/>
<wire x1="5" y1="2" x2="5" y2="3" width="0.254" layer="21"/>
<wire x1="5" y1="3" x2="-5" y2="3" width="0.254" layer="21"/>
<wire x1="-5" y1="3" x2="-5" y2="2" width="0.254" layer="21"/>
<smd name="P$1" x="4.75" y="0" dx="5.5" dy="2" layer="1"/>
<smd name="P$2" x="-4.75" y="0" dx="5.5" dy="2" layer="1"/>
<text x="-5.08" y="-5.08" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="1.27" y="-5.08" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-2" y1="-1" x2="2" y2="1" layer="41"/>
<rectangle x1="-6" y1="1" x2="6" y2="3" layer="41"/>
<rectangle x1="-6" y1="-3" x2="6" y2="-1" layer="41"/>
</package>
<package name="SMD_KC7050B">
<wire x1="-1.5" y1="3" x2="1.5" y2="3" width="0.254" layer="21"/>
<wire x1="3.3" y1="0.9" x2="3.3" y2="-0.9" width="0.254" layer="21"/>
<wire x1="-3.3" y1="0.9" x2="-3.3" y2="-0.9" width="0.254" layer="21"/>
<wire x1="-1.5" y1="-3" x2="1.5" y2="-3" width="0.254" layer="21"/>
<smd name="3" x="2.54" y="2.1" dx="1.8" dy="2" layer="1"/>
<smd name="2" x="2.54" y="-2.1" dx="1.8" dy="2" layer="1"/>
<smd name="4" x="-2.54" y="2.1" dx="1.8" dy="2" layer="1"/>
<smd name="1" x="-2.54" y="-2.1" dx="1.8" dy="2" layer="1"/>
<text x="-3.175" y="-6.35" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<text x="-3.175" y="-5.08" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<rectangle x1="-3.44" y1="-1.1" x2="3.44" y2="1.1" layer="41"/>
<rectangle x1="-1.64" y1="-3.1" x2="1.64" y2="-1.1" layer="41"/>
<rectangle x1="-1.64" y1="1.1" x2="1.64" y2="3.1" layer="41"/>
</package>
<package name="SMD_SOP14">
<description>&lt;b&gt;&lt;/b&gt;&lt;p&gt;
Auto generated by &lt;i&gt;make-symbol-device-package-bsdl.ulp Rev. 34&lt;/i&gt;&lt;br&gt;</description>
<wire x1="-0.2" y1="5.2984" x2="0.2" y2="5.2984" width="0.254" layer="21" curve="180"/>
<wire x1="-2.5484" y1="-5.2984" x2="2.5484" y2="-5.2984" width="0.254" layer="21"/>
<wire x1="2.5484" y1="-5.2984" x2="2.5484" y2="5.2984" width="0.254" layer="21"/>
<wire x1="2.5484" y1="5.2984" x2="-2.5484" y2="5.2984" width="0.254" layer="21"/>
<wire x1="-2.5484" y1="5.2984" x2="-2.5484" y2="-5.2984" width="0.254" layer="21"/>
<smd name="1" x="-3.95" y="3.81" dx="1.5" dy="0.76" layer="1"/>
<smd name="2" x="-3.95" y="2.54" dx="1.5" dy="0.76" layer="1"/>
<smd name="3" x="-3.95" y="1.27" dx="1.5" dy="0.76" layer="1"/>
<smd name="4" x="-3.95" y="0" dx="1.5" dy="0.76" layer="1"/>
<smd name="5" x="-3.95" y="-1.27" dx="1.5" dy="0.76" layer="1"/>
<smd name="6" x="-3.95" y="-2.54" dx="1.5" dy="0.76" layer="1"/>
<smd name="7" x="-3.95" y="-3.81" dx="1.5" dy="0.76" layer="1"/>
<smd name="8" x="3.95" y="-3.81" dx="1.5" dy="0.76" layer="1"/>
<smd name="9" x="3.95" y="-2.54" dx="1.5" dy="0.76" layer="1"/>
<smd name="10" x="3.95" y="-1.27" dx="1.5" dy="0.76" layer="1"/>
<smd name="11" x="3.95" y="0" dx="1.5" dy="0.76" layer="1"/>
<smd name="12" x="3.95" y="1.27" dx="1.5" dy="0.76" layer="1"/>
<smd name="13" x="3.95" y="2.54" dx="1.5" dy="0.76" layer="1"/>
<smd name="14" x="3.95" y="3.81" dx="1.5" dy="0.76" layer="1"/>
<text x="-0.635" y="3.9" size="1.27" layer="25" font="vector" ratio="20" rot="R270">&gt;NAME</text>
<text x="-3.81" y="-6.985" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-3.9" y1="3.43" x2="-2.65" y2="4.19" layer="51"/>
<rectangle x1="-3.9" y1="2.16" x2="-2.65" y2="2.92" layer="51"/>
<rectangle x1="-3.9" y1="0.89" x2="-2.65" y2="1.65" layer="51"/>
<rectangle x1="-3.9" y1="-0.38" x2="-2.65" y2="0.38" layer="51"/>
<rectangle x1="-3.9" y1="-1.65" x2="-2.65" y2="-0.89" layer="51"/>
<rectangle x1="-3.9" y1="-2.92" x2="-2.65" y2="-2.16" layer="51"/>
<rectangle x1="-3.9" y1="-4.19" x2="-2.65" y2="-3.43" layer="51"/>
<rectangle x1="2.65" y1="-4.19" x2="3.9" y2="-3.43" layer="51"/>
<rectangle x1="2.65" y1="-2.92" x2="3.9" y2="-2.16" layer="51"/>
<rectangle x1="2.65" y1="-1.65" x2="3.9" y2="-0.89" layer="51"/>
<rectangle x1="2.65" y1="-0.38" x2="3.9" y2="0.38" layer="51"/>
<rectangle x1="2.65" y1="0.89" x2="3.9" y2="1.65" layer="51"/>
<rectangle x1="2.65" y1="2.16" x2="3.9" y2="2.92" layer="51"/>
<rectangle x1="2.65" y1="3.43" x2="3.9" y2="4.19" layer="51"/>
</package>
<package name="SMD_SOP8">
<description>&lt;b&gt;&lt;/b&gt;&lt;p&gt;
Auto generated by &lt;i&gt;make-symbol-device-package-bsdl.ulp Rev. 34&lt;/i&gt;&lt;br&gt;</description>
<wire x1="-0.2" y1="2.3484" x2="0.2" y2="2.3484" width="0.2032" layer="21" curve="180"/>
<wire x1="-1.7484" y1="-2.3484" x2="1.7484" y2="-2.3484" width="0.254" layer="21"/>
<wire x1="1.7484" y1="-2.3484" x2="1.7484" y2="2.3484" width="0.254" layer="21"/>
<wire x1="1.7484" y1="2.3484" x2="-1.7484" y2="2.3484" width="0.254" layer="21"/>
<wire x1="-1.7484" y1="2.3484" x2="-1.7484" y2="-2.3484" width="0.254" layer="21"/>
<smd name="1" x="-2.85" y="1.905" dx="1.8" dy="0.6" layer="1"/>
<smd name="2" x="-2.85" y="0.635" dx="1.8" dy="0.6" layer="1"/>
<smd name="3" x="-2.85" y="-0.635" dx="1.8" dy="0.6" layer="1"/>
<smd name="4" x="-2.85" y="-1.905" dx="1.8" dy="0.6" layer="1"/>
<smd name="5" x="2.85" y="-1.905" dx="1.8" dy="0.6" layer="1"/>
<smd name="6" x="2.85" y="-0.635" dx="1.8" dy="0.6" layer="1"/>
<smd name="7" x="2.85" y="0.635" dx="1.8" dy="0.6" layer="1"/>
<smd name="8" x="2.85" y="1.905" dx="1.8" dy="0.6" layer="1"/>
<text x="-0.635" y="1.68" size="1.27" layer="25" font="vector" ratio="20" rot="R270">&gt;NAME</text>
<text x="-3.81" y="-4.445" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-2.95" y1="1.655" x2="-1.95" y2="2.155" layer="51"/>
<rectangle x1="-2.95" y1="0.385" x2="-1.95" y2="0.885" layer="51"/>
<rectangle x1="-2.95" y1="-0.885" x2="-1.95" y2="-0.385" layer="51"/>
<rectangle x1="-2.95" y1="-2.155" x2="-1.95" y2="-1.655" layer="51"/>
<rectangle x1="1.95" y1="-2.155" x2="2.95" y2="-1.655" layer="51"/>
<rectangle x1="1.95" y1="-0.885" x2="2.95" y2="-0.385" layer="51"/>
<rectangle x1="1.95" y1="0.385" x2="2.95" y2="0.885" layer="51"/>
<rectangle x1="1.95" y1="1.655" x2="2.95" y2="2.155" layer="51"/>
</package>
<package name="SMD_1605_POL">
<wire x1="-1.5" y1="1.2" x2="-0.5" y2="1.2" width="0.254" layer="21"/>
<wire x1="-1" y1="0.8" x2="-1" y2="1.7" width="0.254" layer="21"/>
<wire x1="1.016" y1="1.016" x2="0" y2="1.016" width="0.254" layer="21"/>
<wire x1="0" y1="1.016" x2="0" y2="-1.016" width="0.254" layer="21"/>
<wire x1="0" y1="-1.016" x2="-1.27" y2="-1.016" width="0.254" layer="21"/>
<wire x1="0" y1="-1.016" x2="1.016" y2="-1.016" width="0.254" layer="21"/>
<smd name="C" x="1" y="0" dx="1" dy="1" layer="1"/>
<smd name="A" x="-1" y="0" dx="1" dy="1" layer="1"/>
<text x="1.905" y="-0.635" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="-1.905" y="-2.54" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-0.5" y1="-0.5" x2="0.5" y2="0.5" layer="41"/>
</package>
<package name="PAD_POW_M">
<wire x1="-4" y1="6" x2="-4" y2="-6" width="0.254" layer="21"/>
<wire x1="-4" y1="6" x2="-3.5" y2="6" width="0.254" layer="21"/>
<wire x1="-3.5" y1="6" x2="-3.5" y2="-6" width="0.254" layer="21"/>
<wire x1="-3.5" y1="-6" x2="-4" y2="-6" width="0.254" layer="21"/>
<pad name="6" x="1.27" y="0" drill="1" diameter="2"/>
<pad name="4" x="1.27" y="2.54" drill="1" diameter="2"/>
<pad name="2" x="1.27" y="5.08" drill="1" diameter="2"/>
<pad name="8" x="1.27" y="-2.54" drill="1" diameter="2"/>
<pad name="10" x="1.27" y="-5.08" drill="1" diameter="2"/>
<pad name="9" x="-1.27" y="-5.08" drill="1" diameter="2"/>
<pad name="7" x="-1.27" y="-2.54" drill="1" diameter="2"/>
<pad name="5" x="-1.27" y="0" drill="1" diameter="2"/>
<pad name="3" x="-1.27" y="2.54" drill="1" diameter="2"/>
<pad name="1" x="-1.27" y="5.08" drill="1" diameter="2"/>
<text x="-3.175" y="6.985" size="1.27" layer="21" font="vector" ratio="20">&gt;NAME</text>
<text x="-3.175" y="-8.255" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
</package>
<package name="PAD_VPORT">
<wire x1="10" y1="-4.2" x2="10" y2="4.2" width="0.254" layer="21"/>
<wire x1="10" y1="4.2" x2="-10" y2="4.2" width="0.254" layer="21"/>
<wire x1="-10" y1="4.2" x2="-10" y2="-4.2" width="0.254" layer="21"/>
<wire x1="-10" y1="-4.2" x2="-1.27" y2="-4.2" width="0.254" layer="21"/>
<wire x1="1.27" y1="-4.2" x2="10" y2="-4.2" width="0.254" layer="21"/>
<wire x1="-9.5" y1="3.5" x2="9.5" y2="3.5" width="0.254" layer="21"/>
<wire x1="9.5" y1="3.5" x2="9.5" y2="-3.5" width="0.254" layer="21"/>
<wire x1="9.5" y1="-3.5" x2="1.27" y2="-3.5" width="0.254" layer="21"/>
<wire x1="1.27" y1="-3.5" x2="1.27" y2="-4.2" width="0.254" layer="21"/>
<wire x1="-9.5" y1="3.5" x2="-9.5" y2="-3.5" width="0.254" layer="21"/>
<wire x1="-9.5" y1="-3.5" x2="-1.27" y2="-3.5" width="0.254" layer="21"/>
<wire x1="-1.27" y1="-3.5" x2="-1.27" y2="-4.2" width="0.254" layer="21"/>
<pad name="6" x="0" y="1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="4" x="-2.54" y="1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="2" x="-5.08" y="1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="8" x="2.54" y="1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="10" x="5.08" y="1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="9" x="5.08" y="-1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="7" x="2.54" y="-1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="5" x="0" y="-1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="3" x="-2.54" y="-1.27" drill="1" diameter="1.8" rot="R90"/>
<pad name="1" x="-5.08" y="-1.27" drill="1" diameter="1.8" rot="R90"/>
<text x="-12" y="3" size="1.27" layer="21" font="vector" ratio="20" rot="R270">&gt;NAME</text>
<text x="12.065" y="-3.81" size="1.27" layer="27" font="vector" ratio="20" rot="R90">&gt;VALUE</text>
</package>
<package name="SMD_SW_PUSH">
<wire x1="2.35" y1="1.6" x2="2.35" y2="-1.6" width="0.254" layer="21"/>
<wire x1="-2.35" y1="1.6" x2="-2.35" y2="-1.6" width="0.254" layer="21"/>
<wire x1="-1" y1="2.3" x2="1" y2="2.3" width="0.254" layer="21"/>
<wire x1="-1" y1="-2.3" x2="1" y2="-2.3" width="0.254" layer="21"/>
<circle x="0" y="0" radius="1.15" width="0.254" layer="21"/>
<circle x="0" y="0" radius="1.95" width="0.254" layer="21"/>
<smd name="2" x="-2.2125" y="-2.925" dx="1.825" dy="2.15" layer="1"/>
<smd name="4" x="2.2125" y="-2.925" dx="1.825" dy="2.15" layer="1"/>
<smd name="1" x="-2.2125" y="2.925" dx="1.825" dy="2.15" layer="1"/>
<smd name="3" x="2.2125" y="2.925" dx="1.825" dy="2.15" layer="1"/>
<text x="2.54" y="4.445" size="1.27" layer="25" font="vector" ratio="20" rot="R180">&gt;NAME</text>
<text x="-3.175" y="-4.445" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-2" y1="-2" x2="2" y2="2" layer="41"/>
</package>
<package name="LOGO">
<wire x1="0" y1="0" x2="-6.5" y2="-1.5" width="0.254" layer="21"/>
<wire x1="-6.5" y1="-1.5" x2="-5.75" y2="-1.5" width="0.254" layer="21"/>
<wire x1="-5.75" y1="-1.5" x2="-5.75" y2="-3" width="0.254" layer="21"/>
<wire x1="-5.75" y1="-3" x2="-6.5" y2="-3" width="0.254" layer="21"/>
<wire x1="-6.5" y1="-3" x2="0" y2="-4.5" width="0.254" layer="21"/>
<wire x1="0" y1="-4.5" x2="6.5" y2="-3" width="0.254" layer="21"/>
<wire x1="6.5" y1="-3" x2="5.75" y2="-3" width="0.254" layer="21"/>
<wire x1="5.75" y1="-3" x2="5.75" y2="-1.5" width="0.254" layer="21"/>
<wire x1="5.75" y1="-1.5" x2="6.5" y2="-1.5" width="0.254" layer="21"/>
<wire x1="6.5" y1="-1.5" x2="0" y2="0" width="0.254" layer="21"/>
<wire x1="-5.75" y1="-1.35" x2="-5.75" y2="-3.15" width="0.254" layer="21"/>
<wire x1="5.75" y1="-1.35" x2="5.75" y2="-3.15" width="0.254" layer="21"/>
<text x="-5.5" y="-2.75" size="0.8204" layer="21" font="vector" ratio="31">VPort with Power</text>
</package>
<package name="SMD_USB">
<wire x1="4" y1="-3.15" x2="4" y2="-0.65" width="0.254" layer="21"/>
<wire x1="-4" y1="-3.15" x2="-4" y2="-0.65" width="0.254" layer="21"/>
<wire x1="-2.7" y1="1.35" x2="-2.1" y2="1.35" width="0.254" layer="21"/>
<wire x1="2.7" y1="1.35" x2="2.1" y2="1.35" width="0.254" layer="21"/>
<smd name="3" x="0" y="1" dx="0.5" dy="2" layer="1"/>
<smd name="4" x="0.8" y="1" dx="0.5" dy="2" layer="1"/>
<smd name="5" x="1.6" y="1" dx="0.5" dy="2" layer="1"/>
<smd name="2" x="-0.8" y="1" dx="0.5" dy="2" layer="1"/>
<smd name="1" x="-1.6" y="1" dx="0.5" dy="2" layer="1"/>
<smd name="PAD@3" x="4.2" y="0.75" dx="2.5" dy="2.2" layer="1"/>
<smd name="PAD@1" x="-4.2" y="0.75" dx="2.5" dy="2.2" layer="1"/>
<smd name="PAD@4" x="4.2" y="-4.55" dx="2.5" dy="2.2" layer="1"/>
<smd name="PAD@2" x="-4.2" y="-4.55" dx="2.5" dy="2.2" layer="1"/>
<text x="2.54" y="2.51" size="1.27" layer="21" font="vector" ratio="20">&gt;NAME</text>
<text x="-5.715" y="2.51" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<polygon width="0.127" layer="41">
<vertex x="-4" y="-7.65"/>
<vertex x="4" y="-7.65"/>
<vertex x="4" y="-5.75"/>
<vertex x="2.8" y="-5.75"/>
<vertex x="2.8" y="-3.35"/>
<vertex x="4" y="-3.35"/>
<vertex x="4" y="-1.65"/>
<vertex x="1.8" y="-1.65"/>
<vertex x="1.8" y="-0.15"/>
<vertex x="-1.8" y="-0.15"/>
<vertex x="-1.8" y="-1.65"/>
<vertex x="-4" y="-1.65"/>
<vertex x="-4" y="-3.35"/>
<vertex x="-2.8" y="-3.35"/>
<vertex x="-2.8" y="-5.75"/>
<vertex x="-4" y="-5.75"/>
</polygon>
</package>
<package name="SMD_VQ100">
<description>&lt;b&gt;&lt;/b&gt;&lt;p&gt;
Auto generated by &lt;i&gt;make-symbol-device-package-bsdl.ulp Rev. 34&lt;/i&gt;&lt;br&gt;</description>
<wire x1="-6.8984" y1="6.2" x2="-6.2" y2="6.8984" width="0.254" layer="21"/>
<wire x1="-6.8984" y1="-6.8984" x2="6.8984" y2="-6.8984" width="0.254" layer="21"/>
<wire x1="6.8984" y1="-6.8984" x2="6.8984" y2="6.8984" width="0.254" layer="21"/>
<wire x1="6.8984" y1="6.8984" x2="-6.2" y2="6.8984" width="0.254" layer="21"/>
<wire x1="-6.2" y1="6.8984" x2="-6.8984" y2="6.8984" width="0.254" layer="21"/>
<wire x1="-6.8984" y1="6.8984" x2="-6.8984" y2="6.2" width="0.254" layer="21"/>
<wire x1="-6.8984" y1="6.2" x2="-6.8984" y2="-6.8984" width="0.254" layer="21"/>
<smd name="1" x="-8.3" y="6" dx="2" dy="0.25" layer="1"/>
<smd name="2" x="-8.3" y="5.5" dx="2" dy="0.25" layer="1"/>
<smd name="3" x="-8.3" y="5" dx="2" dy="0.25" layer="1"/>
<smd name="4" x="-8.3" y="4.5" dx="2" dy="0.25" layer="1"/>
<smd name="5" x="-8.3" y="4" dx="2" dy="0.25" layer="1"/>
<smd name="6" x="-8.3" y="3.5" dx="2" dy="0.25" layer="1"/>
<smd name="7" x="-8.3" y="3" dx="2" dy="0.25" layer="1"/>
<smd name="8" x="-8.3" y="2.5" dx="2" dy="0.25" layer="1"/>
<smd name="9" x="-8.3" y="2" dx="2" dy="0.25" layer="1"/>
<smd name="10" x="-8.3" y="1.5" dx="2" dy="0.25" layer="1"/>
<smd name="11" x="-8.3" y="1" dx="2" dy="0.25" layer="1"/>
<smd name="12" x="-8.3" y="0.5" dx="2" dy="0.25" layer="1"/>
<smd name="13" x="-8.3" y="0" dx="2" dy="0.25" layer="1"/>
<smd name="14" x="-8.3" y="-0.5" dx="2" dy="0.25" layer="1"/>
<smd name="15" x="-8.3" y="-1" dx="2" dy="0.25" layer="1"/>
<smd name="16" x="-8.3" y="-1.5" dx="2" dy="0.25" layer="1"/>
<smd name="17" x="-8.3" y="-2" dx="2" dy="0.25" layer="1"/>
<smd name="18" x="-8.3" y="-2.5" dx="2" dy="0.25" layer="1"/>
<smd name="19" x="-8.3" y="-3" dx="2" dy="0.25" layer="1"/>
<smd name="20" x="-8.3" y="-3.5" dx="2" dy="0.25" layer="1"/>
<smd name="21" x="-8.3" y="-4" dx="2" dy="0.25" layer="1"/>
<smd name="22" x="-8.3" y="-4.5" dx="2" dy="0.25" layer="1"/>
<smd name="23" x="-8.3" y="-5" dx="2" dy="0.25" layer="1"/>
<smd name="24" x="-8.3" y="-5.5" dx="2" dy="0.25" layer="1"/>
<smd name="25" x="-8.3" y="-6" dx="2" dy="0.25" layer="1"/>
<smd name="26" x="-6" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="27" x="-5.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="28" x="-5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="29" x="-4.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="30" x="-4" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="31" x="-3.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="32" x="-3" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="33" x="-2.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="34" x="-2" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="35" x="-1.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="36" x="-1" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="37" x="-0.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="38" x="0" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="39" x="0.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="40" x="1" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="41" x="1.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="42" x="2" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="43" x="2.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="44" x="3" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="45" x="3.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="46" x="4" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="47" x="4.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="48" x="5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="49" x="5.5" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="50" x="6" y="-8.3" dx="0.25" dy="2" layer="1"/>
<smd name="51" x="8.3" y="-6" dx="2" dy="0.25" layer="1"/>
<smd name="52" x="8.3" y="-5.5" dx="2" dy="0.25" layer="1"/>
<smd name="53" x="8.3" y="-5" dx="2" dy="0.25" layer="1"/>
<smd name="54" x="8.3" y="-4.5" dx="2" dy="0.25" layer="1"/>
<smd name="55" x="8.3" y="-4" dx="2" dy="0.25" layer="1"/>
<smd name="56" x="8.3" y="-3.5" dx="2" dy="0.25" layer="1"/>
<smd name="57" x="8.3" y="-3" dx="2" dy="0.25" layer="1"/>
<smd name="58" x="8.3" y="-2.5" dx="2" dy="0.25" layer="1"/>
<smd name="59" x="8.3" y="-2" dx="2" dy="0.25" layer="1"/>
<smd name="60" x="8.3" y="-1.5" dx="2" dy="0.25" layer="1"/>
<smd name="61" x="8.3" y="-1" dx="2" dy="0.25" layer="1"/>
<smd name="62" x="8.3" y="-0.5" dx="2" dy="0.25" layer="1"/>
<smd name="63" x="8.3" y="0" dx="2" dy="0.25" layer="1"/>
<smd name="64" x="8.3" y="0.5" dx="2" dy="0.25" layer="1"/>
<smd name="65" x="8.3" y="1" dx="2" dy="0.25" layer="1"/>
<smd name="66" x="8.3" y="1.5" dx="2" dy="0.25" layer="1"/>
<smd name="67" x="8.3" y="2" dx="2" dy="0.25" layer="1"/>
<smd name="68" x="8.3" y="2.5" dx="2" dy="0.25" layer="1"/>
<smd name="69" x="8.3" y="3" dx="2" dy="0.25" layer="1"/>
<smd name="70" x="8.3" y="3.5" dx="2" dy="0.25" layer="1"/>
<smd name="71" x="8.3" y="4" dx="2" dy="0.25" layer="1"/>
<smd name="72" x="8.3" y="4.5" dx="2" dy="0.25" layer="1"/>
<smd name="73" x="8.3" y="5" dx="2" dy="0.25" layer="1"/>
<smd name="74" x="8.3" y="5.5" dx="2" dy="0.25" layer="1"/>
<smd name="75" x="8.3" y="6" dx="2" dy="0.25" layer="1"/>
<smd name="76" x="6" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="77" x="5.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="78" x="5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="79" x="4.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="80" x="4" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="81" x="3.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="82" x="3" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="83" x="2.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="84" x="2" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="85" x="1.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="86" x="1" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="87" x="0.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="88" x="0" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="89" x="-0.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="90" x="-1" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="91" x="-1.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="92" x="-2" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="93" x="-2.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="94" x="-3" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="95" x="-3.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="96" x="-4" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="97" x="-4.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="98" x="-5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="99" x="-5.5" y="8.3" dx="0.25" dy="2" layer="1"/>
<smd name="100" x="-6" y="8.3" dx="0.25" dy="2" layer="1"/>
<text x="0.635" y="2.54" size="1.27" layer="25" font="vector" ratio="20" rot="R270">&gt;NAME</text>
<text x="-2.54" y="2.54" size="1.27" layer="27" font="vector" ratio="20" rot="R270">&gt;VALUE</text>
<rectangle x1="-8" y1="5.89" x2="-7" y2="6.11" layer="51"/>
<rectangle x1="-8" y1="5.39" x2="-7" y2="5.61" layer="51"/>
<rectangle x1="-8" y1="4.89" x2="-7" y2="5.11" layer="51"/>
<rectangle x1="-8" y1="4.39" x2="-7" y2="4.61" layer="51"/>
<rectangle x1="-8" y1="3.89" x2="-7" y2="4.11" layer="51"/>
<rectangle x1="-8" y1="3.39" x2="-7" y2="3.61" layer="51"/>
<rectangle x1="-8" y1="2.89" x2="-7" y2="3.11" layer="51"/>
<rectangle x1="-8" y1="2.39" x2="-7" y2="2.61" layer="51"/>
<rectangle x1="-8" y1="1.89" x2="-7" y2="2.11" layer="51"/>
<rectangle x1="-8" y1="1.39" x2="-7" y2="1.61" layer="51"/>
<rectangle x1="-8" y1="0.89" x2="-7" y2="1.11" layer="51"/>
<rectangle x1="-8" y1="0.39" x2="-7" y2="0.61" layer="51"/>
<rectangle x1="-8" y1="-0.11" x2="-7" y2="0.11" layer="51"/>
<rectangle x1="-8" y1="-0.61" x2="-7" y2="-0.39" layer="51"/>
<rectangle x1="-8" y1="-1.11" x2="-7" y2="-0.89" layer="51"/>
<rectangle x1="-8" y1="-1.61" x2="-7" y2="-1.39" layer="51"/>
<rectangle x1="-8" y1="-2.11" x2="-7" y2="-1.89" layer="51"/>
<rectangle x1="-8" y1="-2.61" x2="-7" y2="-2.39" layer="51"/>
<rectangle x1="-8" y1="-3.11" x2="-7" y2="-2.89" layer="51"/>
<rectangle x1="-8" y1="-3.61" x2="-7" y2="-3.39" layer="51"/>
<rectangle x1="-8" y1="-4.11" x2="-7" y2="-3.89" layer="51"/>
<rectangle x1="-8" y1="-4.61" x2="-7" y2="-4.39" layer="51"/>
<rectangle x1="-8" y1="-5.11" x2="-7" y2="-4.89" layer="51"/>
<rectangle x1="-8" y1="-5.61" x2="-7" y2="-5.39" layer="51"/>
<rectangle x1="-8" y1="-6.11" x2="-7" y2="-5.89" layer="51"/>
<rectangle x1="-6.11" y1="-8" x2="-5.89" y2="-7" layer="51"/>
<rectangle x1="-5.61" y1="-8" x2="-5.39" y2="-7" layer="51"/>
<rectangle x1="-5.11" y1="-8" x2="-4.89" y2="-7" layer="51"/>
<rectangle x1="-4.61" y1="-8" x2="-4.39" y2="-7" layer="51"/>
<rectangle x1="-4.11" y1="-8" x2="-3.89" y2="-7" layer="51"/>
<rectangle x1="-3.61" y1="-8" x2="-3.39" y2="-7" layer="51"/>
<rectangle x1="-3.11" y1="-8" x2="-2.89" y2="-7" layer="51"/>
<rectangle x1="-2.61" y1="-8" x2="-2.39" y2="-7" layer="51"/>
<rectangle x1="-2.11" y1="-8" x2="-1.89" y2="-7" layer="51"/>
<rectangle x1="-1.61" y1="-8" x2="-1.39" y2="-7" layer="51"/>
<rectangle x1="-1.11" y1="-8" x2="-0.89" y2="-7" layer="51"/>
<rectangle x1="-0.61" y1="-8" x2="-0.39" y2="-7" layer="51"/>
<rectangle x1="-0.11" y1="-8" x2="0.11" y2="-7" layer="51"/>
<rectangle x1="0.39" y1="-8" x2="0.61" y2="-7" layer="51"/>
<rectangle x1="0.89" y1="-8" x2="1.11" y2="-7" layer="51"/>
<rectangle x1="1.39" y1="-8" x2="1.61" y2="-7" layer="51"/>
<rectangle x1="1.89" y1="-8" x2="2.11" y2="-7" layer="51"/>
<rectangle x1="2.39" y1="-8" x2="2.61" y2="-7" layer="51"/>
<rectangle x1="2.89" y1="-8" x2="3.11" y2="-7" layer="51"/>
<rectangle x1="3.39" y1="-8" x2="3.61" y2="-7" layer="51"/>
<rectangle x1="3.89" y1="-8" x2="4.11" y2="-7" layer="51"/>
<rectangle x1="4.39" y1="-8" x2="4.61" y2="-7" layer="51"/>
<rectangle x1="4.89" y1="-8" x2="5.11" y2="-7" layer="51"/>
<rectangle x1="5.39" y1="-8" x2="5.61" y2="-7" layer="51"/>
<rectangle x1="5.89" y1="-8" x2="6.11" y2="-7" layer="51"/>
<rectangle x1="7" y1="-6.11" x2="8" y2="-5.89" layer="51"/>
<rectangle x1="7" y1="-5.61" x2="8" y2="-5.39" layer="51"/>
<rectangle x1="7" y1="-5.11" x2="8" y2="-4.89" layer="51"/>
<rectangle x1="7" y1="-4.61" x2="8" y2="-4.39" layer="51"/>
<rectangle x1="7" y1="-4.11" x2="8" y2="-3.89" layer="51"/>
<rectangle x1="7" y1="-3.61" x2="8" y2="-3.39" layer="51"/>
<rectangle x1="7" y1="-3.11" x2="8" y2="-2.89" layer="51"/>
<rectangle x1="7" y1="-2.61" x2="8" y2="-2.39" layer="51"/>
<rectangle x1="7" y1="-2.11" x2="8" y2="-1.89" layer="51"/>
<rectangle x1="7" y1="-1.61" x2="8" y2="-1.39" layer="51"/>
<rectangle x1="7" y1="-1.11" x2="8" y2="-0.89" layer="51"/>
<rectangle x1="7" y1="-0.61" x2="8" y2="-0.39" layer="51"/>
<rectangle x1="7" y1="-0.11" x2="8" y2="0.11" layer="51"/>
<rectangle x1="7" y1="0.39" x2="8" y2="0.61" layer="51"/>
<rectangle x1="7" y1="0.89" x2="8" y2="1.11" layer="51"/>
<rectangle x1="7" y1="1.39" x2="8" y2="1.61" layer="51"/>
<rectangle x1="7" y1="1.89" x2="8" y2="2.11" layer="51"/>
<rectangle x1="7" y1="2.39" x2="8" y2="2.61" layer="51"/>
<rectangle x1="7" y1="2.89" x2="8" y2="3.11" layer="51"/>
<rectangle x1="7" y1="3.39" x2="8" y2="3.61" layer="51"/>
<rectangle x1="7" y1="3.89" x2="8" y2="4.11" layer="51"/>
<rectangle x1="7" y1="4.39" x2="8" y2="4.61" layer="51"/>
<rectangle x1="7" y1="4.89" x2="8" y2="5.11" layer="51"/>
<rectangle x1="7" y1="5.39" x2="8" y2="5.61" layer="51"/>
<rectangle x1="7" y1="5.89" x2="8" y2="6.11" layer="51"/>
<rectangle x1="5.89" y1="7" x2="6.11" y2="8" layer="51"/>
<rectangle x1="5.39" y1="7" x2="5.61" y2="8" layer="51"/>
<rectangle x1="4.89" y1="7" x2="5.11" y2="8" layer="51"/>
<rectangle x1="4.39" y1="7" x2="4.61" y2="8" layer="51"/>
<rectangle x1="3.89" y1="7" x2="4.11" y2="8" layer="51"/>
<rectangle x1="3.39" y1="7" x2="3.61" y2="8" layer="51"/>
<rectangle x1="2.89" y1="7" x2="3.11" y2="8" layer="51"/>
<rectangle x1="2.39" y1="7" x2="2.61" y2="8" layer="51"/>
<rectangle x1="1.89" y1="7" x2="2.11" y2="8" layer="51"/>
<rectangle x1="1.39" y1="7" x2="1.61" y2="8" layer="51"/>
<rectangle x1="0.89" y1="7" x2="1.11" y2="8" layer="51"/>
<rectangle x1="0.39" y1="7" x2="0.61" y2="8" layer="51"/>
<rectangle x1="-0.11" y1="7" x2="0.11" y2="8" layer="51"/>
<rectangle x1="-0.61" y1="7" x2="-0.39" y2="8" layer="51"/>
<rectangle x1="-1.11" y1="7" x2="-0.89" y2="8" layer="51"/>
<rectangle x1="-1.61" y1="7" x2="-1.39" y2="8" layer="51"/>
<rectangle x1="-2.11" y1="7" x2="-1.89" y2="8" layer="51"/>
<rectangle x1="-2.61" y1="7" x2="-2.39" y2="8" layer="51"/>
<rectangle x1="-3.11" y1="7" x2="-2.89" y2="8" layer="51"/>
<rectangle x1="-3.61" y1="7" x2="-3.39" y2="8" layer="51"/>
<rectangle x1="-4.11" y1="7" x2="-3.89" y2="8" layer="51"/>
<rectangle x1="-4.61" y1="7" x2="-4.39" y2="8" layer="51"/>
<rectangle x1="-5.11" y1="7" x2="-4.89" y2="8" layer="51"/>
<rectangle x1="-5.61" y1="7" x2="-5.39" y2="8" layer="51"/>
<rectangle x1="-6.11" y1="7" x2="-5.89" y2="8" layer="51"/>
</package>
<package name="SMD_VO20C">
<description>&lt;b&gt;VO20C&lt;/b&gt;&lt;p&gt;
Auto generated by &lt;i&gt;make-symbol-device-package-bsdl.ulp Rev. 34&lt;/i&gt;&lt;br&gt;</description>
<wire x1="-0.2" y1="3.1484" x2="0.2" y2="3.1484" width="0.254" layer="21" curve="180"/>
<wire x1="-2.0984" y1="-3.1484" x2="2.0984" y2="-3.1484" width="0.254" layer="21"/>
<wire x1="2.0984" y1="-3.1484" x2="2.0984" y2="3.1484" width="0.254" layer="21"/>
<wire x1="2.0984" y1="3.1484" x2="-2.0984" y2="3.1484" width="0.254" layer="21"/>
<wire x1="-2.0984" y1="3.1484" x2="-2.0984" y2="-3.1484" width="0.254" layer="21"/>
<smd name="1" x="-3.25" y="2.925" dx="1.5" dy="0.35" layer="1"/>
<smd name="2" x="-3.25" y="2.275" dx="1.5" dy="0.35" layer="1"/>
<smd name="3" x="-3.25" y="1.625" dx="1.5" dy="0.35" layer="1"/>
<smd name="4" x="-3.25" y="0.975" dx="1.5" dy="0.35" layer="1"/>
<smd name="5" x="-3.25" y="0.325" dx="1.5" dy="0.35" layer="1"/>
<smd name="6" x="-3.25" y="-0.325" dx="1.5" dy="0.35" layer="1"/>
<smd name="7" x="-3.25" y="-0.975" dx="1.5" dy="0.35" layer="1"/>
<smd name="8" x="-3.25" y="-1.625" dx="1.5" dy="0.35" layer="1"/>
<smd name="9" x="-3.25" y="-2.275" dx="1.5" dy="0.35" layer="1"/>
<smd name="10" x="-3.25" y="-2.925" dx="1.5" dy="0.35" layer="1"/>
<smd name="11" x="3.25" y="-2.925" dx="1.5" dy="0.35" layer="1"/>
<smd name="12" x="3.25" y="-2.275" dx="1.5" dy="0.35" layer="1"/>
<smd name="13" x="3.25" y="-1.625" dx="1.5" dy="0.35" layer="1"/>
<smd name="14" x="3.25" y="-0.975" dx="1.5" dy="0.35" layer="1"/>
<smd name="15" x="3.25" y="-0.325" dx="1.5" dy="0.35" layer="1"/>
<smd name="16" x="3.25" y="0.325" dx="1.5" dy="0.35" layer="1"/>
<smd name="17" x="3.25" y="0.975" dx="1.5" dy="0.35" layer="1"/>
<smd name="18" x="3.25" y="1.625" dx="1.5" dy="0.35" layer="1"/>
<smd name="19" x="3.25" y="2.275" dx="1.5" dy="0.35" layer="1"/>
<smd name="20" x="3.25" y="2.925" dx="1.5" dy="0.35" layer="1"/>
<text x="-0.635" y="1.93" size="1.27" layer="25" font="vector" ratio="20" rot="R270">&gt;NAME</text>
<text x="-3.175" y="-5.08" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-3.2" y1="2.8" x2="-2.2" y2="3.05" layer="51"/>
<rectangle x1="-3.2" y1="2.15" x2="-2.2" y2="2.4" layer="51"/>
<rectangle x1="-3.2" y1="1.5" x2="-2.2" y2="1.75" layer="51"/>
<rectangle x1="-3.2" y1="0.85" x2="-2.2" y2="1.1" layer="51"/>
<rectangle x1="-3.2" y1="0.2" x2="-2.2" y2="0.45" layer="51"/>
<rectangle x1="-3.2" y1="-0.45" x2="-2.2" y2="-0.2" layer="51"/>
<rectangle x1="-3.2" y1="-1.1" x2="-2.2" y2="-0.85" layer="51"/>
<rectangle x1="-3.2" y1="-1.75" x2="-2.2" y2="-1.5" layer="51"/>
<rectangle x1="-3.2" y1="-2.4" x2="-2.2" y2="-2.15" layer="51"/>
<rectangle x1="-3.2" y1="-3.05" x2="-2.2" y2="-2.8" layer="51"/>
<rectangle x1="2.2" y1="-3.05" x2="3.2" y2="-2.8" layer="51"/>
<rectangle x1="2.2" y1="-2.4" x2="3.2" y2="-2.15" layer="51"/>
<rectangle x1="2.2" y1="-1.75" x2="3.2" y2="-1.5" layer="51"/>
<rectangle x1="2.2" y1="-1.1" x2="3.2" y2="-0.85" layer="51"/>
<rectangle x1="2.2" y1="-0.45" x2="3.2" y2="-0.2" layer="51"/>
<rectangle x1="2.2" y1="0.2" x2="3.2" y2="0.45" layer="51"/>
<rectangle x1="2.2" y1="0.85" x2="3.2" y2="1.1" layer="51"/>
<rectangle x1="2.2" y1="1.5" x2="3.2" y2="1.75" layer="51"/>
<rectangle x1="2.2" y1="2.15" x2="3.2" y2="2.4" layer="51"/>
<rectangle x1="2.2" y1="2.8" x2="3.2" y2="3.05" layer="51"/>
</package>
<package name="SMD_LQFP48">
<description>&lt;b&gt;&lt;/b&gt;&lt;p&gt;
Auto generated by &lt;i&gt;make-symbol-device-package-bsdl.ulp Rev. 34&lt;/i&gt;&lt;br&gt;</description>
<wire x1="-3.3984" y1="2.5" x2="-2.5" y2="3.3984" width="0.254" layer="21"/>
<wire x1="-3.3984" y1="-3.3984" x2="3.3984" y2="-3.3984" width="0.254" layer="21"/>
<wire x1="3.3984" y1="-3.3984" x2="3.3984" y2="3.3984" width="0.254" layer="21"/>
<wire x1="3.3984" y1="3.3984" x2="-3.3984" y2="3.3984" width="0.254" layer="21"/>
<wire x1="-3.3984" y1="3.3984" x2="-3.3984" y2="-3.3984" width="0.254" layer="21"/>
<smd name="1" x="-4.55" y="2.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="2" x="-4.55" y="2.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="3" x="-4.55" y="1.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="4" x="-4.55" y="1.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="5" x="-4.55" y="0.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="6" x="-4.55" y="0.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="7" x="-4.55" y="-0.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="8" x="-4.55" y="-0.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="9" x="-4.55" y="-1.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="10" x="-4.55" y="-1.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="11" x="-4.55" y="-2.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="12" x="-4.55" y="-2.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="13" x="-2.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="14" x="-2.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="15" x="-1.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="16" x="-1.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="17" x="-0.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="18" x="-0.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="19" x="0.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="20" x="0.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="21" x="1.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="22" x="1.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="23" x="2.25" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="24" x="2.75" y="-4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="25" x="4.55" y="-2.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="26" x="4.55" y="-2.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="27" x="4.55" y="-1.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="28" x="4.55" y="-1.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="29" x="4.55" y="-0.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="30" x="4.55" y="-0.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="31" x="4.55" y="0.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="32" x="4.55" y="0.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="33" x="4.55" y="1.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="34" x="4.55" y="1.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="35" x="4.55" y="2.25" dx="1.5" dy="0.25" layer="1"/>
<smd name="36" x="4.55" y="2.75" dx="1.5" dy="0.25" layer="1"/>
<smd name="37" x="2.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="38" x="2.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="39" x="1.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="40" x="1.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="41" x="0.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="42" x="0.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="43" x="-0.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="44" x="-0.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="45" x="-1.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="46" x="-1.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="47" x="-2.25" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<smd name="48" x="-2.75" y="4.55" dx="0.25" dy="1.5" layer="1"/>
<text x="-2.5" y="-0.365" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="-2.54" y="-1.905" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<rectangle x1="-4.5" y1="2.64" x2="-3.5" y2="2.86" layer="51"/>
<rectangle x1="-4.5" y1="2.14" x2="-3.5" y2="2.36" layer="51"/>
<rectangle x1="-4.5" y1="1.64" x2="-3.5" y2="1.86" layer="51"/>
<rectangle x1="-4.5" y1="1.14" x2="-3.5" y2="1.36" layer="51"/>
<rectangle x1="-4.5" y1="0.64" x2="-3.5" y2="0.86" layer="51"/>
<rectangle x1="-4.5" y1="0.14" x2="-3.5" y2="0.36" layer="51"/>
<rectangle x1="-4.5" y1="-0.36" x2="-3.5" y2="-0.14" layer="51"/>
<rectangle x1="-4.5" y1="-0.86" x2="-3.5" y2="-0.64" layer="51"/>
<rectangle x1="-4.5" y1="-1.36" x2="-3.5" y2="-1.14" layer="51"/>
<rectangle x1="-4.5" y1="-1.86" x2="-3.5" y2="-1.64" layer="51"/>
<rectangle x1="-4.5" y1="-2.36" x2="-3.5" y2="-2.14" layer="51"/>
<rectangle x1="-4.5" y1="-2.86" x2="-3.5" y2="-2.64" layer="51"/>
<rectangle x1="-2.86" y1="-4.5" x2="-2.64" y2="-3.5" layer="51"/>
<rectangle x1="-2.36" y1="-4.5" x2="-2.14" y2="-3.5" layer="51"/>
<rectangle x1="-1.86" y1="-4.5" x2="-1.64" y2="-3.5" layer="51"/>
<rectangle x1="-1.36" y1="-4.5" x2="-1.14" y2="-3.5" layer="51"/>
<rectangle x1="-0.86" y1="-4.5" x2="-0.64" y2="-3.5" layer="51"/>
<rectangle x1="-0.36" y1="-4.5" x2="-0.14" y2="-3.5" layer="51"/>
<rectangle x1="0.14" y1="-4.5" x2="0.36" y2="-3.5" layer="51"/>
<rectangle x1="0.64" y1="-4.5" x2="0.86" y2="-3.5" layer="51"/>
<rectangle x1="1.14" y1="-4.5" x2="1.36" y2="-3.5" layer="51"/>
<rectangle x1="1.64" y1="-4.5" x2="1.86" y2="-3.5" layer="51"/>
<rectangle x1="2.14" y1="-4.5" x2="2.36" y2="-3.5" layer="51"/>
<rectangle x1="2.64" y1="-4.5" x2="2.86" y2="-3.5" layer="51"/>
<rectangle x1="3.5" y1="-2.86" x2="4.5" y2="-2.64" layer="51"/>
<rectangle x1="3.5" y1="-2.36" x2="4.5" y2="-2.14" layer="51"/>
<rectangle x1="3.5" y1="-1.86" x2="4.5" y2="-1.64" layer="51"/>
<rectangle x1="3.5" y1="-1.36" x2="4.5" y2="-1.14" layer="51"/>
<rectangle x1="3.5" y1="-0.86" x2="4.5" y2="-0.64" layer="51"/>
<rectangle x1="3.5" y1="-0.36" x2="4.5" y2="-0.14" layer="51"/>
<rectangle x1="3.5" y1="0.14" x2="4.5" y2="0.36" layer="51"/>
<rectangle x1="3.5" y1="0.64" x2="4.5" y2="0.86" layer="51"/>
<rectangle x1="3.5" y1="1.14" x2="4.5" y2="1.36" layer="51"/>
<rectangle x1="3.5" y1="1.64" x2="4.5" y2="1.86" layer="51"/>
<rectangle x1="3.5" y1="2.14" x2="4.5" y2="2.36" layer="51"/>
<rectangle x1="3.5" y1="2.64" x2="4.5" y2="2.86" layer="51"/>
<rectangle x1="2.64" y1="3.5" x2="2.86" y2="4.5" layer="51"/>
<rectangle x1="2.14" y1="3.5" x2="2.36" y2="4.5" layer="51"/>
<rectangle x1="1.64" y1="3.5" x2="1.86" y2="4.5" layer="51"/>
<rectangle x1="1.14" y1="3.5" x2="1.36" y2="4.5" layer="51"/>
<rectangle x1="0.64" y1="3.5" x2="0.86" y2="4.5" layer="51"/>
<rectangle x1="0.14" y1="3.5" x2="0.36" y2="4.5" layer="51"/>
<rectangle x1="-0.36" y1="3.5" x2="-0.14" y2="4.5" layer="51"/>
<rectangle x1="-0.86" y1="3.5" x2="-0.64" y2="4.5" layer="51"/>
<rectangle x1="-1.36" y1="3.5" x2="-1.14" y2="4.5" layer="51"/>
<rectangle x1="-1.86" y1="3.5" x2="-1.64" y2="4.5" layer="51"/>
<rectangle x1="-2.36" y1="3.5" x2="-2.14" y2="4.5" layer="51"/>
<rectangle x1="-2.86" y1="3.5" x2="-2.64" y2="4.5" layer="51"/>
</package>
<package name="SMD_LED_7SEG">
<wire x1="-3.6" y1="5.2" x2="3.6" y2="5.2" width="0.254" layer="21"/>
<wire x1="-3.6" y1="-5.2" x2="3.6" y2="-5.2" width="0.254" layer="21"/>
<wire x1="-3.6" y1="-5.2" x2="-3.6" y2="5.2" width="0.254" layer="21"/>
<wire x1="3.6" y1="5.2" x2="3.6" y2="-5.2" width="0.254" layer="21"/>
<wire x1="-1.65" y1="4" x2="-2.35" y2="-4" width="0.127" layer="51"/>
<wire x1="2.35" y1="4" x2="1.65" y2="-4" width="0.127" layer="51"/>
<wire x1="-1.65" y1="4" x2="-0.65" y2="4" width="0.127" layer="51"/>
<wire x1="-0.65" y1="4" x2="1.35" y2="4" width="0.127" layer="51"/>
<wire x1="1.35" y1="4" x2="2.35" y2="4" width="0.127" layer="51"/>
<wire x1="-2.35" y1="-4" x2="-1.35" y2="-4" width="0.127" layer="51"/>
<wire x1="-1.35" y1="-4" x2="1.65" y2="-4" width="0.127" layer="51"/>
<wire x1="-1.75" y1="3" x2="2.25" y2="3" width="0.127" layer="51"/>
<wire x1="-2.25" y1="-3" x2="1.75" y2="-3" width="0.127" layer="51"/>
<wire x1="-0.65" y1="4" x2="-1.35" y2="-4" width="0.127" layer="51"/>
<wire x1="1.35" y1="4" x2="0.65" y2="-4" width="0.127" layer="51"/>
<wire x1="-1.95" y1="0.5" x2="2.05" y2="0.5" width="0.127" layer="51"/>
<wire x1="-2.05" y1="-0.5" x2="1.95" y2="-0.5" width="0.127" layer="51"/>
<circle x="2.7" y="-3.8" radius="0.2828" width="0.127" layer="51"/>
<circle x="2.5" y="-3.6" radius="0.2" width="0.4" layer="21"/>
<smd name="6" x="2.54" y="6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="7" x="1.27" y="6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="8" x="0" y="6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="9" x="-1.27" y="6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="10" x="-2.54" y="6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="5" x="2.54" y="-6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="4" x="1.27" y="-6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="3" x="0" y="-6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="2" x="-1.27" y="-6.25" dx="0.7" dy="1.5" layer="1"/>
<smd name="1" x="-2.54" y="-6.25" dx="0.7" dy="1.5" layer="1"/>
<text x="-3.175" y="3.175" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="-3.175" y="1.27" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
<polygon width="0.127" layer="21">
<vertex x="-0.56" y="0.5"/>
<vertex x="0.74" y="0.5"/>
<vertex x="1.15" y="0"/>
<vertex x="0.56" y="-0.5"/>
<vertex x="-0.74" y="-0.5"/>
<vertex x="-1.15" y="0"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="1.44" y="-0.35"/>
<vertex x="1.85" y="-0.85"/>
<vertex x="1.44" y="-3.15"/>
<vertex x="1.15" y="-3.4"/>
<vertex x="0.53" y="-2.65"/>
<vertex x="0.85" y="-0.85"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="-1.18" y="-3"/>
<vertex x="0.12" y="-3"/>
<vertex x="0.74" y="-3.75"/>
<vertex x="0.44" y="-4"/>
<vertex x="-1.86" y="-4"/>
<vertex x="-2.06" y="-3.75"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="-1.56" y="-0.35"/>
<vertex x="-1.15" y="-0.85"/>
<vertex x="-1.47" y="-2.65"/>
<vertex x="-2.35" y="-3.4"/>
<vertex x="-2.56" y="-3.15"/>
<vertex x="-2.15" y="-0.85"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="2.35" y="3.4"/>
<vertex x="2.56" y="3.15"/>
<vertex x="2.15" y="0.85"/>
<vertex x="1.56" y="0.35"/>
<vertex x="1.15" y="0.85"/>
<vertex x="1.47" y="2.65"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="-1.15" y="3.4"/>
<vertex x="-0.53" y="2.65"/>
<vertex x="-0.85" y="0.85"/>
<vertex x="-1.44" y="0.35"/>
<vertex x="-1.85" y="0.85"/>
<vertex x="-1.44" y="3.15"/>
</polygon>
<polygon width="0.127" layer="21">
<vertex x="-0.44" y="4"/>
<vertex x="1.86" y="4"/>
<vertex x="2.06" y="3.75"/>
<vertex x="1.18" y="3"/>
<vertex x="-0.12" y="3"/>
<vertex x="-0.74" y="3.75"/>
</polygon>
</package>
<package name="HOLE_32">
<wire x1="-2.159" y1="0" x2="0" y2="-2.159" width="2.4892" layer="51" curve="90" cap="flat"/>
<wire x1="0" y1="2.159" x2="2.159" y2="0" width="2.4892" layer="51" curve="-90" cap="flat"/>
<circle x="0" y="0" radius="3.5" width="0.254" layer="21"/>
<circle x="0" y="0" radius="0.762" width="0.4572" layer="51"/>
<circle x="0" y="0" radius="3.683" width="1.27" layer="39"/>
<circle x="0" y="0" radius="3.683" width="1.27" layer="40"/>
<circle x="0" y="0" radius="3.556" width="1.016" layer="43"/>
<pad name="B3,2" x="0" y="0" drill="3.2" diameter="5.5"/>
<text x="-1.27" y="-3.81" size="1.27" layer="48">3.2</text>
</package>
<package name="PAD_JMP_2P">
<wire x1="-1.27" y1="1.27" x2="1.27" y2="1.27" width="0.254" layer="21"/>
<wire x1="1.27" y1="1.27" x2="1.27" y2="-3.81" width="0.254" layer="21"/>
<wire x1="1.27" y1="-3.81" x2="-1.27" y2="-3.81" width="0.254" layer="21"/>
<wire x1="-1.27" y1="-3.81" x2="-1.27" y2="1.27" width="0.254" layer="21"/>
<pad name="P$1" x="0" y="-2.54" drill="0.8" diameter="1.8"/>
<pad name="P$2" x="0" y="0" drill="0.8" diameter="1.8"/>
<text x="1.27" y="0" size="1.27" layer="25" font="vector" ratio="20">&gt;NAME</text>
<text x="1.27" y="-1.27" size="1.27" layer="27" font="vector" ratio="20">&gt;VALUE</text>
</package>
</packages>
<symbols>
<symbol name="VCCAUX">
<wire x1="0" y1="0" x2="0" y2="2.54" width="0.254" layer="94"/>
<wire x1="-2.54" y1="2.54" x2="2.54" y2="2.54" width="0.254" layer="94"/>
<wire x1="2.54" y1="2.54" x2="0" y2="7.62" width="0.254" layer="94"/>
<wire x1="0" y1="7.62" x2="-2.54" y2="2.54" width="0.254" layer="94"/>
<text x="2.54" y="0" size="1.778" layer="96">&gt;VALUE</text>
<pin name="VCCAUX" x="0" y="0" visible="pad" length="point" direction="sup" rot="R90"/>
</symbol>
<symbol name="VCCINT">
<wire x1="0" y1="0" x2="0" y2="2.54" width="0.254" layer="94"/>
<wire x1="0" y1="2.54" x2="-2.54" y2="2.54" width="0.254" layer="94"/>
<wire x1="-2.54" y1="2.54" x2="0" y2="7.62" width="0.254" layer="94"/>
<wire x1="0" y1="7.62" x2="2.54" y2="2.54" width="0.254" layer="94"/>
<wire x1="2.54" y1="2.54" x2="0" y2="2.54" width="0.254" layer="94"/>
<text x="2.54" y="0" size="1.778" layer="96">&gt;VALUE</text>
<pin name="VCCINT" x="0" y="0" visible="pad" length="point" direction="sup" rot="R90"/>
</symbol>
<symbol name="VCCO">
<wire x1="0" y1="0" x2="0" y2="2.54" width="0.254" layer="94"/>
<wire x1="0" y1="2.54" x2="-2.54" y2="2.54" width="0.254" layer="94"/>
<wire x1="-2.54" y1="2.54" x2="0" y2="7.62" width="0.254" layer="94"/>
<wire x1="0" y1="7.62" x2="2.54" y2="2.54" width="0.254" layer="94"/>
<wire x1="2.54" y1="2.54" x2="0" y2="2.54" width="0.254" layer="94"/>
<text x="2.54" y="0" size="1.778" layer="96">&gt;VALUE</text>
<pin name="VCCO" x="0" y="0" visible="pad" length="point" direction="sup" rot="R90"/>
</symbol>
<symbol name="UNIT_M51957B">
<wire x1="-10.16" y1="-12.7" x2="10.16" y2="-12.7" width="0.254" layer="94"/>
<wire x1="10.16" y1="-12.7" x2="10.16" y2="0" width="0.254" layer="94"/>
<wire x1="10.16" y1="0" x2="-10.16" y2="0" width="0.254" layer="94"/>
<wire x1="-10.16" y1="0" x2="-10.16" y2="-12.7" width="0.254" layer="94"/>
<text x="-7.62" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="-7.62" y="-15.24" size="1.778" layer="96">&gt;VALUE</text>
<pin name="NC@1" x="-15.24" y="-2.54" length="middle"/>
<pin name="IN" x="-15.24" y="-5.08" length="middle"/>
<pin name="NC@2" x="-15.24" y="-7.62" length="middle"/>
<pin name="GND" x="-15.24" y="-10.16" length="middle" direction="pwr"/>
<pin name="DLY" x="15.24" y="-10.16" length="middle" rot="R180"/>
<pin name="OUT" x="15.24" y="-7.62" length="middle" rot="R180"/>
<pin name="VCCO" x="15.24" y="-5.08" length="middle" direction="pwr" rot="R180"/>
<pin name="NC@3" x="15.24" y="-2.54" length="middle" rot="R180"/>
</symbol>
<symbol name="GND">
<wire x1="0" y1="0" x2="0" y2="-2.54" width="0.254" layer="94"/>
<wire x1="2.54" y1="-2.54" x2="-2.54" y2="-2.54" width="0.254" layer="94"/>
<wire x1="-2.54" y1="-2.54" x2="0" y2="-7.62" width="0.254" layer="94"/>
<wire x1="0" y1="-7.62" x2="2.54" y2="-2.54" width="0.254" layer="94"/>
<text x="2.54" y="0" size="1.778" layer="96">&gt;VALUE</text>
<pin name="GND" x="0" y="0" visible="pad" length="point" direction="sup" rot="R270"/>
</symbol>
<symbol name="LED">
<wire x1="1.27" y1="0" x2="0" y2="-2.54" width="0.254" layer="94"/>
<wire x1="0" y1="-2.54" x2="-1.27" y2="0" width="0.254" layer="94"/>
<wire x1="1.27" y1="-2.54" x2="0" y2="-2.54" width="0.254" layer="94"/>
<wire x1="0" y1="-2.54" x2="-1.27" y2="-2.54" width="0.254" layer="94"/>
<wire x1="1.27" y1="0" x2="0" y2="0" width="0.254" layer="94"/>
<wire x1="0" y1="0" x2="-1.27" y2="0" width="0.254" layer="94"/>
<wire x1="0" y1="0" x2="0" y2="-2.54" width="0.1524" layer="94"/>
<wire x1="-2.032" y1="-0.762" x2="-3.429" y2="-2.159" width="0.1524" layer="94"/>
<wire x1="-1.905" y1="-1.905" x2="-3.302" y2="-3.302" width="0.1524" layer="94"/>
<text x="3.81" y="-3.81" size="1.778" layer="95" rot="R90">&gt;NAME</text>
<text x="6.35" y="-3.81" size="1.778" layer="97" rot="R90">&gt;VALUE</text>
<pin name="C" x="0" y="-5.08" visible="off" length="short" direction="pas" rot="R90"/>
<pin name="A" x="0" y="2.54" visible="off" length="short" direction="pas" rot="R270"/>
<polygon width="0.1524" layer="94">
<vertex x="-3.429" y="-2.159"/>
<vertex x="-3.048" y="-1.27"/>
<vertex x="-2.54" y="-1.778"/>
</polygon>
<polygon width="0.1524" layer="94">
<vertex x="-3.302" y="-3.302"/>
<vertex x="-2.921" y="-2.413"/>
<vertex x="-2.413" y="-2.921"/>
</polygon>
</symbol>
<symbol name="REG">
<wire x1="0" y1="0" x2="1.016" y2="0" width="0.254" layer="94"/>
<wire x1="5.08" y1="0" x2="4.064" y2="0" width="0.254" layer="94"/>
<wire x1="1.016" y1="0" x2="1.524" y2="0.508" width="0.254" layer="94"/>
<wire x1="1.524" y1="0.508" x2="2.032" y2="-0.508" width="0.254" layer="94"/>
<wire x1="2.032" y1="-0.508" x2="2.54" y2="0.508" width="0.254" layer="94"/>
<wire x1="2.54" y1="0.508" x2="3.048" y2="-0.508" width="0.254" layer="94"/>
<wire x1="3.048" y1="-0.508" x2="3.556" y2="0.508" width="0.254" layer="94"/>
<wire x1="3.556" y1="0.508" x2="4.064" y2="0" width="0.254" layer="94"/>
<text x="-5.08" y="0" size="1.778" layer="95">&gt;NAME</text>
<text x="5.08" y="0" size="1.778" layer="95">&gt;VALUE</text>
<pin name="1" x="0" y="0" visible="off" length="point"/>
<pin name="2" x="5.08" y="0" visible="off" length="point" rot="R180"/>
</symbol>
<symbol name="UNIT_XCF02S">
<wire x1="-17.78" y1="-48.26" x2="17.78" y2="-48.26" width="0.254" layer="94"/>
<wire x1="17.78" y1="-48.26" x2="17.78" y2="0" width="0.254" layer="94"/>
<wire x1="17.78" y1="0" x2="-17.78" y2="0" width="0.254" layer="94"/>
<wire x1="-17.78" y1="0" x2="-17.78" y2="-48.26" width="0.254" layer="94"/>
<text x="-17.78" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="2.54" y="-50.8" size="1.778" layer="96">&gt;VALUE</text>
<pin name="CLK" x="22.86" y="-7.62" length="middle" rot="R180"/>
<pin name="CE_" x="22.86" y="-10.16" length="middle" function="dot" rot="R180"/>
<pin name="CEO_" x="22.86" y="-12.7" length="middle" function="dot" rot="R180"/>
<pin name="OE/NRST" x="22.86" y="-15.24" length="middle" rot="R180"/>
<pin name="CF_" x="22.86" y="-17.78" length="middle" rot="R180"/>
<pin name="D0" x="22.86" y="-20.32" length="middle" rot="R180"/>
<pin name="NC1" x="22.86" y="-27.94" length="middle" direction="nc" rot="R180"/>
<pin name="NC2" x="22.86" y="-30.48" length="middle" direction="nc" rot="R180"/>
<pin name="NC3" x="22.86" y="-33.02" length="middle" direction="nc" rot="R180"/>
<pin name="NC4" x="22.86" y="-35.56" length="middle" direction="nc" rot="R180"/>
<pin name="NC5" x="22.86" y="-38.1" length="middle" direction="nc" rot="R180"/>
<pin name="NC6" x="22.86" y="-40.64" length="middle" direction="nc" rot="R180"/>
<pin name="VCCI" x="-22.86" y="-7.62" length="middle" direction="pwr"/>
<pin name="VCCO" x="-22.86" y="-12.7" length="middle" direction="pwr"/>
<pin name="VCCJ" x="-22.86" y="-17.78" length="middle" direction="pwr"/>
<pin name="TCK" x="-22.86" y="-27.94" length="middle"/>
<pin name="TDI" x="-22.86" y="-30.48" length="middle"/>
<pin name="TMS" x="-22.86" y="-33.02" length="middle"/>
<pin name="TDO" x="-22.86" y="-35.56" length="middle"/>
<pin name="GND1" x="-22.86" y="-40.64" length="middle" direction="pwr"/>
</symbol>
<symbol name="CAP">
<wire x1="-1.27" y1="-0.889" x2="0" y2="-0.889" width="0.254" layer="94"/>
<wire x1="0" y1="-0.889" x2="1.27" y2="-0.889" width="0.254" layer="94"/>
<wire x1="-1.27" y1="-1.524" x2="0" y2="-1.524" width="0.254" layer="94"/>
<wire x1="0" y1="-1.524" x2="1.27" y2="-1.524" width="0.254" layer="94"/>
<wire x1="0" y1="-1.524" x2="0" y2="-2.54" width="0.254" layer="94"/>
<wire x1="0" y1="-0.889" x2="0" y2="0" width="0.254" layer="94"/>
<text x="1.27" y="-2.54" size="1.778" layer="96">&gt;VALUE</text>
<text x="1.27" y="0" size="1.778" layer="95">&gt;NAME</text>
<pin name="1" x="0" y="0" visible="off" length="point" rot="R270"/>
<pin name="2" x="0" y="-2.54" visible="off" length="point" rot="R90"/>
</symbol>
<symbol name="CON_POW_BRD">
<wire x1="0" y1="12.7" x2="0" y2="-15.24" width="0.254" layer="94"/>
<wire x1="0" y1="-15.24" x2="-15.24" y2="-15.24" width="0.254" layer="94"/>
<wire x1="-15.24" y1="-15.24" x2="-15.24" y2="12.7" width="0.254" layer="94"/>
<wire x1="-15.24" y1="12.7" x2="0" y2="12.7" width="0.254" layer="94"/>
<text x="-15.24" y="12.7" size="1.778" layer="95">&gt;NAME</text>
<text x="-15.24" y="-17.78" size="1.778" layer="96">&gt;VALUE</text>
<pin name="VCCO@1" x="2.54" y="10.16" length="short" direction="pwr" rot="R180"/>
<pin name="VCCO@2" x="2.54" y="7.62" length="short" direction="pwr" rot="R180"/>
<pin name="VCCAUX@1" x="2.54" y="5.08" length="short" direction="pwr" rot="R180"/>
<pin name="VCCAUX@2" x="2.54" y="2.54" length="short" direction="pwr" rot="R180"/>
<pin name="VCCINT@1" x="2.54" y="0" length="short" direction="pwr" rot="R180"/>
<pin name="VCCINT@2" x="2.54" y="-2.54" length="short" direction="pwr" rot="R180"/>
<pin name="GND@1" x="2.54" y="-5.08" length="short" direction="pwr" rot="R180"/>
<pin name="GND@2" x="2.54" y="-7.62" length="short" direction="pwr" rot="R180"/>
<pin name="GND@3" x="2.54" y="-10.16" length="short" direction="pwr" rot="R180"/>
<pin name="GND@4" x="2.54" y="-12.7" length="short" direction="pwr" rot="R180"/>
</symbol>
<symbol name="CON_USB">
<wire x1="0" y1="7.62" x2="0" y2="-7.62" width="0.254" layer="94"/>
<wire x1="0" y1="-7.62" x2="2.54" y2="-7.62" width="0.254" layer="94"/>
<wire x1="2.54" y1="-7.62" x2="5.08" y2="-7.62" width="0.254" layer="94"/>
<wire x1="5.08" y1="-7.62" x2="5.08" y2="7.62" width="0.254" layer="94"/>
<wire x1="5.08" y1="7.62" x2="0" y2="7.62" width="0.254" layer="94"/>
<wire x1="2.54" y1="-7.62" x2="2.54" y2="-17.78" width="0.254" layer="94"/>
<text x="-5.08" y="10.16" size="1.778" layer="97">&gt;NAME</text>
<text x="-5.08" y="7.62" size="1.778" layer="96">&gt;VALUE</text>
<pin name="1" x="-2.54" y="5.08" visible="pin" length="short"/>
<pin name="2" x="-2.54" y="2.54" visible="pin" length="short"/>
<pin name="3" x="-2.54" y="0" visible="pin" length="short"/>
<pin name="4" x="-2.54" y="-2.54" visible="pin" length="short"/>
<pin name="5" x="-2.54" y="-5.08" visible="pin" length="short"/>
<pin name="PAD1" x="0" y="-10.16" visible="off" length="short" direction="pwr"/>
<pin name="PAD2" x="0" y="-12.7" visible="off" length="short" direction="pwr"/>
<pin name="PAD3" x="0" y="-15.24" visible="off" length="short" direction="pwr"/>
<pin name="PAD4" x="0" y="-17.78" visible="off" length="short" direction="pwr"/>
</symbol>
<symbol name="CON_VPORT">
<wire x1="-5.08" y1="12.7" x2="5.08" y2="12.7" width="0.254" layer="94"/>
<wire x1="5.08" y1="12.7" x2="5.08" y2="-15.24" width="0.254" layer="94"/>
<wire x1="5.08" y1="-15.24" x2="-5.08" y2="-15.24" width="0.254" layer="94"/>
<wire x1="-5.08" y1="-15.24" x2="-5.08" y2="12.7" width="0.254" layer="94"/>
<text x="-5.08" y="12.7" size="1.778" layer="95">&gt;NAME</text>
<text x="-5.08" y="-17.78" size="1.778" layer="97">&gt;VALUE</text>
<pin name="1" x="-7.62" y="10.16" visible="pin" length="short" direction="pwr"/>
<pin name="2" x="-7.62" y="7.62" visible="pin" length="short" direction="pwr"/>
<pin name="3" x="-7.62" y="5.08" visible="pin" length="short" direction="pwr"/>
<pin name="4" x="-7.62" y="2.54" visible="pin" length="short" direction="pwr"/>
<pin name="5" x="-7.62" y="0" visible="pin" length="short" direction="pwr"/>
<pin name="6" x="-7.62" y="-2.54" visible="pin" length="short" direction="pwr"/>
<pin name="7" x="-7.62" y="-5.08" visible="pin" length="short" direction="pwr"/>
<pin name="8" x="-7.62" y="-7.62" visible="pin" length="short" direction="pwr"/>
<pin name="9" x="-7.62" y="-10.16" visible="pin" length="short" direction="pwr"/>
<pin name="10" x="-7.62" y="-12.7" visible="pin" length="short" direction="pwr"/>
</symbol>
<symbol name="FB">
<wire x1="0" y1="0" x2="1.27" y2="0" width="0.254" layer="94"/>
<wire x1="7.62" y1="0" x2="6.35" y2="0" width="0.254" layer="94"/>
<wire x1="1.27" y1="1.27" x2="1.27" y2="-1.27" width="0.254" layer="94"/>
<wire x1="1.27" y1="-1.27" x2="6.35" y2="-1.27" width="0.254" layer="94"/>
<wire x1="6.35" y1="-1.27" x2="6.35" y2="1.27" width="0.254" layer="94"/>
<wire x1="6.35" y1="1.27" x2="1.27" y2="1.27" width="0.254" layer="94"/>
<text x="0" y="2.54" size="1.778" layer="95">&gt;NAME</text>
<text x="0" y="-3.81" size="1.778" layer="97">&gt;VALUE</text>
<pin name="P$1" x="0" y="0" visible="off" length="point"/>
<pin name="P$2" x="7.62" y="0" visible="off" length="point" rot="R180"/>
</symbol>
<symbol name="HOLE">
<wire x1="0.254" y1="2.032" x2="2.032" y2="0.254" width="1.016" layer="94" curve="-75.749967" cap="flat"/>
<wire x1="-2.032" y1="0.254" x2="-0.254" y2="2.032" width="1.016" layer="94" curve="-75.749967" cap="flat"/>
<wire x1="-2.032" y1="-0.254" x2="-0.254" y2="-2.032" width="1.016" layer="94" curve="75.749967" cap="flat"/>
<wire x1="0.254" y1="-2.032" x2="2.032" y2="-0.254" width="1.016" layer="94" curve="75.749967" cap="flat"/>
<circle x="0" y="0" radius="1.524" width="0.0508" layer="94"/>
<text x="2.794" y="0.5842" size="1.778" layer="95">&gt;NAME</text>
<text x="2.794" y="-2.4638" size="1.778" layer="97">&gt;VALUE</text>
<pin name="MOUNT" x="-2.54" y="0" visible="off" length="short" direction="pas"/>
</symbol>
<symbol name="LED_7SEG">
<wire x1="0" y1="0" x2="0" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="2.54" y1="0" x2="2.54" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="5.08" y1="0" x2="5.08" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="-2.54" y1="0" x2="-2.54" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="-5.08" y1="0" x2="-5.08" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="-7.62" y1="0" x2="-7.62" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="7.62" y1="0" x2="7.62" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="10.16" y1="0" x2="10.16" y2="-7.62" width="0.1524" layer="94"/>
<wire x1="-7.62" y1="0" x2="0" y2="0" width="0.1524" layer="94"/>
<wire x1="0" y1="0" x2="10.16" y2="0" width="0.1524" layer="94"/>
<wire x1="0" y1="0" x2="0" y2="5.08" width="0.1524" layer="94"/>
<wire x1="0" y1="5.08" x2="0" y2="10.16" width="0.1524" layer="94"/>
<wire x1="0" y1="5.08" x2="2.54" y2="5.08" width="0.1524" layer="94"/>
<wire x1="2.54" y1="5.08" x2="2.54" y2="10.16" width="0.1524" layer="94"/>
<circle x="0" y="0" radius="0.2794" width="0.4572" layer="94"/>
<circle x="0" y="5.08" radius="0.2794" width="0.4572" layer="94"/>
<circle x="-5.08" y="0" radius="0.2794" width="0.4572" layer="94"/>
<circle x="-2.54" y="0" radius="0.2794" width="0.4572" layer="94"/>
<circle x="2.54" y="0" radius="0.2794" width="0.4572" layer="94"/>
<circle x="5.08" y="0" radius="0.2794" width="0.4572" layer="94"/>
<circle x="7.62" y="0" radius="0.2794" width="0.4572" layer="94"/>
<text x="5.08" y="5.08" size="1.27" layer="95">&gt;NAME</text>
<text x="5.08" y="2.54" size="1.27" layer="97">&gt;VALUE</text>
<pin name="A" x="-7.62" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="B" x="-5.08" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="C" x="-2.54" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="D" x="0" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="E" x="2.54" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="F" x="5.08" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="G" x="7.62" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="D.P" x="10.16" y="-7.62" visible="pad" length="point" rot="R90"/>
<pin name="COMMON8" x="0" y="10.16" visible="pad" length="point" rot="R270"/>
<pin name="COMMON3" x="2.54" y="10.16" visible="pad" length="point" rot="R270"/>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-0.762" y="-3.3655"/>
<vertex x="0.762" y="-3.3655"/>
<vertex x="0" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-0.762" y="-4.6355"/>
<vertex x="0.762" y="-4.6355"/>
<vertex x="0.762" y="-4.3815"/>
<vertex x="-0.762" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="4.318" y="-3.3655"/>
<vertex x="5.842" y="-3.3655"/>
<vertex x="5.08" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="6.858" y="-3.3655"/>
<vertex x="8.382" y="-3.3655"/>
<vertex x="7.62" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="9.398" y="-3.3655"/>
<vertex x="10.922" y="-3.3655"/>
<vertex x="10.16" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="1.778" y="-3.3655"/>
<vertex x="3.302" y="-3.3655"/>
<vertex x="2.54" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-3.302" y="-3.3655"/>
<vertex x="-1.778" y="-3.3655"/>
<vertex x="-2.54" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-5.842" y="-3.3655"/>
<vertex x="-4.318" y="-3.3655"/>
<vertex x="-5.08" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-8.382" y="-3.3655"/>
<vertex x="-6.858" y="-3.3655"/>
<vertex x="-7.62" y="-4.6355"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-3.302" y="-4.6355"/>
<vertex x="-1.778" y="-4.6355"/>
<vertex x="-1.778" y="-4.3815"/>
<vertex x="-3.302" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-5.842" y="-4.6355"/>
<vertex x="-4.318" y="-4.6355"/>
<vertex x="-4.318" y="-4.3815"/>
<vertex x="-5.842" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="-8.382" y="-4.6355"/>
<vertex x="-6.858" y="-4.6355"/>
<vertex x="-6.858" y="-4.3815"/>
<vertex x="-8.382" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="1.778" y="-4.6355"/>
<vertex x="3.302" y="-4.6355"/>
<vertex x="3.302" y="-4.3815"/>
<vertex x="1.778" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="4.318" y="-4.6355"/>
<vertex x="5.842" y="-4.6355"/>
<vertex x="5.842" y="-4.3815"/>
<vertex x="4.318" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="6.858" y="-4.6355"/>
<vertex x="8.382" y="-4.6355"/>
<vertex x="8.382" y="-4.3815"/>
<vertex x="6.858" y="-4.3815"/>
</polygon>
<polygon width="0" layer="94" spacing="0.254">
<vertex x="9.398" y="-4.6355"/>
<vertex x="10.922" y="-4.6355"/>
<vertex x="10.922" y="-4.3815"/>
<vertex x="9.398" y="-4.3815"/>
</polygon>
</symbol>
<symbol name="SW_PUSH">
<wire x1="0" y1="2.54" x2="0.508" y2="2.54" width="0.1524" layer="94"/>
<wire x1="0.508" y1="2.54" x2="0.508" y2="0" width="0.1524" layer="94"/>
<wire x1="0.508" y1="0" x2="0.508" y2="-2.54" width="0.1524" layer="94"/>
<wire x1="0.508" y1="-2.54" x2="0" y2="-2.54" width="0.1524" layer="94"/>
<wire x1="5.08" y1="2.54" x2="4.572" y2="2.54" width="0.1524" layer="94"/>
<wire x1="4.572" y1="2.54" x2="4.572" y2="0" width="0.1524" layer="94"/>
<wire x1="4.572" y1="0" x2="4.572" y2="-2.54" width="0.1524" layer="94"/>
<wire x1="4.572" y1="-2.54" x2="5.08" y2="-2.54" width="0.1524" layer="94"/>
<wire x1="1.27" y1="0" x2="0.508" y2="0" width="0.1524" layer="94"/>
<wire x1="3.81" y1="0" x2="4.572" y2="0" width="0.1524" layer="94"/>
<wire x1="1.778" y1="0.1778" x2="3.302" y2="0.762" width="0.1524" layer="94"/>
<circle x="1.524" y="0" radius="0.254" width="0.1524" layer="94"/>
<circle x="3.556" y="0" radius="0.254" width="0.1524" layer="94"/>
<text x="-1.27" y="-5.08" size="1.778" layer="97">&gt;VALUE</text>
<text x="3.81" y="3.81" size="1.778" layer="95" rot="R90">&gt;NAME</text>
<pin name="1" x="0" y="2.54" visible="off" length="point"/>
<pin name="2" x="0" y="-2.54" visible="off" length="point"/>
<pin name="3" x="5.08" y="2.54" visible="off" length="point" rot="R180"/>
<pin name="4" x="5.08" y="-2.54" visible="off" length="point" rot="R180"/>
</symbol>
<symbol name="UNIT_74AC125">
<wire x1="-7.62" y1="0" x2="7.62" y2="0" width="0.254" layer="94"/>
<wire x1="7.62" y1="0" x2="7.62" y2="-5.08" width="0.254" layer="94"/>
<wire x1="7.62" y1="-5.08" x2="7.62" y2="-12.7" width="0.254" layer="94"/>
<wire x1="7.62" y1="-12.7" x2="7.62" y2="-20.32" width="0.254" layer="94"/>
<wire x1="7.62" y1="-20.32" x2="7.62" y2="-27.94" width="0.254" layer="94"/>
<wire x1="7.62" y1="-27.94" x2="7.62" y2="-35.56" width="0.254" layer="94"/>
<wire x1="7.62" y1="-35.56" x2="7.62" y2="-40.64" width="0.254" layer="94"/>
<wire x1="7.62" y1="-40.64" x2="-7.62" y2="-40.64" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-40.64" x2="-7.62" y2="-35.56" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-35.56" x2="-7.62" y2="-27.94" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-27.94" x2="-7.62" y2="-20.32" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-20.32" x2="-7.62" y2="-12.7" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-12.7" x2="-7.62" y2="-5.08" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-5.08" x2="-7.62" y2="0" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-5.08" x2="7.62" y2="-5.08" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-12.7" x2="7.62" y2="-12.7" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-20.32" x2="7.62" y2="-20.32" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-27.94" x2="7.62" y2="-27.94" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-35.56" x2="7.62" y2="-35.56" width="0.254" layer="94"/>
<text x="-7.62" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="-7.62" y="-43.18" size="1.778" layer="96">&gt;VALUE</text>
<pin name="VCCO" x="-12.7" y="-2.54" length="middle" direction="pwr"/>
<pin name="1G_" x="-12.7" y="-7.62" length="middle"/>
<pin name="1A" x="-12.7" y="-10.16" length="middle"/>
<pin name="2G_" x="-12.7" y="-15.24" length="middle"/>
<pin name="2A" x="-12.7" y="-17.78" length="middle"/>
<pin name="3G_" x="-12.7" y="-22.86" length="middle"/>
<pin name="3A" x="-12.7" y="-25.4" length="middle"/>
<pin name="4G_" x="-12.7" y="-30.48" length="middle"/>
<pin name="4A" x="-12.7" y="-33.02" length="middle"/>
<pin name="GND" x="-12.7" y="-38.1" length="middle"/>
<pin name="1Y" x="12.7" y="-10.16" length="middle" rot="R180"/>
<pin name="2Y" x="12.7" y="-17.78" length="middle" rot="R180"/>
<pin name="3Y" x="12.7" y="-25.4" length="middle" rot="R180"/>
<pin name="4Y" x="12.7" y="-33.02" length="middle" rot="R180"/>
</symbol>
<symbol name="UNIT_74HC14">
<wire x1="-7.62" y1="0" x2="7.62" y2="0" width="0.254" layer="94"/>
<wire x1="7.62" y1="0" x2="7.62" y2="-40.64" width="0.254" layer="94"/>
<wire x1="7.62" y1="-40.64" x2="-7.62" y2="-40.64" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-40.64" x2="-7.62" y2="0" width="0.254" layer="94"/>
<text x="-7.62" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="-7.62" y="-43.18" size="1.778" layer="96">&gt;VALUE</text>
<pin name="IN1" x="-12.7" y="-7.62" length="middle"/>
<pin name="IN2" x="-12.7" y="-12.7" length="middle"/>
<pin name="IN3" x="-12.7" y="-17.78" length="middle"/>
<pin name="IN4" x="-12.7" y="-22.86" length="middle"/>
<pin name="IN5" x="-12.7" y="-27.94" length="middle"/>
<pin name="IN6" x="-12.7" y="-33.02" length="middle"/>
<pin name="VCCO" x="-12.7" y="-2.54" length="middle" direction="pwr"/>
<pin name="GND" x="-12.7" y="-38.1" length="middle" direction="pwr"/>
<pin name="OUT6" x="12.7" y="-33.02" length="middle" rot="R180"/>
<pin name="OUT5" x="12.7" y="-27.94" length="middle" rot="R180"/>
<pin name="OUT4" x="12.7" y="-22.86" length="middle" rot="R180"/>
<pin name="OUT3" x="12.7" y="-17.78" length="middle" rot="R180"/>
<pin name="OUT2" x="12.7" y="-12.7" length="middle" rot="R180"/>
<pin name="OUT1" x="12.7" y="-7.62" length="middle" rot="R180"/>
</symbol>
<symbol name="UNIT_93CX6">
<wire x1="-7.62" y1="-12.7" x2="7.62" y2="-12.7" width="0.254" layer="94"/>
<wire x1="7.62" y1="-12.7" x2="7.62" y2="0" width="0.254" layer="94"/>
<wire x1="7.62" y1="0" x2="-7.62" y2="0" width="0.254" layer="94"/>
<wire x1="-7.62" y1="0" x2="-7.62" y2="-12.7" width="0.254" layer="94"/>
<text x="-7.62" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="-5.08" y="-15.24" size="1.778" layer="96">&gt;VALUE</text>
<pin name="CS" x="-12.7" y="-2.54" length="middle"/>
<pin name="SK" x="-12.7" y="-5.08" length="middle"/>
<pin name="DI" x="-12.7" y="-7.62" length="middle"/>
<pin name="DO" x="-12.7" y="-10.16" length="middle"/>
<pin name="GND" x="12.7" y="-10.16" length="middle" rot="R180"/>
<pin name="ORG" x="12.7" y="-7.62" length="middle" rot="R180"/>
<pin name="NC" x="12.7" y="-5.08" length="middle" rot="R180"/>
<pin name="VCC" x="12.7" y="-2.54" length="middle" rot="R180"/>
</symbol>
<symbol name="UNIT_CRYSTAL">
<wire x1="-2.54" y1="0" x2="2.54" y2="0" width="0.254" layer="94"/>
<wire x1="-2.54" y1="-5.08" x2="2.54" y2="-5.08" width="0.254" layer="94"/>
<wire x1="-2.54" y1="-0.762" x2="2.54" y2="-0.762" width="0.254" layer="94"/>
<wire x1="2.54" y1="-0.762" x2="2.54" y2="-4.318" width="0.254" layer="94"/>
<wire x1="2.54" y1="-4.318" x2="-2.54" y2="-4.318" width="0.254" layer="94"/>
<wire x1="-2.54" y1="-4.318" x2="-2.54" y2="-0.762" width="0.254" layer="94"/>
<text x="-3.81" y="-6.35" size="1.778" layer="95" rot="R90">&gt;NAME</text>
<text x="6.35" y="-6.35" size="1.778" layer="97" rot="R90">&gt;VALUE</text>
<pin name="P$1" x="0" y="2.54" visible="off" length="short" rot="R270"/>
<pin name="P$2" x="0" y="-7.62" visible="off" length="short" rot="R90"/>
</symbol>
<symbol name="UNIT_FT2232L">
<wire x1="-25.4" y1="-86.36" x2="25.4" y2="-86.36" width="0.254" layer="94"/>
<wire x1="25.4" y1="-86.36" x2="25.4" y2="0" width="0.254" layer="94"/>
<wire x1="25.4" y1="0" x2="-25.4" y2="0" width="0.254" layer="94"/>
<wire x1="-25.4" y1="0" x2="-25.4" y2="-86.36" width="0.254" layer="94"/>
<text x="-25.4" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="8.89" y="-88.9" size="1.778" layer="96">&gt;VALUE</text>
<pin name="ADBUS0" x="30.48" y="-2.54" length="middle" rot="R180"/>
<pin name="ADBUS1" x="30.48" y="-5.08" length="middle" rot="R180"/>
<pin name="ADBUS2" x="30.48" y="-7.62" length="middle" rot="R180"/>
<pin name="ADBUS3" x="30.48" y="-10.16" length="middle" rot="R180"/>
<pin name="ADBUS4" x="30.48" y="-12.7" length="middle" rot="R180"/>
<pin name="ADBUS5" x="30.48" y="-15.24" length="middle" rot="R180"/>
<pin name="ADBUS6" x="30.48" y="-17.78" length="middle" rot="R180"/>
<pin name="ADBUS7" x="30.48" y="-20.32" length="middle" rot="R180"/>
<pin name="ACBUS0" x="30.48" y="-25.4" length="middle" rot="R180"/>
<pin name="ACBUS1" x="30.48" y="-27.94" length="middle" rot="R180"/>
<pin name="ACBUS2" x="30.48" y="-30.48" length="middle" rot="R180"/>
<pin name="ACBUS3" x="30.48" y="-33.02" length="middle" rot="R180"/>
<pin name="SI/WUA" x="30.48" y="-35.56" length="middle" rot="R180"/>
<pin name="BDBUS0" x="30.48" y="-43.18" length="middle" rot="R180"/>
<pin name="BDBUS1" x="30.48" y="-45.72" length="middle" rot="R180"/>
<pin name="BDBUS2" x="30.48" y="-48.26" length="middle" rot="R180"/>
<pin name="BDBUS3" x="30.48" y="-50.8" length="middle" rot="R180"/>
<pin name="BDBUS4" x="30.48" y="-53.34" length="middle" rot="R180"/>
<pin name="BDBUS5" x="30.48" y="-55.88" length="middle" rot="R180"/>
<pin name="BDBUS6" x="30.48" y="-58.42" length="middle" rot="R180"/>
<pin name="BDBUS7" x="30.48" y="-60.96" length="middle" rot="R180"/>
<pin name="BCBUS0" x="30.48" y="-66.04" length="middle" rot="R180"/>
<pin name="BCBUS1" x="30.48" y="-68.58" length="middle" rot="R180"/>
<pin name="BCBUS2" x="30.48" y="-71.12" length="middle" rot="R180"/>
<pin name="BCBUS3" x="30.48" y="-73.66" length="middle" rot="R180"/>
<pin name="SI/WUB" x="30.48" y="-76.2" length="middle" rot="R180"/>
<pin name="PWREN_" x="30.48" y="-83.82" length="middle" rot="R180"/>
<pin name="AVCC" x="-30.48" y="-2.54" length="middle" direction="pwr"/>
<pin name="VCC@1" x="-30.48" y="-7.62" length="middle" direction="pwr"/>
<pin name="VCC@2" x="-30.48" y="-10.16" length="middle" direction="pwr"/>
<pin name="VCCIOA" x="-30.48" y="-15.24" length="middle" direction="pwr"/>
<pin name="VCCIOB" x="-30.48" y="-17.78" length="middle" direction="pwr"/>
<pin name="GND@4" x="-30.48" y="-83.82" length="middle" direction="pwr"/>
<pin name="GND@3" x="-30.48" y="-81.28" length="middle" direction="pwr"/>
<pin name="GND@2" x="-30.48" y="-78.74" length="middle" direction="pwr"/>
<pin name="GND@1" x="-30.48" y="-76.2" length="middle" direction="pwr"/>
<pin name="AGND" x="-30.48" y="-71.12" length="middle" direction="pwr"/>
<pin name="TEST" x="-30.48" y="-66.04" length="middle"/>
<pin name="EEDATA" x="-30.48" y="-63.5" length="middle"/>
<pin name="EESK" x="-30.48" y="-60.96" length="middle"/>
<pin name="EECS" x="-30.48" y="-58.42" length="middle"/>
<pin name="XTOUT" x="-30.48" y="-53.34" length="middle"/>
<pin name="XTIN" x="-30.48" y="-48.26" length="middle"/>
<pin name="RESET_" x="-30.48" y="-43.18" length="middle"/>
<pin name="RSTOUT_" x="-30.48" y="-40.64" length="middle"/>
<pin name="USBDP" x="-30.48" y="-35.56" length="middle"/>
<pin name="USBDM" x="-30.48" y="-30.48" length="middle"/>
<pin name="3V3OUT" x="-30.48" y="-25.4" length="middle"/>
</symbol>
<symbol name="UNIT_XC3S250E">
<wire x1="-25.4" y1="-172.72" x2="25.4" y2="-172.72" width="0.254" layer="94"/>
<wire x1="25.4" y1="-172.72" x2="25.4" y2="0" width="0.254" layer="94"/>
<wire x1="25.4" y1="0" x2="-25.4" y2="0" width="0.254" layer="94"/>
<wire x1="-25.4" y1="0" x2="-25.4" y2="-172.72" width="0.254" layer="94"/>
<text x="-25.4" y="1.27" size="1.778" layer="95">&gt;NAME</text>
<text x="6.35" y="-175.26" size="1.778" layer="96">&gt;VALUE</text>
<pin name="IO_L01P_3" x="30.48" y="-5.08" length="middle" rot="R180"/>
<pin name="IO_L01N_3" x="30.48" y="-7.62" length="middle" rot="R180"/>
<pin name="IO_L02P_3" x="30.48" y="-10.16" length="middle" rot="R180"/>
<pin name="IO_L02N_3/VREF_3" x="30.48" y="-12.7" length="middle" rot="R180"/>
<pin name="IO_L03P_3/LHCLK0" x="30.48" y="-15.24" length="middle" rot="R180"/>
<pin name="IO_L03N_3/LHCLK1" x="30.48" y="-17.78" length="middle" rot="R180"/>
<pin name="IO_L04P_3/LHCLK2" x="30.48" y="-20.32" length="middle" rot="R180"/>
<pin name="IO_L04N_3/LHCLK3" x="30.48" y="-22.86" length="middle" rot="R180"/>
<pin name="IP_BK3" x="30.48" y="-25.4" length="middle" direction="in" rot="R180"/>
<pin name="IO_L05P_3/LHCLK4" x="30.48" y="-27.94" length="middle" rot="R180"/>
<pin name="IO_L05N_3/LHCLK5" x="30.48" y="-30.48" length="middle" rot="R180"/>
<pin name="IO_L06P_3/LHCLK6" x="30.48" y="-33.02" length="middle" rot="R180"/>
<pin name="IO_L06N_3/LHCLK7" x="30.48" y="-35.56" length="middle" rot="R180"/>
<pin name="IO_L07P_3" x="30.48" y="-38.1" length="middle" rot="R180"/>
<pin name="IO_L07N_3" x="30.48" y="-40.64" length="middle" rot="R180"/>
<pin name="CSO_B/IO_BK2" x="30.48" y="-48.26" length="middle" rot="R180"/>
<pin name="DOUT/IO_BK2" x="30.48" y="-50.8" length="middle" rot="R180"/>
<pin name="CSI_B/IO_BK2" x="30.48" y="-53.34" length="middle" rot="R180"/>
<pin name="IP/VREF_2" x="30.48" y="-55.88" length="middle" direction="in" rot="R180"/>
<pin name="IO_L03P_2/D7/GCLK12" x="30.48" y="-58.42" length="middle" rot="R180"/>
<pin name="IO_L03N_2/D6/GCLK13" x="30.48" y="-60.96" length="middle" rot="R180"/>
<pin name="IO_BK2/D5" x="30.48" y="-63.5" length="middle" rot="R180"/>
<pin name="IO_L04P_2/D4/GCLK14" x="30.48" y="-66.04" length="middle" rot="R180"/>
<pin name="IO_L04N_2/D3/GCLK15" x="30.48" y="-68.58" length="middle" rot="R180"/>
<pin name="IP_BK2/GCLK0/NWR" x="30.48" y="-71.12" length="middle" direction="in" rot="R180"/>
<pin name="IO_L06P_2/D2/GCLK2" x="30.48" y="-73.66" length="middle" rot="R180"/>
<pin name="IO_L06N_2/D1/GCLK3" x="30.48" y="-76.2" length="middle" rot="R180"/>
<pin name="IO_L08P_2/VS2" x="30.48" y="-78.74" length="middle" rot="R180"/>
<pin name="IO_L08N_2/VS1" x="30.48" y="-81.28" length="middle" rot="R180"/>
<pin name="IO_L09P_2/VS0" x="30.48" y="-83.82" length="middle" rot="R180"/>
<pin name="IO_L01P_1" x="30.48" y="-91.44" length="middle" rot="R180"/>
<pin name="IO_L01N_1" x="30.48" y="-93.98" length="middle" rot="R180"/>
<pin name="IO_L02P_1" x="30.48" y="-96.52" length="middle" rot="R180"/>
<pin name="IO_L02N_1" x="30.48" y="-99.06" length="middle" rot="R180"/>
<pin name="IO_L03P_1/RHCLK0" x="30.48" y="-101.6" length="middle" rot="R180"/>
<pin name="IO_L03N_1/RHCLK1" x="30.48" y="-104.14" length="middle" rot="R180"/>
<pin name="IO_L04P_1/RHCLK2" x="30.48" y="-106.68" length="middle" rot="R180"/>
<pin name="IO_L04N_1/RHCLK3" x="30.48" y="-109.22" length="middle" rot="R180"/>
<pin name="IO_L05P_1/RHCLK4" x="30.48" y="-111.76" length="middle" rot="R180"/>
<pin name="IO_L05N_1/RHCLK5" x="30.48" y="-114.3" length="middle" rot="R180"/>
<pin name="IO_L06P_1/RHCLK6" x="30.48" y="-116.84" length="middle" rot="R180"/>
<pin name="IO_L06N_1/RHCLK7" x="30.48" y="-119.38" length="middle" rot="R180"/>
<pin name="IP/VREF_1" x="30.48" y="-121.92" length="middle" direction="in" rot="R180"/>
<pin name="IO_L07P_1" x="30.48" y="-124.46" length="middle" rot="R180"/>
<pin name="IO_L07N_1" x="30.48" y="-127" length="middle" rot="R180"/>
<pin name="IO_L01P_0" x="30.48" y="-134.62" length="middle" rot="R180"/>
<pin name="IO_L01N_0" x="30.48" y="-137.16" length="middle" rot="R180"/>
<pin name="IO_L02P_0/GCLK4" x="30.48" y="-139.7" length="middle" rot="R180"/>
<pin name="IO_L02N_0/GCLK5" x="30.48" y="-142.24" length="middle" rot="R180"/>
<pin name="IO_L03P_0/GCLK6" x="30.48" y="-144.78" length="middle" rot="R180"/>
<pin name="IO_L03N_0/GCLK7" x="30.48" y="-147.32" length="middle" rot="R180"/>
<pin name="IP_L04P_0/GCLK8" x="30.48" y="-149.86" length="middle" direction="in" rot="R180"/>
<pin name="IP_L04N_0/GCLK9" x="30.48" y="-152.4" length="middle" direction="in" rot="R180"/>
<pin name="IO_L05P_0/GCLK10" x="30.48" y="-154.94" length="middle" rot="R180"/>
<pin name="IO_L05N_0/GCLK11" x="30.48" y="-157.48" length="middle" rot="R180"/>
<pin name="IO_BK0" x="30.48" y="-160.02" length="middle" rot="R180"/>
<pin name="IO_L06P_0" x="30.48" y="-162.56" length="middle" rot="R180"/>
<pin name="IO_L06N_0/VREF_0" x="30.48" y="-165.1" length="middle" rot="R180"/>
<pin name="IO_L07P_0" x="30.48" y="-167.64" length="middle" rot="R180"/>
<pin name="M0" x="-30.48" y="-5.08" length="middle"/>
<pin name="M1" x="-30.48" y="-7.62" length="middle"/>
<pin name="M2" x="-30.48" y="-10.16" length="middle"/>
<pin name="CCLK" x="-30.48" y="-17.78" length="middle"/>
<pin name="DONE" x="-30.48" y="-20.32" length="middle"/>
<pin name="INIT_B" x="-30.48" y="-25.4" length="middle"/>
<pin name="PROG_B" x="-30.48" y="-27.94" length="middle"/>
<pin name="DIN" x="-30.48" y="-30.48" length="middle"/>
<pin name="HSWAP" x="-30.48" y="-38.1" length="middle"/>
<pin name="TCK" x="-30.48" y="-45.72" length="middle"/>
<pin name="TMS" x="-30.48" y="-48.26" length="middle"/>
<pin name="TDI" x="-30.48" y="-50.8" length="middle"/>
<pin name="TDO" x="-30.48" y="-53.34" length="middle"/>
<pin name="GND_12" x="-30.48" y="-167.64" length="middle" direction="pwr"/>
<pin name="GND_11" x="-30.48" y="-165.1" length="middle" direction="pwr"/>
<pin name="GND_10" x="-30.48" y="-162.56" length="middle" direction="pwr"/>
<pin name="GND_9" x="-30.48" y="-160.02" length="middle" direction="pwr"/>
<pin name="GND_8" x="-30.48" y="-157.48" length="middle" direction="pwr"/>
<pin name="GND_7" x="-30.48" y="-154.94" length="middle" direction="pwr"/>
<pin name="GND_6" x="-30.48" y="-152.4" length="middle" direction="pwr"/>
<pin name="GND_5" x="-30.48" y="-149.86" length="middle" direction="pwr"/>
<pin name="GND_4" x="-30.48" y="-147.32" length="middle" direction="pwr"/>
<pin name="GND_3" x="-30.48" y="-144.78" length="middle" direction="pwr"/>
<pin name="GND_2" x="-30.48" y="-142.24" length="middle" direction="pwr"/>
<pin name="GND_1" x="-30.48" y="-139.7" length="middle" direction="pwr"/>
<pin name="VCCO_BK3_2" x="-30.48" y="-132.08" length="middle" direction="pwr"/>
<pin name="VCCO_BK3_1" x="-30.48" y="-129.54" length="middle" direction="pwr"/>
<pin name="VCCO_BK2_2" x="-30.48" y="-124.46" length="middle" direction="pwr"/>
<pin name="VCCO_BK2_1" x="-30.48" y="-121.92" length="middle" direction="pwr"/>
<pin name="VCCO_BK1_2" x="-30.48" y="-116.84" length="middle" direction="pwr"/>
<pin name="VCCO_BK1_1" x="-30.48" y="-114.3" length="middle" direction="pwr"/>
<pin name="VCCO_BK0_2" x="-30.48" y="-109.22" length="middle" direction="pwr"/>
<pin name="VCCO_BK0_1" x="-30.48" y="-106.68" length="middle" direction="pwr"/>
<pin name="VCCAUX_4" x="-30.48" y="-99.06" length="middle" direction="pwr"/>
<pin name="VCCAUX_3" x="-30.48" y="-96.52" length="middle" direction="pwr"/>
<pin name="VCCAUX_2" x="-30.48" y="-93.98" length="middle" direction="pwr"/>
<pin name="VCCAUX_1" x="-30.48" y="-91.44" length="middle" direction="pwr"/>
<pin name="VCCINT_4" x="-30.48" y="-83.82" length="middle" direction="pwr"/>
<pin name="VCCINT_3" x="-30.48" y="-81.28" length="middle" direction="pwr"/>
<pin name="VCCINT_2" x="-30.48" y="-78.74" length="middle" direction="pwr"/>
<pin name="VCCINT_1" x="-30.48" y="-76.2" length="middle" direction="pwr"/>
</symbol>
<symbol name="UNIT_KC7050B">
<wire x1="-5.08" y1="5.08" x2="-5.08" y2="-5.08" width="0.254" layer="94"/>
<wire x1="-5.08" y1="-5.08" x2="5.08" y2="-5.08" width="0.254" layer="94"/>
<wire x1="5.08" y1="-5.08" x2="5.08" y2="5.08" width="0.254" layer="94"/>
<wire x1="5.08" y1="5.08" x2="-5.08" y2="5.08" width="0.254" layer="94"/>
<text x="-5.08" y="-7.62" size="1.778" layer="96">&gt;VALUE</text>
<text x="-5.08" y="6.35" size="1.778" layer="95">&gt;NAME</text>
<pin name="4" x="-10.16" y="2.54" length="middle" direction="pwr"/>
<pin name="1" x="-10.16" y="-2.54" length="middle"/>
<pin name="3" x="10.16" y="2.54" length="middle" direction="out" rot="R180"/>
<pin name="2" x="10.16" y="-2.54" length="middle" direction="pwr" rot="R180"/>
</symbol>
<symbol name="LOGO">
<wire x1="0" y1="3.81" x2="10.16" y2="1.27" width="0.254" layer="94"/>
<wire x1="10.16" y1="1.27" x2="7.62" y2="1.27" width="0.254" layer="94"/>
<wire x1="7.62" y1="1.27" x2="7.62" y2="-2.54" width="0.254" layer="94"/>
<wire x1="7.62" y1="-2.54" x2="10.16" y2="-2.54" width="0.254" layer="94"/>
<wire x1="10.16" y1="-2.54" x2="0" y2="-5.08" width="0.254" layer="94"/>
<wire x1="0" y1="-5.08" x2="-10.16" y2="-2.54" width="0.254" layer="94"/>
<wire x1="-10.16" y1="-2.54" x2="-7.62" y2="-2.54" width="0.254" layer="94"/>
<wire x1="-7.62" y1="-2.54" x2="-7.62" y2="1.27" width="0.254" layer="94"/>
<wire x1="-7.62" y1="1.27" x2="-10.16" y2="1.27" width="0.254" layer="94"/>
<wire x1="-10.16" y1="1.27" x2="0" y2="3.81" width="0.254" layer="94"/>
<text x="-6.35" y="-1.27" size="1.27" layer="94">VPort with Power</text>
<text x="-10.16" y="5.08" size="1.778" layer="95">&gt;NAME</text>
<text x="0" y="-7.62" size="1.778" layer="97">&gt;VALUE</text>
</symbol>
<symbol name="A3L-LOC">
<wire x1="288.29" y1="3.81" x2="342.265" y2="3.81" width="0.1016" layer="94"/>
<wire x1="342.265" y1="3.81" x2="373.38" y2="3.81" width="0.1016" layer="94"/>
<wire x1="373.38" y1="3.81" x2="383.54" y2="3.81" width="0.1016" layer="94"/>
<wire x1="383.54" y1="3.81" x2="383.54" y2="8.89" width="0.1016" layer="94"/>
<wire x1="383.54" y1="8.89" x2="383.54" y2="13.97" width="0.1016" layer="94"/>
<wire x1="383.54" y1="13.97" x2="383.54" y2="19.05" width="0.1016" layer="94"/>
<wire x1="383.54" y1="19.05" x2="383.54" y2="24.13" width="0.1016" layer="94"/>
<wire x1="288.29" y1="3.81" x2="288.29" y2="24.13" width="0.1016" layer="94"/>
<wire x1="288.29" y1="24.13" x2="342.265" y2="24.13" width="0.1016" layer="94"/>
<wire x1="342.265" y1="24.13" x2="383.54" y2="24.13" width="0.1016" layer="94"/>
<wire x1="373.38" y1="3.81" x2="373.38" y2="8.89" width="0.1016" layer="94"/>
<wire x1="373.38" y1="8.89" x2="383.54" y2="8.89" width="0.1016" layer="94"/>
<wire x1="373.38" y1="8.89" x2="342.265" y2="8.89" width="0.1016" layer="94"/>
<wire x1="342.265" y1="8.89" x2="342.265" y2="3.81" width="0.1016" layer="94"/>
<wire x1="342.265" y1="8.89" x2="342.265" y2="13.97" width="0.1016" layer="94"/>
<wire x1="342.265" y1="13.97" x2="383.54" y2="13.97" width="0.1016" layer="94"/>
<wire x1="342.265" y1="13.97" x2="342.265" y2="19.05" width="0.1016" layer="94"/>
<wire x1="342.265" y1="19.05" x2="383.54" y2="19.05" width="0.1016" layer="94"/>
<wire x1="342.265" y1="19.05" x2="342.265" y2="24.13" width="0.1016" layer="94"/>
<text x="344.17" y="15.24" size="2.54" layer="94" font="vector">&gt;DRAWING_NAME</text>
<text x="344.17" y="10.16" size="2.286" layer="94" font="vector">&gt;LAST_DATE_TIME</text>
<text x="357.505" y="5.08" size="2.54" layer="94" font="vector">&gt;NAME</text>
<text x="343.916" y="4.953" size="2.54" layer="94" font="vector">Sheet:</text>
<text x="344.17" y="20.32" size="2.286" layer="94" font="vector">&gt;VALUE</text>
<frame x1="0" y1="0" x2="387.35" y2="260.35" columns="8" rows="5" layer="94"/>
</symbol>
<symbol name="JMP_2P">
<wire x1="-1.27" y1="2.54" x2="1.27" y2="2.54" width="0.254" layer="94"/>
<wire x1="1.27" y1="2.54" x2="1.27" y2="-2.54" width="0.254" layer="94"/>
<wire x1="1.27" y1="-2.54" x2="-1.27" y2="-2.54" width="0.254" layer="94"/>
<wire x1="-1.27" y1="-2.54" x2="-1.27" y2="2.54" width="0.254" layer="94"/>
<circle x="0" y="-1.27" radius="0.635" width="0.254" layer="94"/>
<circle x="0" y="1.27" radius="0.635" width="0.254" layer="94"/>
<text x="2.54" y="0.635" size="1.778" layer="95" ratio="20">&gt;NAME</text>
<text x="2.54" y="-1.905" size="1.778" layer="97" ratio="20">&gt;VALUE</text>
<pin name="P$1" x="0" y="5.08" visible="off" length="short" rot="R270"/>
<pin name="P$2" x="0" y="-5.08" visible="off" length="short" rot="R90"/>
</symbol>
</symbols>
<devicesets>
<deviceset name="SW_PUSH" uservalue="yes">
<gates>
<gate name="LS6J2M-T" symbol="SW_PUSH" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_SW_PUSH">
<connects>
<connect gate="LS6J2M-T" pin="1" pad="1"/>
<connect gate="LS6J2M-T" pin="2" pad="2"/>
<connect gate="LS6J2M-T" pin="3" pad="3"/>
<connect gate="LS6J2M-T" pin="4" pad="4"/>
</connects>
<technologies>
<technology name="">
<attribute name="PUSH-SW" value="PUSH-WS" constant="no"/>
</technology>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_74AC125" prefix="IC">
<gates>
<gate name="G$1" symbol="UNIT_74AC125" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_SOP14">
<connects>
<connect gate="G$1" pin="1A" pad="2"/>
<connect gate="G$1" pin="1G_" pad="1"/>
<connect gate="G$1" pin="1Y" pad="3"/>
<connect gate="G$1" pin="2A" pad="5"/>
<connect gate="G$1" pin="2G_" pad="4"/>
<connect gate="G$1" pin="2Y" pad="6"/>
<connect gate="G$1" pin="3A" pad="9"/>
<connect gate="G$1" pin="3G_" pad="10"/>
<connect gate="G$1" pin="3Y" pad="8"/>
<connect gate="G$1" pin="4A" pad="12"/>
<connect gate="G$1" pin="4G_" pad="13"/>
<connect gate="G$1" pin="4Y" pad="11"/>
<connect gate="G$1" pin="GND" pad="7"/>
<connect gate="G$1" pin="VCCO" pad="14"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_93CX6" prefix="IC">
<gates>
<gate name="UNIT_93CX6" symbol="UNIT_93CX6" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_SOP8">
<connects>
<connect gate="UNIT_93CX6" pin="CS" pad="1"/>
<connect gate="UNIT_93CX6" pin="DI" pad="3"/>
<connect gate="UNIT_93CX6" pin="DO" pad="4"/>
<connect gate="UNIT_93CX6" pin="GND" pad="5"/>
<connect gate="UNIT_93CX6" pin="NC" pad="7"/>
<connect gate="UNIT_93CX6" pin="ORG" pad="6"/>
<connect gate="UNIT_93CX6" pin="SK" pad="2"/>
<connect gate="UNIT_93CX6" pin="VCC" pad="8"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="A3L-LOC" prefix="FRAME" uservalue="yes">
<description>&lt;b&gt;FRAME&lt;/b&gt;&lt;p&gt;
DIN A3, landscape with location and doc. field</description>
<gates>
<gate name="G$1" symbol="A3L-LOC" x="0" y="0"/>
</gates>
<devices>
<device name="">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_XC3S250E" prefix="IC">
<description>&lt;H1&gt;XC3S250E-VQ100&lt;/H1&gt;
Xilinx Spartan-3E FPGA device</description>
<gates>
<gate name="XC3S250E" symbol="UNIT_XC3S250E" x="25.4" y="0"/>
</gates>
<devices>
<device name="" package="SMD_VQ100">
<connects>
<connect gate="XC3S250E" pin="CCLK" pad="50"/>
<connect gate="XC3S250E" pin="CSI_B/IO_BK2" pad="27"/>
<connect gate="XC3S250E" pin="CSO_B/IO_BK2" pad="24"/>
<connect gate="XC3S250E" pin="DIN" pad="44"/>
<connect gate="XC3S250E" pin="DONE" pad="51"/>
<connect gate="XC3S250E" pin="DOUT/IO_BK2" pad="26"/>
<connect gate="XC3S250E" pin="GND_1" pad="7"/>
<connect gate="XC3S250E" pin="GND_10" pad="81"/>
<connect gate="XC3S250E" pin="GND_11" pad="87"/>
<connect gate="XC3S250E" pin="GND_12" pad="93"/>
<connect gate="XC3S250E" pin="GND_2" pad="14"/>
<connect gate="XC3S250E" pin="GND_3" pad="19"/>
<connect gate="XC3S250E" pin="GND_4" pad="29"/>
<connect gate="XC3S250E" pin="GND_5" pad="37"/>
<connect gate="XC3S250E" pin="GND_6" pad="52"/>
<connect gate="XC3S250E" pin="GND_7" pad="59"/>
<connect gate="XC3S250E" pin="GND_8" pad="64"/>
<connect gate="XC3S250E" pin="GND_9" pad="72"/>
<connect gate="XC3S250E" pin="HSWAP" pad="99"/>
<connect gate="XC3S250E" pin="INIT_B" pad="25"/>
<connect gate="XC3S250E" pin="IO_BK0" pad="92"/>
<connect gate="XC3S250E" pin="IO_BK2/D5" pad="34"/>
<connect gate="XC3S250E" pin="IO_L01N_0" pad="79"/>
<connect gate="XC3S250E" pin="IO_L01N_1" pad="54"/>
<connect gate="XC3S250E" pin="IO_L01N_3" pad="3"/>
<connect gate="XC3S250E" pin="IO_L01P_0" pad="78"/>
<connect gate="XC3S250E" pin="IO_L01P_1" pad="53"/>
<connect gate="XC3S250E" pin="IO_L01P_3" pad="2"/>
<connect gate="XC3S250E" pin="IO_L02N_0/GCLK5" pad="84"/>
<connect gate="XC3S250E" pin="IO_L02N_1" pad="58"/>
<connect gate="XC3S250E" pin="IO_L02N_3/VREF_3" pad="5"/>
<connect gate="XC3S250E" pin="IO_L02P_0/GCLK4" pad="83"/>
<connect gate="XC3S250E" pin="IO_L02P_1" pad="57"/>
<connect gate="XC3S250E" pin="IO_L02P_3" pad="4"/>
<connect gate="XC3S250E" pin="IO_L03N_0/GCLK7" pad="86"/>
<connect gate="XC3S250E" pin="IO_L03N_1/RHCLK1" pad="61"/>
<connect gate="XC3S250E" pin="IO_L03N_2/D6/GCLK13" pad="33"/>
<connect gate="XC3S250E" pin="IO_L03N_3/LHCLK1" pad="10"/>
<connect gate="XC3S250E" pin="IO_L03P_0/GCLK6" pad="85"/>
<connect gate="XC3S250E" pin="IO_L03P_1/RHCLK0" pad="60"/>
<connect gate="XC3S250E" pin="IO_L03P_2/D7/GCLK12" pad="32"/>
<connect gate="XC3S250E" pin="IO_L03P_3/LHCLK0" pad="9"/>
<connect gate="XC3S250E" pin="IO_L04N_1/RHCLK3" pad="63"/>
<connect gate="XC3S250E" pin="IO_L04N_2/D3/GCLK15" pad="36"/>
<connect gate="XC3S250E" pin="IO_L04N_3/LHCLK3" pad="12"/>
<connect gate="XC3S250E" pin="IO_L04P_1/RHCLK2" pad="62"/>
<connect gate="XC3S250E" pin="IO_L04P_2/D4/GCLK14" pad="35"/>
<connect gate="XC3S250E" pin="IO_L04P_3/LHCLK2" pad="11"/>
<connect gate="XC3S250E" pin="IO_L05N_0/GCLK11" pad="91"/>
<connect gate="XC3S250E" pin="IO_L05N_1/RHCLK5" pad="66"/>
<connect gate="XC3S250E" pin="IO_L05N_3/LHCLK5" pad="16"/>
<connect gate="XC3S250E" pin="IO_L05P_0/GCLK10" pad="90"/>
<connect gate="XC3S250E" pin="IO_L05P_1/RHCLK4" pad="65"/>
<connect gate="XC3S250E" pin="IO_L05P_3/LHCLK4" pad="15"/>
<connect gate="XC3S250E" pin="IO_L06N_0/VREF_0" pad="95"/>
<connect gate="XC3S250E" pin="IO_L06N_1/RHCLK7" pad="68"/>
<connect gate="XC3S250E" pin="IO_L06N_2/D1/GCLK3" pad="41"/>
<connect gate="XC3S250E" pin="IO_L06N_3/LHCLK7" pad="18"/>
<connect gate="XC3S250E" pin="IO_L06P_0" pad="94"/>
<connect gate="XC3S250E" pin="IO_L06P_1/RHCLK6" pad="67"/>
<connect gate="XC3S250E" pin="IO_L06P_2/D2/GCLK2" pad="40"/>
<connect gate="XC3S250E" pin="IO_L06P_3/LHCLK6" pad="17"/>
<connect gate="XC3S250E" pin="IO_L07N_1" pad="71"/>
<connect gate="XC3S250E" pin="IO_L07N_3" pad="23"/>
<connect gate="XC3S250E" pin="IO_L07P_0" pad="98"/>
<connect gate="XC3S250E" pin="IO_L07P_1" pad="70"/>
<connect gate="XC3S250E" pin="IO_L07P_3" pad="22"/>
<connect gate="XC3S250E" pin="IO_L08N_2/VS1" pad="48"/>
<connect gate="XC3S250E" pin="IO_L08P_2/VS2" pad="47"/>
<connect gate="XC3S250E" pin="IO_L09P_2/VS0" pad="49"/>
<connect gate="XC3S250E" pin="IP/VREF_1" pad="69"/>
<connect gate="XC3S250E" pin="IP/VREF_2" pad="30"/>
<connect gate="XC3S250E" pin="IP_BK2/GCLK0/NWR" pad="38"/>
<connect gate="XC3S250E" pin="IP_BK3" pad="13"/>
<connect gate="XC3S250E" pin="IP_L04N_0/GCLK9" pad="89"/>
<connect gate="XC3S250E" pin="IP_L04P_0/GCLK8" pad="88"/>
<connect gate="XC3S250E" pin="M0" pad="43"/>
<connect gate="XC3S250E" pin="M1" pad="42"/>
<connect gate="XC3S250E" pin="M2" pad="39"/>
<connect gate="XC3S250E" pin="PROG_B" pad="1"/>
<connect gate="XC3S250E" pin="TCK" pad="77"/>
<connect gate="XC3S250E" pin="TDI" pad="100"/>
<connect gate="XC3S250E" pin="TDO" pad="76"/>
<connect gate="XC3S250E" pin="TMS" pad="75"/>
<connect gate="XC3S250E" pin="VCCAUX_1" pad="21"/>
<connect gate="XC3S250E" pin="VCCAUX_2" pad="46"/>
<connect gate="XC3S250E" pin="VCCAUX_3" pad="74"/>
<connect gate="XC3S250E" pin="VCCAUX_4" pad="96"/>
<connect gate="XC3S250E" pin="VCCINT_1" pad="6"/>
<connect gate="XC3S250E" pin="VCCINT_2" pad="28"/>
<connect gate="XC3S250E" pin="VCCINT_3" pad="56"/>
<connect gate="XC3S250E" pin="VCCINT_4" pad="80"/>
<connect gate="XC3S250E" pin="VCCO_BK0_1" pad="82"/>
<connect gate="XC3S250E" pin="VCCO_BK0_2" pad="97"/>
<connect gate="XC3S250E" pin="VCCO_BK1_1" pad="55"/>
<connect gate="XC3S250E" pin="VCCO_BK1_2" pad="73"/>
<connect gate="XC3S250E" pin="VCCO_BK2_1" pad="31"/>
<connect gate="XC3S250E" pin="VCCO_BK2_2" pad="45"/>
<connect gate="XC3S250E" pin="VCCO_BK3_1" pad="8"/>
<connect gate="XC3S250E" pin="VCCO_BK3_2" pad="20"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_XCF02S" prefix="IC">
<description>&lt;H1&gt;XCF02S&lt;/H1&gt;
Xilinx FPGA Configration ROM XCF02S VO20C</description>
<gates>
<gate name="XCF02S" symbol="UNIT_XCF02S" x="17.78" y="0"/>
</gates>
<devices>
<device name="" package="SMD_VO20C">
<connects>
<connect gate="XCF02S" pin="CEO_" pad="13"/>
<connect gate="XCF02S" pin="CE_" pad="10"/>
<connect gate="XCF02S" pin="CF_" pad="7"/>
<connect gate="XCF02S" pin="CLK" pad="3"/>
<connect gate="XCF02S" pin="D0" pad="1"/>
<connect gate="XCF02S" pin="GND1" pad="11"/>
<connect gate="XCF02S" pin="NC1" pad="2"/>
<connect gate="XCF02S" pin="NC2" pad="9"/>
<connect gate="XCF02S" pin="NC3" pad="12"/>
<connect gate="XCF02S" pin="NC4" pad="14"/>
<connect gate="XCF02S" pin="NC5" pad="15"/>
<connect gate="XCF02S" pin="NC6" pad="16"/>
<connect gate="XCF02S" pin="OE/NRST" pad="8"/>
<connect gate="XCF02S" pin="TCK" pad="6"/>
<connect gate="XCF02S" pin="TDI" pad="4"/>
<connect gate="XCF02S" pin="TDO" pad="17"/>
<connect gate="XCF02S" pin="TMS" pad="5"/>
<connect gate="XCF02S" pin="VCCI" pad="18"/>
<connect gate="XCF02S" pin="VCCJ" pad="20"/>
<connect gate="XCF02S" pin="VCCO" pad="19"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_FT2232L" prefix="IC">
<gates>
<gate name="FT2232L" symbol="UNIT_FT2232L" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_LQFP48">
<connects>
<connect gate="FT2232L" pin="3V3OUT" pad="6"/>
<connect gate="FT2232L" pin="ACBUS0" pad="15"/>
<connect gate="FT2232L" pin="ACBUS1" pad="13"/>
<connect gate="FT2232L" pin="ACBUS2" pad="12"/>
<connect gate="FT2232L" pin="ACBUS3" pad="11"/>
<connect gate="FT2232L" pin="ADBUS0" pad="24"/>
<connect gate="FT2232L" pin="ADBUS1" pad="23"/>
<connect gate="FT2232L" pin="ADBUS2" pad="22"/>
<connect gate="FT2232L" pin="ADBUS3" pad="21"/>
<connect gate="FT2232L" pin="ADBUS4" pad="20"/>
<connect gate="FT2232L" pin="ADBUS5" pad="19"/>
<connect gate="FT2232L" pin="ADBUS6" pad="17"/>
<connect gate="FT2232L" pin="ADBUS7" pad="16"/>
<connect gate="FT2232L" pin="AGND" pad="45"/>
<connect gate="FT2232L" pin="AVCC" pad="46"/>
<connect gate="FT2232L" pin="BCBUS0" pad="30"/>
<connect gate="FT2232L" pin="BCBUS1" pad="29"/>
<connect gate="FT2232L" pin="BCBUS2" pad="28"/>
<connect gate="FT2232L" pin="BCBUS3" pad="27"/>
<connect gate="FT2232L" pin="BDBUS0" pad="40"/>
<connect gate="FT2232L" pin="BDBUS1" pad="39"/>
<connect gate="FT2232L" pin="BDBUS2" pad="38"/>
<connect gate="FT2232L" pin="BDBUS3" pad="37"/>
<connect gate="FT2232L" pin="BDBUS4" pad="36"/>
<connect gate="FT2232L" pin="BDBUS5" pad="35"/>
<connect gate="FT2232L" pin="BDBUS6" pad="33"/>
<connect gate="FT2232L" pin="BDBUS7" pad="32"/>
<connect gate="FT2232L" pin="EECS" pad="48"/>
<connect gate="FT2232L" pin="EEDATA" pad="2"/>
<connect gate="FT2232L" pin="EESK" pad="1"/>
<connect gate="FT2232L" pin="GND@1" pad="9"/>
<connect gate="FT2232L" pin="GND@2" pad="18"/>
<connect gate="FT2232L" pin="GND@3" pad="25"/>
<connect gate="FT2232L" pin="GND@4" pad="34"/>
<connect gate="FT2232L" pin="PWREN_" pad="41"/>
<connect gate="FT2232L" pin="RESET_" pad="4"/>
<connect gate="FT2232L" pin="RSTOUT_" pad="5"/>
<connect gate="FT2232L" pin="SI/WUA" pad="10"/>
<connect gate="FT2232L" pin="SI/WUB" pad="26"/>
<connect gate="FT2232L" pin="TEST" pad="47"/>
<connect gate="FT2232L" pin="USBDM" pad="8"/>
<connect gate="FT2232L" pin="USBDP" pad="7"/>
<connect gate="FT2232L" pin="VCC@1" pad="3"/>
<connect gate="FT2232L" pin="VCC@2" pad="42"/>
<connect gate="FT2232L" pin="VCCIOA" pad="14"/>
<connect gate="FT2232L" pin="VCCIOB" pad="31"/>
<connect gate="FT2232L" pin="XTIN" pad="43"/>
<connect gate="FT2232L" pin="XTOUT" pad="44"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="VCCAUX">
<description>VCC 2.5V</description>
<gates>
<gate name="G$1" symbol="VCCAUX" x="0" y="0"/>
</gates>
<devices>
<device name="">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="VCCINT" prefix="SUPPLY">
<description>VCC 1.25V</description>
<gates>
<gate name="VCCINT" symbol="VCCINT" x="0" y="0"/>
</gates>
<devices>
<device name="">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="VCCO">
<description>VCC 3.3V</description>
<gates>
<gate name="G$1" symbol="VCCO" x="0" y="0"/>
</gates>
<devices>
<device name="">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="GND">
<description>GND</description>
<gates>
<gate name="G$1" symbol="GND" x="0" y="0"/>
</gates>
<devices>
<device name="">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_M51957B" prefix="IC">
<description>M51957B&lt;br&gt;</description>
<gates>
<gate name="G$1" symbol="UNIT_M51957B" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_SOP8">
<connects>
<connect gate="G$1" pin="DLY" pad="5"/>
<connect gate="G$1" pin="GND" pad="4"/>
<connect gate="G$1" pin="IN" pad="2"/>
<connect gate="G$1" pin="NC@1" pad="1"/>
<connect gate="G$1" pin="NC@2" pad="3"/>
<connect gate="G$1" pin="NC@3" pad="8"/>
<connect gate="G$1" pin="OUT" pad="6"/>
<connect gate="G$1" pin="VCCO" pad="7"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="LED_SMD_1605_POL">
<description>LED&lt;br&gt;
- SMD package&lt;br&gt;
- 0.1inch interval</description>
<gates>
<gate name="G$1" symbol="LED" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_1605_POL">
<connects>
<connect gate="G$1" pin="A" pad="A"/>
<connect gate="G$1" pin="C" pad="C"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="REG_SMD_1605" prefix="R" uservalue="yes">
<description>resistor&lt;br&gt;
SMD package&lt;br&gt;
1605 or 2125</description>
<gates>
<gate name="G$1" symbol="REG" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_1605_R">
<connects>
<connect gate="G$1" pin="1" pad="1"/>
<connect gate="G$1" pin="2" pad="2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CAP_SMD_1605" prefix="C" uservalue="yes">
<description>resistor&lt;br&gt;
SMD package&lt;br&gt;
1605 or 2125 size</description>
<gates>
<gate name="G$1" symbol="CAP" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_1605">
<connects>
<connect gate="G$1" pin="1" pad="1"/>
<connect gate="G$1" pin="2" pad="2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CAP_SMDPAD_1605_04INCH" prefix="C" uservalue="yes">
<description>ceramic capacitor&lt;br&gt;
- SMD or PAD package&lt;br&gt;
- 1605 or 2125&lt;br&gt;
- 0.4-inch interval</description>
<gates>
<gate name="G$1" symbol="CAP" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMDPAD_1605_04INCH">
<connects>
<connect gate="G$1" pin="1" pad="P$1"/>
<connect gate="G$1" pin="2" pad="P$2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CAP_SMDPAD_1605_02INCH" prefix="C" uservalue="yes">
<description>ceramic capacitor&lt;br&gt;
- SMD or PAD package&lt;br&gt;
- 1605 or 2125&lt;br&gt;
- 0.2-inch interval</description>
<gates>
<gate name="G$1" symbol="CAP" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMDPAD_1605_02INCH">
<connects>
<connect gate="G$1" pin="1" pad="P$1"/>
<connect gate="G$1" pin="2" pad="P$2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CAP_SMDPAD_1605_03INCH" prefix="C" uservalue="yes">
<description>ceramic capacitor&lt;br&gt;
- SMD or PAD package&lt;br&gt;
- 1605 or 2125&lt;br&gt;
- 0.3-inch interval</description>
<gates>
<gate name="G$1" symbol="CAP" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMDPAD_1605_03INCH">
<connects>
<connect gate="G$1" pin="1" pad="P$1"/>
<connect gate="G$1" pin="2" pad="P$2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CON_POW_M">
<description>pin header&lt;br&gt;
- 2*5
- L angle</description>
<gates>
<gate name="G$1" symbol="CON_POW_BRD" x="0" y="0"/>
</gates>
<devices>
<device name="" package="PAD_POW_M">
<connects>
<connect gate="G$1" pin="GND@1" pad="1"/>
<connect gate="G$1" pin="GND@2" pad="2"/>
<connect gate="G$1" pin="GND@3" pad="9"/>
<connect gate="G$1" pin="GND@4" pad="10"/>
<connect gate="G$1" pin="VCCAUX@1" pad="5"/>
<connect gate="G$1" pin="VCCAUX@2" pad="6"/>
<connect gate="G$1" pin="VCCINT@1" pad="3"/>
<connect gate="G$1" pin="VCCINT@2" pad="4"/>
<connect gate="G$1" pin="VCCO@1" pad="7"/>
<connect gate="G$1" pin="VCCO@2" pad="8"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CON_VPORT" prefix="VP">
<description>BOX header&lt;br&gt;
- 2*5</description>
<gates>
<gate name="G$1" symbol="CON_VPORT" x="0" y="0"/>
</gates>
<devices>
<device name="" package="PAD_VPORT">
<connects>
<connect gate="G$1" pin="1" pad="1"/>
<connect gate="G$1" pin="10" pad="10"/>
<connect gate="G$1" pin="2" pad="2"/>
<connect gate="G$1" pin="3" pad="3"/>
<connect gate="G$1" pin="4" pad="4"/>
<connect gate="G$1" pin="5" pad="5"/>
<connect gate="G$1" pin="6" pad="6"/>
<connect gate="G$1" pin="7" pad="7"/>
<connect gate="G$1" pin="8" pad="8"/>
<connect gate="G$1" pin="9" pad="9"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="FB_SMD_1605">
<description>ferrite beads&lt;br&gt;
SMD package&lt;br&gt;
1605 or 2125 size</description>
<gates>
<gate name="G$1" symbol="FB" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_1605">
<connects>
<connect gate="G$1" pin="P$1" pad="1"/>
<connect gate="G$1" pin="P$2" pad="2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_74HC14" prefix="IC">
<description>74VHC14</description>
<gates>
<gate name="G$1" symbol="UNIT_74HC14" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_SOP14">
<connects>
<connect gate="G$1" pin="GND" pad="7"/>
<connect gate="G$1" pin="IN1" pad="1"/>
<connect gate="G$1" pin="IN2" pad="3"/>
<connect gate="G$1" pin="IN3" pad="5"/>
<connect gate="G$1" pin="IN4" pad="9"/>
<connect gate="G$1" pin="IN5" pad="11"/>
<connect gate="G$1" pin="IN6" pad="13"/>
<connect gate="G$1" pin="OUT1" pad="2"/>
<connect gate="G$1" pin="OUT2" pad="4"/>
<connect gate="G$1" pin="OUT3" pad="6"/>
<connect gate="G$1" pin="OUT4" pad="8"/>
<connect gate="G$1" pin="OUT5" pad="10"/>
<connect gate="G$1" pin="OUT6" pad="12"/>
<connect gate="G$1" pin="VCCO" pad="14"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_CRYSTAL" prefix="U">
<description>Crystal Units&lt;br&gt;
SMD package</description>
<gates>
<gate name="UNIT_CRYSTAL" symbol="UNIT_CRYSTAL" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_HC-49_S">
<connects>
<connect gate="UNIT_CRYSTAL" pin="P$1" pad="P$1"/>
<connect gate="UNIT_CRYSTAL" pin="P$2" pad="P$2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="UNIT_KC7050B" prefix="U">
<description>KC7050B Clock Crystal Oscillators&lt;br&gt;
SMD pakage</description>
<gates>
<gate name="UNIT_KC7050B" symbol="UNIT_KC7050B" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_KC7050B">
<connects>
<connect gate="UNIT_KC7050B" pin="1" pad="1"/>
<connect gate="UNIT_KC7050B" pin="2" pad="2"/>
<connect gate="UNIT_KC7050B" pin="3" pad="3"/>
<connect gate="UNIT_KC7050B" pin="4" pad="4"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="LOGO" prefix="LOGO">
<description>VPort logo</description>
<gates>
<gate name="G$1" symbol="LOGO" x="0" y="0"/>
</gates>
<devices>
<device name="" package="LOGO">
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="CON_USB">
<description>UX60A-MB-5ST USB mini-B connector&lt;br&gt;
HRS</description>
<gates>
<gate name="CON_USB" symbol="CON_USB" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_USB">
<connects>
<connect gate="CON_USB" pin="1" pad="1"/>
<connect gate="CON_USB" pin="2" pad="2"/>
<connect gate="CON_USB" pin="3" pad="3"/>
<connect gate="CON_USB" pin="4" pad="4"/>
<connect gate="CON_USB" pin="5" pad="5"/>
<connect gate="CON_USB" pin="PAD1" pad="PAD@1"/>
<connect gate="CON_USB" pin="PAD2" pad="PAD@2"/>
<connect gate="CON_USB" pin="PAD3" pad="PAD@3"/>
<connect gate="CON_USB" pin="PAD4" pad="PAD@4"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="HOLE_32" prefix="H">
<description>Mounting PAD&lt;br&gt;
3.2mm</description>
<gates>
<gate name="G$1" symbol="HOLE" x="0" y="0"/>
</gates>
<devices>
<device name="" package="HOLE_32">
<connects>
<connect gate="G$1" pin="MOUNT" pad="B3,2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="LED_7SEG" prefix="7SEG">
<description>7seg led LF-301VA&lt;br&gt;
- anode common&lt;br&gt;
- SMD</description>
<gates>
<gate name="7SEG" symbol="LED_7SEG" x="0" y="0"/>
</gates>
<devices>
<device name="" package="SMD_LED_7SEG">
<connects>
<connect gate="7SEG" pin="A" pad="7"/>
<connect gate="7SEG" pin="B" pad="6"/>
<connect gate="7SEG" pin="C" pad="4"/>
<connect gate="7SEG" pin="COMMON3" pad="3"/>
<connect gate="7SEG" pin="COMMON8" pad="8"/>
<connect gate="7SEG" pin="D" pad="2"/>
<connect gate="7SEG" pin="D.P" pad="5"/>
<connect gate="7SEG" pin="E" pad="1"/>
<connect gate="7SEG" pin="F" pad="9"/>
<connect gate="7SEG" pin="G" pad="10"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
<deviceset name="JMP_2P" prefix="JP">
<gates>
<gate name="G$1" symbol="JMP_2P" x="0" y="0"/>
</gates>
<devices>
<device name="" package="PAD_JMP_2P">
<connects>
<connect gate="G$1" pin="P$1" pad="P$1"/>
<connect gate="G$1" pin="P$2" pad="P$2"/>
</connects>
<technologies>
<technology name=""/>
</technologies>
</device>
</devices>
</deviceset>
</devicesets>
</library>
</libraries>
<attributes>
</attributes>
<variantdefs>
</variantdefs>
<classes>
<class number="0" name="default" width="0" drill="0">
</class>
</classes>
<parts>
<part name="IC1" library="00_Respon" deviceset="UNIT_XC3S250E" device=""/>
<part name="POWER" library="00_Respon" deviceset="CON_POW_M" device=""/>
<part name="U$7" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="U$10" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$2" library="00_Respon" deviceset="VCCO" device=""/>
<part name="SUPPLY2" library="00_Respon" deviceset="VCCINT" device=""/>
<part name="SUPPLY1" library="00_Respon" deviceset="VCCINT" device=""/>
<part name="SUPPLY3" library="00_Respon" deviceset="VCCINT" device=""/>
<part name="C1" library="00_Respon" deviceset="CAP_SMDPAD_1605_04INCH" device="" value="0.1u"/>
<part name="C2" library="00_Respon" deviceset="CAP_SMDPAD_1605_04INCH" device="" value="0.1u"/>
<part name="C3" library="00_Respon" deviceset="CAP_SMDPAD_1605_04INCH" device="" value="0.1u"/>
<part name="C4" library="00_Respon" deviceset="CAP_SMDPAD_1605_04INCH" device="" value="0.1u"/>
<part name="U$9" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="U$11" library="00_Respon" deviceset="VCCO" device=""/>
<part name="C5" library="00_Respon" deviceset="CAP_SMDPAD_1605_03INCH" device="" value="0.1u"/>
<part name="C6" library="00_Respon" deviceset="CAP_SMDPAD_1605_03INCH" device="" value="0.1u"/>
<part name="C7" library="00_Respon" deviceset="CAP_SMDPAD_1605_03INCH" device="" value="0.1u"/>
<part name="C8" library="00_Respon" deviceset="CAP_SMDPAD_1605_03INCH" device="" value="0.1u"/>
<part name="C9" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C10" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C11" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C12" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C13" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C14" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C15" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C16" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="IC6" library="00_Respon" deviceset="UNIT_FT2232L" device=""/>
<part name="USB" library="00_Respon" deviceset="CON_USB" device=""/>
<part name="IC2" library="00_Respon" deviceset="UNIT_XCF02S" device=""/>
<part name="H1" library="00_Respon" deviceset="HOLE_32" device="" value="MOUNT-PAD-ROUND3.0"/>
<part name="H2" library="00_Respon" deviceset="HOLE_32" device="" value="MOUNT-PAD-ROUND3.0"/>
<part name="H3" library="00_Respon" deviceset="HOLE_32" device="" value="MOUNT-PAD-ROUND3.0"/>
<part name="H4" library="00_Respon" deviceset="HOLE_32" device="" value="MOUNT-PAD-ROUND3.0"/>
<part name="7SEG2" library="00_Respon" deviceset="LED_7SEG" device=""/>
<part name="7SEG1" library="00_Respon" deviceset="LED_7SEG" device=""/>
<part name="SW1" library="00_Respon" deviceset="SW_PUSH" device=""/>
<part name="SW2" library="00_Respon" deviceset="SW_PUSH" device=""/>
<part name="SW3" library="00_Respon" deviceset="SW_PUSH" device=""/>
<part name="SW4" library="00_Respon" deviceset="SW_PUSH" device=""/>
<part name="VPD" library="00_Respon" deviceset="CON_VPORT" device=""/>
<part name="VPE" library="00_Respon" deviceset="CON_VPORT" device=""/>
<part name="IC3" library="00_Respon" deviceset="UNIT_74AC125" device=""/>
<part name="IC7" library="00_Respon" deviceset="UNIT_93CX6" device=""/>
<part name="U2" library="00_Respon" deviceset="UNIT_CRYSTAL" device=""/>
<part name="DONE" library="00_Respon" deviceset="LED_SMD_1605_POL" device=""/>
<part name="R1" library="00_Respon" deviceset="REG_SMD_1605" device="" value="330"/>
<part name="R2" library="00_Respon" deviceset="REG_SMD_1605" device="" value="4.7K"/>
<part name="R3" library="00_Respon" deviceset="REG_SMD_1605" device="" value="4.7K"/>
<part name="C29" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="33p"/>
<part name="C30" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="33p"/>
<part name="R46" library="00_Respon" deviceset="REG_SMD_1605" device="" value="10K"/>
<part name="R43" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="R18" library="00_Respon" deviceset="REG_SMD_1605" device="" value="33"/>
<part name="R19" library="00_Respon" deviceset="REG_SMD_1605" device="" value="33"/>
<part name="R26" library="00_Respon" deviceset="REG_SMD_1605" device="" value="1.5K"/>
<part name="GND" library="00_Respon" deviceset="GND" device=""/>
<part name="R17" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="U1" library="00_Respon" deviceset="UNIT_KC7050B" device=""/>
<part name="IC4" library="00_Respon" deviceset="UNIT_M51957B" device=""/>
<part name="RESET" library="00_Respon" deviceset="SW_PUSH" device=""/>
<part name="C21" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="1u"/>
<part name="R9" library="00_Respon" deviceset="REG_SMD_1605" device="" value="10K"/>
<part name="R12" library="00_Respon" deviceset="REG_SMD_1605" device="" value="7.5K"/>
<part name="R10" library="00_Respon" deviceset="REG_SMD_1605" device="" value="1K"/>
<part name="R16" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="R15" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="R14" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="GND1" library="00_Respon" deviceset="GND" device=""/>
<part name="GND3" library="00_Respon" deviceset="GND" device=""/>
<part name="GND5" library="00_Respon" deviceset="GND" device=""/>
<part name="1/2" library="00_Respon" deviceset="A3L-LOC" device="" value="Drawing: yone2"/>
<part name="2/2" library="00_Respon" deviceset="A3L-LOC" device="" value="Drawing: yone2"/>
<part name="U$48" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$89" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="GND6" library="00_Respon" deviceset="GND" device=""/>
<part name="GND7" library="00_Respon" deviceset="GND" device=""/>
<part name="GND8" library="00_Respon" deviceset="GND" device=""/>
<part name="GND9" library="00_Respon" deviceset="GND" device=""/>
<part name="GND10" library="00_Respon" deviceset="GND" device=""/>
<part name="R13" library="00_Respon" deviceset="REG_SMD_1605" device="" value="470"/>
<part name="U$40" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND11" library="00_Respon" deviceset="GND" device=""/>
<part name="C28" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.033u"/>
<part name="C27" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.01u"/>
<part name="U$57" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$58" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND12" library="00_Respon" deviceset="GND" device=""/>
<part name="U$59" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND13" library="00_Respon" deviceset="GND" device=""/>
<part name="U$60" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND14" library="00_Respon" deviceset="GND" device=""/>
<part name="U$63" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="U$64" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$65" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND15" library="00_Respon" deviceset="GND" device=""/>
<part name="U$66" library="00_Respon" deviceset="VCCO" device=""/>
<part name="R4" library="00_Respon" deviceset="REG_SMD_1605" device="" value="4.7K"/>
<part name="U$39" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="U$68" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND16" library="00_Respon" deviceset="GND" device=""/>
<part name="GND17" library="00_Respon" deviceset="GND" device=""/>
<part name="C17" library="00_Respon" deviceset="CAP_SMDPAD_1605_02INCH" device="" value="0.1u"/>
<part name="C18" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.1u"/>
<part name="GND18" library="00_Respon" deviceset="GND" device=""/>
<part name="U$71" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND19" library="00_Respon" deviceset="GND" device=""/>
<part name="GND20" library="00_Respon" deviceset="GND" device=""/>
<part name="C20" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.1u"/>
<part name="C19" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.1u"/>
<part name="U$92" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$93" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND21" library="00_Respon" deviceset="GND" device=""/>
<part name="GND22" library="00_Respon" deviceset="GND" device=""/>
<part name="R36" library="00_Respon" deviceset="REG_SMD_1605" device="" value="330"/>
<part name="LED2" library="00_Respon" deviceset="LED_SMD_1605_POL" device=""/>
<part name="U$97" library="00_Respon" deviceset="VCCO" device=""/>
<part name="R28" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R30" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R32" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R34" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R38" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R40" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R42" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R45" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R27" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R29" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R31" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R33" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R37" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R39" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R41" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R44" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R5" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R7" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R8" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="C24" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="1u"/>
<part name="C25" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="1u"/>
<part name="C26" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="1u"/>
<part name="C23" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="1u"/>
<part name="R23" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R24" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R25" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="R22" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="IC5" library="00_Respon" deviceset="UNIT_74HC14" device=""/>
<part name="R20" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="R21" library="00_Respon" deviceset="REG_SMD_1605" device="" value="2.2K"/>
<part name="R35" library="00_Respon" deviceset="REG_SMD_1605" device="" value="330"/>
<part name="LED1" library="00_Respon" deviceset="LED_SMD_1605_POL" device=""/>
<part name="U$117" library="00_Respon" deviceset="VCCO" device=""/>
<part name="GND23" library="00_Respon" deviceset="GND" device=""/>
<part name="FB1" library="00_Respon" deviceset="FB_SMD_1605" device=""/>
<part name="R11" library="00_Respon" deviceset="REG_SMD_1605" device="" value="100"/>
<part name="U$1" library="00_Respon" deviceset="VCCO" device=""/>
<part name="U$3" library="00_Respon" deviceset="VCCO" device=""/>
<part name="LOGO" library="00_Respon" deviceset="LOGO" device=""/>
<part name="JP1" library="00_Respon" deviceset="JMP_2P" device=""/>
<part name="R6" library="00_Respon" deviceset="REG_SMD_1605" device="" value="4.7K"/>
<part name="GND2" library="00_Respon" deviceset="GND" device=""/>
<part name="U$4" library="00_Respon" deviceset="VCCAUX" device=""/>
<part name="C22" library="00_Respon" deviceset="CAP_SMD_1605" device="" value="0.1u"/>
</parts>
<sheets>
<sheet>
<plain>
<text x="210.82" y="424.18" size="7.62" layer="100">FPGA</text>
<text x="88.9" y="231.14" size="7.62" layer="100">USB_IC</text>
<text x="251.46" y="241.3" size="7.62" layer="100">PUSH_SW</text>
<text x="213.36" y="78.74" size="7.62" layer="100">VPORT</text>
<text x="93.98" y="58.42" size="7.62" layer="100">LED</text>
<text x="254" y="147.32" size="7.62" layer="100">7SEG</text>
<text x="17.78" y="274.32" size="7.62" layer="100">CLOCK</text>
<text x="132.08" y="330.2" size="7.62" layer="100">RESET</text>
<text x="45.72" y="449.58" size="7.62" layer="100">POWER</text>
</plain>
<instances>
<instance part="IC1" gate="XC3S250E" x="299.72" y="485.14"/>
<instance part="POWER" gate="G$1" x="33.02" y="480.06"/>
<instance part="U$7" gate="G$1" x="251.46" y="396.24"/>
<instance part="U$10" gate="G$1" x="251.46" y="381"/>
<instance part="U$2" gate="G$1" x="40.64" y="495.3"/>
<instance part="SUPPLY2" gate="VCCINT" x="73.66" y="495.3"/>
<instance part="SUPPLY1" gate="VCCINT" x="251.46" y="411.48"/>
<instance part="SUPPLY3" gate="VCCINT" x="20.32" y="441.96" smashed="yes">
<attribute name="VALUE" x="25.4" y="447.04" size="1.778" layer="96"/>
</instance>
<instance part="C1" gate="G$1" x="20.32" y="436.88"/>
<instance part="C2" gate="G$1" x="27.94" y="436.88"/>
<instance part="C3" gate="G$1" x="35.56" y="436.88"/>
<instance part="C4" gate="G$1" x="43.18" y="436.88"/>
<instance part="U$9" gate="G$1" x="20.32" y="408.94" smashed="yes">
<attribute name="VALUE" x="25.4" y="414.02" size="1.778" layer="96"/>
</instance>
<instance part="U$11" gate="G$1" x="20.32" y="375.92" smashed="yes">
<attribute name="VALUE" x="25.4" y="381" size="1.778" layer="96"/>
</instance>
<instance part="C5" gate="G$1" x="20.32" y="403.86"/>
<instance part="C6" gate="G$1" x="27.94" y="403.86"/>
<instance part="C7" gate="G$1" x="35.56" y="403.86"/>
<instance part="C8" gate="G$1" x="43.18" y="403.86"/>
<instance part="C9" gate="G$1" x="20.32" y="370.84"/>
<instance part="C10" gate="G$1" x="27.94" y="370.84"/>
<instance part="C11" gate="G$1" x="35.56" y="370.84"/>
<instance part="C12" gate="G$1" x="43.18" y="370.84"/>
<instance part="C13" gate="G$1" x="50.8" y="370.84"/>
<instance part="C14" gate="G$1" x="58.42" y="370.84"/>
<instance part="C15" gate="G$1" x="66.04" y="370.84"/>
<instance part="C16" gate="G$1" x="73.66" y="370.84"/>
<instance part="IC6" gate="FT2232L" x="111.76" y="213.36"/>
<instance part="USB" gate="CON_USB" x="20.32" y="198.12" rot="MR0"/>
<instance part="IC2" gate="XCF02S" x="139.7" y="474.98"/>
<instance part="H1" gate="G$1" x="27.94" y="55.88"/>
<instance part="H2" gate="G$1" x="27.94" y="48.26"/>
<instance part="H3" gate="G$1" x="27.94" y="40.64"/>
<instance part="H4" gate="G$1" x="27.94" y="33.02"/>
<instance part="7SEG2" gate="7SEG" x="358.14" y="142.24"/>
<instance part="7SEG1" gate="7SEG" x="299.72" y="142.24"/>
<instance part="SW1" gate="LS6J2M-T" x="203.2" y="187.96" rot="R270"/>
<instance part="SW2" gate="LS6J2M-T" x="228.6" y="187.96" rot="R270"/>
<instance part="SW3" gate="LS6J2M-T" x="254" y="187.96" rot="R270"/>
<instance part="SW4" gate="LS6J2M-T" x="279.4" y="187.96" rot="R270"/>
<instance part="VPD" gate="G$1" x="276.86" y="81.28" smashed="yes" rot="R90">
<attribute name="NAME" x="264.16" y="86.36" size="1.778" layer="95"/>
<attribute name="VALUE" x="264.16" y="88.9" size="1.778" layer="97"/>
</instance>
<instance part="VPE" gate="G$1" x="335.28" y="81.28" smashed="yes" rot="R90">
<attribute name="NAME" x="322.58" y="86.36" size="1.778" layer="95"/>
<attribute name="VALUE" x="322.58" y="88.9" size="1.778" layer="97"/>
</instance>
<instance part="IC3" gate="G$1" x="154.94" y="396.24"/>
<instance part="IC7" gate="UNIT_93CX6" x="93.98" y="106.68"/>
<instance part="U2" gate="UNIT_CRYSTAL" x="50.8" y="160.02" smashed="yes" rot="R90">
<attribute name="NAME" x="50.8" y="154.94" size="1.778" layer="95"/>
<attribute name="VALUE" x="57.15" y="166.37" size="1.778" layer="97" rot="R180"/>
</instance>
<instance part="DONE" gate="G$1" x="182.88" y="464.82" smashed="yes" rot="MR90">
<attribute name="NAME" x="184.15" y="463.55" size="1.778" layer="95" rot="MR180"/>
<attribute name="VALUE" x="179.07" y="471.17" size="1.778" layer="97" rot="MR180"/>
</instance>
<instance part="R1" gate="G$1" x="226.06" y="482.6" rot="R270"/>
<instance part="R2" gate="G$1" x="231.14" y="482.6" rot="R270"/>
<instance part="R3" gate="G$1" x="236.22" y="482.6" rot="R270"/>
<instance part="C29" gate="G$1" x="45.72" y="144.78"/>
<instance part="C30" gate="G$1" x="60.96" y="144.78" smashed="yes">
<attribute name="NAME" x="54.61" y="144.78" size="1.778" layer="95"/>
<attribute name="VALUE" x="54.61" y="142.24" size="1.778" layer="96"/>
</instance>
<instance part="R46" gate="G$1" x="55.88" y="96.52" smashed="yes" rot="R180">
<attribute name="NAME" x="55.88" y="93.98" size="1.778" layer="95" rot="R180"/>
<attribute name="VALUE" x="55.88" y="91.44" size="1.778" layer="95" rot="R180"/>
</instance>
<instance part="R43" gate="G$1" x="60.96" y="106.68" smashed="yes" rot="R270">
<attribute name="NAME" x="53.34" y="104.14" size="1.778" layer="95"/>
<attribute name="VALUE" x="53.34" y="101.6" size="1.778" layer="95"/>
</instance>
<instance part="R18" gate="G$1" x="43.18" y="200.66"/>
<instance part="R19" gate="G$1" x="43.18" y="198.12"/>
<instance part="R26" gate="G$1" x="66.04" y="172.72"/>
<instance part="GND" gate="G$1" x="261.62" y="312.42"/>
<instance part="R17" gate="G$1" x="281.94" y="203.2" smashed="yes" rot="R270">
<attribute name="NAME" x="274.32" y="200.66" size="1.778" layer="95"/>
<attribute name="VALUE" x="274.32" y="198.12" size="1.778" layer="95"/>
</instance>
<instance part="U1" gate="UNIT_KC7050B" x="66.04" y="294.64"/>
<instance part="IC4" gate="G$1" x="144.78" y="312.42"/>
<instance part="RESET" gate="LS6J2M-T" x="167.64" y="297.18" rot="R270"/>
<instance part="C21" gate="G$1" x="182.88" y="297.18"/>
<instance part="R9" gate="G$1" x="119.38" y="314.96" smashed="yes" rot="R270">
<attribute name="NAME" x="111.76" y="312.42" size="1.778" layer="95"/>
<attribute name="VALUE" x="111.76" y="309.88" size="1.778" layer="95"/>
</instance>
<instance part="R12" gate="G$1" x="119.38" y="299.72" smashed="yes" rot="R270">
<attribute name="NAME" x="111.76" y="297.18" size="1.778" layer="95"/>
<attribute name="VALUE" x="111.76" y="294.64" size="1.778" layer="95"/>
</instance>
<instance part="R10" gate="G$1" x="170.18" y="314.96" smashed="yes" rot="R270">
<attribute name="NAME" x="172.72" y="312.42" size="1.778" layer="95"/>
<attribute name="VALUE" x="172.72" y="309.88" size="1.778" layer="95"/>
</instance>
<instance part="R16" gate="G$1" x="256.54" y="203.2" smashed="yes" rot="R270">
<attribute name="NAME" x="248.92" y="200.66" size="1.778" layer="95"/>
<attribute name="VALUE" x="248.92" y="198.12" size="1.778" layer="95"/>
</instance>
<instance part="R15" gate="G$1" x="231.14" y="203.2" smashed="yes" rot="R270">
<attribute name="NAME" x="223.52" y="200.66" size="1.778" layer="95"/>
<attribute name="VALUE" x="223.52" y="198.12" size="1.778" layer="95"/>
</instance>
<instance part="R14" gate="G$1" x="205.74" y="203.2" smashed="yes" rot="R270">
<attribute name="NAME" x="198.12" y="200.66" size="1.778" layer="95"/>
<attribute name="VALUE" x="198.12" y="198.12" size="1.778" layer="95"/>
</instance>
<instance part="GND1" gate="G$1" x="111.76" y="429.26"/>
<instance part="GND3" gate="G$1" x="210.82" y="172.72"/>
<instance part="GND5" gate="G$1" x="40.64" y="464.82"/>
<instance part="1/2" gate="G$1" x="0" y="261.62"/>
<instance part="2/2" gate="G$1" x="0" y="0"/>
<instance part="U$48" gate="G$1" x="205.74" y="241.3"/>
<instance part="U$89" gate="G$1" x="55.88" y="495.3"/>
<instance part="GND6" gate="G$1" x="20.32" y="429.26" smashed="yes">
<attribute name="VALUE" x="25.4" y="424.18" size="1.778" layer="96"/>
</instance>
<instance part="GND7" gate="G$1" x="20.32" y="396.24" smashed="yes">
<attribute name="VALUE" x="25.4" y="391.16" size="1.778" layer="96"/>
</instance>
<instance part="GND8" gate="G$1" x="20.32" y="363.22" smashed="yes">
<attribute name="VALUE" x="25.4" y="358.14" size="1.778" layer="96"/>
</instance>
<instance part="GND9" gate="G$1" x="76.2" y="124.46"/>
<instance part="GND10" gate="G$1" x="45.72" y="132.08"/>
<instance part="R13" gate="G$1" x="71.12" y="210.82" smashed="yes">
<attribute name="NAME" x="71.12" y="215.9" size="1.778" layer="95"/>
<attribute name="VALUE" x="71.12" y="213.36" size="1.778" layer="95"/>
</instance>
<instance part="U$40" gate="G$1" x="66.04" y="223.52"/>
<instance part="GND11" gate="G$1" x="27.94" y="132.08"/>
<instance part="C28" gate="G$1" x="45.72" y="182.88"/>
<instance part="C27" gate="G$1" x="33.02" y="182.88"/>
<instance part="U$57" gate="G$1" x="45.72" y="109.22"/>
<instance part="U$58" gate="G$1" x="111.76" y="109.22"/>
<instance part="GND12" gate="G$1" x="111.76" y="91.44"/>
<instance part="U$59" gate="G$1" x="360.68" y="76.2"/>
<instance part="GND13" gate="G$1" x="345.44" y="40.64"/>
<instance part="U$60" gate="G$1" x="297.18" y="76.2"/>
<instance part="GND14" gate="G$1" x="287.02" y="38.1"/>
<instance part="U$63" gate="G$1" x="96.52" y="480.06"/>
<instance part="U$64" gate="G$1" x="111.76" y="480.06"/>
<instance part="U$65" gate="G$1" x="119.38" y="398.78"/>
<instance part="GND15" gate="G$1" x="119.38" y="353.06"/>
<instance part="U$66" gate="G$1" x="119.38" y="325.12"/>
<instance part="R4" gate="G$1" x="241.3" y="482.6" rot="R270"/>
<instance part="U$39" gate="G$1" x="231.14" y="497.84"/>
<instance part="U$68" gate="G$1" x="48.26" y="304.8"/>
<instance part="GND16" gate="G$1" x="83.82" y="287.02"/>
<instance part="GND17" gate="G$1" x="119.38" y="281.94"/>
<instance part="C17" gate="G$1" x="20.32" y="335.28"/>
<instance part="C18" gate="G$1" x="27.94" y="335.28"/>
<instance part="GND18" gate="G$1" x="251.46" y="477.52"/>
<instance part="U$71" gate="G$1" x="20.32" y="340.36" smashed="yes">
<attribute name="VALUE" x="25.4" y="345.44" size="1.778" layer="96"/>
</instance>
<instance part="GND19" gate="G$1" x="20.32" y="327.66" smashed="yes">
<attribute name="VALUE" x="25.4" y="322.58" size="1.778" layer="96"/>
</instance>
<instance part="GND20" gate="G$1" x="22.86" y="22.86"/>
<instance part="C20" gate="G$1" x="58.42" y="335.28" smashed="yes">
<attribute name="NAME" x="60.96" y="335.28" size="1.778" layer="95"/>
<attribute name="VALUE" x="60.96" y="332.74" size="1.778" layer="96"/>
</instance>
<instance part="C19" gate="G$1" x="38.1" y="335.28" smashed="yes">
<attribute name="NAME" x="40.64" y="335.28" size="1.778" layer="95"/>
<attribute name="VALUE" x="40.64" y="332.74" size="1.778" layer="96"/>
</instance>
<instance part="U$92" gate="G$1" x="58.42" y="340.36" smashed="yes">
<attribute name="VALUE" x="63.5" y="345.44" size="1.778" layer="96"/>
</instance>
<instance part="U$93" gate="G$1" x="38.1" y="340.36" smashed="yes">
<attribute name="VALUE" x="43.18" y="345.44" size="1.778" layer="96"/>
</instance>
<instance part="GND21" gate="G$1" x="58.42" y="327.66" smashed="yes">
<attribute name="VALUE" x="63.5" y="322.58" size="1.778" layer="96"/>
</instance>
<instance part="GND22" gate="G$1" x="38.1" y="327.66" smashed="yes">
<attribute name="VALUE" x="43.18" y="322.58" size="1.778" layer="96"/>
</instance>
<instance part="R36" gate="G$1" x="147.32" y="33.02" smashed="yes" rot="R270">
<attribute name="NAME" x="149.86" y="30.48" size="1.778" layer="95"/>
<attribute name="VALUE" x="149.86" y="27.94" size="1.778" layer="95"/>
</instance>
<instance part="LED2" gate="G$1" x="147.32" y="40.64" smashed="yes">
<attribute name="NAME" x="149.86" y="38.1" size="1.778" layer="95"/>
<attribute name="VALUE" x="153.67" y="36.83" size="1.778" layer="97" rot="R90"/>
</instance>
<instance part="U$97" gate="G$1" x="147.32" y="45.72"/>
<instance part="R28" gate="G$1" x="335.28" y="132.08"/>
<instance part="R30" gate="G$1" x="335.28" y="129.54"/>
<instance part="R32" gate="G$1" x="335.28" y="127"/>
<instance part="R34" gate="G$1" x="335.28" y="124.46"/>
<instance part="R38" gate="G$1" x="335.28" y="114.3"/>
<instance part="R40" gate="G$1" x="335.28" y="111.76"/>
<instance part="R42" gate="G$1" x="335.28" y="109.22"/>
<instance part="R45" gate="G$1" x="335.28" y="106.68"/>
<instance part="R27" gate="G$1" x="276.86" y="132.08"/>
<instance part="R29" gate="G$1" x="276.86" y="129.54"/>
<instance part="R31" gate="G$1" x="276.86" y="127"/>
<instance part="R33" gate="G$1" x="276.86" y="124.46"/>
<instance part="R37" gate="G$1" x="276.86" y="114.3"/>
<instance part="R39" gate="G$1" x="276.86" y="111.76"/>
<instance part="R41" gate="G$1" x="276.86" y="109.22"/>
<instance part="R44" gate="G$1" x="276.86" y="106.68"/>
<instance part="R5" gate="G$1" x="129.54" y="378.46"/>
<instance part="R7" gate="G$1" x="129.54" y="370.84"/>
<instance part="R8" gate="G$1" x="129.54" y="363.22"/>
<instance part="C24" gate="G$1" x="241.3" y="187.96" smashed="yes">
<attribute name="NAME" x="243.84" y="187.96" size="1.778" layer="95"/>
<attribute name="VALUE" x="243.84" y="185.42" size="1.778" layer="96"/>
</instance>
<instance part="C25" gate="G$1" x="266.7" y="187.96" smashed="yes">
<attribute name="NAME" x="269.24" y="187.96" size="1.778" layer="95"/>
<attribute name="VALUE" x="269.24" y="185.42" size="1.778" layer="96"/>
</instance>
<instance part="C26" gate="G$1" x="292.1" y="187.96" smashed="yes">
<attribute name="NAME" x="294.64" y="187.96" size="1.778" layer="95"/>
<attribute name="VALUE" x="294.64" y="185.42" size="1.778" layer="96"/>
</instance>
<instance part="C23" gate="G$1" x="215.9" y="187.96" smashed="yes">
<attribute name="NAME" x="218.44" y="187.96" size="1.778" layer="95"/>
<attribute name="VALUE" x="218.44" y="185.42" size="1.778" layer="96"/>
</instance>
<instance part="R23" gate="G$1" x="233.68" y="190.5" smashed="yes">
<attribute name="NAME" x="233.68" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="233.68" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="R24" gate="G$1" x="259.08" y="190.5" smashed="yes">
<attribute name="NAME" x="259.08" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="259.08" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="R25" gate="G$1" x="284.48" y="190.5" smashed="yes">
<attribute name="NAME" x="284.48" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="284.48" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="R22" gate="G$1" x="208.28" y="190.5" smashed="yes">
<attribute name="NAME" x="208.28" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="208.28" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="IC5" gate="G$1" x="345.44" y="238.76"/>
<instance part="R20" gate="G$1" x="307.34" y="198.12" smashed="yes" rot="R270">
<attribute name="NAME" x="309.88" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="309.88" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="R21" gate="G$1" x="317.5" y="198.12" smashed="yes" rot="R270">
<attribute name="NAME" x="320.04" y="195.58" size="1.778" layer="95"/>
<attribute name="VALUE" x="320.04" y="193.04" size="1.778" layer="95"/>
</instance>
<instance part="R35" gate="G$1" x="114.3" y="33.02" smashed="yes" rot="R270">
<attribute name="NAME" x="116.84" y="30.48" size="1.778" layer="95"/>
<attribute name="VALUE" x="116.84" y="27.94" size="1.778" layer="95"/>
</instance>
<instance part="LED1" gate="G$1" x="114.3" y="40.64" smashed="yes">
<attribute name="NAME" x="116.84" y="38.1" size="1.778" layer="95"/>
<attribute name="VALUE" x="120.65" y="36.83" size="1.778" layer="97" rot="R90"/>
</instance>
<instance part="U$117" gate="G$1" x="114.3" y="45.72"/>
<instance part="GND23" gate="G$1" x="172.72" y="447.04"/>
<instance part="FB1" gate="G$1" x="53.34" y="203.2"/>
<instance part="R11" gate="G$1" x="81.28" y="297.18" smashed="yes">
<attribute name="NAME" x="81.28" y="302.26" size="1.778" layer="95"/>
<attribute name="VALUE" x="81.28" y="299.72" size="1.778" layer="95"/>
</instance>
<instance part="U$1" gate="G$1" x="299.72" y="157.48"/>
<instance part="U$3" gate="G$1" x="358.14" y="157.48"/>
<instance part="LOGO" gate="G$1" x="223.52" y="58.42"/>
<instance part="JP1" gate="G$1" x="88.9" y="398.78"/>
<instance part="R6" gate="G$1" x="88.9" y="375.92" smashed="yes" rot="R270">
<attribute name="NAME" x="91.44" y="373.38" size="1.778" layer="95"/>
<attribute name="VALUE" x="91.44" y="370.84" size="1.778" layer="95"/>
</instance>
<instance part="GND2" gate="G$1" x="88.9" y="353.06"/>
<instance part="U$4" gate="G$1" x="88.9" y="408.94"/>
<instance part="C22" gate="G$1" x="78.74" y="193.04" smashed="yes">
<attribute name="VALUE" x="72.39" y="190.5" size="1.778" layer="96"/>
<attribute name="NAME" x="72.39" y="193.04" size="1.778" layer="95"/>
</instance>
</instances>
<busses>
</busses>
<nets>
<net name="VCCINT" class="0">
<segment>
<wire x1="35.56" y1="477.52" x2="73.66" y2="477.52" width="0.1524" layer="91"/>
<wire x1="73.66" y1="477.52" x2="73.66" y2="480.06" width="0.1524" layer="91"/>
<wire x1="73.66" y1="480.06" x2="73.66" y2="495.3" width="0.1524" layer="91"/>
<wire x1="35.56" y1="480.06" x2="73.66" y2="480.06" width="0.1524" layer="91"/>
<junction x="73.66" y="480.06"/>
<pinref part="POWER" gate="G$1" pin="VCCINT@2"/>
<pinref part="SUPPLY2" gate="VCCINT" pin="VCCINT"/>
<pinref part="POWER" gate="G$1" pin="VCCINT@1"/>
</segment>
<segment>
<wire x1="251.46" y1="411.48" x2="251.46" y2="408.94" width="0.1524" layer="91"/>
<wire x1="251.46" y1="408.94" x2="261.62" y2="408.94" width="0.1524" layer="91"/>
<wire x1="269.24" y1="408.94" x2="261.62" y2="408.94" width="0.1524" layer="91"/>
<wire x1="269.24" y1="403.86" x2="261.62" y2="403.86" width="0.1524" layer="91"/>
<wire x1="261.62" y1="403.86" x2="261.62" y2="406.4" width="0.1524" layer="91"/>
<wire x1="261.62" y1="406.4" x2="261.62" y2="408.94" width="0.1524" layer="91"/>
<wire x1="269.24" y1="401.32" x2="261.62" y2="401.32" width="0.1524" layer="91"/>
<wire x1="261.62" y1="401.32" x2="261.62" y2="403.86" width="0.1524" layer="91"/>
<wire x1="269.24" y1="406.4" x2="261.62" y2="406.4" width="0.1524" layer="91"/>
<junction x="261.62" y="408.94"/>
<junction x="261.62" y="403.86"/>
<junction x="261.62" y="406.4"/>
<pinref part="SUPPLY1" gate="VCCINT" pin="VCCINT"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCINT_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCINT_3"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCINT_4"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCINT_2"/>
</segment>
<segment>
<wire x1="20.32" y1="441.96" x2="20.32" y2="439.42" width="0.1524" layer="91"/>
<wire x1="20.32" y1="436.88" x2="20.32" y2="439.42" width="0.1524" layer="91"/>
<wire x1="27.94" y1="436.88" x2="27.94" y2="439.42" width="0.1524" layer="91"/>
<wire x1="35.56" y1="436.88" x2="35.56" y2="439.42" width="0.1524" layer="91"/>
<wire x1="27.94" y1="439.42" x2="20.32" y2="439.42" width="0.1524" layer="91"/>
<wire x1="20.32" y1="439.42" x2="35.56" y2="439.42" width="0.1524" layer="91"/>
<wire x1="43.18" y1="436.88" x2="43.18" y2="439.42" width="0.1524" layer="91"/>
<wire x1="35.56" y1="439.42" x2="43.18" y2="439.42" width="0.1524" layer="91"/>
<junction x="27.94" y="439.42"/>
<junction x="20.32" y="439.42"/>
<junction x="35.56" y="439.42"/>
<pinref part="SUPPLY3" gate="VCCINT" pin="VCCINT"/>
<pinref part="C1" gate="G$1" pin="1"/>
<pinref part="C2" gate="G$1" pin="1"/>
<pinref part="C3" gate="G$1" pin="1"/>
<pinref part="C4" gate="G$1" pin="1"/>
</segment>
</net>
<net name="VCCAUX" class="0">
<segment>
<wire x1="261.62" y1="393.7" x2="269.24" y2="393.7" width="0.1524" layer="91"/>
<wire x1="269.24" y1="391.16" x2="261.62" y2="391.16" width="0.1524" layer="91"/>
<wire x1="261.62" y1="393.7" x2="261.62" y2="391.16" width="0.1524" layer="91"/>
<wire x1="261.62" y1="388.62" x2="269.24" y2="388.62" width="0.1524" layer="91"/>
<wire x1="261.62" y1="391.16" x2="261.62" y2="388.62" width="0.1524" layer="91"/>
<wire x1="269.24" y1="386.08" x2="261.62" y2="386.08" width="0.1524" layer="91"/>
<wire x1="261.62" y1="388.62" x2="261.62" y2="386.08" width="0.1524" layer="91"/>
<wire x1="251.46" y1="396.24" x2="251.46" y2="393.7" width="0.1524" layer="91"/>
<wire x1="251.46" y1="393.7" x2="261.62" y2="393.7" width="0.1524" layer="91"/>
<junction x="261.62" y="391.16"/>
<junction x="261.62" y="388.62"/>
<junction x="261.62" y="393.7"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCAUX_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCAUX_2"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCAUX_3"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCAUX_4"/>
<pinref part="U$7" gate="G$1" pin="VCCAUX"/>
</segment>
<segment>
<wire x1="20.32" y1="403.86" x2="20.32" y2="406.4" width="0.1524" layer="91"/>
<wire x1="27.94" y1="403.86" x2="27.94" y2="406.4" width="0.1524" layer="91"/>
<wire x1="35.56" y1="403.86" x2="35.56" y2="406.4" width="0.1524" layer="91"/>
<wire x1="27.94" y1="406.4" x2="20.32" y2="406.4" width="0.1524" layer="91"/>
<wire x1="20.32" y1="406.4" x2="35.56" y2="406.4" width="0.1524" layer="91"/>
<wire x1="43.18" y1="403.86" x2="43.18" y2="406.4" width="0.1524" layer="91"/>
<wire x1="35.56" y1="406.4" x2="43.18" y2="406.4" width="0.1524" layer="91"/>
<wire x1="20.32" y1="408.94" x2="20.32" y2="406.4" width="0.1524" layer="91"/>
<junction x="27.94" y="406.4"/>
<junction x="35.56" y="406.4"/>
<junction x="20.32" y="406.4"/>
<pinref part="C5" gate="G$1" pin="1"/>
<pinref part="C6" gate="G$1" pin="1"/>
<pinref part="C7" gate="G$1" pin="1"/>
<pinref part="C8" gate="G$1" pin="1"/>
<pinref part="U$9" gate="G$1" pin="VCCAUX"/>
</segment>
<segment>
<wire x1="55.88" y1="495.3" x2="55.88" y2="485.14" width="0.1524" layer="91"/>
<wire x1="55.88" y1="485.14" x2="55.88" y2="482.6" width="0.1524" layer="91"/>
<wire x1="55.88" y1="482.6" x2="35.56" y2="482.6" width="0.1524" layer="91"/>
<wire x1="35.56" y1="485.14" x2="55.88" y2="485.14" width="0.1524" layer="91"/>
<junction x="55.88" y="485.14"/>
<pinref part="U$89" gate="G$1" pin="VCCAUX"/>
<pinref part="POWER" gate="G$1" pin="VCCAUX@2"/>
<pinref part="POWER" gate="G$1" pin="VCCAUX@1"/>
</segment>
<segment>
<wire x1="96.52" y1="480.06" x2="96.52" y2="462.28" width="0.1524" layer="91"/>
<wire x1="96.52" y1="462.28" x2="96.52" y2="457.2" width="0.1524" layer="91"/>
<wire x1="96.52" y1="457.2" x2="116.84" y2="457.2" width="0.1524" layer="91"/>
<wire x1="116.84" y1="462.28" x2="96.52" y2="462.28" width="0.1524" layer="91"/>
<junction x="96.52" y="462.28"/>
<pinref part="U$63" gate="G$1" pin="VCCAUX"/>
<pinref part="IC2" gate="XCF02S" pin="VCCJ"/>
<pinref part="IC2" gate="XCF02S" pin="VCCO"/>
</segment>
<segment>
<wire x1="241.3" y1="492.76" x2="241.3" y2="482.6" width="0.1524" layer="91"/>
<wire x1="236.22" y1="492.76" x2="236.22" y2="482.6" width="0.1524" layer="91"/>
<wire x1="236.22" y1="492.76" x2="241.3" y2="492.76" width="0.1524" layer="91"/>
<wire x1="231.14" y1="497.84" x2="231.14" y2="492.76" width="0.1524" layer="91"/>
<wire x1="231.14" y1="492.76" x2="236.22" y2="492.76" width="0.1524" layer="91"/>
<wire x1="231.14" y1="492.76" x2="231.14" y2="482.6" width="0.1524" layer="91"/>
<wire x1="226.06" y1="482.6" x2="226.06" y2="492.76" width="0.1524" layer="91"/>
<wire x1="226.06" y1="492.76" x2="231.14" y2="492.76" width="0.1524" layer="91"/>
<junction x="236.22" y="492.76"/>
<junction x="231.14" y="492.76"/>
<pinref part="R4" gate="G$1" pin="1"/>
<pinref part="R3" gate="G$1" pin="1"/>
<pinref part="U$39" gate="G$1" pin="VCCAUX"/>
<pinref part="R2" gate="G$1" pin="1"/>
<pinref part="R1" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="88.9" y1="408.94" x2="88.9" y2="403.86" width="0.1524" layer="91"/>
<pinref part="U$4" gate="G$1" pin="VCCAUX"/>
<pinref part="JP1" gate="G$1" pin="P$1"/>
</segment>
</net>
<net name="VCCO" class="0">
<segment>
<wire x1="269.24" y1="368.3" x2="261.62" y2="368.3" width="0.1524" layer="91"/>
<wire x1="269.24" y1="378.46" x2="261.62" y2="378.46" width="0.1524" layer="91"/>
<wire x1="269.24" y1="375.92" x2="261.62" y2="375.92" width="0.1524" layer="91"/>
<wire x1="261.62" y1="378.46" x2="261.62" y2="375.92" width="0.1524" layer="91"/>
<wire x1="269.24" y1="370.84" x2="261.62" y2="370.84" width="0.1524" layer="91"/>
<wire x1="261.62" y1="375.92" x2="261.62" y2="370.84" width="0.1524" layer="91"/>
<wire x1="261.62" y1="368.3" x2="261.62" y2="370.84" width="0.1524" layer="91"/>
<wire x1="269.24" y1="363.22" x2="261.62" y2="363.22" width="0.1524" layer="91"/>
<wire x1="261.62" y1="368.3" x2="261.62" y2="363.22" width="0.1524" layer="91"/>
<wire x1="269.24" y1="360.68" x2="261.62" y2="360.68" width="0.1524" layer="91"/>
<wire x1="261.62" y1="363.22" x2="261.62" y2="360.68" width="0.1524" layer="91"/>
<wire x1="269.24" y1="355.6" x2="261.62" y2="355.6" width="0.1524" layer="91"/>
<wire x1="261.62" y1="360.68" x2="261.62" y2="355.6" width="0.1524" layer="91"/>
<wire x1="269.24" y1="353.06" x2="261.62" y2="353.06" width="0.1524" layer="91"/>
<wire x1="261.62" y1="355.6" x2="261.62" y2="353.06" width="0.1524" layer="91"/>
<wire x1="251.46" y1="381" x2="251.46" y2="378.46" width="0.1524" layer="91"/>
<wire x1="251.46" y1="378.46" x2="261.62" y2="378.46" width="0.1524" layer="91"/>
<junction x="261.62" y="375.92"/>
<junction x="261.62" y="370.84"/>
<junction x="261.62" y="368.3"/>
<junction x="261.62" y="363.22"/>
<junction x="261.62" y="360.68"/>
<junction x="261.62" y="355.6"/>
<junction x="261.62" y="378.46"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK1_2"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK0_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK0_2"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK1_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK2_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK2_2"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK3_1"/>
<pinref part="IC1" gate="XC3S250E" pin="VCCO_BK3_2"/>
<pinref part="U$10" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="20.32" y1="375.92" x2="20.32" y2="373.38" width="0.1524" layer="91"/>
<wire x1="27.94" y1="370.84" x2="27.94" y2="373.38" width="0.1524" layer="91"/>
<wire x1="20.32" y1="370.84" x2="20.32" y2="373.38" width="0.1524" layer="91"/>
<wire x1="27.94" y1="373.38" x2="20.32" y2="373.38" width="0.1524" layer="91"/>
<wire x1="35.56" y1="370.84" x2="35.56" y2="373.38" width="0.1524" layer="91"/>
<wire x1="20.32" y1="373.38" x2="35.56" y2="373.38" width="0.1524" layer="91"/>
<wire x1="43.18" y1="370.84" x2="43.18" y2="373.38" width="0.1524" layer="91"/>
<wire x1="35.56" y1="373.38" x2="43.18" y2="373.38" width="0.1524" layer="91"/>
<wire x1="50.8" y1="370.84" x2="50.8" y2="373.38" width="0.1524" layer="91"/>
<wire x1="43.18" y1="373.38" x2="50.8" y2="373.38" width="0.1524" layer="91"/>
<wire x1="50.8" y1="373.38" x2="58.42" y2="373.38" width="0.1524" layer="91"/>
<wire x1="58.42" y1="373.38" x2="58.42" y2="370.84" width="0.1524" layer="91"/>
<wire x1="58.42" y1="373.38" x2="66.04" y2="373.38" width="0.1524" layer="91"/>
<wire x1="66.04" y1="373.38" x2="66.04" y2="370.84" width="0.1524" layer="91"/>
<wire x1="73.66" y1="370.84" x2="73.66" y2="373.38" width="0.1524" layer="91"/>
<wire x1="66.04" y1="373.38" x2="73.66" y2="373.38" width="0.1524" layer="91"/>
<junction x="27.94" y="373.38"/>
<junction x="20.32" y="373.38"/>
<junction x="35.56" y="373.38"/>
<junction x="43.18" y="373.38"/>
<junction x="50.8" y="373.38"/>
<junction x="58.42" y="373.38"/>
<junction x="66.04" y="373.38"/>
<pinref part="U$11" gate="G$1" pin="VCCO"/>
<pinref part="C10" gate="G$1" pin="1"/>
<pinref part="C9" gate="G$1" pin="1"/>
<pinref part="C11" gate="G$1" pin="1"/>
<pinref part="C12" gate="G$1" pin="1"/>
<pinref part="C13" gate="G$1" pin="1"/>
<pinref part="C14" gate="G$1" pin="1"/>
<pinref part="C15" gate="G$1" pin="1"/>
<pinref part="C16" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="205.74" y1="241.3" x2="205.74" y2="236.22" width="0.1524" layer="91"/>
<wire x1="205.74" y1="236.22" x2="205.74" y2="203.2" width="0.1524" layer="91"/>
<wire x1="281.94" y1="236.22" x2="256.54" y2="236.22" width="0.1524" layer="91"/>
<wire x1="256.54" y1="236.22" x2="231.14" y2="236.22" width="0.1524" layer="91"/>
<wire x1="231.14" y1="236.22" x2="205.74" y2="236.22" width="0.1524" layer="91"/>
<wire x1="231.14" y1="203.2" x2="231.14" y2="236.22" width="0.1524" layer="91"/>
<wire x1="256.54" y1="203.2" x2="256.54" y2="236.22" width="0.1524" layer="91"/>
<wire x1="281.94" y1="203.2" x2="281.94" y2="236.22" width="0.1524" layer="91"/>
<wire x1="332.74" y1="236.22" x2="281.94" y2="236.22" width="0.1524" layer="91"/>
<junction x="205.74" y="236.22"/>
<junction x="231.14" y="236.22"/>
<junction x="256.54" y="236.22"/>
<junction x="281.94" y="236.22"/>
<pinref part="R14" gate="G$1" pin="1"/>
<pinref part="U$48" gate="G$1" pin="VCCO"/>
<pinref part="R15" gate="G$1" pin="1"/>
<pinref part="R16" gate="G$1" pin="1"/>
<pinref part="R17" gate="G$1" pin="1"/>
<pinref part="IC5" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="40.64" y1="495.3" x2="40.64" y2="490.22" width="0.1524" layer="91"/>
<wire x1="40.64" y1="490.22" x2="40.64" y2="487.68" width="0.1524" layer="91"/>
<wire x1="40.64" y1="487.68" x2="35.56" y2="487.68" width="0.1524" layer="91"/>
<wire x1="35.56" y1="490.22" x2="40.64" y2="490.22" width="0.1524" layer="91"/>
<junction x="40.64" y="490.22"/>
<pinref part="U$2" gate="G$1" pin="VCCO"/>
<pinref part="POWER" gate="G$1" pin="VCCO@2"/>
<pinref part="POWER" gate="G$1" pin="VCCO@1"/>
</segment>
<segment>
<wire x1="66.04" y1="223.52" x2="66.04" y2="198.12" width="0.1524" layer="91"/>
<wire x1="66.04" y1="198.12" x2="66.04" y2="195.58" width="0.1524" layer="91"/>
<wire x1="66.04" y1="195.58" x2="81.28" y2="195.58" width="0.1524" layer="91"/>
<wire x1="66.04" y1="198.12" x2="81.28" y2="198.12" width="0.1524" layer="91"/>
<junction x="66.04" y="198.12"/>
<pinref part="U$40" gate="G$1" pin="VCCO"/>
<pinref part="IC6" gate="FT2232L" pin="VCCIOB"/>
<pinref part="IC6" gate="FT2232L" pin="VCCIOA"/>
</segment>
<segment>
<wire x1="50.8" y1="96.52" x2="45.72" y2="96.52" width="0.1524" layer="91"/>
<wire x1="45.72" y1="96.52" x2="45.72" y2="109.22" width="0.1524" layer="91"/>
<pinref part="R46" gate="G$1" pin="2"/>
<pinref part="U$57" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="106.68" y1="104.14" x2="111.76" y2="104.14" width="0.1524" layer="91"/>
<wire x1="111.76" y1="104.14" x2="111.76" y2="109.22" width="0.1524" layer="91"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="VCC"/>
<pinref part="U$58" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="347.98" y1="73.66" x2="347.98" y2="45.72" width="0.1524" layer="91"/>
<wire x1="347.98" y1="45.72" x2="360.68" y2="45.72" width="0.1524" layer="91"/>
<wire x1="360.68" y1="45.72" x2="360.68" y2="76.2" width="0.1524" layer="91"/>
<pinref part="VPE" gate="G$1" pin="10"/>
<pinref part="U$59" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="289.56" y1="73.66" x2="289.56" y2="45.72" width="0.1524" layer="91"/>
<wire x1="289.56" y1="45.72" x2="297.18" y2="45.72" width="0.1524" layer="91"/>
<wire x1="297.18" y1="45.72" x2="297.18" y2="76.2" width="0.1524" layer="91"/>
<pinref part="VPD" gate="G$1" pin="10"/>
<pinref part="U$60" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="111.76" y1="480.06" x2="111.76" y2="467.36" width="0.1524" layer="91"/>
<wire x1="111.76" y1="467.36" x2="116.84" y2="467.36" width="0.1524" layer="91"/>
<pinref part="U$64" gate="G$1" pin="VCCO"/>
<pinref part="IC2" gate="XCF02S" pin="VCCI"/>
</segment>
<segment>
<wire x1="119.38" y1="398.78" x2="119.38" y2="393.7" width="0.1524" layer="91"/>
<wire x1="119.38" y1="393.7" x2="142.24" y2="393.7" width="0.1524" layer="91"/>
<pinref part="U$65" gate="G$1" pin="VCCO"/>
<pinref part="IC3" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="48.26" y1="304.8" x2="48.26" y2="297.18" width="0.1524" layer="91"/>
<wire x1="48.26" y1="297.18" x2="55.88" y2="297.18" width="0.1524" layer="91"/>
<pinref part="U$68" gate="G$1" pin="VCCO"/>
<pinref part="U1" gate="UNIT_KC7050B" pin="4"/>
</segment>
<segment>
<wire x1="119.38" y1="320.04" x2="119.38" y2="325.12" width="0.1524" layer="91"/>
<wire x1="119.38" y1="320.04" x2="165.1" y2="320.04" width="0.1524" layer="91"/>
<wire x1="165.1" y1="320.04" x2="170.18" y2="320.04" width="0.1524" layer="91"/>
<wire x1="160.02" y1="307.34" x2="165.1" y2="307.34" width="0.1524" layer="91"/>
<wire x1="165.1" y1="307.34" x2="165.1" y2="320.04" width="0.1524" layer="91"/>
<wire x1="119.38" y1="320.04" x2="119.38" y2="314.96" width="0.1524" layer="91"/>
<wire x1="170.18" y1="320.04" x2="170.18" y2="314.96" width="0.1524" layer="91"/>
<junction x="165.1" y="320.04"/>
<junction x="119.38" y="320.04"/>
<pinref part="U$66" gate="G$1" pin="VCCO"/>
<pinref part="IC4" gate="G$1" pin="VCCO"/>
<pinref part="R9" gate="G$1" pin="1"/>
<pinref part="R10" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="20.32" y1="340.36" x2="20.32" y2="337.82" width="0.1524" layer="91"/>
<wire x1="20.32" y1="337.82" x2="20.32" y2="335.28" width="0.1524" layer="91"/>
<wire x1="20.32" y1="337.82" x2="27.94" y2="337.82" width="0.1524" layer="91"/>
<wire x1="27.94" y1="337.82" x2="27.94" y2="335.28" width="0.1524" layer="91"/>
<junction x="20.32" y="337.82"/>
<pinref part="U$71" gate="G$1" pin="VCCO"/>
<pinref part="C17" gate="G$1" pin="1"/>
<pinref part="C18" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="58.42" y1="340.36" x2="58.42" y2="335.28" width="0.1524" layer="91"/>
<pinref part="U$92" gate="G$1" pin="VCCO"/>
<pinref part="C20" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="38.1" y1="335.28" x2="38.1" y2="340.36" width="0.1524" layer="91"/>
<pinref part="C19" gate="G$1" pin="1"/>
<pinref part="U$93" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="147.32" y1="45.72" x2="147.32" y2="43.18" width="0.1524" layer="91"/>
<pinref part="U$97" gate="G$1" pin="VCCO"/>
<pinref part="LED2" gate="G$1" pin="A"/>
</segment>
<segment>
<wire x1="114.3" y1="45.72" x2="114.3" y2="43.18" width="0.1524" layer="91"/>
<pinref part="U$117" gate="G$1" pin="VCCO"/>
<pinref part="LED1" gate="G$1" pin="A"/>
</segment>
<segment>
<wire x1="299.72" y1="152.4" x2="299.72" y2="154.94" width="0.1524" layer="91"/>
<wire x1="302.26" y1="152.4" x2="302.26" y2="154.94" width="0.1524" layer="91"/>
<wire x1="302.26" y1="154.94" x2="299.72" y2="154.94" width="0.1524" layer="91"/>
<wire x1="299.72" y1="154.94" x2="299.72" y2="157.48" width="0.1524" layer="91"/>
<junction x="299.72" y="154.94"/>
<pinref part="7SEG1" gate="7SEG" pin="COMMON8"/>
<pinref part="7SEG1" gate="7SEG" pin="COMMON3"/>
<pinref part="U$1" gate="G$1" pin="VCCO"/>
</segment>
<segment>
<wire x1="358.14" y1="152.4" x2="358.14" y2="154.94" width="0.1524" layer="91"/>
<wire x1="360.68" y1="152.4" x2="360.68" y2="154.94" width="0.1524" layer="91"/>
<wire x1="360.68" y1="154.94" x2="358.14" y2="154.94" width="0.1524" layer="91"/>
<wire x1="358.14" y1="154.94" x2="358.14" y2="157.48" width="0.1524" layer="91"/>
<junction x="358.14" y="154.94"/>
<pinref part="7SEG2" gate="7SEG" pin="COMMON8"/>
<pinref part="7SEG2" gate="7SEG" pin="COMMON3"/>
<pinref part="U$3" gate="G$1" pin="VCCO"/>
</segment>
</net>
<net name="TDIO" class="0">
<segment>
<wire x1="116.84" y1="439.42" x2="104.14" y2="439.42" width="0.1524" layer="91"/>
<label x="104.14" y="439.42" size="1.778" layer="95"/>
<pinref part="IC2" gate="XCF02S" pin="TDO"/>
</segment>
<segment>
<wire x1="269.24" y1="434.34" x2="259.08" y2="434.34" width="0.1524" layer="91"/>
<label x="259.08" y="434.34" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="TDI"/>
</segment>
</net>
<net name="TCK_FTDI" class="0">
<segment>
<wire x1="142.24" y1="210.82" x2="157.48" y2="210.82" width="0.1524" layer="91"/>
<label x="144.78" y="210.82" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="ADBUS0"/>
</segment>
<segment>
<wire x1="129.54" y1="370.84" x2="101.6" y2="370.84" width="0.1524" layer="91"/>
<label x="101.6" y="370.84" size="1.778" layer="95"/>
<pinref part="R7" gate="G$1" pin="1"/>
</segment>
</net>
<net name="TDI_FTDI" class="0">
<segment>
<wire x1="142.24" y1="208.28" x2="157.48" y2="208.28" width="0.1524" layer="91"/>
<label x="144.78" y="208.28" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="ADBUS1"/>
</segment>
<segment>
<wire x1="129.54" y1="363.22" x2="101.6" y2="363.22" width="0.1524" layer="91"/>
<label x="101.6" y="363.22" size="1.778" layer="95"/>
<pinref part="R8" gate="G$1" pin="1"/>
</segment>
</net>
<net name="TDO_FTDI" class="0">
<segment>
<wire x1="142.24" y1="205.74" x2="157.48" y2="205.74" width="0.1524" layer="91"/>
<label x="144.78" y="205.74" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="ADBUS2"/>
</segment>
<segment>
<wire x1="167.64" y1="386.08" x2="187.96" y2="386.08" width="0.1524" layer="91"/>
<label x="175.26" y="386.08" size="1.778" layer="95"/>
<pinref part="IC3" gate="G$1" pin="1Y"/>
</segment>
</net>
<net name="TMS_FTDI" class="0">
<segment>
<wire x1="142.24" y1="203.2" x2="157.48" y2="203.2" width="0.1524" layer="91"/>
<label x="144.78" y="203.2" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="ADBUS3"/>
</segment>
<segment>
<wire x1="129.54" y1="378.46" x2="101.6" y2="378.46" width="0.1524" layer="91"/>
<label x="101.6" y="378.46" size="1.778" layer="95"/>
<pinref part="R5" gate="G$1" pin="1"/>
</segment>
</net>
<net name="TCK" class="0">
<segment>
<wire x1="269.24" y1="439.42" x2="259.08" y2="439.42" width="0.1524" layer="91"/>
<label x="259.08" y="439.42" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="TCK"/>
</segment>
<segment>
<wire x1="116.84" y1="447.04" x2="104.14" y2="447.04" width="0.1524" layer="91"/>
<label x="104.14" y="447.04" size="1.778" layer="95"/>
<pinref part="IC2" gate="XCF02S" pin="TCK"/>
</segment>
<segment>
<wire x1="167.64" y1="370.84" x2="187.96" y2="370.84" width="0.1524" layer="91"/>
<label x="175.26" y="370.84" size="1.778" layer="95"/>
<pinref part="IC3" gate="G$1" pin="3Y"/>
</segment>
</net>
<net name="TDI" class="0">
<segment>
<wire x1="116.84" y1="444.5" x2="104.14" y2="444.5" width="0.1524" layer="91"/>
<label x="104.14" y="444.5" size="1.778" layer="95"/>
<pinref part="IC2" gate="XCF02S" pin="TDI"/>
</segment>
<segment>
<wire x1="167.64" y1="363.22" x2="187.96" y2="363.22" width="0.1524" layer="91"/>
<label x="175.26" y="363.22" size="1.778" layer="95"/>
<pinref part="IC3" gate="G$1" pin="4Y"/>
</segment>
</net>
<net name="TDO" class="0">
<segment>
<wire x1="269.24" y1="431.8" x2="259.08" y2="431.8" width="0.1524" layer="91"/>
<label x="259.08" y="431.8" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="TDO"/>
</segment>
<segment>
<wire x1="142.24" y1="386.08" x2="101.6" y2="386.08" width="0.1524" layer="91"/>
<label x="101.6" y="386.08" size="1.778" layer="95"/>
<pinref part="IC3" gate="G$1" pin="1A"/>
</segment>
</net>
<net name="TMS" class="0">
<segment>
<wire x1="116.84" y1="441.96" x2="104.14" y2="441.96" width="0.1524" layer="91"/>
<label x="104.14" y="441.96" size="1.778" layer="95"/>
<pinref part="IC2" gate="XCF02S" pin="TMS"/>
</segment>
<segment>
<wire x1="269.24" y1="436.88" x2="259.08" y2="436.88" width="0.1524" layer="91"/>
<label x="259.08" y="436.88" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="TMS"/>
</segment>
<segment>
<wire x1="167.64" y1="378.46" x2="187.96" y2="378.46" width="0.1524" layer="91"/>
<label x="175.26" y="378.46" size="1.778" layer="95"/>
<pinref part="IC3" gate="G$1" pin="2Y"/>
</segment>
</net>
<net name="GND" class="0">
<segment>
<wire x1="269.24" y1="317.5" x2="261.62" y2="317.5" width="0.1524" layer="91"/>
<wire x1="261.62" y1="317.5" x2="261.62" y2="312.42" width="0.1524" layer="91"/>
<wire x1="269.24" y1="320.04" x2="261.62" y2="320.04" width="0.1524" layer="91"/>
<wire x1="261.62" y1="320.04" x2="261.62" y2="317.5" width="0.1524" layer="91"/>
<wire x1="269.24" y1="345.44" x2="261.62" y2="345.44" width="0.1524" layer="91"/>
<wire x1="261.62" y1="345.44" x2="261.62" y2="342.9" width="0.1524" layer="91"/>
<wire x1="261.62" y1="342.9" x2="261.62" y2="340.36" width="0.1524" layer="91"/>
<wire x1="261.62" y1="340.36" x2="261.62" y2="337.82" width="0.1524" layer="91"/>
<wire x1="261.62" y1="337.82" x2="261.62" y2="335.28" width="0.1524" layer="91"/>
<wire x1="261.62" y1="335.28" x2="261.62" y2="332.74" width="0.1524" layer="91"/>
<wire x1="261.62" y1="332.74" x2="261.62" y2="330.2" width="0.1524" layer="91"/>
<wire x1="261.62" y1="330.2" x2="261.62" y2="327.66" width="0.1524" layer="91"/>
<wire x1="261.62" y1="327.66" x2="261.62" y2="325.12" width="0.1524" layer="91"/>
<wire x1="261.62" y1="325.12" x2="261.62" y2="322.58" width="0.1524" layer="91"/>
<wire x1="261.62" y1="322.58" x2="261.62" y2="320.04" width="0.1524" layer="91"/>
<wire x1="269.24" y1="342.9" x2="261.62" y2="342.9" width="0.1524" layer="91"/>
<wire x1="261.62" y1="340.36" x2="269.24" y2="340.36" width="0.1524" layer="91"/>
<wire x1="269.24" y1="337.82" x2="261.62" y2="337.82" width="0.1524" layer="91"/>
<wire x1="261.62" y1="335.28" x2="269.24" y2="335.28" width="0.1524" layer="91"/>
<wire x1="269.24" y1="332.74" x2="261.62" y2="332.74" width="0.1524" layer="91"/>
<wire x1="261.62" y1="330.2" x2="269.24" y2="330.2" width="0.1524" layer="91"/>
<wire x1="269.24" y1="327.66" x2="261.62" y2="327.66" width="0.1524" layer="91"/>
<wire x1="261.62" y1="325.12" x2="269.24" y2="325.12" width="0.1524" layer="91"/>
<wire x1="269.24" y1="322.58" x2="261.62" y2="322.58" width="0.1524" layer="91"/>
<junction x="261.62" y="317.5"/>
<junction x="261.62" y="320.04"/>
<junction x="261.62" y="342.9"/>
<junction x="261.62" y="340.36"/>
<junction x="261.62" y="337.82"/>
<junction x="261.62" y="335.28"/>
<junction x="261.62" y="332.74"/>
<junction x="261.62" y="330.2"/>
<junction x="261.62" y="327.66"/>
<junction x="261.62" y="325.12"/>
<junction x="261.62" y="322.58"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_12"/>
<pinref part="GND" gate="G$1" pin="GND"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_11"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_1"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_2"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_3"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_4"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_5"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_6"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_7"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_8"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_9"/>
<pinref part="IC1" gate="XC3S250E" pin="GND_10"/>
</segment>
<segment>
<wire x1="35.56" y1="469.9" x2="40.64" y2="469.9" width="0.1524" layer="91"/>
<wire x1="35.56" y1="474.98" x2="40.64" y2="474.98" width="0.1524" layer="91"/>
<wire x1="40.64" y1="474.98" x2="40.64" y2="472.44" width="0.1524" layer="91"/>
<wire x1="40.64" y1="472.44" x2="40.64" y2="469.9" width="0.1524" layer="91"/>
<wire x1="35.56" y1="472.44" x2="40.64" y2="472.44" width="0.1524" layer="91"/>
<wire x1="40.64" y1="464.82" x2="40.64" y2="467.36" width="0.1524" layer="91"/>
<wire x1="40.64" y1="467.36" x2="40.64" y2="469.9" width="0.1524" layer="91"/>
<wire x1="35.56" y1="467.36" x2="40.64" y2="467.36" width="0.1524" layer="91"/>
<junction x="40.64" y="472.44"/>
<junction x="40.64" y="469.9"/>
<junction x="40.64" y="467.36"/>
<pinref part="POWER" gate="G$1" pin="GND@3"/>
<pinref part="POWER" gate="G$1" pin="GND@1"/>
<pinref part="POWER" gate="G$1" pin="GND@2"/>
<pinref part="GND5" gate="G$1" pin="GND"/>
<pinref part="POWER" gate="G$1" pin="GND@4"/>
</segment>
<segment>
<wire x1="20.32" y1="429.26" x2="20.32" y2="431.8" width="0.1524" layer="91"/>
<wire x1="20.32" y1="431.8" x2="20.32" y2="434.34" width="0.1524" layer="91"/>
<wire x1="20.32" y1="431.8" x2="27.94" y2="431.8" width="0.1524" layer="91"/>
<wire x1="27.94" y1="431.8" x2="35.56" y2="431.8" width="0.1524" layer="91"/>
<wire x1="35.56" y1="431.8" x2="43.18" y2="431.8" width="0.1524" layer="91"/>
<wire x1="43.18" y1="431.8" x2="43.18" y2="434.34" width="0.1524" layer="91"/>
<wire x1="35.56" y1="434.34" x2="35.56" y2="431.8" width="0.1524" layer="91"/>
<wire x1="27.94" y1="434.34" x2="27.94" y2="431.8" width="0.1524" layer="91"/>
<junction x="20.32" y="431.8"/>
<junction x="35.56" y="431.8"/>
<junction x="27.94" y="431.8"/>
<pinref part="GND6" gate="G$1" pin="GND"/>
<pinref part="C1" gate="G$1" pin="2"/>
<pinref part="C4" gate="G$1" pin="2"/>
<pinref part="C3" gate="G$1" pin="2"/>
<pinref part="C2" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="20.32" y1="363.22" x2="20.32" y2="365.76" width="0.1524" layer="91"/>
<wire x1="20.32" y1="365.76" x2="20.32" y2="368.3" width="0.1524" layer="91"/>
<wire x1="20.32" y1="365.76" x2="27.94" y2="365.76" width="0.1524" layer="91"/>
<wire x1="27.94" y1="365.76" x2="35.56" y2="365.76" width="0.1524" layer="91"/>
<wire x1="35.56" y1="365.76" x2="43.18" y2="365.76" width="0.1524" layer="91"/>
<wire x1="43.18" y1="365.76" x2="50.8" y2="365.76" width="0.1524" layer="91"/>
<wire x1="50.8" y1="365.76" x2="58.42" y2="365.76" width="0.1524" layer="91"/>
<wire x1="58.42" y1="365.76" x2="66.04" y2="365.76" width="0.1524" layer="91"/>
<wire x1="66.04" y1="365.76" x2="73.66" y2="365.76" width="0.1524" layer="91"/>
<wire x1="73.66" y1="365.76" x2="73.66" y2="368.3" width="0.1524" layer="91"/>
<wire x1="66.04" y1="368.3" x2="66.04" y2="365.76" width="0.1524" layer="91"/>
<wire x1="58.42" y1="368.3" x2="58.42" y2="365.76" width="0.1524" layer="91"/>
<wire x1="50.8" y1="368.3" x2="50.8" y2="365.76" width="0.1524" layer="91"/>
<wire x1="43.18" y1="368.3" x2="43.18" y2="365.76" width="0.1524" layer="91"/>
<wire x1="35.56" y1="368.3" x2="35.56" y2="365.76" width="0.1524" layer="91"/>
<wire x1="27.94" y1="368.3" x2="27.94" y2="365.76" width="0.1524" layer="91"/>
<junction x="20.32" y="365.76"/>
<junction x="66.04" y="365.76"/>
<junction x="58.42" y="365.76"/>
<junction x="50.8" y="365.76"/>
<junction x="43.18" y="365.76"/>
<junction x="35.56" y="365.76"/>
<junction x="27.94" y="365.76"/>
<pinref part="GND8" gate="G$1" pin="GND"/>
<pinref part="C9" gate="G$1" pin="2"/>
<pinref part="C16" gate="G$1" pin="2"/>
<pinref part="C15" gate="G$1" pin="2"/>
<pinref part="C14" gate="G$1" pin="2"/>
<pinref part="C13" gate="G$1" pin="2"/>
<pinref part="C12" gate="G$1" pin="2"/>
<pinref part="C11" gate="G$1" pin="2"/>
<pinref part="C10" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="20.32" y1="396.24" x2="20.32" y2="398.78" width="0.1524" layer="91"/>
<wire x1="20.32" y1="398.78" x2="20.32" y2="401.32" width="0.1524" layer="91"/>
<wire x1="43.18" y1="401.32" x2="43.18" y2="398.78" width="0.1524" layer="91"/>
<wire x1="43.18" y1="398.78" x2="35.56" y2="398.78" width="0.1524" layer="91"/>
<wire x1="35.56" y1="398.78" x2="27.94" y2="398.78" width="0.1524" layer="91"/>
<wire x1="27.94" y1="398.78" x2="20.32" y2="398.78" width="0.1524" layer="91"/>
<wire x1="27.94" y1="401.32" x2="27.94" y2="398.78" width="0.1524" layer="91"/>
<wire x1="35.56" y1="401.32" x2="35.56" y2="398.78" width="0.1524" layer="91"/>
<junction x="20.32" y="398.78"/>
<junction x="27.94" y="398.78"/>
<junction x="35.56" y="398.78"/>
<pinref part="GND7" gate="G$1" pin="GND"/>
<pinref part="C5" gate="G$1" pin="2"/>
<pinref part="C8" gate="G$1" pin="2"/>
<pinref part="C6" gate="G$1" pin="2"/>
<pinref part="C7" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="81.28" y1="142.24" x2="76.2" y2="142.24" width="0.1524" layer="91"/>
<wire x1="76.2" y1="142.24" x2="76.2" y2="137.16" width="0.1524" layer="91"/>
<wire x1="76.2" y1="137.16" x2="76.2" y2="134.62" width="0.1524" layer="91"/>
<wire x1="76.2" y1="134.62" x2="76.2" y2="132.08" width="0.1524" layer="91"/>
<wire x1="76.2" y1="132.08" x2="76.2" y2="129.54" width="0.1524" layer="91"/>
<wire x1="76.2" y1="129.54" x2="76.2" y2="124.46" width="0.1524" layer="91"/>
<wire x1="81.28" y1="129.54" x2="76.2" y2="129.54" width="0.1524" layer="91"/>
<wire x1="81.28" y1="132.08" x2="76.2" y2="132.08" width="0.1524" layer="91"/>
<wire x1="81.28" y1="134.62" x2="76.2" y2="134.62" width="0.1524" layer="91"/>
<wire x1="81.28" y1="137.16" x2="76.2" y2="137.16" width="0.1524" layer="91"/>
<wire x1="81.28" y1="147.32" x2="76.2" y2="147.32" width="0.1524" layer="91"/>
<wire x1="76.2" y1="147.32" x2="76.2" y2="142.24" width="0.1524" layer="91"/>
<junction x="76.2" y="129.54"/>
<junction x="76.2" y="132.08"/>
<junction x="76.2" y="134.62"/>
<junction x="76.2" y="137.16"/>
<junction x="76.2" y="142.24"/>
<pinref part="IC6" gate="FT2232L" pin="AGND"/>
<pinref part="GND9" gate="G$1" pin="GND"/>
<pinref part="IC6" gate="FT2232L" pin="GND@4"/>
<pinref part="IC6" gate="FT2232L" pin="GND@3"/>
<pinref part="IC6" gate="FT2232L" pin="GND@2"/>
<pinref part="IC6" gate="FT2232L" pin="GND@1"/>
<pinref part="IC6" gate="FT2232L" pin="TEST"/>
</segment>
<segment>
<wire x1="45.72" y1="132.08" x2="45.72" y2="137.16" width="0.1524" layer="91"/>
<wire x1="45.72" y1="137.16" x2="45.72" y2="142.24" width="0.1524" layer="91"/>
<wire x1="45.72" y1="137.16" x2="60.96" y2="137.16" width="0.1524" layer="91"/>
<wire x1="60.96" y1="137.16" x2="60.96" y2="142.24" width="0.1524" layer="91"/>
<junction x="45.72" y="137.16"/>
<pinref part="GND10" gate="G$1" pin="GND"/>
<pinref part="C29" gate="G$1" pin="2"/>
<pinref part="C30" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="22.86" y1="195.58" x2="27.94" y2="195.58" width="0.1524" layer="91"/>
<wire x1="27.94" y1="195.58" x2="27.94" y2="193.04" width="0.1524" layer="91"/>
<wire x1="27.94" y1="193.04" x2="27.94" y2="187.96" width="0.1524" layer="91"/>
<wire x1="27.94" y1="187.96" x2="27.94" y2="185.42" width="0.1524" layer="91"/>
<wire x1="27.94" y1="185.42" x2="27.94" y2="182.88" width="0.1524" layer="91"/>
<wire x1="27.94" y1="182.88" x2="27.94" y2="180.34" width="0.1524" layer="91"/>
<wire x1="27.94" y1="180.34" x2="27.94" y2="170.18" width="0.1524" layer="91"/>
<wire x1="27.94" y1="170.18" x2="27.94" y2="132.08" width="0.1524" layer="91"/>
<wire x1="22.86" y1="193.04" x2="27.94" y2="193.04" width="0.1524" layer="91"/>
<wire x1="20.32" y1="187.96" x2="27.94" y2="187.96" width="0.1524" layer="91"/>
<wire x1="20.32" y1="185.42" x2="27.94" y2="185.42" width="0.1524" layer="91"/>
<wire x1="20.32" y1="182.88" x2="27.94" y2="182.88" width="0.1524" layer="91"/>
<wire x1="20.32" y1="180.34" x2="27.94" y2="180.34" width="0.1524" layer="91"/>
<wire x1="45.72" y1="180.34" x2="45.72" y2="170.18" width="0.1524" layer="91"/>
<wire x1="45.72" y1="170.18" x2="33.02" y2="170.18" width="0.1524" layer="91"/>
<wire x1="33.02" y1="170.18" x2="27.94" y2="170.18" width="0.1524" layer="91"/>
<wire x1="33.02" y1="180.34" x2="33.02" y2="170.18" width="0.1524" layer="91"/>
<junction x="27.94" y="193.04"/>
<junction x="27.94" y="187.96"/>
<junction x="27.94" y="185.42"/>
<junction x="27.94" y="182.88"/>
<junction x="27.94" y="180.34"/>
<junction x="27.94" y="170.18"/>
<junction x="33.02" y="170.18"/>
<pinref part="USB" gate="CON_USB" pin="4"/>
<pinref part="GND11" gate="G$1" pin="GND"/>
<pinref part="USB" gate="CON_USB" pin="5"/>
<pinref part="USB" gate="CON_USB" pin="PAD1"/>
<pinref part="USB" gate="CON_USB" pin="PAD2"/>
<pinref part="USB" gate="CON_USB" pin="PAD3"/>
<pinref part="USB" gate="CON_USB" pin="PAD4"/>
<pinref part="C28" gate="G$1" pin="2"/>
<pinref part="C27" gate="G$1" pin="2"/>
<pinref part="C22" gate="G$1" pin="2"/>
<wire x1="45.72" y1="170.18" x2="78.74" y2="170.18" width="0.1524" layer="91"/>
<wire x1="78.74" y1="170.18" x2="78.74" y2="190.5" width="0.1524" layer="91"/>
<junction x="45.72" y="170.18"/>
</segment>
<segment>
<wire x1="106.68" y1="96.52" x2="111.76" y2="96.52" width="0.1524" layer="91"/>
<wire x1="111.76" y1="96.52" x2="111.76" y2="91.44" width="0.1524" layer="91"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="GND"/>
<pinref part="GND12" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="345.44" y1="73.66" x2="345.44" y2="40.64" width="0.1524" layer="91"/>
<pinref part="VPE" gate="G$1" pin="9"/>
<pinref part="GND13" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="287.02" y1="38.1" x2="287.02" y2="73.66" width="0.1524" layer="91"/>
<pinref part="GND14" gate="G$1" pin="GND"/>
<pinref part="VPD" gate="G$1" pin="9"/>
</segment>
<segment>
<wire x1="116.84" y1="434.34" x2="111.76" y2="434.34" width="0.1524" layer="91"/>
<wire x1="111.76" y1="434.34" x2="111.76" y2="429.26" width="0.1524" layer="91"/>
<pinref part="IC2" gate="XCF02S" pin="GND1"/>
<pinref part="GND1" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="76.2" y1="292.1" x2="83.82" y2="292.1" width="0.1524" layer="91"/>
<wire x1="83.82" y1="292.1" x2="83.82" y2="287.02" width="0.1524" layer="91"/>
<pinref part="U1" gate="UNIT_KC7050B" pin="2"/>
<pinref part="GND16" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="119.38" y1="287.02" x2="119.38" y2="281.94" width="0.1524" layer="91"/>
<wire x1="182.88" y1="294.64" x2="182.88" y2="287.02" width="0.1524" layer="91"/>
<wire x1="182.88" y1="287.02" x2="170.18" y2="287.02" width="0.1524" layer="91"/>
<wire x1="170.18" y1="287.02" x2="165.1" y2="287.02" width="0.1524" layer="91"/>
<wire x1="165.1" y1="292.1" x2="165.1" y2="287.02" width="0.1524" layer="91"/>
<wire x1="170.18" y1="292.1" x2="170.18" y2="287.02" width="0.1524" layer="91"/>
<wire x1="165.1" y1="287.02" x2="124.46" y2="287.02" width="0.1524" layer="91"/>
<wire x1="124.46" y1="287.02" x2="119.38" y2="287.02" width="0.1524" layer="91"/>
<wire x1="119.38" y1="294.64" x2="119.38" y2="287.02" width="0.1524" layer="91"/>
<wire x1="129.54" y1="302.26" x2="124.46" y2="302.26" width="0.1524" layer="91"/>
<wire x1="124.46" y1="302.26" x2="124.46" y2="287.02" width="0.1524" layer="91"/>
<junction x="165.1" y="287.02"/>
<junction x="170.18" y="287.02"/>
<junction x="119.38" y="287.02"/>
<junction x="124.46" y="287.02"/>
<pinref part="GND17" gate="G$1" pin="GND"/>
<pinref part="C21" gate="G$1" pin="2"/>
<pinref part="RESET" gate="LS6J2M-T" pin="4"/>
<pinref part="RESET" gate="LS6J2M-T" pin="3"/>
<pinref part="R12" gate="G$1" pin="2"/>
<pinref part="IC4" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="269.24" y1="480.06" x2="261.62" y2="480.06" width="0.1524" layer="91"/>
<wire x1="261.62" y1="480.06" x2="251.46" y2="480.06" width="0.1524" layer="91"/>
<wire x1="251.46" y1="480.06" x2="251.46" y2="477.52" width="0.1524" layer="91"/>
<wire x1="269.24" y1="474.98" x2="261.62" y2="474.98" width="0.1524" layer="91"/>
<wire x1="261.62" y1="474.98" x2="261.62" y2="477.52" width="0.1524" layer="91"/>
<wire x1="261.62" y1="477.52" x2="261.62" y2="480.06" width="0.1524" layer="91"/>
<wire x1="261.62" y1="477.52" x2="269.24" y2="477.52" width="0.1524" layer="91"/>
<junction x="261.62" y="480.06"/>
<junction x="261.62" y="477.52"/>
<pinref part="IC1" gate="XC3S250E" pin="M0"/>
<pinref part="GND18" gate="G$1" pin="GND"/>
<pinref part="IC1" gate="XC3S250E" pin="M2"/>
<pinref part="IC1" gate="XC3S250E" pin="M1"/>
</segment>
<segment>
<wire x1="20.32" y1="332.74" x2="20.32" y2="330.2" width="0.1524" layer="91"/>
<wire x1="20.32" y1="330.2" x2="20.32" y2="327.66" width="0.1524" layer="91"/>
<wire x1="20.32" y1="330.2" x2="27.94" y2="330.2" width="0.1524" layer="91"/>
<wire x1="27.94" y1="330.2" x2="27.94" y2="332.74" width="0.1524" layer="91"/>
<junction x="20.32" y="330.2"/>
<pinref part="C17" gate="G$1" pin="2"/>
<pinref part="GND19" gate="G$1" pin="GND"/>
<pinref part="C18" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="25.4" y1="55.88" x2="22.86" y2="55.88" width="0.1524" layer="91"/>
<wire x1="22.86" y1="55.88" x2="22.86" y2="48.26" width="0.1524" layer="91"/>
<wire x1="22.86" y1="48.26" x2="22.86" y2="40.64" width="0.1524" layer="91"/>
<wire x1="22.86" y1="40.64" x2="22.86" y2="33.02" width="0.1524" layer="91"/>
<wire x1="22.86" y1="33.02" x2="22.86" y2="22.86" width="0.1524" layer="91"/>
<wire x1="25.4" y1="33.02" x2="22.86" y2="33.02" width="0.1524" layer="91"/>
<wire x1="25.4" y1="40.64" x2="22.86" y2="40.64" width="0.1524" layer="91"/>
<wire x1="25.4" y1="48.26" x2="22.86" y2="48.26" width="0.1524" layer="91"/>
<junction x="22.86" y="33.02"/>
<junction x="22.86" y="40.64"/>
<junction x="22.86" y="48.26"/>
<pinref part="H1" gate="G$1" pin="MOUNT"/>
<pinref part="GND20" gate="G$1" pin="GND"/>
<pinref part="H4" gate="G$1" pin="MOUNT"/>
<pinref part="H3" gate="G$1" pin="MOUNT"/>
<pinref part="H2" gate="G$1" pin="MOUNT"/>
</segment>
<segment>
<wire x1="58.42" y1="332.74" x2="58.42" y2="327.66" width="0.1524" layer="91"/>
<pinref part="C20" gate="G$1" pin="2"/>
<pinref part="GND21" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="38.1" y1="327.66" x2="38.1" y2="332.74" width="0.1524" layer="91"/>
<pinref part="GND22" gate="G$1" pin="GND"/>
<pinref part="C19" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="205.74" y1="182.88" x2="205.74" y2="177.8" width="0.1524" layer="91"/>
<wire x1="200.66" y1="182.88" x2="200.66" y2="177.8" width="0.1524" layer="91"/>
<wire x1="200.66" y1="177.8" x2="205.74" y2="177.8" width="0.1524" layer="91"/>
<wire x1="205.74" y1="177.8" x2="210.82" y2="177.8" width="0.1524" layer="91"/>
<wire x1="210.82" y1="172.72" x2="210.82" y2="177.8" width="0.1524" layer="91"/>
<wire x1="210.82" y1="177.8" x2="215.9" y2="177.8" width="0.1524" layer="91"/>
<wire x1="215.9" y1="177.8" x2="215.9" y2="185.42" width="0.1524" layer="91"/>
<wire x1="215.9" y1="177.8" x2="226.06" y2="177.8" width="0.1524" layer="91"/>
<wire x1="226.06" y1="182.88" x2="226.06" y2="177.8" width="0.1524" layer="91"/>
<wire x1="226.06" y1="177.8" x2="231.14" y2="177.8" width="0.1524" layer="91"/>
<wire x1="231.14" y1="177.8" x2="231.14" y2="182.88" width="0.1524" layer="91"/>
<wire x1="231.14" y1="177.8" x2="241.3" y2="177.8" width="0.1524" layer="91"/>
<wire x1="241.3" y1="177.8" x2="241.3" y2="185.42" width="0.1524" layer="91"/>
<wire x1="241.3" y1="177.8" x2="251.46" y2="177.8" width="0.1524" layer="91"/>
<wire x1="251.46" y1="182.88" x2="251.46" y2="177.8" width="0.1524" layer="91"/>
<wire x1="251.46" y1="177.8" x2="256.54" y2="177.8" width="0.1524" layer="91"/>
<wire x1="256.54" y1="177.8" x2="256.54" y2="182.88" width="0.1524" layer="91"/>
<wire x1="256.54" y1="177.8" x2="266.7" y2="177.8" width="0.1524" layer="91"/>
<wire x1="266.7" y1="177.8" x2="266.7" y2="185.42" width="0.1524" layer="91"/>
<wire x1="266.7" y1="177.8" x2="276.86" y2="177.8" width="0.1524" layer="91"/>
<wire x1="276.86" y1="182.88" x2="276.86" y2="177.8" width="0.1524" layer="91"/>
<wire x1="276.86" y1="177.8" x2="281.94" y2="177.8" width="0.1524" layer="91"/>
<wire x1="281.94" y1="177.8" x2="281.94" y2="182.88" width="0.1524" layer="91"/>
<wire x1="281.94" y1="177.8" x2="292.1" y2="177.8" width="0.1524" layer="91"/>
<wire x1="292.1" y1="177.8" x2="292.1" y2="185.42" width="0.1524" layer="91"/>
<wire x1="307.34" y1="193.04" x2="307.34" y2="177.8" width="0.1524" layer="91"/>
<wire x1="307.34" y1="177.8" x2="292.1" y2="177.8" width="0.1524" layer="91"/>
<wire x1="317.5" y1="193.04" x2="317.5" y2="177.8" width="0.1524" layer="91"/>
<wire x1="317.5" y1="177.8" x2="327.66" y2="177.8" width="0.1524" layer="91"/>
<wire x1="317.5" y1="177.8" x2="307.34" y2="177.8" width="0.1524" layer="91"/>
<wire x1="327.66" y1="200.66" x2="332.74" y2="200.66" width="0.1524" layer="91"/>
<wire x1="327.66" y1="200.66" x2="327.66" y2="177.8" width="0.1524" layer="91"/>
<junction x="205.74" y="177.8"/>
<junction x="210.82" y="177.8"/>
<junction x="215.9" y="177.8"/>
<junction x="226.06" y="177.8"/>
<junction x="231.14" y="177.8"/>
<junction x="241.3" y="177.8"/>
<junction x="251.46" y="177.8"/>
<junction x="256.54" y="177.8"/>
<junction x="266.7" y="177.8"/>
<junction x="276.86" y="177.8"/>
<junction x="281.94" y="177.8"/>
<junction x="292.1" y="177.8"/>
<junction x="317.5" y="177.8"/>
<junction x="307.34" y="177.8"/>
<pinref part="SW4" gate="LS6J2M-T" pin="3"/>
<pinref part="SW3" gate="LS6J2M-T" pin="3"/>
<pinref part="SW2" gate="LS6J2M-T" pin="3"/>
<pinref part="SW1" gate="LS6J2M-T" pin="3"/>
<pinref part="SW1" gate="LS6J2M-T" pin="4"/>
<pinref part="GND3" gate="G$1" pin="GND"/>
<pinref part="SW2" gate="LS6J2M-T" pin="4"/>
<pinref part="SW3" gate="LS6J2M-T" pin="4"/>
<pinref part="SW4" gate="LS6J2M-T" pin="4"/>
<pinref part="C24" gate="G$1" pin="2"/>
<pinref part="C25" gate="G$1" pin="2"/>
<pinref part="C26" gate="G$1" pin="2"/>
<pinref part="C23" gate="G$1" pin="2"/>
<pinref part="R20" gate="G$1" pin="2"/>
<pinref part="R21" gate="G$1" pin="2"/>
<pinref part="IC5" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="162.56" y1="464.82" x2="172.72" y2="464.82" width="0.1524" layer="91"/>
<wire x1="172.72" y1="464.82" x2="177.8" y2="464.82" width="0.1524" layer="91"/>
<wire x1="172.72" y1="447.04" x2="172.72" y2="464.82" width="0.1524" layer="91"/>
<junction x="172.72" y="464.82"/>
<pinref part="IC2" gate="XCF02S" pin="CE_"/>
<pinref part="DONE" gate="G$1" pin="C"/>
<pinref part="GND23" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="119.38" y1="353.06" x2="119.38" y2="358.14" width="0.1524" layer="91"/>
<wire x1="142.24" y1="358.14" x2="119.38" y2="358.14" width="0.1524" layer="91"/>
<pinref part="GND15" gate="G$1" pin="GND"/>
<pinref part="IC3" gate="G$1" pin="GND"/>
</segment>
<segment>
<wire x1="88.9" y1="370.84" x2="88.9" y2="353.06" width="0.1524" layer="91"/>
<pinref part="R6" gate="G$1" pin="2"/>
<pinref part="GND2" gate="G$1" pin="GND"/>
</segment>
</net>
<net name="JTAG_CCLK" class="0">
<segment>
<wire x1="162.56" y1="467.36" x2="269.24" y2="467.36" width="0.1524" layer="91"/>
<pinref part="IC2" gate="XCF02S" pin="CLK"/>
<pinref part="IC1" gate="XC3S250E" pin="CCLK"/>
</segment>
</net>
<net name="JTAG_INIT_B" class="0">
<segment>
<wire x1="162.56" y1="459.74" x2="236.22" y2="459.74" width="0.1524" layer="91"/>
<wire x1="236.22" y1="459.74" x2="269.24" y2="459.74" width="0.1524" layer="91"/>
<wire x1="236.22" y1="459.74" x2="236.22" y2="477.52" width="0.1524" layer="91"/>
<junction x="236.22" y="459.74"/>
<pinref part="IC2" gate="XCF02S" pin="OE/NRST"/>
<pinref part="IC1" gate="XC3S250E" pin="INIT_B"/>
<pinref part="R3" gate="G$1" pin="2"/>
</segment>
</net>
<net name="JTAG_PROG_B" class="0">
<segment>
<wire x1="269.24" y1="457.2" x2="231.14" y2="457.2" width="0.1524" layer="91"/>
<wire x1="231.14" y1="457.2" x2="162.56" y2="457.2" width="0.1524" layer="91"/>
<wire x1="231.14" y1="457.2" x2="231.14" y2="477.52" width="0.1524" layer="91"/>
<junction x="231.14" y="457.2"/>
<pinref part="IC1" gate="XC3S250E" pin="PROG_B"/>
<pinref part="IC2" gate="XCF02S" pin="CF_"/>
<pinref part="R2" gate="G$1" pin="2"/>
</segment>
</net>
<net name="JTAG_DIN" class="0">
<segment>
<wire x1="162.56" y1="454.66" x2="269.24" y2="454.66" width="0.1524" layer="91"/>
<pinref part="IC2" gate="XCF02S" pin="D0"/>
<pinref part="IC1" gate="XC3S250E" pin="DIN"/>
</segment>
</net>
<net name="JTAG_DONE" class="0">
<segment>
<wire x1="185.42" y1="464.82" x2="226.06" y2="464.82" width="0.1524" layer="91"/>
<wire x1="226.06" y1="464.82" x2="269.24" y2="464.82" width="0.1524" layer="91"/>
<wire x1="226.06" y1="464.82" x2="226.06" y2="477.52" width="0.1524" layer="91"/>
<junction x="226.06" y="464.82"/>
<pinref part="DONE" gate="G$1" pin="A"/>
<pinref part="IC1" gate="XC3S250E" pin="DONE"/>
<pinref part="R1" gate="G$1" pin="2"/>
</segment>
</net>
<net name="JTAG_IOEN" class="0">
<segment>
<wire x1="142.24" y1="200.66" x2="157.48" y2="200.66" width="0.1524" layer="91"/>
<label x="144.78" y="200.66" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="ADBUS4"/>
</segment>
</net>
<net name="EECS" class="0">
<segment>
<wire x1="81.28" y1="154.94" x2="71.12" y2="154.94" width="0.1524" layer="91"/>
<wire x1="71.12" y1="154.94" x2="71.12" y2="104.14" width="0.1524" layer="91"/>
<wire x1="81.28" y1="104.14" x2="71.12" y2="104.14" width="0.1524" layer="91"/>
<pinref part="IC6" gate="FT2232L" pin="EECS"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="CS"/>
</segment>
</net>
<net name="EESK" class="0">
<segment>
<wire x1="68.58" y1="152.4" x2="81.28" y2="152.4" width="0.1524" layer="91"/>
<wire x1="68.58" y1="152.4" x2="68.58" y2="101.6" width="0.1524" layer="91"/>
<wire x1="68.58" y1="101.6" x2="81.28" y2="101.6" width="0.1524" layer="91"/>
<pinref part="IC6" gate="FT2232L" pin="EESK"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="SK"/>
</segment>
</net>
<net name="EEDATA" class="0">
<segment>
<wire x1="81.28" y1="149.86" x2="66.04" y2="149.86" width="0.1524" layer="91"/>
<wire x1="66.04" y1="149.86" x2="66.04" y2="111.76" width="0.1524" layer="91"/>
<wire x1="66.04" y1="111.76" x2="66.04" y2="99.06" width="0.1524" layer="91"/>
<wire x1="66.04" y1="99.06" x2="81.28" y2="99.06" width="0.1524" layer="91"/>
<wire x1="60.96" y1="106.68" x2="60.96" y2="111.76" width="0.1524" layer="91"/>
<wire x1="60.96" y1="111.76" x2="66.04" y2="111.76" width="0.1524" layer="91"/>
<junction x="66.04" y="111.76"/>
<pinref part="IC6" gate="FT2232L" pin="EEDATA"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="DI"/>
<pinref part="R43" gate="G$1" pin="1"/>
</segment>
</net>
<net name="USB_POW" class="0">
<segment>
<wire x1="68.58" y1="210.82" x2="71.12" y2="210.82" width="0.1524" layer="91"/>
<wire x1="68.58" y1="205.74" x2="81.28" y2="205.74" width="0.1524" layer="91"/>
<wire x1="68.58" y1="205.74" x2="68.58" y2="210.82" width="0.1524" layer="91"/>
<wire x1="68.58" y1="203.2" x2="81.28" y2="203.2" width="0.1524" layer="91"/>
<wire x1="68.58" y1="203.2" x2="68.58" y2="205.74" width="0.1524" layer="91"/>
<wire x1="60.96" y1="203.2" x2="68.58" y2="203.2" width="0.1524" layer="91"/>
<junction x="68.58" y="205.74"/>
<junction x="68.58" y="203.2"/>
<pinref part="R13" gate="G$1" pin="1"/>
<pinref part="IC6" gate="FT2232L" pin="VCC@1"/>
<pinref part="IC6" gate="FT2232L" pin="VCC@2"/>
<pinref part="FB1" gate="G$1" pin="P$2"/>
</segment>
</net>
<net name="RX" class="0">
<segment>
<wire x1="142.24" y1="170.18" x2="157.48" y2="170.18" width="0.1524" layer="91"/>
<label x="144.78" y="170.18" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="BDBUS0"/>
</segment>
<segment>
<wire x1="330.2" y1="360.68" x2="353.06" y2="360.68" width="0.1524" layer="91"/>
<label x="337.82" y="360.68" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L07P_1"/>
</segment>
</net>
<net name="TX" class="0">
<segment>
<wire x1="142.24" y1="167.64" x2="157.48" y2="167.64" width="0.1524" layer="91"/>
<label x="144.78" y="167.64" size="1.778" layer="95"/>
<pinref part="IC6" gate="FT2232L" pin="BDBUS1"/>
</segment>
<segment>
<wire x1="330.2" y1="358.14" x2="353.06" y2="358.14" width="0.1524" layer="91"/>
<label x="337.82" y="358.14" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L07N_1"/>
</segment>
</net>
<net name="VPORT_1_2" class="0">
<segment>
<wire x1="269.24" y1="66.04" x2="269.24" y2="73.66" width="0.1524" layer="91"/>
<wire x1="269.24" y1="66.04" x2="246.38" y2="66.04" width="0.1524" layer="91"/>
<label x="246.38" y="66.04" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="330.2" y1="381" x2="353.06" y2="381" width="0.1524" layer="91"/>
<label x="337.82" y="381" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03N_1/RHCLK1"/>
</segment>
</net>
<net name="VPORT_1_1" class="0">
<segment>
<wire x1="266.7" y1="73.66" x2="266.7" y2="68.58" width="0.1524" layer="91"/>
<wire x1="266.7" y1="68.58" x2="246.38" y2="68.58" width="0.1524" layer="91"/>
<label x="246.38" y="68.58" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="383.54" x2="330.2" y2="383.54" width="0.1524" layer="91"/>
<label x="337.82" y="383.54" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03P_1/RHCLK0"/>
</segment>
</net>
<net name="VPORT_1_3" class="0">
<segment>
<wire x1="271.78" y1="63.5" x2="271.78" y2="73.66" width="0.1524" layer="91"/>
<wire x1="271.78" y1="63.5" x2="246.38" y2="63.5" width="0.1524" layer="91"/>
<label x="246.38" y="63.5" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="3"/>
</segment>
<segment>
<wire x1="330.2" y1="378.46" x2="353.06" y2="378.46" width="0.1524" layer="91"/>
<label x="337.82" y="378.46" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04P_1/RHCLK2"/>
</segment>
</net>
<net name="VPORT_1_4" class="0">
<segment>
<wire x1="274.32" y1="60.96" x2="274.32" y2="73.66" width="0.1524" layer="91"/>
<wire x1="274.32" y1="60.96" x2="246.38" y2="60.96" width="0.1524" layer="91"/>
<label x="246.38" y="60.96" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="4"/>
</segment>
<segment>
<wire x1="330.2" y1="375.92" x2="353.06" y2="375.92" width="0.1524" layer="91"/>
<label x="337.82" y="375.92" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04N_1/RHCLK3"/>
</segment>
</net>
<net name="VPORT_1_5" class="0">
<segment>
<wire x1="276.86" y1="58.42" x2="276.86" y2="73.66" width="0.1524" layer="91"/>
<wire x1="276.86" y1="58.42" x2="246.38" y2="58.42" width="0.1524" layer="91"/>
<label x="246.38" y="58.42" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="5"/>
</segment>
<segment>
<wire x1="330.2" y1="373.38" x2="353.06" y2="373.38" width="0.1524" layer="91"/>
<label x="337.82" y="373.38" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05P_1/RHCLK4"/>
</segment>
</net>
<net name="VPORT_1_6" class="0">
<segment>
<wire x1="279.4" y1="55.88" x2="279.4" y2="73.66" width="0.1524" layer="91"/>
<wire x1="279.4" y1="55.88" x2="246.38" y2="55.88" width="0.1524" layer="91"/>
<label x="246.38" y="55.88" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="6"/>
</segment>
<segment>
<wire x1="330.2" y1="370.84" x2="353.06" y2="370.84" width="0.1524" layer="91"/>
<label x="337.82" y="370.84" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05N_1/RHCLK5"/>
</segment>
</net>
<net name="VPORT_1_7" class="0">
<segment>
<wire x1="281.94" y1="53.34" x2="281.94" y2="73.66" width="0.1524" layer="91"/>
<wire x1="281.94" y1="53.34" x2="246.38" y2="53.34" width="0.1524" layer="91"/>
<label x="246.38" y="53.34" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="7"/>
</segment>
<segment>
<wire x1="330.2" y1="368.3" x2="353.06" y2="368.3" width="0.1524" layer="91"/>
<label x="337.82" y="368.3" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06P_1/RHCLK6"/>
</segment>
</net>
<net name="VPORT_1_8" class="0">
<segment>
<wire x1="284.48" y1="50.8" x2="284.48" y2="73.66" width="0.1524" layer="91"/>
<wire x1="284.48" y1="50.8" x2="246.38" y2="50.8" width="0.1524" layer="91"/>
<label x="246.38" y="50.8" size="1.778" layer="95"/>
<pinref part="VPD" gate="G$1" pin="8"/>
</segment>
<segment>
<wire x1="330.2" y1="365.76" x2="353.06" y2="365.76" width="0.1524" layer="91"/>
<label x="337.82" y="365.76" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06N_1/RHCLK7"/>
</segment>
</net>
<net name="CLOCK" class="0">
<segment>
<wire x1="330.2" y1="345.44" x2="353.06" y2="345.44" width="0.1524" layer="91"/>
<label x="337.82" y="345.44" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L02P_0/GCLK4"/>
</segment>
<segment>
<wire x1="86.36" y1="297.18" x2="99.06" y2="297.18" width="0.1524" layer="91"/>
<label x="88.9" y="297.18" size="1.778" layer="95"/>
<pinref part="R11" gate="G$1" pin="2"/>
</segment>
</net>
<net name="RESET" class="0">
<segment>
<wire x1="160.02" y1="304.8" x2="170.18" y2="304.8" width="0.1524" layer="91"/>
<wire x1="170.18" y1="304.8" x2="187.96" y2="304.8" width="0.1524" layer="91"/>
<wire x1="170.18" y1="309.88" x2="170.18" y2="304.8" width="0.1524" layer="91"/>
<junction x="170.18" y="304.8"/>
<label x="177.8" y="304.8" size="1.778" layer="95"/>
<pinref part="IC4" gate="G$1" pin="OUT"/>
<pinref part="R10" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="330.2" y1="340.36" x2="353.06" y2="340.36" width="0.1524" layer="91"/>
<label x="337.82" y="340.36" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03P_0/GCLK6"/>
</segment>
</net>
<net name="RESET_IN" class="0">
<segment>
<wire x1="129.54" y1="307.34" x2="119.38" y2="307.34" width="0.1524" layer="91"/>
<wire x1="119.38" y1="309.88" x2="119.38" y2="307.34" width="0.1524" layer="91"/>
<wire x1="119.38" y1="307.34" x2="119.38" y2="299.72" width="0.1524" layer="91"/>
<junction x="119.38" y="307.34"/>
<pinref part="IC4" gate="G$1" pin="IN"/>
<pinref part="R9" gate="G$1" pin="2"/>
<pinref part="R12" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_1_IN" class="0">
<segment>
<wire x1="358.14" y1="231.14" x2="378.46" y2="231.14" width="0.1524" layer="91"/>
<label x="363.22" y="231.14" size="1.778" layer="95"/>
<pinref part="IC5" gate="G$1" pin="OUT1"/>
</segment>
<segment>
<wire x1="330.2" y1="447.04" x2="353.06" y2="447.04" width="0.1524" layer="91"/>
<label x="337.82" y="447.04" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L07P_3"/>
</segment>
</net>
<net name="SW_2_IN" class="0">
<segment>
<wire x1="358.14" y1="226.06" x2="378.46" y2="226.06" width="0.1524" layer="91"/>
<label x="363.22" y="226.06" size="1.778" layer="95"/>
<pinref part="IC5" gate="G$1" pin="OUT2"/>
</segment>
<segment>
<wire x1="330.2" y1="444.5" x2="353.06" y2="444.5" width="0.1524" layer="91"/>
<label x="337.82" y="444.5" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L07N_3"/>
</segment>
</net>
<net name="SW_3_IN" class="0">
<segment>
<wire x1="378.46" y1="220.98" x2="358.14" y2="220.98" width="0.1524" layer="91"/>
<label x="363.22" y="220.98" size="1.778" layer="95"/>
<pinref part="IC5" gate="G$1" pin="OUT3"/>
</segment>
<segment>
<wire x1="330.2" y1="436.88" x2="353.06" y2="436.88" width="0.1524" layer="91"/>
<label x="337.82" y="436.88" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="CSO_B/IO_BK2"/>
</segment>
</net>
<net name="SW_4_IN" class="0">
<segment>
<wire x1="378.46" y1="215.9" x2="358.14" y2="215.9" width="0.1524" layer="91"/>
<label x="363.22" y="215.9" size="1.778" layer="95"/>
<pinref part="IC5" gate="G$1" pin="OUT4"/>
</segment>
<segment>
<wire x1="330.2" y1="434.34" x2="353.06" y2="434.34" width="0.1524" layer="91"/>
<label x="337.82" y="434.34" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="DOUT/IO_BK2"/>
</segment>
</net>
<net name="SW_1_IN_T2" class="0">
<segment>
<wire x1="205.74" y1="198.12" x2="205.74" y2="190.5" width="0.1524" layer="91"/>
<wire x1="205.74" y1="190.5" x2="205.74" y2="187.96" width="0.1524" layer="91"/>
<wire x1="200.66" y1="187.96" x2="200.66" y2="190.5" width="0.1524" layer="91"/>
<wire x1="200.66" y1="190.5" x2="205.74" y2="190.5" width="0.1524" layer="91"/>
<wire x1="208.28" y1="190.5" x2="205.74" y2="190.5" width="0.1524" layer="91"/>
<junction x="205.74" y="190.5"/>
<pinref part="R14" gate="G$1" pin="2"/>
<pinref part="SW1" gate="LS6J2M-T" pin="1"/>
<pinref part="SW1" gate="LS6J2M-T" pin="2"/>
<pinref part="R22" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_1_IN_T" class="0">
<segment>
<wire x1="215.9" y1="231.14" x2="215.9" y2="190.5" width="0.1524" layer="91"/>
<wire x1="215.9" y1="190.5" x2="215.9" y2="187.96" width="0.1524" layer="91"/>
<wire x1="215.9" y1="190.5" x2="213.36" y2="190.5" width="0.1524" layer="91"/>
<wire x1="215.9" y1="231.14" x2="332.74" y2="231.14" width="0.1524" layer="91"/>
<junction x="215.9" y="190.5"/>
<pinref part="C23" gate="G$1" pin="1"/>
<pinref part="R22" gate="G$1" pin="2"/>
<pinref part="IC5" gate="G$1" pin="IN1"/>
</segment>
</net>
<net name="SW_2_IN_T2" class="0">
<segment>
<wire x1="231.14" y1="198.12" x2="231.14" y2="190.5" width="0.1524" layer="91"/>
<wire x1="231.14" y1="190.5" x2="231.14" y2="187.96" width="0.1524" layer="91"/>
<wire x1="231.14" y1="190.5" x2="233.68" y2="190.5" width="0.1524" layer="91"/>
<wire x1="226.06" y1="187.96" x2="226.06" y2="190.5" width="0.1524" layer="91"/>
<wire x1="226.06" y1="190.5" x2="231.14" y2="190.5" width="0.1524" layer="91"/>
<junction x="231.14" y="190.5"/>
<pinref part="R15" gate="G$1" pin="2"/>
<pinref part="SW2" gate="LS6J2M-T" pin="1"/>
<pinref part="SW2" gate="LS6J2M-T" pin="2"/>
<pinref part="R23" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_2_IN_T" class="0">
<segment>
<wire x1="238.76" y1="190.5" x2="241.3" y2="190.5" width="0.1524" layer="91"/>
<wire x1="241.3" y1="190.5" x2="241.3" y2="187.96" width="0.1524" layer="91"/>
<wire x1="241.3" y1="190.5" x2="241.3" y2="226.06" width="0.1524" layer="91"/>
<wire x1="241.3" y1="226.06" x2="332.74" y2="226.06" width="0.1524" layer="91"/>
<junction x="241.3" y="190.5"/>
<pinref part="C24" gate="G$1" pin="1"/>
<pinref part="R23" gate="G$1" pin="2"/>
<pinref part="IC5" gate="G$1" pin="IN2"/>
</segment>
</net>
<net name="SW_3_IN_T2" class="0">
<segment>
<wire x1="256.54" y1="198.12" x2="256.54" y2="190.5" width="0.1524" layer="91"/>
<wire x1="256.54" y1="190.5" x2="256.54" y2="187.96" width="0.1524" layer="91"/>
<wire x1="259.08" y1="190.5" x2="256.54" y2="190.5" width="0.1524" layer="91"/>
<wire x1="251.46" y1="187.96" x2="251.46" y2="190.5" width="0.1524" layer="91"/>
<wire x1="251.46" y1="190.5" x2="256.54" y2="190.5" width="0.1524" layer="91"/>
<junction x="256.54" y="190.5"/>
<pinref part="SW3" gate="LS6J2M-T" pin="1"/>
<pinref part="R16" gate="G$1" pin="2"/>
<pinref part="SW3" gate="LS6J2M-T" pin="2"/>
<pinref part="R24" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_3_IN_T" class="0">
<segment>
<wire x1="264.16" y1="190.5" x2="266.7" y2="190.5" width="0.1524" layer="91"/>
<wire x1="266.7" y1="190.5" x2="266.7" y2="187.96" width="0.1524" layer="91"/>
<wire x1="266.7" y1="190.5" x2="266.7" y2="220.98" width="0.1524" layer="91"/>
<wire x1="266.7" y1="220.98" x2="332.74" y2="220.98" width="0.1524" layer="91"/>
<junction x="266.7" y="190.5"/>
<pinref part="C25" gate="G$1" pin="1"/>
<pinref part="R24" gate="G$1" pin="2"/>
<pinref part="IC5" gate="G$1" pin="IN3"/>
</segment>
</net>
<net name="SW_4_IN_T2" class="0">
<segment>
<wire x1="281.94" y1="198.12" x2="281.94" y2="190.5" width="0.1524" layer="91"/>
<wire x1="281.94" y1="190.5" x2="281.94" y2="187.96" width="0.1524" layer="91"/>
<wire x1="281.94" y1="190.5" x2="284.48" y2="190.5" width="0.1524" layer="91"/>
<wire x1="276.86" y1="187.96" x2="276.86" y2="190.5" width="0.1524" layer="91"/>
<wire x1="276.86" y1="190.5" x2="281.94" y2="190.5" width="0.1524" layer="91"/>
<junction x="281.94" y="190.5"/>
<pinref part="R17" gate="G$1" pin="2"/>
<pinref part="SW4" gate="LS6J2M-T" pin="1"/>
<pinref part="SW4" gate="LS6J2M-T" pin="2"/>
<pinref part="R25" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_4_IN_T" class="0">
<segment>
<wire x1="289.56" y1="190.5" x2="292.1" y2="190.5" width="0.1524" layer="91"/>
<wire x1="292.1" y1="190.5" x2="292.1" y2="187.96" width="0.1524" layer="91"/>
<wire x1="292.1" y1="190.5" x2="292.1" y2="215.9" width="0.1524" layer="91"/>
<wire x1="292.1" y1="215.9" x2="332.74" y2="215.9" width="0.1524" layer="91"/>
<junction x="292.1" y="190.5"/>
<pinref part="C26" gate="G$1" pin="1"/>
<pinref part="R25" gate="G$1" pin="2"/>
<pinref part="IC5" gate="G$1" pin="IN4"/>
</segment>
</net>
<net name="EETMP" class="0">
<segment>
<wire x1="60.96" y1="101.6" x2="60.96" y2="96.52" width="0.1524" layer="91"/>
<wire x1="60.96" y1="96.52" x2="81.28" y2="96.52" width="0.1524" layer="91"/>
<wire x1="55.88" y1="96.52" x2="60.96" y2="96.52" width="0.1524" layer="91"/>
<junction x="60.96" y="96.52"/>
<pinref part="R43" gate="G$1" pin="2"/>
<pinref part="IC7" gate="UNIT_93CX6" pin="DO"/>
<pinref part="R46" gate="G$1" pin="1"/>
</segment>
</net>
<net name="XTOUT" class="0">
<segment>
<wire x1="81.28" y1="160.02" x2="60.96" y2="160.02" width="0.1524" layer="91"/>
<wire x1="60.96" y1="160.02" x2="58.42" y2="160.02" width="0.1524" layer="91"/>
<wire x1="60.96" y1="144.78" x2="60.96" y2="160.02" width="0.1524" layer="91"/>
<junction x="60.96" y="160.02"/>
<pinref part="IC6" gate="FT2232L" pin="XTOUT"/>
<pinref part="U2" gate="UNIT_CRYSTAL" pin="P$2"/>
<pinref part="C30" gate="G$1" pin="1"/>
</segment>
</net>
<net name="XTIN" class="0">
<segment>
<wire x1="81.28" y1="165.1" x2="45.72" y2="165.1" width="0.1524" layer="91"/>
<wire x1="45.72" y1="165.1" x2="45.72" y2="160.02" width="0.1524" layer="91"/>
<wire x1="45.72" y1="160.02" x2="45.72" y2="144.78" width="0.1524" layer="91"/>
<wire x1="48.26" y1="160.02" x2="45.72" y2="160.02" width="0.1524" layer="91"/>
<junction x="45.72" y="160.02"/>
<pinref part="IC6" gate="FT2232L" pin="XTIN"/>
<pinref part="C29" gate="G$1" pin="1"/>
<pinref part="U2" gate="UNIT_CRYSTAL" pin="P$1"/>
</segment>
</net>
<net name="USB_DP" class="0">
<segment>
<wire x1="48.26" y1="198.12" x2="58.42" y2="198.12" width="0.1524" layer="91"/>
<wire x1="58.42" y1="198.12" x2="58.42" y2="177.8" width="0.1524" layer="91"/>
<wire x1="58.42" y1="177.8" x2="58.42" y2="172.72" width="0.1524" layer="91"/>
<wire x1="58.42" y1="172.72" x2="66.04" y2="172.72" width="0.1524" layer="91"/>
<wire x1="58.42" y1="177.8" x2="81.28" y2="177.8" width="0.1524" layer="91"/>
<junction x="58.42" y="177.8"/>
<pinref part="R19" gate="G$1" pin="2"/>
<pinref part="R26" gate="G$1" pin="1"/>
<pinref part="IC6" gate="FT2232L" pin="USBDP"/>
</segment>
</net>
<net name="USB_RSTOUT_" class="0">
<segment>
<wire x1="71.12" y1="172.72" x2="81.28" y2="172.72" width="0.1524" layer="91"/>
<pinref part="R26" gate="G$1" pin="2"/>
<pinref part="IC6" gate="FT2232L" pin="RSTOUT_"/>
</segment>
</net>
<net name="USB_DM" class="0">
<segment>
<wire x1="48.26" y1="200.66" x2="60.96" y2="200.66" width="0.1524" layer="91"/>
<wire x1="60.96" y1="200.66" x2="60.96" y2="182.88" width="0.1524" layer="91"/>
<wire x1="60.96" y1="182.88" x2="81.28" y2="182.88" width="0.1524" layer="91"/>
<pinref part="R18" gate="G$1" pin="2"/>
<pinref part="IC6" gate="FT2232L" pin="USBDM"/>
</segment>
</net>
<net name="USB_CON2" class="0">
<segment>
<wire x1="22.86" y1="200.66" x2="43.18" y2="200.66" width="0.1524" layer="91"/>
<pinref part="USB" gate="CON_USB" pin="2"/>
<pinref part="R18" gate="G$1" pin="1"/>
</segment>
</net>
<net name="USB_CON3" class="0">
<segment>
<wire x1="22.86" y1="198.12" x2="43.18" y2="198.12" width="0.1524" layer="91"/>
<pinref part="USB" gate="CON_USB" pin="3"/>
<pinref part="R19" gate="G$1" pin="1"/>
</segment>
</net>
<net name="USB_AVCC" class="0">
<segment>
<pinref part="R13" gate="G$1" pin="2"/>
<pinref part="IC6" gate="FT2232L" pin="AVCC"/>
<wire x1="81.28" y1="210.82" x2="78.74" y2="210.82" width="0.1524" layer="91"/>
<pinref part="C22" gate="G$1" pin="1"/>
<wire x1="78.74" y1="210.82" x2="76.2" y2="210.82" width="0.1524" layer="91"/>
<wire x1="78.74" y1="193.04" x2="78.74" y2="210.82" width="0.1524" layer="91"/>
<junction x="78.74" y="210.82"/>
</segment>
</net>
<net name="USB_33OUT" class="0">
<segment>
<wire x1="45.72" y1="187.96" x2="81.28" y2="187.96" width="0.1524" layer="91"/>
<wire x1="45.72" y1="187.96" x2="45.72" y2="182.88" width="0.1524" layer="91"/>
<pinref part="IC6" gate="FT2232L" pin="3V3OUT"/>
<pinref part="C28" gate="G$1" pin="1"/>
</segment>
</net>
<net name="SW_PD_1" class="0">
<segment>
<wire x1="307.34" y1="210.82" x2="307.34" y2="198.12" width="0.1524" layer="91"/>
<wire x1="332.74" y1="210.82" x2="307.34" y2="210.82" width="0.1524" layer="91"/>
<pinref part="R20" gate="G$1" pin="1"/>
<pinref part="IC5" gate="G$1" pin="IN5"/>
</segment>
</net>
<net name="SW_PD_2" class="0">
<segment>
<wire x1="317.5" y1="198.12" x2="317.5" y2="205.74" width="0.1524" layer="91"/>
<wire x1="317.5" y1="205.74" x2="332.74" y2="205.74" width="0.1524" layer="91"/>
<pinref part="R21" gate="G$1" pin="1"/>
<pinref part="IC5" gate="G$1" pin="IN6"/>
</segment>
</net>
<net name="JTAG_HSWAP" class="0">
<segment>
<wire x1="241.3" y1="477.52" x2="241.3" y2="447.04" width="0.1524" layer="91"/>
<wire x1="241.3" y1="447.04" x2="269.24" y2="447.04" width="0.1524" layer="91"/>
<pinref part="R4" gate="G$1" pin="2"/>
<pinref part="IC1" gate="XC3S250E" pin="HSWAP"/>
</segment>
</net>
<net name="RESET_SW" class="0">
<segment>
<wire x1="160.02" y1="302.26" x2="165.1" y2="302.26" width="0.1524" layer="91"/>
<wire x1="165.1" y1="302.26" x2="170.18" y2="302.26" width="0.1524" layer="91"/>
<wire x1="170.18" y1="302.26" x2="182.88" y2="302.26" width="0.1524" layer="91"/>
<wire x1="182.88" y1="302.26" x2="182.88" y2="297.18" width="0.1524" layer="91"/>
<wire x1="165.1" y1="302.26" x2="165.1" y2="297.18" width="0.1524" layer="91"/>
<wire x1="170.18" y1="302.26" x2="170.18" y2="297.18" width="0.1524" layer="91"/>
<junction x="165.1" y="302.26"/>
<junction x="170.18" y="302.26"/>
<pinref part="IC4" gate="G$1" pin="DLY"/>
<pinref part="C21" gate="G$1" pin="1"/>
<pinref part="RESET" gate="LS6J2M-T" pin="2"/>
<pinref part="RESET" gate="LS6J2M-T" pin="1"/>
</segment>
</net>
<net name="VPORT_2_1" class="0">
<segment>
<wire x1="325.12" y1="73.66" x2="325.12" y2="68.58" width="0.1524" layer="91"/>
<wire x1="325.12" y1="68.58" x2="309.88" y2="68.58" width="0.1524" layer="91"/>
<label x="309.88" y="68.58" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="424.18" x2="330.2" y2="424.18" width="0.1524" layer="91"/>
<label x="337.82" y="424.18" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03N_2/D6/GCLK13"/>
</segment>
</net>
<net name="VPORT_2_2" class="0">
<segment>
<wire x1="309.88" y1="66.04" x2="327.66" y2="66.04" width="0.1524" layer="91"/>
<wire x1="327.66" y1="66.04" x2="327.66" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="66.04" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="330.2" y1="421.64" x2="353.06" y2="421.64" width="0.1524" layer="91"/>
<label x="337.82" y="421.64" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_BK2/D5"/>
</segment>
</net>
<net name="VPORT_2_3" class="0">
<segment>
<wire x1="309.88" y1="63.5" x2="330.2" y2="63.5" width="0.1524" layer="91"/>
<wire x1="330.2" y1="63.5" x2="330.2" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="63.5" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="3"/>
</segment>
<segment>
<wire x1="353.06" y1="419.1" x2="330.2" y2="419.1" width="0.1524" layer="91"/>
<label x="337.82" y="419.1" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04P_2/D4/GCLK14"/>
</segment>
</net>
<net name="VPORT_2_4" class="0">
<segment>
<wire x1="309.88" y1="60.96" x2="332.74" y2="60.96" width="0.1524" layer="91"/>
<wire x1="332.74" y1="60.96" x2="332.74" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="60.96" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="4"/>
</segment>
<segment>
<wire x1="353.06" y1="416.56" x2="330.2" y2="416.56" width="0.1524" layer="91"/>
<label x="337.82" y="416.56" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04N_2/D3/GCLK15"/>
</segment>
</net>
<net name="VPORT_2_5" class="0">
<segment>
<wire x1="309.88" y1="58.42" x2="335.28" y2="58.42" width="0.1524" layer="91"/>
<wire x1="335.28" y1="58.42" x2="335.28" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="58.42" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="5"/>
</segment>
<segment>
<wire x1="353.06" y1="411.48" x2="330.2" y2="411.48" width="0.1524" layer="91"/>
<label x="337.82" y="411.48" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06P_2/D2/GCLK2"/>
</segment>
</net>
<net name="VPORT_2_6" class="0">
<segment>
<wire x1="309.88" y1="55.88" x2="337.82" y2="55.88" width="0.1524" layer="91"/>
<wire x1="337.82" y1="55.88" x2="337.82" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="55.88" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="6"/>
</segment>
<segment>
<wire x1="353.06" y1="408.94" x2="330.2" y2="408.94" width="0.1524" layer="91"/>
<label x="337.82" y="408.94" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06N_2/D1/GCLK3"/>
</segment>
</net>
<net name="VPORT_2_7" class="0">
<segment>
<wire x1="309.88" y1="53.34" x2="340.36" y2="53.34" width="0.1524" layer="91"/>
<wire x1="340.36" y1="53.34" x2="340.36" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="53.34" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="7"/>
</segment>
<segment>
<wire x1="330.2" y1="388.62" x2="353.06" y2="388.62" width="0.1524" layer="91"/>
<label x="337.82" y="388.62" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L02P_1"/>
</segment>
</net>
<net name="VPORT_2_8" class="0">
<segment>
<wire x1="309.88" y1="50.8" x2="342.9" y2="50.8" width="0.1524" layer="91"/>
<wire x1="342.9" y1="50.8" x2="342.9" y2="73.66" width="0.1524" layer="91"/>
<label x="309.88" y="50.8" size="1.778" layer="95"/>
<pinref part="VPE" gate="G$1" pin="8"/>
</segment>
<segment>
<wire x1="353.06" y1="386.08" x2="330.2" y2="386.08" width="0.1524" layer="91"/>
<label x="337.82" y="386.08" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L02N_1"/>
</segment>
</net>
<net name="T_7SEG_1_6" class="0">
<segment>
<wire x1="294.64" y1="134.62" x2="294.64" y2="114.3" width="0.1524" layer="91"/>
<wire x1="294.64" y1="114.3" x2="281.94" y2="114.3" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="B"/>
<pinref part="R37" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_7" class="0">
<segment>
<wire x1="292.1" y1="134.62" x2="292.1" y2="111.76" width="0.1524" layer="91"/>
<wire x1="292.1" y1="111.76" x2="281.94" y2="111.76" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="A"/>
<pinref part="R39" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_9" class="0">
<segment>
<wire x1="304.8" y1="134.62" x2="304.8" y2="109.22" width="0.1524" layer="91"/>
<wire x1="304.8" y1="109.22" x2="281.94" y2="109.22" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="F"/>
<pinref part="R41" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_10" class="0">
<segment>
<wire x1="307.34" y1="134.62" x2="307.34" y2="106.68" width="0.1524" layer="91"/>
<wire x1="307.34" y1="106.68" x2="281.94" y2="106.68" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="G"/>
<pinref part="R44" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_6" class="0">
<segment>
<wire x1="353.06" y1="134.62" x2="353.06" y2="114.3" width="0.1524" layer="91"/>
<wire x1="353.06" y1="114.3" x2="340.36" y2="114.3" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="B"/>
<pinref part="R38" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_7" class="0">
<segment>
<wire x1="350.52" y1="134.62" x2="350.52" y2="111.76" width="0.1524" layer="91"/>
<wire x1="350.52" y1="111.76" x2="340.36" y2="111.76" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="A"/>
<pinref part="R40" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_9" class="0">
<segment>
<wire x1="363.22" y1="134.62" x2="363.22" y2="109.22" width="0.1524" layer="91"/>
<wire x1="363.22" y1="109.22" x2="340.36" y2="109.22" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="F"/>
<pinref part="R42" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_10" class="0">
<segment>
<wire x1="365.76" y1="134.62" x2="365.76" y2="106.68" width="0.1524" layer="91"/>
<wire x1="365.76" y1="106.68" x2="340.36" y2="106.68" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="G"/>
<pinref part="R45" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_1" class="0">
<segment>
<wire x1="302.26" y1="134.62" x2="302.26" y2="124.46" width="0.1524" layer="91"/>
<wire x1="302.26" y1="124.46" x2="281.94" y2="124.46" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="E"/>
<pinref part="R33" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_2" class="0">
<segment>
<wire x1="299.72" y1="134.62" x2="299.72" y2="127" width="0.1524" layer="91"/>
<wire x1="299.72" y1="127" x2="281.94" y2="127" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="D"/>
<pinref part="R31" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_4" class="0">
<segment>
<wire x1="297.18" y1="134.62" x2="297.18" y2="129.54" width="0.1524" layer="91"/>
<wire x1="297.18" y1="129.54" x2="281.94" y2="129.54" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="C"/>
<pinref part="R29" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_1_5" class="0">
<segment>
<wire x1="309.88" y1="134.62" x2="309.88" y2="132.08" width="0.1524" layer="91"/>
<wire x1="309.88" y1="132.08" x2="281.94" y2="132.08" width="0.1524" layer="91"/>
<pinref part="7SEG1" gate="7SEG" pin="D.P"/>
<pinref part="R27" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_1" class="0">
<segment>
<wire x1="360.68" y1="134.62" x2="360.68" y2="124.46" width="0.1524" layer="91"/>
<wire x1="360.68" y1="124.46" x2="340.36" y2="124.46" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="E"/>
<pinref part="R34" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_2" class="0">
<segment>
<wire x1="358.14" y1="134.62" x2="358.14" y2="127" width="0.1524" layer="91"/>
<wire x1="358.14" y1="127" x2="340.36" y2="127" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="D"/>
<pinref part="R32" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_4" class="0">
<segment>
<wire x1="355.6" y1="134.62" x2="355.6" y2="129.54" width="0.1524" layer="91"/>
<wire x1="355.6" y1="129.54" x2="340.36" y2="129.54" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="C"/>
<pinref part="R30" gate="G$1" pin="2"/>
</segment>
</net>
<net name="T_7SEG_2_5" class="0">
<segment>
<wire x1="368.3" y1="134.62" x2="368.3" y2="132.08" width="0.1524" layer="91"/>
<wire x1="368.3" y1="132.08" x2="340.36" y2="132.08" width="0.1524" layer="91"/>
<pinref part="7SEG2" gate="7SEG" pin="D.P"/>
<pinref part="R28" gate="G$1" pin="2"/>
</segment>
</net>
<net name="7SEG_1_5" class="0">
<segment>
<wire x1="276.86" y1="132.08" x2="256.54" y2="132.08" width="0.1524" layer="91"/>
<label x="256.54" y="132.08" size="1.778" layer="95"/>
<pinref part="R27" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="464.82" x2="330.2" y2="464.82" width="0.1524" layer="91"/>
<label x="337.82" y="464.82" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04P_3/LHCLK2"/>
</segment>
</net>
<net name="7SEG_1_4" class="0">
<segment>
<wire x1="256.54" y1="129.54" x2="276.86" y2="129.54" width="0.1524" layer="91"/>
<label x="256.54" y="129.54" size="1.778" layer="95"/>
<pinref part="R29" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="462.28" x2="330.2" y2="462.28" width="0.1524" layer="91"/>
<label x="337.82" y="462.28" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L04N_3/LHCLK3"/>
</segment>
</net>
<net name="7SEG_1_2" class="0">
<segment>
<wire x1="256.54" y1="127" x2="276.86" y2="127" width="0.1524" layer="91"/>
<label x="256.54" y="127" size="1.778" layer="95"/>
<pinref part="R31" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="457.2" x2="330.2" y2="457.2" width="0.1524" layer="91"/>
<label x="337.82" y="457.2" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05P_3/LHCLK4"/>
</segment>
</net>
<net name="7SEG_1_1" class="0">
<segment>
<wire x1="256.54" y1="124.46" x2="276.86" y2="124.46" width="0.1524" layer="91"/>
<label x="256.54" y="124.46" size="1.778" layer="95"/>
<pinref part="R33" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="454.66" x2="330.2" y2="454.66" width="0.1524" layer="91"/>
<label x="337.82" y="454.66" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05N_3/LHCLK5"/>
</segment>
</net>
<net name="7SEG_1_6" class="0">
<segment>
<wire x1="256.54" y1="114.3" x2="276.86" y2="114.3" width="0.1524" layer="91"/>
<label x="256.54" y="114.3" size="1.778" layer="95"/>
<pinref part="R37" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="325.12" x2="353.06" y2="325.12" width="0.1524" layer="91"/>
<label x="337.82" y="325.12" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_BK0"/>
</segment>
</net>
<net name="7SEG_1_9" class="0">
<segment>
<wire x1="256.54" y1="109.22" x2="276.86" y2="109.22" width="0.1524" layer="91"/>
<label x="256.54" y="109.22" size="1.778" layer="95"/>
<pinref part="R41" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="330.2" x2="353.06" y2="330.2" width="0.1524" layer="91"/>
<label x="337.82" y="330.2" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05P_0/GCLK10"/>
</segment>
</net>
<net name="7SEG_1_10" class="0">
<segment>
<wire x1="256.54" y1="106.68" x2="276.86" y2="106.68" width="0.1524" layer="91"/>
<label x="256.54" y="106.68" size="1.778" layer="95"/>
<pinref part="R44" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="337.82" x2="353.06" y2="337.82" width="0.1524" layer="91"/>
<label x="337.82" y="337.82" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03N_0/GCLK7"/>
</segment>
</net>
<net name="7SEG_2_5" class="0">
<segment>
<wire x1="314.96" y1="132.08" x2="335.28" y2="132.08" width="0.1524" layer="91"/>
<label x="314.96" y="132.08" size="1.778" layer="95"/>
<pinref part="R28" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="474.98" x2="330.2" y2="474.98" width="0.1524" layer="91"/>
<label x="337.82" y="474.98" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L02P_3"/>
</segment>
</net>
<net name="7SEG_2_4" class="0">
<segment>
<wire x1="314.96" y1="129.54" x2="335.28" y2="129.54" width="0.1524" layer="91"/>
<label x="314.96" y="129.54" size="1.778" layer="95"/>
<pinref part="R30" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="472.44" x2="330.2" y2="472.44" width="0.1524" layer="91"/>
<label x="337.82" y="472.44" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L02N_3/VREF_3"/>
</segment>
</net>
<net name="7SEG_2_2" class="0">
<segment>
<wire x1="314.96" y1="127" x2="335.28" y2="127" width="0.1524" layer="91"/>
<label x="314.96" y="127" size="1.778" layer="95"/>
<pinref part="R32" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="469.9" x2="330.2" y2="469.9" width="0.1524" layer="91"/>
<label x="337.82" y="469.9" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03P_3/LHCLK0"/>
</segment>
</net>
<net name="7SEG_2_1" class="0">
<segment>
<wire x1="314.96" y1="124.46" x2="335.28" y2="124.46" width="0.1524" layer="91"/>
<label x="314.96" y="124.46" size="1.778" layer="95"/>
<pinref part="R34" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="353.06" y1="467.36" x2="330.2" y2="467.36" width="0.1524" layer="91"/>
<label x="337.82" y="467.36" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L03N_3/LHCLK1"/>
</segment>
</net>
<net name="7SEG_2_6" class="0">
<segment>
<wire x1="314.96" y1="114.3" x2="335.28" y2="114.3" width="0.1524" layer="91"/>
<label x="314.96" y="114.3" size="1.778" layer="95"/>
<pinref part="R38" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="480.06" x2="353.06" y2="480.06" width="0.1524" layer="91"/>
<label x="337.82" y="480.06" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L01P_3"/>
</segment>
</net>
<net name="7SEG_2_7" class="0">
<segment>
<wire x1="314.96" y1="111.76" x2="335.28" y2="111.76" width="0.1524" layer="91"/>
<label x="314.96" y="111.76" size="1.778" layer="95"/>
<pinref part="R40" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="477.52" x2="353.06" y2="477.52" width="0.1524" layer="91"/>
<label x="337.82" y="477.52" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L01N_3"/>
</segment>
</net>
<net name="7SEG_2_9" class="0">
<segment>
<wire x1="314.96" y1="109.22" x2="335.28" y2="109.22" width="0.1524" layer="91"/>
<label x="314.96" y="109.22" size="1.778" layer="95"/>
<pinref part="R42" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="320.04" x2="353.06" y2="320.04" width="0.1524" layer="91"/>
<label x="337.82" y="320.04" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06N_0/VREF_0"/>
</segment>
</net>
<net name="7SEG_2_10" class="0">
<segment>
<wire x1="314.96" y1="106.68" x2="335.28" y2="106.68" width="0.1524" layer="91"/>
<label x="314.96" y="106.68" size="1.778" layer="95"/>
<pinref part="R45" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="322.58" x2="353.06" y2="322.58" width="0.1524" layer="91"/>
<label x="337.82" y="322.58" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L06P_0"/>
</segment>
</net>
<net name="7SEG_1_7" class="0">
<segment>
<wire x1="256.54" y1="111.76" x2="276.86" y2="111.76" width="0.1524" layer="91"/>
<label x="256.54" y="111.76" size="1.778" layer="95"/>
<pinref part="R39" gate="G$1" pin="1"/>
</segment>
<segment>
<wire x1="330.2" y1="327.66" x2="353.06" y2="327.66" width="0.1524" layer="91"/>
<label x="337.82" y="327.66" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L05N_0/GCLK11"/>
</segment>
</net>
<net name="USB_CON1" class="0">
<segment>
<wire x1="33.02" y1="203.2" x2="33.02" y2="182.88" width="0.1524" layer="91"/>
<wire x1="22.86" y1="203.2" x2="33.02" y2="203.2" width="0.1524" layer="91"/>
<wire x1="33.02" y1="203.2" x2="53.34" y2="203.2" width="0.1524" layer="91"/>
<junction x="33.02" y="203.2"/>
<pinref part="C27" gate="G$1" pin="1"/>
<pinref part="USB" gate="CON_USB" pin="1"/>
<pinref part="FB1" gate="G$1" pin="P$1"/>
</segment>
</net>
<net name="N$11" class="0">
<segment>
<wire x1="147.32" y1="35.56" x2="147.32" y2="33.02" width="0.1524" layer="91"/>
<pinref part="LED2" gate="G$1" pin="C"/>
<pinref part="R36" gate="G$1" pin="1"/>
</segment>
</net>
<net name="LED_2" class="0">
<segment>
<wire x1="147.32" y1="27.94" x2="147.32" y2="25.4" width="0.1524" layer="91"/>
<wire x1="147.32" y1="25.4" x2="129.54" y2="25.4" width="0.1524" layer="91"/>
<label x="129.54" y="25.4" size="1.778" layer="95"/>
<pinref part="R36" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="330.2" y1="393.7" x2="353.06" y2="393.7" width="0.1524" layer="91"/>
<label x="337.82" y="393.7" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L01P_1"/>
</segment>
</net>
<net name="N$2" class="0">
<segment>
<wire x1="114.3" y1="35.56" x2="114.3" y2="33.02" width="0.1524" layer="91"/>
<pinref part="LED1" gate="G$1" pin="C"/>
<pinref part="R35" gate="G$1" pin="1"/>
</segment>
</net>
<net name="N$3" class="0">
<segment>
<wire x1="134.62" y1="378.46" x2="142.24" y2="378.46" width="0.1524" layer="91"/>
<pinref part="R5" gate="G$1" pin="2"/>
<pinref part="IC3" gate="G$1" pin="2A"/>
</segment>
</net>
<net name="N$5" class="0">
<segment>
<wire x1="142.24" y1="370.84" x2="134.62" y2="370.84" width="0.1524" layer="91"/>
<pinref part="IC3" gate="G$1" pin="3A"/>
<pinref part="R7" gate="G$1" pin="2"/>
</segment>
</net>
<net name="N$6" class="0">
<segment>
<wire x1="142.24" y1="363.22" x2="134.62" y2="363.22" width="0.1524" layer="91"/>
<pinref part="IC3" gate="G$1" pin="4A"/>
<pinref part="R8" gate="G$1" pin="2"/>
</segment>
</net>
<net name="LED_1" class="0">
<segment>
<wire x1="114.3" y1="27.94" x2="114.3" y2="25.4" width="0.1524" layer="91"/>
<wire x1="114.3" y1="25.4" x2="96.52" y2="25.4" width="0.1524" layer="91"/>
<label x="96.52" y="25.4" size="1.778" layer="95"/>
<pinref part="R35" gate="G$1" pin="2"/>
</segment>
<segment>
<wire x1="330.2" y1="391.16" x2="353.06" y2="391.16" width="0.1524" layer="91"/>
<label x="337.82" y="391.16" size="1.778" layer="95"/>
<pinref part="IC1" gate="XC3S250E" pin="IO_L01N_1"/>
</segment>
</net>
<net name="N$58" class="0">
<segment>
<wire x1="76.2" y1="297.18" x2="81.28" y2="297.18" width="0.1524" layer="91"/>
<pinref part="U1" gate="UNIT_KC7050B" pin="3"/>
<pinref part="R11" gate="G$1" pin="1"/>
</segment>
</net>
<net name="N$1" class="0">
<segment>
<wire x1="142.24" y1="388.62" x2="119.38" y2="388.62" width="0.1524" layer="91"/>
<wire x1="119.38" y1="388.62" x2="88.9" y2="388.62" width="0.1524" layer="91"/>
<wire x1="142.24" y1="381" x2="119.38" y2="381" width="0.1524" layer="91"/>
<wire x1="142.24" y1="373.38" x2="119.38" y2="373.38" width="0.1524" layer="91"/>
<wire x1="142.24" y1="365.76" x2="119.38" y2="365.76" width="0.1524" layer="91"/>
<wire x1="119.38" y1="365.76" x2="119.38" y2="373.38" width="0.1524" layer="91"/>
<wire x1="119.38" y1="373.38" x2="119.38" y2="381" width="0.1524" layer="91"/>
<wire x1="119.38" y1="381" x2="119.38" y2="388.62" width="0.1524" layer="91"/>
<wire x1="88.9" y1="393.7" x2="88.9" y2="388.62" width="0.1524" layer="91"/>
<wire x1="88.9" y1="388.62" x2="88.9" y2="375.92" width="0.1524" layer="91"/>
<junction x="119.38" y="388.62"/>
<junction x="119.38" y="373.38"/>
<junction x="119.38" y="381"/>
<junction x="88.9" y="388.62"/>
<pinref part="IC3" gate="G$1" pin="1G_"/>
<pinref part="IC3" gate="G$1" pin="2G_"/>
<pinref part="IC3" gate="G$1" pin="3G_"/>
<pinref part="IC3" gate="G$1" pin="4G_"/>
<pinref part="JP1" gate="G$1" pin="P$2"/>
<pinref part="R6" gate="G$1" pin="1"/>
</segment>
</net>
<net name="N$18" class="0">
<segment>
<pinref part="IC6" gate="FT2232L" pin="BDBUS3"/>
<pinref part="IC6" gate="FT2232L" pin="BDBUS2"/>
<wire x1="142.24" y1="162.56" x2="157.48" y2="162.56" width="0.1524" layer="91"/>
<wire x1="157.48" y1="162.56" x2="157.48" y2="165.1" width="0.1524" layer="91"/>
<wire x1="157.48" y1="165.1" x2="142.24" y2="165.1" width="0.1524" layer="91"/>
</segment>
</net>
</nets>
</sheet>
</sheets>
</schematic>
</drawing>
</eagle>
