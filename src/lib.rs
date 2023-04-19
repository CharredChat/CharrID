use std::{
    thread,
    time::{Duration, SystemTime, UNIX_EPOCH},
};
use ux::u48;

pub struct CharredProcess {
    pub process_id: u8,
    last_time: SystemTime,
    incr: u8,
}

impl CharredProcess {
    pub fn new(id: u8) -> Self {
        Self {
            process_id: id,
            last_time: SystemTime::now(),
            incr: u8::MIN,
        }
    }

    pub fn generate(&mut self) -> CharrID {
        let since = SystemTime::now()
            .duration_since(self.last_time)
            .expect("Time travel is not legal.");
        if since.as_millis() == 0 {
            //Cooldown
            if self.incr == u8::MAX {
                thread::sleep(Duration::from_millis(1));
                self.incr = u8::MIN;
            } else {
                self.incr = self.incr + 1;
            }
        } else {
            self.incr = u8::MIN;
        }
        self.last_time = SystemTime::now();
        CharrID::new(self.process_id, self.incr)
    }
}

#[derive(Copy, Clone)]
pub struct CharrID {
    timestamp: u48,
    process: u8,
    incr: u8,
}

impl CharrID {
    fn new(process: u8, incr: u8) -> Self {
        let now = SystemTime::now();
        let time_since_epoch = now
            .duration_since(UNIX_EPOCH)
            .expect("Time travel is not legal.");
        Self {
            timestamp: u48::new(time_since_epoch.as_millis() as u64),
            process: process,
            incr: incr,
        }
    }
    pub fn as_u64(&self) -> u64 {
        (Into::<u64>::into(self.timestamp) << 16)
            | (Into::<u64>::into(self.process) << 8)
            | Into::<u64>::into(self.incr)
    }
    pub fn get_timestamp(&self) -> u64 {
        Into::<u64>::into(self.timestamp)
    }
}
