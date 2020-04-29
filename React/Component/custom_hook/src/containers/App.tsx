import React, { FC, useState, useEffect } from "react";
import AppComponent from "../components/App";

const useTimer = (
  limitSec: number,
  remain: boolean
): [number, boolean, () => void, () => void] => {
  const [timeLeft, setTimeLeft] = useState(limitSec);
  const [keep, setKeep] = useState(remain);

  const reset = () => {
    setTimeLeft(limitSec);
  };

  const stopOrRestart = () => {
    setKeep(!keep);
  };

  useEffect(() => {
    const tick = () => {
      setTimeLeft((prevTime) => (prevTime === 0 ? limitSec : prevTime - 1));
    };

    let timerId: NodeJS.Timer;
    if (keep) {
      timerId = setInterval(tick, 1000);
    }

    return () => clearInterval(timerId);
  }, [keep, limitSec]);

  return [timeLeft, keep, reset, stopOrRestart];
};

const AppContainer: FC = () => {
  const LIMIT = 120;
  const remain = true;

  const [timeLeft, keep, reset, stopOrRestart] = useTimer(LIMIT, remain);

  return (
    <AppComponent
      timeLeft={timeLeft}
      keep={keep}
      reset={reset}
      stopOrRestart={stopOrRestart}
    />
  );
};

export default AppContainer;
