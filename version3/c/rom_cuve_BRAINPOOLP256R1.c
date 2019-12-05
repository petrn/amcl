#include "arch.h"
#include "ecp_BRAINPOOL.h"

/* Brainpool Curve  */
/* Note that the original curve has been transformed to an isomorphic curve with A=-3 */

#if CHUNK==16

#error Not supported

#endif

#if CHUNK==32

const int CURVE_A_BRAINPOOLP256R1= -3;
const int CURVE_Cof_I_BRAINPOOLP256R1= 1;
const BIG_256_28 CURVE_Cof_BRAINPOOLP256R1= {0x1,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
const int CURVE_B_I_BRAINPOOLP256R1= 0;
const BIG_256_28 CURVE_B_BRAINPOOLP256R1= {0xF8C07B6,0xCCDC18F,0x7E1CE6B,0x16295CF,0xCBF9584,0xD9BBD77,0x4F330B5,0xE94A4B4,0x6DC5C6C,0x2};
const BIG_256_28 CURVE_Order_BRAINPOOLP256R1= {0x74856A7,0x1E0E829,0x1A6F790,0x7AA3B56,0xD718C39,0x909D838,0xC3E660A,0xA1EEA9B,0x9FB57DB,0xA};
const BIG_256_28 CURVE_Gx_BRAINPOOLP256R1= {0xACE3262,0x4453BD9,0xD23C23A,0x27E1E3B,0x7AFB9DE,0x2FFC81B,0xB2C4B48,0xCB7E57C,0xBD2AEB9,0x8};
const BIG_256_28 CURVE_Gy_BRAINPOOLP256R1= {0xF046997,0x1D54C72,0xD8E545C,0x45132DE,0xDC9C277,0x1A14611,0xD97F846,0xC3DAC4F,0x47EF835,0x5};
#endif

#if CHUNK==64

const int CURVE_A_BRAINPOOLP256R1= -3;
const int CURVE_Cof_I_BRAINPOOLP256R1= 1;
const BIG_256_56 CURVE_Cof_BRAINPOOLP256R1= {0x1L,0x0L,0x0L,0x0L,0x0L};
const int CURVE_B_I_BRAINPOOLP256R1= 0;
const BIG_256_56 CURVE_B_BRAINPOOLP256R1= {0xCCDC18FF8C07B6L,0x16295CF7E1CE6BL,0xD9BBD77CBF9584L,0xE94A4B44F330B5L,0x26DC5C6CL};
const BIG_256_56 CURVE_Order_BRAINPOOLP256R1= {0x1E0E82974856A7L,0x7AA3B561A6F790L,0x909D838D718C39L,0xA1EEA9BC3E660AL,0xA9FB57DBL};
const BIG_256_56 CURVE_Gx_BRAINPOOLP256R1= {0x4453BD9ACE3262L,0x27E1E3BD23C23AL,0x2FFC81B7AFB9DEL,0xCB7E57CB2C4B48L,0x8BD2AEB9L};
const BIG_256_56 CURVE_Gy_BRAINPOOLP256R1= {0x1D54C72F046997L,0x45132DED8E545CL,0x1A14611DC9C277L,0xC3DAC4FD97F846L,0x547EF835L};
#endif
