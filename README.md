**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-22 06:20:18 UTC（2026-04-22 14:20:18 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 526ms | ✓ 1137ms | 否 | ✓ 1037ms | ✓ 734ms | http |
| 1.231.81.166:3128 | ✓ 1826ms | ✓ 1311ms | ✓ 1045ms | ✓ 1182ms | ✓ 870ms | http |
| 188.246.224.49:7890 | ✓ 1096ms | ✓ 1551ms | ✓ 1810ms | ✓ 1936ms | ✓ 1597ms | http |
| 81.30.156.115:8080 | ✓ 1067ms | ✓ 1754ms | ✓ 1516ms | ✓ 1891ms | ✓ 1289ms | http |
| 152.42.208.139:8118 | ✓ 1062ms | 否 | ✓ 1040ms | ✓ 1386ms | ✓ 1224ms | http |
| 113.160.132.26:8080 | ✓ 1966ms | 否 | ✓ 1138ms | ✓ 1882ms | ✓ 1174ms | http |
| 78.11.96.22:8888 | 否 | ✓ 1796ms | ✓ 1567ms | 否 | ✓ 1693ms | http |
| 46.101.95.183:8888 | ✓ 880ms | ✓ 1904ms | ✓ 341ms | ✓ 1770ms | ✓ 1162ms | http |
| 177.93.132.244:3128 | ✓ 799ms | 否 | ✓ 660ms | 否 | ✓ 1721ms | http |
| 94.131.118.129:1081 | ✓ 1552ms | 否 | ✓ 1725ms | 否 | ✓ 1974ms | http |
| 115.231.181.40:8128 | ✓ 1127ms | 否 | ✓ 1662ms | 否 | ✓ 1507ms | http |
| 91.233.223.147:3128 | ✓ 1139ms | 否 | ✓ 1657ms | 否 | ✓ 1483ms | http |
| 91.99.15.45:2095 | ✓ 1072ms | ✓ 1368ms | ✓ 1539ms | ✓ 1794ms | 否 | http |
| 210.45.76.58:42992 | ✓ 1207ms | ✓ 1501ms | ✓ 1647ms | ✓ 1520ms | ✓ 1206ms | http |
| 59.46.216.131:30001 | ✓ 1220ms | 否 | 否 | ✓ 1645ms | ✓ 1353ms | http |
| 210.77.29.244:6478 | ✓ 1046ms | 否 | 否 | ✓ 1379ms | ✓ 1107ms | http |
| 152.32.132.190:7890 | ✓ 941ms | ✓ 1340ms | ✓ 891ms | ✓ 1129ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1388ms | 否 | ✓ 1202ms | 否 | ✓ 1501ms | http |
| 35.225.22.61:80 | ✓ 471ms | ✓ 1090ms | 否 | ✓ 1126ms | ✓ 804ms | http |
| 104.248.243.244:3128 | ✓ 1992ms | 否 | ✓ 1475ms | ✓ 1843ms | ✓ 1171ms | http |
| 34.71.229.255:3128 | ✓ 784ms | 否 | ✓ 1533ms | ✓ 1801ms | ✓ 1751ms | http |
| 213.32.85.26:3128 | ✓ 896ms | 否 | ✓ 1073ms | 否 | ✓ 1383ms | http |
| 218.108.131.186:17890 | ✓ 855ms | ✓ 1625ms | ✓ 1308ms | 否 | ✓ 1139ms | http |
| 217.182.195.221:30000 | ✓ 1780ms | 否 | 否 | ✓ 1714ms | ✓ 1551ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1397ms | ✓ 1779ms | ✓ 1487ms | http |
| 45.12.151.226:2829 | ✓ 932ms | 否 | ✓ 1285ms | ✓ 1929ms | 否 | http |
| 103.93.93.95:8181 | 否 | 否 | ✓ 1565ms | ✓ 1948ms | ✓ 1725ms | http |
| 84.47.150.126:1080 | ✓ 932ms | 否 | 否 | ✓ 1775ms | ✓ 1454ms | http |
| 161.97.184.191:8080 | ✓ 798ms | ✓ 1956ms | 否 | 否 | ✓ 1952ms | http |
| 45.153.231.229:8080 | ✓ 1746ms | 否 | ✓ 1203ms | ✓ 1840ms | 否 | http |
| 185.191.236.162:3128 | ✓ 1100ms | ✓ 1610ms | ✓ 1469ms | 否 | 否 | http |
| 89.208.106.138:10808 | ✓ 576ms | ✓ 1494ms | 否 | 否 | ✓ 1366ms | http |
| 118.193.44.243:3128 | ✓ 855ms | 否 | ✓ 1030ms | ✓ 1135ms | ✓ 827ms | http |
| 160.250.134.143:3128 | 否 | 否 | ✓ 1098ms | ✓ 1493ms | ✓ 1238ms | http |
| 121.230.8.138:1080 | ✓ 1340ms | 否 | ✓ 1221ms | ✓ 1883ms | ✓ 1320ms | http |
| 114.237.77.245:1080 | ✓ 1112ms | ✓ 1479ms | ✓ 1528ms | ✓ 1448ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1636ms | ✓ 1629ms | ✓ 1367ms | http |
| 43.155.86.199:8080 | ✓ 837ms | 否 | ✓ 927ms | ✓ 1033ms | ✓ 827ms | http |
| 82.114.228.67:1080 | ✓ 1347ms | 否 | ✓ 1015ms | ✓ 1621ms | ✓ 1124ms | http |
| 34.101.184.164:3128 | ✓ 1795ms | 否 | ✓ 1285ms | ✓ 1465ms | ✓ 1201ms | http |
| 41.196.16.231:1981 | ✓ 1069ms | 否 | ✓ 1666ms | 否 | ✓ 1931ms | http |
| 45.140.147.82:1082 | ✓ 940ms | ✓ 1069ms | ✓ 1357ms | ✓ 1445ms | ✓ 1276ms | http |
| 45.140.147.82:1081 | ✓ 1704ms | 否 | ✓ 403ms | ✓ 1152ms | ✓ 963ms | http |
| 51.92.173.133:30118 | ✓ 1066ms | 否 | ✓ 1402ms | ✓ 1983ms | ✓ 1382ms | http |
| 63.179.134.206:36851 | 否 | 否 | ✓ 861ms | ✓ 1883ms | ✓ 1885ms | http |
| 85.190.99.143:443 | ✓ 1974ms | ✓ 1661ms | ✓ 1866ms | 否 | ✓ 1744ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1528ms | ✓ 848ms | 否 | ✓ 885ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1104ms | ✓ 780ms | ✓ 1026ms | ✓ 840ms | http |
| 20.78.26.206:8561 | 否 | ✓ 1453ms | ✓ 694ms | ✓ 1072ms | ✓ 786ms | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 677ms | ✓ 995ms | ✓ 824ms | http |
| 91.193.240.157:9877 | ✓ 1080ms | 否 | ✓ 1729ms | 否 | ✓ 1591ms | http |
| 144.31.25.69:21064 | ✓ 1093ms | 否 | ✓ 999ms | 否 | ✓ 1999ms | http |
| 168.144.75.9:3128 | ✓ 974ms | 否 | 否 | ✓ 1814ms | ✓ 1649ms | http |
| 20.127.128.70:8080 | ✓ 1512ms | ✓ 1746ms | ✓ 1418ms | ✓ 1789ms | ✓ 1481ms | http |
| 45.186.6.104:3128 | ✓ 1625ms | ✓ 1727ms | ✓ 1945ms | 否 | 否 | http |
| 43.129.182.137:1080 | ✓ 822ms | ✓ 1732ms | ✓ 922ms | ✓ 1059ms | ✓ 823ms | http |
| 103.85.113.66:9999 | ✓ 584ms | 否 | ✓ 1243ms | ✓ 1975ms | 否 | http |
| 52.210.57.38:14825 | ✓ 590ms | 否 | ✓ 1294ms | 否 | ✓ 1528ms | http |
| 45.129.141.143:3128 | ✓ 1186ms | ✓ 1838ms | ✓ 1408ms | ✓ 1982ms | ✓ 1379ms | http |
| 2.27.40.180:1080 | ✓ 1048ms | ✓ 1881ms | ✓ 1227ms | ✓ 1894ms | ✓ 1417ms | http |
| 47.84.73.61:1080 | 否 | 否 | ✓ 967ms | ✓ 1314ms | ✓ 1013ms | http |
| 152.70.91.193:40000 | ✓ 1834ms | 否 | 否 | ✓ 1819ms | ✓ 1715ms | http |
| 135.125.232.193:3128 | ✓ 1813ms | ✓ 1534ms | ✓ 1539ms | 否 | 否 | http |
| 108.181.201.118:1234 | ✓ 398ms | ✓ 1160ms | ✓ 202ms | 否 | 否 | http |
| 159.223.225.118:8888 | ✓ 1997ms | ✓ 1556ms | ✓ 1461ms | 否 | ✓ 1196ms | http |
| 43.132.188.134:443 | ✓ 844ms | 否 | ✓ 1152ms | 否 | ✓ 1898ms | http |
| 103.113.70.189:1081 | ✓ 330ms | ✓ 1571ms | ✓ 553ms | ✓ 1350ms | ✓ 1098ms | http |
| 147.45.186.28:3128 | 否 | ✓ 1584ms | ✓ 685ms | ✓ 1876ms | 否 | http |
| 103.113.70.189:1082 | ✓ 462ms | ✓ 1166ms | ✓ 211ms | 否 | ✓ 821ms | http |
| 102.212.44.151:12354 | ✓ 1919ms | 否 | ✓ 1407ms | 否 | ✓ 1983ms | http |
| 180.103.19.99:1080 | ✓ 1197ms | ✓ 1485ms | ✓ 1274ms | 否 | 否 | http |
| 54.37.72.89:80 | ✓ 1517ms | 否 | ✓ 1754ms | 否 | ✓ 1948ms | http |
| 61.52.131.172:8443 | ✓ 954ms | ✓ 1152ms | ✓ 984ms | ✓ 1175ms | ✓ 969ms | http |
| 104.247.51.76:3128 | ✓ 306ms | 否 | ✓ 1192ms | ✓ 1117ms | ✓ 761ms | http |
| 223.16.170.103:80 | ✓ 1137ms | 否 | ✓ 1318ms | ✓ 1678ms | 否 | http |
| 103.191.92.157:1009 | 否 | 否 | ✓ 1460ms | ✓ 1485ms | ✓ 1243ms | http |
| 8.140.104.98:3128 | 否 | ✓ 1301ms | ✓ 970ms | ✓ 1261ms | ✓ 1048ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
