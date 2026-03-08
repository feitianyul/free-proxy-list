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

最后更新：2026-03-08 14:44:29 UTC（2026-03-08 22:44:29 UTC+8）

**代理总数：48**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 699ms | ✓ 1887ms | ✓ 611ms | ✓ 1188ms | ✓ 1011ms | http |
| 103.84.95.54:7890 | ✓ 1960ms | 否 | ✓ 1035ms | 否 | ✓ 762ms | http |
| 1.231.81.166:3128 | ✓ 1919ms | ✓ 1551ms | ✓ 1147ms | ✓ 1066ms | ✓ 963ms | http |
| 20.27.15.49:8561 | ✓ 765ms | ✓ 1081ms | ✓ 714ms | ✓ 968ms | ✓ 782ms | http |
| 20.210.76.175:8561 | ✓ 767ms | ✓ 1537ms | ✓ 593ms | ✓ 1049ms | ✓ 818ms | http |
| 20.210.76.104:8561 | ✓ 770ms | 否 | ✓ 666ms | ✓ 1015ms | ✓ 803ms | http |
| 20.210.76.178:8561 | ✓ 767ms | 否 | ✓ 668ms | ✓ 1018ms | ✓ 807ms | http |
| 205.209.118.30:3138 | ✓ 1731ms | 否 | ✓ 1645ms | ✓ 1893ms | ✓ 863ms | http |
| 114.4.251.26:8080 | 否 | 否 | ✓ 1376ms | ✓ 1603ms | ✓ 1555ms | http |
| 120.92.212.16:7890 | ✓ 1826ms | ✓ 1350ms | 否 | 否 | ✓ 1622ms | http |
| 8.210.193.14:59394 | ✓ 754ms | ✓ 1377ms | ✓ 774ms | ✓ 967ms | ✓ 771ms | http |
| 5.101.0.233:3128 | ✓ 756ms | ✓ 1985ms | 否 | 否 | ✓ 1754ms | http |
| 47.110.42.192:9003 | ✓ 1845ms | ✓ 1737ms | 否 | ✓ 1762ms | ✓ 1553ms | http |
| 81.70.169.194:80 | ✓ 1057ms | ✓ 1750ms | 否 | 否 | ✓ 1062ms | http |
| 39.104.201.40:7890 | ✓ 1326ms | 否 | 否 | ✓ 1352ms | ✓ 1077ms | http |
| 129.213.162.27:17777 | ✓ 232ms | ✓ 1205ms | 否 | 否 | ✓ 1208ms | http |
| 168.235.110.63:3128 | ✓ 403ms | ✓ 1235ms | ✓ 826ms | ✓ 1054ms | ✓ 1098ms | http |
| 20.78.118.91:8561 | ✓ 1190ms | ✓ 1436ms | ✓ 640ms | ✓ 1371ms | ✓ 729ms | http |
| 20.210.39.153:8561 | ✓ 1186ms | ✓ 1264ms | ✓ 756ms | ✓ 1397ms | ✓ 731ms | http |
| 20.78.26.206:8561 | ✓ 1186ms | 否 | ✓ 680ms | ✓ 1013ms | ✓ 754ms | http |
| 202.155.12.161:443 | ✓ 1780ms | 否 | ✓ 1558ms | 否 | ✓ 1395ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1587ms | ✓ 1566ms | 否 | ✓ 1325ms | http |
| 165.227.5.10:8888 | ✓ 740ms | 否 | ✓ 1887ms | 否 | ✓ 1255ms | http |
| 46.183.25.8:443 | ✓ 1512ms | 否 | ✓ 941ms | ✓ 1406ms | ✓ 1258ms | http |
| 62.113.119.14:8080 | ✓ 952ms | ✓ 1918ms | ✓ 735ms | ✓ 1552ms | ✓ 1087ms | http |
| 59.46.216.131:30001 | ✓ 1161ms | ✓ 1806ms | 否 | ✓ 1478ms | 否 | http |
| 185.243.218.43:49153 | ✓ 1852ms | 否 | ✓ 1580ms | 否 | ✓ 1684ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1359ms | ✓ 1859ms | ✓ 1024ms | http |
| 121.230.8.153:1080 | ✓ 1159ms | 否 | ✓ 1198ms | ✓ 1593ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1757ms | 否 | ✓ 1267ms | ✓ 1181ms | ✓ 1103ms | http |
| 20.27.15.111:8561 | ✓ 1225ms | 否 | ✓ 749ms | ✓ 984ms | ✓ 787ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 1787ms | ✓ 1992ms | ✓ 1481ms | http |
| 109.234.38.35:3128 | ✓ 731ms | 否 | ✓ 1935ms | ✓ 1945ms | ✓ 1610ms | http |
| 45.136.198.40:3128 | ✓ 1118ms | ✓ 1627ms | 否 | 否 | ✓ 1775ms | http |
| 45.140.147.155:1082 | ✓ 559ms | 否 | ✓ 1259ms | ✓ 1669ms | ✓ 1075ms | http |
| 47.77.193.180:1080 | ✓ 778ms | ✓ 1603ms | ✓ 573ms | ✓ 896ms | ✓ 700ms | http |
| 54.222.174.194:80 | 否 | ✓ 1746ms | ✓ 1817ms | ✓ 1983ms | ✓ 1581ms | http |
| 103.215.36.88:13884 | ✓ 1007ms | 否 | ✓ 1082ms | 否 | ✓ 1001ms | http |
| 180.127.149.247:1080 | ✓ 1019ms | 否 | 否 | ✓ 1346ms | ✓ 1477ms | http |
| 103.82.23.118:5196 | ✓ 1975ms | 否 | ✓ 1452ms | 否 | ✓ 1590ms | http |
| 103.215.36.88:16474 | ✓ 1081ms | 否 | ✓ 1110ms | 否 | ✓ 1513ms | http |
| 101.43.255.96:80 | ✓ 1093ms | 否 | ✓ 1183ms | 否 | ✓ 1399ms | http |
| 88.80.150.82:8080 | ✓ 1416ms | ✓ 1843ms | 否 | 否 | ✓ 1870ms | https |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1109ms | ✓ 1386ms | ✓ 1305ms | http |
| 192.71.213.85:9091 | ✓ 752ms | 否 | ✓ 1866ms | ✓ 1524ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1467ms | 否 | ✓ 926ms | ✓ 1079ms | ✓ 847ms | http |
| 178.236.245.59:3128 | ✓ 622ms | 否 | ✓ 755ms | 否 | ✓ 1796ms | http |
| 178.236.245.17:3128 | ✓ 622ms | ✓ 1671ms | ✓ 1082ms | 否 | 否 | http |

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
