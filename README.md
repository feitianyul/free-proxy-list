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

最后更新：2026-05-04 22:53:39 UTC（2026-05-05 06:53:39 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 893ms | ✓ 1518ms | ✓ 1411ms | ✓ 1205ms | ✓ 990ms | http |
| 113.160.132.26:8080 | ✓ 1911ms | ✓ 1924ms | 否 | ✓ 1511ms | ✓ 942ms | http |
| 165.225.113.220:8800 | ✓ 1826ms | 否 | ✓ 755ms | ✓ 1025ms | ✓ 1184ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1733ms | ✓ 1049ms | ✓ 1492ms | 否 | http |
| 38.180.192.119:3128 | ✓ 1411ms | ✓ 869ms | ✓ 593ms | ✓ 847ms | ✓ 1060ms | http |
| 8.211.166.184:8081 | 否 | 否 | ✓ 1207ms | ✓ 1012ms | ✓ 703ms | http |
| 38.180.121.135:10808 | ✓ 1718ms | ✓ 1628ms | ✓ 1616ms | 否 | 否 | http |
| 218.108.131.186:17890 | ✓ 981ms | 否 | 否 | ✓ 1038ms | ✓ 860ms | http |
| 116.80.96.162:3172 | 否 | 否 | ✓ 1944ms | ✓ 1760ms | ✓ 1626ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 1183ms | ✓ 993ms | ✓ 754ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1768ms | ✓ 1767ms | ✓ 1377ms | http |
| 1.231.81.166:3128 | ✓ 896ms | ✓ 1176ms | ✓ 1695ms | ✓ 898ms | ✓ 788ms | http |
| 120.92.212.16:7890 | ✓ 931ms | ✓ 1198ms | ✓ 926ms | ✓ 1189ms | ✓ 1002ms | http |
| 15.223.237.12:3247 | ✓ 1253ms | 否 | ✓ 1392ms | 否 | ✓ 1786ms | http |
| 185.191.236.162:3128 | ✓ 1113ms | ✓ 1837ms | ✓ 1406ms | 否 | ✓ 1630ms | http |
| 120.92.212.16:8890 | ✓ 1199ms | ✓ 1356ms | ✓ 1139ms | ✓ 1457ms | ✓ 978ms | http |
| 18.170.25.193:54929 | ✓ 1844ms | 否 | ✓ 967ms | 否 | ✓ 1647ms | http |
| 59.46.216.131:30001 | ✓ 939ms | 否 | ✓ 977ms | 否 | ✓ 1098ms | http |
| 86.104.72.220:1081 | 否 | 否 | ✓ 1093ms | ✓ 1633ms | ✓ 1100ms | http |
| 103.35.190.69:1082 | 否 | ✓ 1167ms | ✓ 1045ms | 否 | ✓ 1143ms | http |
| 154.64.232.35:8080 | ✓ 1424ms | 否 | ✓ 1282ms | ✓ 1799ms | 否 | http |
| 193.123.250.39:1080 | ✓ 1339ms | 否 | 否 | ✓ 1590ms | ✓ 779ms | http |
| 103.157.200.126:3128 | ✓ 1243ms | 否 | 否 | ✓ 1953ms | ✓ 1609ms | http |
| 47.77.216.82:1080 | ✓ 395ms | 否 | ✓ 1208ms | ✓ 828ms | ✓ 498ms | http |
| 80.92.204.47:1082 | 否 | ✓ 1611ms | ✓ 1921ms | ✓ 1944ms | ✓ 1611ms | http |
| 80.92.204.47:1081 | 否 | ✓ 1837ms | ✓ 1700ms | ✓ 1837ms | ✓ 1721ms | http |
| 45.59.122.132:80 | ✓ 1921ms | 否 | ✓ 1599ms | ✓ 1999ms | 否 | http |
| 101.32.243.189:80 | ✓ 1192ms | ✓ 1493ms | ✓ 1201ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 1211ms | 否 | ✓ 841ms | 否 | ✓ 1868ms | http |
| 52.47.115.41:443 | ✓ 1426ms | 否 | 否 | ✓ 1856ms | ✓ 1695ms | http |
| 141.11.93.27:8080 | ✓ 131ms | ✓ 1100ms | ✓ 366ms | ✓ 1058ms | ✓ 728ms | http |
| 104.129.194.44:8800 | ✓ 711ms | ✓ 1245ms | ✓ 1249ms | 否 | 否 | http |
| 45.125.67.37:8443 | ✓ 981ms | 否 | ✓ 958ms | ✓ 959ms | ✓ 1023ms | http |
| 62.133.60.126:24558 | ✓ 819ms | ✓ 1972ms | ✓ 1932ms | 否 | 否 | http |
| 86.104.72.219:1082 | 否 | 否 | ✓ 1366ms | ✓ 1528ms | ✓ 1963ms | http |
| 101.6.65.112:10080 | ✓ 929ms | ✓ 1134ms | ✓ 1020ms | ✓ 1420ms | ✓ 942ms | http |
| 3.101.133.120:80 | 否 | ✓ 1199ms | ✓ 1421ms | ✓ 924ms | ✓ 731ms | http |
| 59.11.138.152:3128 | ✓ 662ms | ✓ 791ms | ✓ 864ms | ✓ 928ms | ✓ 723ms | http |
| 77.232.142.164:3128 | ✓ 1528ms | 否 | ✓ 1411ms | ✓ 1986ms | ✓ 1397ms | http |
| 13.53.139.178:43924 | ✓ 1658ms | 否 | ✓ 1826ms | 否 | ✓ 1804ms | http |
| 3.71.26.7:47562 | ✓ 807ms | 否 | 否 | ✓ 1863ms | ✓ 1704ms | http |
| 148.230.4.241:999 | ✓ 722ms | ✓ 1518ms | ✓ 522ms | 否 | 否 | http |
| 103.162.16.113:8080 | ✓ 1707ms | 否 | ✓ 1198ms | ✓ 1461ms | 否 | http |
| 86.104.72.220:1082 | ✓ 1580ms | ✓ 1120ms | ✓ 1823ms | ✓ 1419ms | 否 | http |
| 163.61.55.103:1234 | ✓ 1707ms | ✓ 1999ms | ✓ 1752ms | 否 | 否 | http |
| 137.59.47.73:3128 | ✓ 1973ms | 否 | ✓ 1782ms | 否 | ✓ 936ms | http |
| 45.140.147.155:1082 | ✓ 848ms | ✓ 1799ms | ✓ 1226ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1469ms | 否 | ✓ 1413ms | ✓ 1551ms | ✓ 898ms | http |
| 62.113.119.14:8080 | ✓ 996ms | ✓ 1769ms | ✓ 1355ms | 否 | 否 | http |
| 103.120.76.182:8070 | ✓ 1157ms | 否 | ✓ 1668ms | ✓ 1331ms | ✓ 1314ms | http |
| 61.52.131.172:8443 | ✓ 915ms | 否 | ✓ 998ms | ✓ 1213ms | ✓ 1630ms | http |
| 168.194.0.249:252 | 否 | 否 | ✓ 1530ms | ✓ 1848ms | ✓ 1856ms | http |
| 121.230.8.136:1080 | ✓ 973ms | ✓ 1325ms | ✓ 1081ms | ✓ 1488ms | ✓ 1110ms | http |
| 223.16.170.103:80 | ✓ 1288ms | ✓ 1855ms | ✓ 1007ms | ✓ 1259ms | ✓ 1027ms | http |
| 103.39.51.207:8080 | ✓ 1889ms | 否 | ✓ 1743ms | ✓ 1544ms | 否 | http |
| 195.26.243.76:3128 | ✓ 851ms | 否 | ✓ 1271ms | ✓ 1311ms | ✓ 1067ms | http |

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
