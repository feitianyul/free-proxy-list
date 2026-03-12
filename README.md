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

最后更新：2026-03-12 14:00:00 UTC（2026-03-12 22:00:00 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 918ms | ✓ 868ms | ✓ 244ms | ✓ 801ms | ✓ 516ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 980ms | ✓ 1321ms | ✓ 1188ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1194ms | ✓ 1240ms | ✓ 1018ms | http |
| 115.231.181.40:8128 | ✓ 1538ms | ✓ 1151ms | 否 | 否 | ✓ 1561ms | http |
| 35.225.22.61:80 | 否 | ✓ 1914ms | ✓ 939ms | ✓ 1168ms | ✓ 995ms | http |
| 45.136.198.40:3128 | ✓ 909ms | 否 | ✓ 1756ms | 否 | ✓ 1653ms | http |
| 178.236.245.59:3128 | ✓ 1134ms | 否 | ✓ 1532ms | ✓ 1965ms | ✓ 1563ms | http |
| 45.136.131.47:8443 | ✓ 1947ms | ✓ 1026ms | ✓ 1376ms | ✓ 683ms | ✓ 497ms | http |
| 45.136.131.63:8443 | ✓ 1948ms | ✓ 642ms | ✓ 1782ms | ✓ 685ms | 否 | http |
| 190.9.109.198:999 | ✓ 1027ms | 否 | ✓ 1868ms | ✓ 1985ms | ✓ 1393ms | http |
| 165.227.5.10:8888 | 否 | ✓ 711ms | 否 | ✓ 1162ms | ✓ 580ms | http |
| 116.80.49.159:3172 | ✓ 1568ms | 否 | 否 | ✓ 1895ms | ✓ 1701ms | http |
| 171.251.172.78:5104 | 否 | 否 | ✓ 1701ms | ✓ 1457ms | ✓ 1569ms | http |
| 46.183.25.8:443 | ✓ 609ms | 否 | ✓ 324ms | ✓ 1139ms | ✓ 828ms | http |
| 45.136.130.191:8443 | ✓ 81ms | ✓ 896ms | ✓ 359ms | ✓ 653ms | ✓ 537ms | http |
| 20.210.76.104:8561 | ✓ 466ms | ✓ 881ms | ✓ 519ms | ✓ 803ms | ✓ 639ms | http |
| 20.78.26.206:8561 | ✓ 530ms | ✓ 1014ms | ✓ 533ms | ✓ 811ms | ✓ 696ms | http |
| 103.84.95.54:7890 | ✓ 1259ms | 否 | ✓ 741ms | ✓ 1260ms | ✓ 743ms | http |
| 152.42.213.210:8080 | ✓ 1398ms | 否 | ✓ 1817ms | ✓ 1818ms | 否 | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1575ms | ✓ 1505ms | ✓ 1363ms | http |
| 45.136.130.188:8443 | ✓ 357ms | ✓ 1430ms | ✓ 76ms | ✓ 1077ms | ✓ 1525ms | http |
| 1.231.81.166:3128 | ✓ 630ms | ✓ 1049ms | ✓ 852ms | ✓ 1016ms | ✓ 801ms | http |
| 101.43.255.96:80 | ✓ 1971ms | 否 | ✓ 1563ms | 否 | ✓ 963ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1449ms | ✓ 1372ms | 否 | ✓ 1770ms | http |
| 120.92.212.16:8890 | ✓ 1185ms | 否 | ✓ 1425ms | 否 | ✓ 1048ms | http |
| 171.251.172.78:5106 | 否 | 否 | ✓ 1570ms | ✓ 1440ms | ✓ 1295ms | http |
| 202.155.12.161:443 | ✓ 1213ms | 否 | ✓ 1315ms | ✓ 973ms | 否 | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 940ms | ✓ 1489ms | ✓ 965ms | http |
| 120.92.212.16:7890 | ✓ 1802ms | ✓ 1240ms | ✓ 1891ms | 否 | ✓ 1381ms | http |
| 201.150.116.32:999 | ✓ 672ms | 否 | ✓ 471ms | ✓ 1335ms | 否 | http |
| 116.236.189.93:29999 | ✓ 1879ms | 否 | ✓ 864ms | ✓ 1093ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1372ms | ✓ 1179ms | ✓ 919ms | ✓ 1229ms | 否 | http |
| 106.14.203.63:3333 | 否 | ✓ 1056ms | 否 | ✓ 1135ms | ✓ 1652ms | http |
| 85.208.108.43:2094 | ✓ 1177ms | 否 | ✓ 1731ms | 否 | ✓ 942ms | http |
| 20.210.39.153:8561 | ✓ 1452ms | ✓ 733ms | ✓ 533ms | ✓ 844ms | ✓ 729ms | http |
| 20.27.11.248:8561 | ✓ 1461ms | ✓ 1169ms | ✓ 636ms | ✓ 811ms | ✓ 599ms | http |
| 20.78.118.91:8561 | ✓ 1451ms | 否 | ✓ 456ms | ✓ 743ms | ✓ 610ms | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 283ms | ✓ 845ms | ✓ 564ms | http |
| 45.168.238.193:8443 | ✓ 1541ms | 否 | ✓ 632ms | ✓ 1098ms | ✓ 850ms | http |
| 86.53.183.16:1080 | ✓ 1894ms | 否 | ✓ 1469ms | 否 | ✓ 1529ms | http |
| 81.70.169.194:80 | ✓ 1019ms | ✓ 1518ms | ✓ 1759ms | 否 | 否 | http |
| 202.141.161.53:10808 | ✓ 1028ms | 否 | ✓ 1917ms | 否 | ✓ 1747ms | http |
| 103.166.185.54:3128 | ✓ 1236ms | ✓ 1618ms | ✓ 1201ms | ✓ 1221ms | ✓ 955ms | http |
| 34.101.184.164:3128 | ✓ 856ms | 否 | ✓ 1578ms | ✓ 1436ms | ✓ 1091ms | http |
| 178.236.245.17:3128 | 否 | 否 | ✓ 864ms | ✓ 1777ms | ✓ 1440ms | http |
| 107.173.52.58:7890 | ✓ 1722ms | 否 | ✓ 1811ms | ✓ 1997ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1927ms | 否 | ✓ 923ms | ✓ 1477ms | 否 | http |
| 116.80.96.105:3172 | ✓ 1494ms | 否 | 否 | ✓ 1802ms | ✓ 1662ms | http |
| 47.101.149.27:9010 | ✓ 1751ms | 否 | 否 | ✓ 1474ms | ✓ 1745ms | http |
| 45.140.147.82:1081 | ✓ 1187ms | 否 | 否 | ✓ 1883ms | ✓ 1188ms | http |
| 121.237.181.137:8888 | ✓ 1675ms | ✓ 1221ms | ✓ 856ms | ✓ 1435ms | ✓ 1127ms | http |
| 91.107.141.42:8081 | ✓ 1520ms | 否 | ✓ 1368ms | 否 | ✓ 1848ms | http |
| 14.225.222.185:7890 | 否 | ✓ 1524ms | 否 | ✓ 1673ms | ✓ 846ms | http |
| 106.117.208.101:7890 | ✓ 1057ms | ✓ 1314ms | ✓ 1402ms | 否 | 否 | http |
| 119.18.145.49:20326 | 否 | 否 | ✓ 1727ms | ✓ 1998ms | ✓ 1524ms | http |
| 45.136.130.223:8443 | ✓ 447ms | 否 | ✓ 341ms | ✓ 1217ms | ✓ 828ms | http |
| 20.27.15.111:8561 | 否 | ✓ 876ms | ✓ 643ms | ✓ 787ms | ✓ 602ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 484ms | ✓ 760ms | ✓ 658ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 496ms | ✓ 824ms | ✓ 677ms | http |
| 152.42.213.210:443 | ✓ 1476ms | 否 | ✓ 1642ms | ✓ 1293ms | ✓ 1066ms | http |
| 121.204.158.249:3128 | ✓ 1023ms | ✓ 1644ms | ✓ 1766ms | ✓ 1292ms | ✓ 1712ms | http |
| 120.55.163.237:10086 | ✓ 1910ms | ✓ 1732ms | ✓ 1072ms | 否 | ✓ 922ms | http |
| 140.238.156.12:1080 | ✓ 1768ms | 否 | 否 | ✓ 1531ms | ✓ 1442ms | http |
| 14.225.212.37:7890 | ✓ 1847ms | 否 | ✓ 992ms | 否 | ✓ 1406ms | http |
| 103.3.246.71:3128 | 否 | 否 | ✓ 1327ms | ✓ 1938ms | ✓ 934ms | http |
| 8.140.104.98:3128 | ✓ 880ms | 否 | ✓ 1978ms | ✓ 1967ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1289ms | 否 | ✓ 1225ms | ✓ 1287ms | ✓ 1333ms | http |
| 137.184.1.87:3128 | ✓ 1471ms | 否 | ✓ 418ms | ✓ 880ms | ✓ 689ms | http |
| 103.82.93.100:3128 | ✓ 1011ms | 否 | ✓ 1305ms | ✓ 1243ms | ✓ 1226ms | http |
| 103.82.93.219:3128 | 否 | 否 | ✓ 1641ms | ✓ 1418ms | ✓ 1002ms | http |
| 103.39.51.113:8099 | ✓ 1330ms | 否 | ✓ 1932ms | ✓ 1682ms | 否 | http |
| 209.38.54.154:8443 | ✓ 1708ms | 否 | 否 | ✓ 1731ms | ✓ 1337ms | http |
| 194.116.191.181:3128 | ✓ 1942ms | 否 | ✓ 1755ms | 否 | ✓ 1342ms | http |
| 107.155.65.87:13428 | ✓ 862ms | 否 | 否 | ✓ 1043ms | ✓ 999ms | http |
| 157.245.194.13:8888 | ✓ 1842ms | 否 | ✓ 1991ms | ✓ 1276ms | ✓ 940ms | http |
| 111.48.191.1:7890 | ✓ 772ms | ✓ 913ms | ✓ 874ms | ✓ 921ms | ✓ 743ms | http |

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
