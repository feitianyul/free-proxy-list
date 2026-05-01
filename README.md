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

最后更新：2026-05-01 14:15:58 UTC（2026-05-01 22:15:58 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 80.92.204.47:1081 | ✓ 1160ms | 否 | ✓ 712ms | ✓ 1611ms | ✓ 1159ms | http |
| 46.101.95.183:8888 | ✓ 1173ms | 否 | ✓ 1575ms | ✓ 1969ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1871ms | ✓ 1814ms | ✓ 1944ms | ✓ 1278ms | ✓ 1011ms | http |
| 113.160.132.26:8080 | ✓ 1664ms | ✓ 1573ms | ✓ 1348ms | ✓ 1544ms | ✓ 1182ms | http |
| 47.85.51.197:1080 | ✓ 1213ms | ✓ 978ms | ✓ 1735ms | 否 | 否 | http |
| 212.58.132.5:8888 | ✓ 1642ms | 否 | ✓ 1559ms | ✓ 1569ms | ✓ 1382ms | http |
| 45.167.124.71:999 | ✓ 1298ms | 否 | ✓ 1357ms | ✓ 1856ms | ✓ 1385ms | http |
| 34.96.238.40:8080 | ✓ 1376ms | ✓ 1768ms | 否 | 否 | ✓ 1262ms | http |
| 217.76.245.80:999 | ✓ 650ms | 否 | ✓ 954ms | ✓ 1693ms | ✓ 1135ms | http |
| 159.223.225.118:8888 | 否 | 否 | ✓ 352ms | ✓ 1485ms | ✓ 1232ms | http |
| 91.184.241.12:443 | ✓ 751ms | 否 | ✓ 1820ms | ✓ 1905ms | ✓ 1876ms | http |
| 20.127.128.70:8080 | ✓ 1624ms | 否 | ✓ 1161ms | 否 | ✓ 1849ms | http |
| 8.219.97.248:80 | ✓ 1531ms | 否 | ✓ 1545ms | ✓ 1725ms | 否 | http |
| 72.11.151.159:6005 | ✓ 986ms | ✓ 982ms | ✓ 720ms | ✓ 1238ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1169ms | 否 | ✓ 1084ms | ✓ 1568ms | ✓ 1800ms | http |
| 20.78.118.91:8561 | ✓ 1472ms | ✓ 1049ms | ✓ 704ms | ✓ 1051ms | ✓ 810ms | http |
| 20.78.26.206:8561 | ✓ 1453ms | ✓ 1512ms | ✓ 638ms | ✓ 1005ms | ✓ 771ms | http |
| 20.210.39.153:8561 | ✓ 1457ms | ✓ 1765ms | ✓ 618ms | ✓ 979ms | ✓ 810ms | http |
| 62.113.119.14:8080 | ✓ 1027ms | 否 | ✓ 901ms | ✓ 1446ms | ✓ 1069ms | http |
| 86.104.72.220:1082 | ✓ 1337ms | 否 | ✓ 1183ms | ✓ 1151ms | 否 | http |
| 86.104.72.220:1081 | ✓ 503ms | 否 | ✓ 461ms | 否 | ✓ 1008ms | http |
| 89.208.106.138:10808 | ✓ 976ms | ✓ 1446ms | ✓ 707ms | 否 | 否 | http |
| 130.61.174.200:1080 | ✓ 990ms | ✓ 1467ms | ✓ 994ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 957ms | 否 | ✓ 957ms | ✓ 1224ms | ✓ 968ms | http |
| 103.35.190.69:1082 | ✓ 353ms | ✓ 988ms | ✓ 91ms | ✓ 1113ms | ✓ 839ms | http |
| 86.104.72.219:1081 | ✓ 213ms | ✓ 976ms | ✓ 83ms | ✓ 1034ms | ✓ 1731ms | http |
| 152.70.91.193:40000 | ✓ 1751ms | 否 | ✓ 1743ms | ✓ 1857ms | ✓ 1653ms | http |
| 150.107.140.238:3128 | ✓ 1703ms | 否 | ✓ 1638ms | ✓ 1437ms | ✓ 1377ms | http |
| 77.110.119.136:3128 | ✓ 406ms | ✓ 1551ms | ✓ 1096ms | ✓ 1119ms | ✓ 1259ms | http |
| 43.133.44.89:8888 | ✓ 1650ms | 否 | 否 | ✓ 1211ms | ✓ 986ms | http |
| 115.231.181.40:8128 | ✓ 1185ms | ✓ 1971ms | ✓ 1135ms | 否 | 否 | http |
| 101.32.244.83:8080 | ✓ 1620ms | 否 | ✓ 1152ms | ✓ 1741ms | ✓ 1522ms | http |
| 121.43.196.210:8222 | ✓ 1098ms | ✓ 1275ms | ✓ 1086ms | ✓ 1347ms | ✓ 1107ms | http |
| 121.43.196.213:8222 | ✓ 1084ms | ✓ 1319ms | ✓ 1054ms | ✓ 1332ms | ✓ 1134ms | http |
| 121.230.8.55:1080 | 否 | 否 | ✓ 1065ms | ✓ 1629ms | ✓ 1179ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 834ms | ✓ 1106ms | ✓ 795ms | http |
| 45.140.147.82:1081 | ✓ 1772ms | ✓ 1870ms | 否 | 否 | ✓ 1740ms | http |
| 101.32.243.189:80 | ✓ 1671ms | ✓ 1669ms | ✓ 1732ms | ✓ 1709ms | ✓ 1865ms | http |
| 91.217.81.131:1080 | ✓ 1046ms | 否 | ✓ 715ms | ✓ 1925ms | 否 | http |
| 183.238.3.150:7897 | ✓ 1084ms | ✓ 1340ms | ✓ 1134ms | ✓ 1323ms | ✓ 1030ms | http |
| 3.101.133.120:80 | ✓ 869ms | ✓ 1635ms | ✓ 1898ms | 否 | ✓ 1748ms | http |
| 152.70.84.108:8080 | ✓ 1932ms | 否 | ✓ 1926ms | ✓ 1552ms | ✓ 1223ms | http |
| 103.82.23.118:5303 | 否 | 否 | ✓ 1601ms | ✓ 1988ms | ✓ 1730ms | http |
| 94.131.118.39:1081 | ✓ 1060ms | ✓ 1243ms | ✓ 719ms | 否 | 否 | http |
| 94.158.219.111:3128 | ✓ 1081ms | 否 | ✓ 703ms | ✓ 1855ms | ✓ 1629ms | http |
| 91.108.243.203:3128 | ✓ 1074ms | 否 | ✓ 842ms | ✓ 1754ms | 否 | http |
| 20.27.11.248:8561 | 否 | ✓ 1072ms | ✓ 848ms | ✓ 960ms | ✓ 798ms | http |
| 20.27.13.35:8561 | ✓ 1624ms | 否 | ✓ 645ms | ✓ 988ms | ✓ 757ms | http |
| 45.153.231.229:8080 | ✓ 1923ms | ✓ 1961ms | ✓ 1164ms | 否 | 否 | http |
| 20.210.76.175:8561 | ✓ 1615ms | ✓ 1467ms | ✓ 1013ms | ✓ 1284ms | ✓ 1410ms | http |
| 20.27.14.220:8561 | ✓ 1414ms | ✓ 1492ms | ✓ 623ms | ✓ 1145ms | ✓ 1005ms | http |
| 45.129.141.143:3128 | ✓ 614ms | 否 | 否 | ✓ 1727ms | ✓ 1724ms | http |
| 72.11.150.178:6005 | ✓ 1332ms | 否 | ✓ 930ms | ✓ 1953ms | ✓ 1241ms | http |
| 210.223.44.230:3128 | ✓ 1111ms | 否 | ✓ 1046ms | 否 | ✓ 1559ms | http |
| 2.27.40.180:1080 | ✓ 1048ms | ✓ 1103ms | ✓ 1218ms | ✓ 1780ms | ✓ 1255ms | http |
| 43.167.237.94:3128 | 否 | ✓ 1185ms | ✓ 1024ms | 否 | ✓ 779ms | http |
| 61.52.131.172:8443 | ✓ 1065ms | ✓ 1388ms | ✓ 1256ms | 否 | ✓ 1114ms | http |
| 183.232.248.73:7890 | ✓ 1080ms | ✓ 1326ms | ✓ 1106ms | ✓ 1290ms | ✓ 1040ms | http |
| 103.176.96.213:8080 | 否 | 否 | ✓ 1816ms | ✓ 1640ms | ✓ 1630ms | http |

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
