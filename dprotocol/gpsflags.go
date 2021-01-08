package dprotocol

// GPSFlags validity bits combination information.
//
// nolint: lll
//
//  +----------+-----------+----------------------+------------------+-------------+-------------------------------------------------------------------------------------------------------------------------------------------+
//  | Validity | Is fix    | Position in          | Satellites used  | Is possible | Additional info                                                                                                                           |
//  | bits 7-4 | valid now | snapshot current/old | for fix          | situation?  |                                                                                                                                           |
//  +----------+-----------+----------------------+------------------+-------------+-------------------------------------------------------------------------------------------------------------------------------------------+
//  |   0000   |    No     |        None          | None             |     Yes     | No fix nor old coordinates available                                                                                                      |
//  |   0001   |    No     |        Old           | Only GPS         |     Yes     | No fix. Snapshot contains old position and for calculation it only GPS satellites were used                                               |
//  |   0010   |    No     |        Old           | Only Glonass     |     Yes     | No fix. Snapshot contains old position and for calculation it only Glonass satellites were used                                           |
//  |   0011   |    No     |        Old           | GPS and Glonass  |     Yes     | No fix. Snapshot contains old position and for calculation it GPS and Glonass satellites were used                                        |
//  |   0100   |    ---    |        ---           | None             |     No      | Not possible                                                                                                                              |
//  |   0101   |    No     |        Old           | Only GPS         |     Yes     | GPS data in the snapshot is valid but when the snapshot was created there was no fix                                                      |
//  |   0110   |    No     |        Old           | Only Glonass     |     Yes     | GPS data in the snapshot is valid but when the snapshot was created there was no fix                                                      |
//  |   0111   |    No     |        Old           | GPS and Glonass  |     Yes     | GPS data in the snapshot is valid but when the snapshot was created there was no fix                                                      |
//  |   1000   |    Yes    |        Old           | None             |     Yes     | Possible at startup when positions are read from COP buffer. There is no position in the snapshot, but device has valid fix at the moment |
//  |   1001   |    Yes    |        Old           | Only GPS         |     Yes     | Possible at startup when positions are read from COP buffer. Position in the snapshot are old, but device has valid fix at the moment     |
//  |   1010   |    Yes    |        Old           | Only Glonass     |     Yes     | Possible at startup when positions are read from COP buffer. Position in the snapshot are old, but device has valid fix at the moment     |
//  |   1011   |    Yes    |        Old           | GPS and Glonass  |     Yes     | Possible at startup when positions are read from COP buffer. Position in the snapshot are old, but device has valid fix at the moment     |
//  |   1100   |    ---    |        ---           | None             |     No      | Not possible                                                                                                                              |
//  |   1101   |    Yes    |        Current       | Only GPS         |     Yes     | Normal valid fix                                                                                                                          |
//  |   1110   |    Yes    |        Current       | Only Glonass     |     Yes     | Normal valid fix                                                                                                                          |
//  |   1111   |    Yes    |        Current       | GPS and Glonass  |     Yes     | Normal valid fix                                                                                                                          |
//  +----------+-----------+----------------------+------------------+-------------+-------------------------------------------------------------------------------------------------------------------------------------------+
//  +---------+---------+------------+-------------------+------------+---------------+-----------+-------+
//  |  Bit7   |  Bit 6  |   Bit 5    |       Bit 4       |   Bit 3    |     Bit 2     |   Bit 1   | Bit 0 |
//  +---------+---------+------------+-------------------+------------+---------------+-----------+-------+
//  | currfix | postfix | glonassfix | gpsFixWhenGlonass | GPSJamming | AGPS validity | max speed | speed |
//  +---------+---------+------------+-------------------+------------+---------------+-----------+-------+
type GPSFlags uint8

// Has returns true if the provided flag is set (logical high).
func (g GPSFlags) Has(flag GPSFlag) bool {
	return GPSFlag(g)&flag == flag
}
