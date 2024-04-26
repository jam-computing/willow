"use client"

import React, { useEffect, useState } from "react";

export default function Animations() {

    const [stuff, setStuff] = useState(0);

    function do_something() {
        console.log('hey')
    }

    do_something();

    return (
        <div>
            Cosmin Sucks
        </div>
    );
}
