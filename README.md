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

最后更新：2026-05-17 10:40:01 UTC（2026-05-17 18:40:01 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 51.161.50.166:3128 | ✓ 1034ms | ✓ 1832ms | ✓ 825ms | ✓ 1205ms | ✓ 1129ms | http |
| 185.200.188.234:10001 | ✓ 1423ms | 否 | ✓ 1210ms | ✓ 1932ms | ✓ 1628ms | http |
| 218.108.131.186:17890 | ✓ 1013ms | ✓ 1271ms | 否 | 否 | ✓ 1080ms | http |
| 114.214.165.78:10810 | ✓ 1612ms | ✓ 1885ms | ✓ 1788ms | ✓ 1636ms | ✓ 1621ms | http |
| 31.172.78.12:3128 | ✓ 1105ms | 否 | ✓ 708ms | ✓ 1739ms | ✓ 1510ms | http |
| 170.106.136.181:31002 | 否 | 否 | ✓ 951ms | ✓ 1405ms | ✓ 1357ms | http |
| 65.109.125.111:8443 | ✓ 641ms | 否 | ✓ 1363ms | ✓ 1889ms | ✓ 1851ms | http |
| 43.156.90.221:10808 | ✓ 1914ms | 否 | ✓ 955ms | ✓ 1272ms | ✓ 1007ms | http |
| 185.40.77.94:1080 | ✓ 893ms | 否 | ✓ 1861ms | 否 | ✓ 1572ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1815ms | ✓ 1613ms | 否 | ✓ 1194ms | http |
| 129.80.217.21:444 | ✓ 321ms | ✓ 1693ms | ✓ 1924ms | ✓ 1123ms | ✓ 770ms | http |
| 84.47.150.125:1080 | ✓ 1195ms | 否 | 否 | ✓ 1995ms | ✓ 1729ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1093ms | ✓ 1330ms | ✓ 1077ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1583ms | ✓ 1168ms | ✓ 1440ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1192ms | 否 | ✓ 1201ms | ✓ 1537ms | 否 | http |
| 129.80.238.83:444 | ✓ 155ms | ✓ 1929ms | ✓ 152ms | ✓ 905ms | ✓ 1726ms | http |
| 196.204.248.140:8080 | ✓ 1239ms | 否 | ✓ 1021ms | ✓ 1811ms | ✓ 1513ms | http |
| 128.199.254.13:9090 | ✓ 1697ms | 否 | 否 | ✓ 1358ms | ✓ 1051ms | http |
| 128.199.114.189:9090 | ✓ 1691ms | 否 | 否 | ✓ 1684ms | ✓ 1219ms | http |
| 210.76.192.50:10808 | ✓ 1028ms | ✓ 1327ms | ✓ 1110ms | ✓ 1565ms | ✓ 1103ms | http |
| 115.231.181.40:8128 | ✓ 1509ms | ✓ 1597ms | ✓ 1983ms | 否 | ✓ 1461ms | http |
| 148.230.4.241:999 | ✓ 835ms | ✓ 1759ms | ✓ 627ms | ✓ 1820ms | ✓ 1388ms | http |
| 91.242.229.129:8092 | ✓ 904ms | 否 | 否 | ✓ 1777ms | ✓ 1627ms | http |
| 146.190.80.158:9090 | ✓ 911ms | 否 | 否 | ✓ 1441ms | ✓ 1034ms | http |
| 152.42.170.187:9090 | ✓ 937ms | 否 | ✓ 1428ms | ✓ 1302ms | ✓ 1062ms | http |
| 128.199.116.219:9090 | ✓ 991ms | 否 | ✓ 1495ms | ✓ 1275ms | ✓ 1078ms | http |
| 178.63.155.151:8898 | ✓ 1726ms | ✓ 1693ms | ✓ 1112ms | ✓ 1876ms | ✓ 1684ms | http |
| 77.110.107.80:1080 | ✓ 946ms | ✓ 1862ms | ✓ 1229ms | 否 | 否 | http |
| 45.125.67.37:8443 | ✓ 1137ms | 否 | 否 | ✓ 1294ms | ✓ 1333ms | http |
| 147.45.78.89:1080 | ✓ 827ms | 否 | ✓ 1530ms | 否 | ✓ 1285ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1311ms | ✓ 1435ms | ✓ 1143ms | http |
| 190.12.150.244:999 | ✓ 939ms | 否 | ✓ 1389ms | 否 | ✓ 1747ms | http |
| 34.96.238.40:8080 | ✓ 1374ms | ✓ 1523ms | 否 | 否 | ✓ 1640ms | http |
| 38.211.245.35:999 | ✓ 1045ms | 否 | ✓ 955ms | 否 | ✓ 1834ms | http |
| 166.88.55.83:7890 | ✓ 1470ms | ✓ 1313ms | ✓ 808ms | ✓ 1020ms | ✓ 819ms | http |
| 133.18.123.225:26021 | 否 | 否 | ✓ 1412ms | ✓ 1545ms | ✓ 1368ms | http |
| 103.13.235.58:3125 | ✓ 1660ms | 否 | ✓ 1625ms | 否 | ✓ 1735ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1321ms | 否 | ✓ 1324ms | ✓ 1043ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 1508ms | ✓ 1600ms | ✓ 1354ms | http |
| 152.70.91.193:40000 | ✓ 1795ms | 否 | 否 | ✓ 1680ms | ✓ 1444ms | http |
| 3.101.133.120:80 | ✓ 555ms | 否 | ✓ 995ms | ✓ 1386ms | ✓ 1040ms | http |
| 38.211.245.18:999 | ✓ 1039ms | 否 | ✓ 994ms | 否 | ✓ 1883ms | http |
| 103.35.190.69:1081 | ✓ 570ms | 否 | ✓ 782ms | 否 | ✓ 721ms | http |
| 8.219.97.248:80 | ✓ 1583ms | 否 | ✓ 979ms | ✓ 1832ms | 否 | http |
| 1.231.81.166:3128 | ✓ 933ms | ✓ 1483ms | ✓ 886ms | ✓ 1129ms | ✓ 947ms | http |
| 128.199.113.85:9090 | ✓ 1569ms | 否 | ✓ 1344ms | ✓ 1897ms | 否 | http |
| 128.199.121.61:9090 | ✓ 1575ms | 否 | ✓ 1598ms | ✓ 1361ms | ✓ 1007ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1123ms | ✓ 1522ms | ✓ 1153ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1736ms | ✓ 1357ms | 否 | ✓ 1370ms | http |
| 185.21.15.206:3128 | ✓ 487ms | ✓ 1451ms | ✓ 725ms | ✓ 1356ms | ✓ 1438ms | http |
| 103.46.186.181:3125 | ✓ 1979ms | 否 | ✓ 1403ms | ✓ 1867ms | ✓ 1539ms | http |
| 158.160.215.167:8124 | 否 | ✓ 1917ms | ✓ 1595ms | 否 | ✓ 1977ms | http |
| 104.248.151.93:9090 | ✓ 939ms | 否 | 否 | ✓ 1337ms | ✓ 1050ms | http |
| 158.160.215.167:8126 | ✓ 1914ms | ✓ 1803ms | 否 | 否 | ✓ 1565ms | http |
| 20.164.75.153:8080 | ✓ 1654ms | 否 | ✓ 1625ms | 否 | ✓ 1650ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1413ms | ✓ 716ms | ✓ 1787ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1739ms | 否 | ✓ 1953ms | ✓ 1092ms | 否 | http |
| 64.188.77.26:3128 | 否 | ✓ 1396ms | ✓ 741ms | 否 | ✓ 1298ms | http |
| 5.252.33.13:2025 | ✓ 1922ms | 否 | ✓ 1282ms | 否 | ✓ 1721ms | http |
| 2.27.32.81:3128 | 否 | ✓ 1959ms | ✓ 800ms | 否 | ✓ 1507ms | http |
| 38.211.245.131:999 | ✓ 808ms | 否 | ✓ 975ms | 否 | ✓ 1895ms | http |
| 62.60.149.161:3128 | ✓ 776ms | ✓ 1799ms | ✓ 1624ms | 否 | 否 | http |
| 57.129.144.178:40000 | ✓ 468ms | 否 | ✓ 1384ms | ✓ 1644ms | ✓ 1207ms | http |
| 168.222.254.136:8888 | 否 | ✓ 1786ms | 否 | ✓ 1846ms | ✓ 1733ms | http |
| 158.160.215.167:8127 | ✓ 1125ms | ✓ 1618ms | 否 | 否 | ✓ 1849ms | http |
| 185.230.191.240:3128 | ✓ 1106ms | ✓ 1622ms | ✓ 1436ms | ✓ 1674ms | ✓ 1435ms | http |
| 159.223.41.216:9090 | 否 | 否 | ✓ 919ms | ✓ 1658ms | ✓ 1047ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1476ms | ✓ 1452ms | 否 | ✓ 1185ms | http |
| 61.52.131.172:8443 | ✓ 1069ms | ✓ 1423ms | ✓ 1187ms | ✓ 1924ms | ✓ 1187ms | http |
| 138.2.239.213:10010 | ✓ 1485ms | 否 | 否 | ✓ 1384ms | ✓ 1403ms | http |
| 113.108.242.106:8181 | ✓ 1025ms | ✓ 1533ms | ✓ 1035ms | 否 | 否 | http |
| 103.147.152.12:1095 | ✓ 943ms | ✓ 1559ms | ✓ 1112ms | 否 | ✓ 1098ms | http |
| 103.147.152.12:1080 | ✓ 960ms | 否 | 否 | ✓ 1463ms | ✓ 1096ms | http |

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
